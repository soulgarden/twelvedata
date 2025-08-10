package twelvedata //nolint: testpackage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/response"
)

func TestNewWS(t *testing.T) {
	type args struct {
		cfg    *Conf
		logger *zerolog.Logger
		dialer *websocket.Dialer
	}

	tests := []struct {
		name string
		args args
		want func(*WS) bool
	}{
		{
			name: "new websocket with default dialer",
			args: args{
				cfg: &Conf{
					BaseWSURL: "ws.twelvedata.com",
					APIKey:    "test-key",
					WebSocket: WebSocket{PriceURL: "/v1/quotes/price"},
				},
				logger: &zerolog.Logger{},
				dialer: nil,
			},
			want: func(ws *WS) bool {
				return ws != nil &&
					ws.url.Host == "ws.twelvedata.com" &&
					ws.url.Path == "/v1/quotes/price" &&
					ws.eventsCh != nil &&
					ws.dialer == websocket.DefaultDialer
			},
		},
		{
			name: "new websocket with custom dialer",
			args: args{
				cfg: &Conf{
					BaseWSURL: "ws.twelvedata.com",
					APIKey:    "custom-key",
					WebSocket: WebSocket{PriceURL: "/v1/quotes/price"},
				},
				logger: &zerolog.Logger{},
				dialer: &websocket.Dialer{},
			},
			want: func(ws *WS) bool {
				return ws != nil &&
					ws.url.RawQuery == "apikey=custom-key" &&
					ws.dialer != nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWS(tt.args.cfg, tt.args.logger, tt.args.dialer)
			if !tt.want(got) {
				t.Errorf("NewWS() validation failed for %v", got)
			}
		})
	}
}

// nolint: gocognit
func TestWS_Subscribe(t *testing.T) {
	t.Parallel()

	type fields struct {
		url      *url.URL
		eventsCh chan response.PriceEvent
		dialer   *websocket.Dialer
		logger   zerolog.Logger
	}

	type args struct {
		ctx                context.Context
		symbols            []string
		subscribeStatusMsg string
		priceEventMsg      string
	}

	tests := []struct {
		name string
		fields
		args
		wantErr       bool
		expectedPrice float64
	}{
		{
			name: "successful subscribe and price event",
			fields: fields{
				url: &url.URL{
					Scheme: "ws",
					Host:   "127.0.0.1",
					Path:   "/quotes/price",
				},
				eventsCh: make(chan response.PriceEvent, 1),
				dialer:   websocket.DefaultDialer,
				logger:   zerolog.New(os.Stdout),
			},
			args: args{
				ctx:     t.Context(),
				symbols: []string{"AAPL"},
				subscribeStatusMsg: `{
					"event":"subscribe-status",
					"status":"ok","success":[
						{
							"symbol":"AAPL",
							"exchange":"NASDAQ",
							"country":"United States",
							"type":"Common Stock"
						}
					],"fails":null
				}`,
				priceEventMsg: `{
					"event":"price",
					"symbol":"AAPL",
					"currency":"USD",
					"exchange":"NASDAQ",
					"type":"Common Stock",
					"timestamp":1643972766,
					"price":172.8700
				}`,
			},
			wantErr:       false,
			expectedPrice: 172.8700,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				upgrader := websocket.Upgrader{
					Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
						http.Error(w, reason.Error(), status)
					},
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				}

				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					t.Errorf("WebSocket upgrade failed: %v", err)

					return
				}
				defer conn.Close()

				// Read the subscribe message
				_, body, err := conn.ReadMessage()
				if err != nil {
					t.Errorf("Failed to read subscribe message: %v", err)

					return
				}

				tt.fields.logger.Debug().Bytes("body", body).Msg("read subscribe message")

				// Send subscribe status
				err = conn.WriteMessage(websocket.TextMessage, []byte(tt.args.subscribeStatusMsg))
				if err != nil {
					t.Errorf("Failed to write subscribe status: %v", err)

					return
				}

				// Wait a bit then send price event
				time.Sleep(10 * time.Millisecond)

				err = conn.WriteMessage(websocket.TextMessage, []byte(tt.args.priceEventMsg))
				if err != nil {
					t.Errorf("Failed to write price event: %v", err)

					return
				}

				// Keep connection alive for a bit
				time.Sleep(100 * time.Millisecond)
			}))
			defer server.Close()

			// Create WebSocket client with test server URL
			wsURL := &url.URL{
				Scheme: "ws",
				Host:   strings.TrimPrefix(server.URL, "http://"),
				Path:   "/quotes/price",
			}

			ws := &WS{
				url:      wsURL,
				eventsCh: tt.fields.eventsCh,
				dialer:   tt.fields.dialer,
				logger:   &tt.fields.logger,
			}

			// Create context with timeout
			ctx, cancel := context.WithTimeout(tt.args.ctx, 2*time.Second)
			defer cancel()

			// Start subscription in goroutine
			errCh := make(chan error, 1)
			go func() {
				errCh <- ws.Subscribe(ctx, tt.args.symbols)
			}()

			// Wait for price event or timeout
			select {
			case resp := <-ws.Consume():
				if resp.Price != tt.expectedPrice {
					t.Errorf("Expected price %f, got %f", tt.expectedPrice, resp.Price)
				}

				if resp.Symbol != "AAPL" {
					t.Errorf("Expected symbol AAPL, got %s", resp.Symbol)
				}
			case err := <-errCh:
				if (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !tt.wantErr {
					t.Error("Subscribe finished before receiving price event")
				}
			case <-time.After(3 * time.Second):
				t.Error("Test timeout - no price event received")
			}

			// Cancel context to stop subscription
			cancel()
		})
	}
}
