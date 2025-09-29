package twelvedata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

// WS represents a WebSocket client for real-time price data streaming from the Twelve Data API.
// It supports comprehensive WebSocket functionality including subscription management,
// heartbeat handling, and multiple event types with graceful shutdown.
type WS struct {
	url    *url.URL
	dialer *websocket.Dialer
	logger *zerolog.Logger

	// Connection state
	conn      *websocket.Conn
	connMu    sync.RWMutex
	connected atomic.Bool

	// Generic event channels (owned by this struct)
	priceEvents  *EventChannel[response.WSPriceEvent]
	statusEvents *EventChannel[response.WSSubscribeStatusEvent]
	errorEvents  *EventChannel[response.WSErrorEvent]

	// Message parsing (optimized single-pass parser)
	parser *wsMessageParser

	g        *errgroup.Group
	ctx      context.Context //nolint:containedctx // Required for goroutine lifecycle management with errgroup
	cancel   context.CancelFunc
	shutdown chan struct{}
	closed   atomic.Bool
}

// NewWS creates a new WebSocket client instance configured for the Twelve Data API.
// If dialer is nil, the default WebSocket dialer will be used.
// The client uses generic event channels for type-safe event handling.
func NewWS(cfg *Conf, logger *zerolog.Logger, dialer *websocket.Dialer) *WS {
	if dialer == nil {
		dialer = websocket.DefaultDialer
	}

	//nolint: varnamelen
	ws := &WS{
		url: &url.URL{
			Scheme:   "wss",
			Host:     cfg.BaseWSURL,
			Path:     cfg.WebSocket.PriceURL,
			RawQuery: "apikey=" + cfg.APIKey,
		},
		dialer: dialer,
		logger: logger,

		// Initialize generic event channels
		priceEvents:  NewEventChannel[response.WSPriceEvent](dictionary.EventsChSize),
		statusEvents: NewEventChannel[response.WSSubscribeStatusEvent](dictionary.EventsChSize),
		errorEvents:  NewEventChannel[response.WSErrorEvent](dictionary.EventsChSize),

		// Initialize optimized message parser
		parser: newWSMessageParser(),
	}

	return ws
}

// Connect establishes a WebSocket connection and starts message handling.
// It uses errgroup for proper goroutine lifecycle management and graceful shutdown.
func (ws *WS) Connect(ctx context.Context) error {
	ws.connMu.Lock()
	defer ws.connMu.Unlock()

	if ws.conn != nil {
		return fmt.Errorf("WebSocket connection already established")
	}

	conn, resp, err := ws.dialer.DialContext(ctx, ws.url.String(), nil)
	if err != nil {
		ws.logger.Err(err).Str("url", ws.url.String()).Msg("dial")

		return &WSConnectionError{
			URL:     ws.url.String(),
			Message: "Failed to establish WebSocket connection",
			Cause:   err,
		}
	}

	ws.conn = conn
	ws.connected.Store(true)

	// Initialize shutdown mechanism
	ws.shutdown = make(chan struct{})

	// Setup errgroup with context for structured concurrency
	ctx, cancel := context.WithCancel(ctx)
	ws.cancel = cancel
	ws.g, ws.ctx = errgroup.WithContext(ctx)

	defer func() {
		if err := resp.Body.Close(); err != nil {
			ws.logger.Warn().Err(err).Msg("failed to close response body")
		}
	}()

	// Start goroutines with errgroup
	ws.g.Go(func() error {
		return ws.messageReader()
	})

	ws.g.Go(func() error {
		return ws.heartbeatSender()
	})

	return nil
}

// Subscribe subscribes to price events for the specified symbols.
// Supports both simple string format and extended format with exchange parameters.
func (ws *WS) Subscribe(symbols []string) error {
	return ws.sendSubscribeMessage(symbols, false)
}

// SubscribeExtended subscribes to price events using extended symbol format.
func (ws *WS) SubscribeExtended(symbols []request.WSSymbolExtended) error {
	return ws.sendSubscribeExtendedMessage(symbols, false)
}

// Unsubscribe removes subscriptions for the specified symbols.
func (ws *WS) Unsubscribe(symbols []string) error {
	return ws.sendSubscribeMessage(symbols, true)
}

// UnsubscribeExtended removes subscriptions using extended symbol format.
func (ws *WS) UnsubscribeExtended(symbols []request.WSSymbolExtended) error {
	return ws.sendSubscribeExtendedMessage(symbols, true)
}

// Reset clears all current subscriptions.
func (ws *WS) Reset() error {
	resetMsg := request.WSResetRequest{
		Action: request.WSActionReset,
	}

	return ws.sendJSONMessage(resetMsg)
}

// SendHeartbeat sends a heartbeat message to maintain connection stability.
func (ws *WS) SendHeartbeat() error {
	heartbeatMsg := request.WSHeartbeatRequest{
		Action: request.WSActionHeartbeat,
	}

	return ws.sendJSONMessage(heartbeatMsg)
}

// Close gracefully closes the WebSocket connection and stops all goroutines.
// It follows Go best practices for clean shutdown with timeout and proper resource cleanup.
func (ws *WS) Close() error {
	// Atomic check-and-set to ensure Close is called only once
	if !ws.closed.CompareAndSwap(false, true) {
		return nil // Already closed
	}

	ws.logger.Debug().Msg("WebSocket closing initiated")

	var closeErr error

	// Step 1: Signal graceful shutdown to goroutines
	if ws.shutdown != nil {
		close(ws.shutdown)
	}

	// Step 2: Wait for goroutines to finish gracefully (with timeout)
	if ws.g != nil {
		done := make(chan error, 1)
		go func() {
			done <- ws.g.Wait()
		}()

		// Wait for graceful shutdown with timeout
		select {
		case err := <-done:
			if err != nil && !errors.Is(err, context.Canceled) {
				ws.logger.Err(err).Msg("goroutine error during graceful shutdown")
				closeErr = err
			}
		case <-time.After(2 * time.Second):
			// Timeout reached, force shutdown
			ws.logger.Warn().Msg("timeout waiting for goroutines, forcing shutdown")
			if ws.cancel != nil {
				ws.cancel()
			}
			// Still wait a bit more for forced shutdown
			select {
			case <-done:
			case <-time.After(1 * time.Second):
				ws.logger.Error().Msg("goroutines did not terminate after force shutdown")
			}
		}
	}

	// Step 3: Close WebSocket connection
	ws.connMu.Lock()
	if ws.conn != nil {
		// Send close frame with timeout
		if err := ws.conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
			ws.logger.Warn().Err(err).Msg("Failed to set write deadline")
		}
		err := ws.conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		if err != nil {
			ws.logger.Warn().Err(err).Msg("failed to send close message")
		}

		// Close the underlying connection
		if err := ws.conn.Close(); err != nil && closeErr == nil {
			closeErr = err
		}

		ws.conn = nil
		ws.connected.Store(false)
	}
	ws.connMu.Unlock()

	// Step 4: Cancel context if not already cancelled
	if ws.cancel != nil {
		ws.cancel()
	}

	// Step 5: Close event channels to unblock consumers
	// This ensures channels are closed even if messageReader never started
	// or exited abnormally. EventChannel.Close() is idempotent (uses sync.Once)
	// so it's safe to call even if messageReader's defer already closed them.
	if ws.priceEvents != nil {
		ws.priceEvents.Close()
	}
	if ws.statusEvents != nil {
		ws.statusEvents.Close()
	}
	if ws.errorEvents != nil {
		ws.errorEvents.Close()
	}

	ws.logger.Debug().Msg("WebSocket closed")
	return closeErr
}

// ConsumePriceEvents returns a read-only channel for receiving price events.
// The channel will be closed when the WebSocket connection is terminated.
func (ws *WS) ConsumePriceEvents() <-chan response.WSPriceEvent {
	return ws.priceEvents.Channel()
}

// ConsumeStatusEvents returns a read-only channel for receiving subscription status events.
// The channel will be closed when the WebSocket connection is terminated.
func (ws *WS) ConsumeStatusEvents() <-chan response.WSSubscribeStatusEvent {
	return ws.statusEvents.Channel()
}

// ConsumeErrorEvents returns a read-only channel for receiving error events.
// The channel will be closed when the WebSocket connection is terminated.
func (ws *WS) ConsumeErrorEvents() <-chan response.WSErrorEvent {
	return ws.errorEvents.Channel()
}

// Consume returns a read-only channel for receiving price events (legacy method for backward compatibility).
// The channel will be closed when the WebSocket connection is terminated.
func (ws *WS) Consume() <-chan response.WSPriceEvent {
	return ws.priceEvents.Channel()
}

// messageReader handles incoming WebSocket messages and routes them to appropriate channels.
// As the owner of the event channels, it's responsible for closing them on exit.
func (ws *WS) messageReader() error {
	// As the channel owner, ensure channels are closed when we exit
	defer func() {
		// Handle panics from websocket library
		if r := recover(); r != nil {
			ws.logger.Debug().Interface("panic", r).Msg("messageReader: recovered from panic")
		}

		ws.priceEvents.Close()
		ws.statusEvents.Close()
		ws.errorEvents.Close()
		ws.logger.Debug().Msg("messageReader: event channels closed")
	}()

	ws.connMu.RLock()
	conn := ws.conn
	ws.connMu.RUnlock()

	if conn == nil {
		ws.logger.Error().Msg("messageReader: connection is nil")
		return fmt.Errorf("connection is nil")
	}

	for {
		select {
		case <-ws.ctx.Done():
			ws.logger.Debug().Msg("messageReader: context cancelled")
			return ws.ctx.Err()
		case <-ws.shutdown:
			ws.logger.Debug().Msg("messageReader: graceful shutdown")
			return nil
		default:
			// Check if connection is still valid
			ws.connMu.RLock()
			currentConn := ws.conn
			ws.connMu.RUnlock()

			if currentConn == nil || currentConn != conn {
				ws.logger.Debug().Msg("messageReader: connection changed or closed")
				return nil
			}

			// Set a read timeout to make reads non-blocking
			if err := conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond)); err != nil {
				ws.logger.Debug().Err(err).Msg("Failed to set read deadline")
			}

			// Protected read with panic recovery
			var message []byte
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						ws.logger.Debug().Interface("panic", r).Msg("messageReader: recovered from read panic")
						err = fmt.Errorf("websocket read panic: %v", r)
					}
				}()
				_, message, err = conn.ReadMessage()
			}()

			if err != nil {
				// Check if it's a timeout - continue reading
				var netErr net.Error
				if errors.As(err, &netErr) {
					continue
				}

				// Check for normal closure
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					ws.logger.Debug().Msg("messageReader: connection closed normally")
					return nil
				}

				// Check for abnormal closure or other websocket errors
				if websocket.IsUnexpectedCloseError(err) {
					ws.logger.Debug().Msg("messageReader: connection closed unexpectedly")
					return nil
				}

				// Check for generic websocket close errors
				if websocket.IsCloseError(err) {
					ws.logger.Debug().Msg("messageReader: websocket close error")
					return nil
				}

				// Catch-all for websocket library errors (like repeated read)
				errStr := err.Error()
				if strings.Contains(errStr, "repeated read") ||
					strings.Contains(errStr, "closed network connection") ||
					strings.Contains(errStr, "use of closed") {
					ws.logger.Debug().Msg("messageReader: connection is closed")
					return nil
				}

				// Real unexpected error - log and return
				ws.logger.Err(err).Msg("read message error")
				return fmt.Errorf("read message: %w", err)
			}

			ws.logger.Debug().Bytes("message", message).Msg("received message")
			ws.routeMessage(message)
		}
	}
}

// routeMessage parses incoming messages and routes them to appropriate event channels.
// Uses optimized single-pass parsing to eliminate double JSON unmarshaling.
// This reduces CPU usage and memory allocations compared to the previous implementation.
func (ws *WS) routeMessage(message []byte) {
	// Check if we're shutting down before parsing
	select {
	case <-ws.shutdown:
		return // Don't process events during shutdown
	case <-ws.ctx.Done():
		return // Don't process events if context is cancelled
	default:
	}

	// Step 1: Parse message in single pass (extracts both event type and all data)
	eventType, err := ws.parser.parseMessage(message)
	if err != nil {
		ws.logger.Err(err).Bytes("message", message).Msg("failed to parse message")
		return
	}

	// Step 2: Route to appropriate channel based on parsed event type
	switch eventType {
	case response.WSEventPrice:
		priceEvent := ws.parser.getPriceEvent()
		if !ws.priceEvents.Send(ws.ctx, priceEvent) {
			ws.logger.Warn().Msg("failed to send price event (channel full or closed)")
		}

	case response.WSEventSubscribeStatus:
		statusEvent := ws.parser.getSubscribeStatusEvent()
		if !ws.statusEvents.Send(ws.ctx, statusEvent) {
			ws.logger.Warn().Msg("failed to send status event (channel full or closed)")
		}

	case response.WSEventError:
		errorEvent := ws.parser.getErrorEvent()
		if !ws.errorEvents.Send(ws.ctx, errorEvent) {
			ws.logger.Warn().Msg("failed to send error event (channel full or closed)")
		}

	default:
		ws.logger.Warn().Str("event", string(eventType)).Bytes("message", message).Msg("unknown event type")
	}

	// Step 3: Reset parser for reuse (clean up cached data)
	ws.parser.reset()
}

// heartbeatSender sends periodic heartbeat messages to maintain connection stability.
func (ws *WS) heartbeatSender() error {
	ticker := time.NewTicker(dictionary.HeartbeatPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ws.ctx.Done():
			ws.logger.Debug().Msg("heartbeatSender: context cancelled")
			return ws.ctx.Err()
		case <-ws.shutdown:
			ws.logger.Debug().Msg("heartbeatSender: graceful shutdown")
			return nil
		case <-ticker.C:
			if err := ws.SendHeartbeat(); err != nil {
				ws.logger.Err(err).Msg("failed to send heartbeat")
				return fmt.Errorf("send heartbeat: %w", err)
			}
			ws.logger.Debug().Msg("heartbeat sent")
		}
	}
}

// sendJSONMessage sends a JSON message over the WebSocket connection.
func (ws *WS) sendJSONMessage(message interface{}) error {
	ws.connMu.Lock()
	defer ws.connMu.Unlock()

	// Check closed state under lock to prevent TOCTOU race condition
	if ws.IsClosed() || ws.conn == nil {
		return fmt.Errorf("WebSocket is closed or not connected")
	}

	data, err := json.Marshal(message)
	if err != nil {
		return &WSMessageError{
			Message: "Failed to marshal JSON message",
			Cause:   err,
		}
	}

	if err := ws.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return &WSMessageError{
			Message: "Failed to write message",
			Data:    data,
			Cause:   err,
		}
	}

	ws.logger.Debug().Bytes("message", data).Msg("sent message")
	return nil
}

// sendSubscribeMessage sends a subscription/unsubscription message for simple string symbols.
func (ws *WS) sendSubscribeMessage(symbols []string, isUnsubscribe bool) error {
	action := request.WSActionSubscribe
	if isUnsubscribe {
		action = request.WSActionUnsubscribe
	}

	req := request.WSSubscribeRequest{
		Action: action,
		Params: &request.WSSubscribeParams{
			Symbols: strings.Join(symbols, ","),
		},
	}

	return ws.sendJSONMessage(req)
}

// sendSubscribeExtendedMessage sends a subscription/unsubscription message for extended symbol format.
func (ws *WS) sendSubscribeExtendedMessage(symbols []request.WSSymbolExtended, isUnsubscribe bool) error {
	action := request.WSActionSubscribe
	if isUnsubscribe {
		action = request.WSActionUnsubscribe
	}

	req := request.WSSubscribeRequest{
		Action: action,
		Params: &request.WSSubscribeParams{
			Symbols: symbols,
		},
	}

	return ws.sendJSONMessage(req)
}

// IsConnected returns true if the WebSocket is connected.
func (ws *WS) IsConnected() bool {
	return ws.connected.Load()
}

// IsClosed returns true if the WebSocket has been closed.
func (ws *WS) IsClosed() bool {
	return ws.closed.Load()
}
