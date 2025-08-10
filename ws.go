package twelvedata

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/response"
)

type WS struct {
	url      *url.URL
	eventsCh chan response.PriceEvent
	dialer   *websocket.Dialer
	logger   *zerolog.Logger
}

func NewWS(cfg *Conf, logger *zerolog.Logger, dialer *websocket.Dialer) *WS {
	//nolint: varnamelen
	ws := &WS{
		url: &url.URL{
			Scheme:   "wss",
			Host:     cfg.BaseWSURL,
			Path:     cfg.WebSocket.PriceURL,
			RawQuery: "apikey=" + cfg.APIKey,
		},
		eventsCh: make(chan response.PriceEvent, dictionary.EventsChSize),
		logger:   logger,
	}

	if dialer == nil {
		ws.dialer = websocket.DefaultDialer
	} else {
		ws.dialer = dialer
	}

	return ws
}

func (ws *WS) Subscribe(ctx context.Context, symbols []string) error {
	conn, resp, err := ws.dialer.DialContext(ctx, ws.url.String(), nil)
	if err != nil {
		ws.logger.Err(err).Str("url", ws.url.String()).Msg("dial")

		return &WSConnectionError{
			URL:     ws.url.String(),
			Message: "Failed to establish WebSocket connection",
			Cause:   err,
		}
	}

	defer conn.Close()

	defer resp.Body.Close()

	ticker := time.NewTicker(dictionary.PingPeriod)
	defer ticker.Stop()

	done := make(chan struct{})

	go ws.read(conn, done)

	err = ws.sendSubscribeMessage(conn, symbols)
	if err != nil {
		return &WSSubscriptionError{
			Symbols: symbols,
			Message: "Failed to send subscription message",
			Cause:   err,
		}
	}

	return ws.ping(ctx, conn, done)
}

func (ws *WS) Consume() <-chan response.PriceEvent {
	return ws.eventsCh
}

func (ws *WS) read(conn *websocket.Conn, done chan<- struct{}) {
	defer close(done)

	for {
		_, message, err := conn.ReadMessage()

		ws.logger.Err(err).Bytes("message", message).Msg("read message")

		if err != nil {
			return
		}

		event := response.PriceEvent{}

		if err := json.Unmarshal(message, &event); err != nil {
			ws.logger.Err(err).Bytes("val", message).Msg("unmarshall")
			// Log but continue processing other messages
			continue
		}

		if event.Event == dictionary.PriceEventType {
			ws.eventsCh <- event
		}
	}
}

func (ws *WS) ping(ctx context.Context, conn *websocket.Conn, done <-chan struct{}) error {
	ticker := time.NewTicker(dictionary.PingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)

			ws.logger.Err(err).Msg("write close connection message")

			if err != nil {
				return &WSMessageError{
					Message: "Failed to write close message",
					Cause:   err,
				}
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
		case <-ticker.C:
			err := conn.SetWriteDeadline(time.Now().Add(dictionary.WriteWait))
			if err != nil {
				ws.logger.Err(err).Msg("set write deadline")
			}

			ws.logger.Debug().Msg("set write deadline")

			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return &WSMessageError{
					Message: "Failed to write ping message",
					Cause:   err,
				}
			}
		}
	}
}

func (ws *WS) sendSubscribeMessage(conn *websocket.Conn, symbols []string) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(`
		{
			"action": "subscribe",
			"params": {
				"symbols": "`+strings.Join(symbols, ",")+`"
			}
		}`,
	))
	if err != nil {
		ws.logger.Err(err).Msg("write message to ws")

		return &WSMessageError{
			Message: "Failed to write subscribe message",
			Cause:   err,
		}
	}

	return nil
}
