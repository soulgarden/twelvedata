package twelvedata

import (
	"context"
	"encoding/json"
	"fmt"
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
	eventsCh chan *response.PriceEvent
	logger   *zerolog.Logger
}

// nolint: exhaustivestruct
func NewWS(cfg *Conf, logger *zerolog.Logger) *WS {
	return &WS{
		url: &url.URL{
			Scheme:   "wss",
			Host:     cfg.BaseWSURL,
			Path:     cfg.WebSocket.PriceURL,
			RawQuery: "apikey=" + cfg.APIKey,
		},
		eventsCh: make(chan *response.PriceEvent, dictionary.EventsChSize),
		logger:   logger,
	}
}

func (w *WS) Subscribe(ctx context.Context, symbols []string) error {
	conn, _, err := websocket.DefaultDialer.Dial(w.url.String(), nil)
	if err != nil {
		w.logger.Err(err).Str("url", w.url.String()).Msg("dial")

		return fmt.Errorf("dial ws: %w", err)
	}

	defer conn.Close()

	ticker := time.NewTicker(dictionary.PingPeriod)
	defer ticker.Stop()

	done := make(chan struct{})

	go w.read(conn, done)

	err = w.sendSubscribeMessage(conn, symbols)
	if err != nil {
		return err
	}

	return w.ping(ctx, conn, done)
}

func (w *WS) Consume() <-chan *response.PriceEvent {
	return w.eventsCh
}

func (w *WS) read(conn *websocket.Conn, done chan<- struct{}) {
	defer close(done)

	for {
		_, message, err := conn.ReadMessage()

		w.logger.Err(err).Bytes("message", message).Msg("read message")

		if err != nil {
			return
		}

		// nolint: exhaustivestruct
		event := &response.PriceEvent{}

		if err := json.Unmarshal(message, event); err != nil {
			w.logger.Err(err).Bytes("val", message).Msg("unmarshall")

			continue
		}

		if event.Event == dictionary.PriceEventType {
			w.eventsCh <- event
		}
	}
}

func (w *WS) ping(ctx context.Context, conn *websocket.Conn, done <-chan struct{}) error {
	ticker := time.NewTicker(dictionary.PingPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)

			w.logger.Err(err).Msg("write close connection message")

			if err != nil {
				return fmt.Errorf("write close message: %w", err)
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
		case <-ticker.C:
			err := conn.SetWriteDeadline(time.Now().Add(dictionary.WriteWait))

			w.logger.Err(err).Msg("set write deadline")

			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return fmt.Errorf("write message: %w", err)
			}
		}
	}
}

func (w *WS) sendSubscribeMessage(conn *websocket.Conn, symbols []string) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(`
		{
			"action": "subscribe",
			"params": {
				"symbols": "`+strings.Join(symbols, ",")+`"
			}
		}`,
	))
	if err != nil {
		w.logger.Err(err).Msg("write message to ws")

		return fmt.Errorf("write message: %w", err)
	}

	return nil
}
