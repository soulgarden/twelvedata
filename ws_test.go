package twelvedata //nolint: testpackage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/request"
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
					ws.priceEvents != nil &&
					ws.statusEvents != nil &&
					ws.errorEvents != nil &&
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
func TestWS_ConnectAndSubscribe(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		symbols            []string
		subscribeStatusMsg string
		priceEventMsg      string
		wantErr            bool
		expectedPrice      float64
	}{
		{
			name:    "successful connect, subscribe and price event",
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
			wantErr:       false,
			expectedPrice: 172.8700,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				upgrader := websocket.Upgrader{
					Error: func(w http.ResponseWriter, _ *http.Request, status int, reason error) {
						http.Error(w, reason.Error(), status)
					},
					CheckOrigin: func(_ *http.Request) bool {
						return true
					},
				}

				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					t.Errorf("WebSocket upgrade failed: %v", err)
					return
				}
				defer func() {
					if err := conn.Close(); err != nil {
						t.Logf("failed to close websocket connection: %v", err)
					}
				}()

				// Read the subscribe message
				_, body, err := conn.ReadMessage()
				if err != nil {
					t.Errorf("Failed to read subscribe message: %v", err)
					return
				}

				t.Logf("Received subscribe message: %s", body)

				// Send subscribe status
				err = conn.WriteMessage(websocket.TextMessage, []byte(tt.subscribeStatusMsg))
				if err != nil {
					t.Errorf("Failed to write subscribe status: %v", err)
					return
				}

				// Wait a bit then send price event
				time.Sleep(10 * time.Millisecond)

				err = conn.WriteMessage(websocket.TextMessage, []byte(tt.priceEventMsg))
				if err != nil {
					t.Errorf("Failed to write price event: %v", err)
					return
				}

				// Keep connection alive for a bit
				time.Sleep(100 * time.Millisecond)
			}))
			defer server.Close()

			// Create WebSocket client
			logger := zerolog.New(os.Stdout)
			wsURL := strings.Replace(server.URL, "http://", "ws://", 1)
			cfg := &Conf{
				BaseWSURL: strings.TrimPrefix(wsURL, "ws://"),
				APIKey:    "test-key",
				WebSocket: WebSocket{PriceURL: "/quotes/price"},
			}

			// Override URL scheme to use ws:// for test server
			ws := NewWS(cfg, &logger, nil)
			ws.url.Scheme = "ws"
			defer func() { _ = ws.Close() }()

			// Connect to server
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			err := ws.Connect(ctx)
			if err != nil {
				t.Fatalf("Failed to connect: %v", err)
			}

			// Subscribe to symbols
			err = ws.Subscribe(tt.symbols)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// Wait for events
			select {
			case statusEvent := <-ws.ConsumeStatusEvents():
				t.Logf("Received status event: %+v", statusEvent)
				if statusEvent.Status != "ok" {
					t.Errorf("Expected status 'ok', got %s", statusEvent.Status)
				}
			case <-time.After(1 * time.Second):
				t.Error("Timeout waiting for status event")
				return
			}

			select {
			case priceEvent := <-ws.ConsumePriceEvents():
				if priceEvent.Price != tt.expectedPrice {
					t.Errorf("Expected price %f, got %f", tt.expectedPrice, priceEvent.Price)
				}
				if priceEvent.Symbol != "AAPL" {
					t.Errorf("Expected symbol AAPL, got %s", priceEvent.Symbol)
				}
			case <-time.After(1 * time.Second):
				t.Error("Timeout waiting for price event")
			}
		})
	}
}

func TestWS_ExtendedSubscription(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		// Read the subscribe message and validate its format
		_, body, err := conn.ReadMessage()
		if err != nil {
			t.Errorf("Failed to read subscribe message: %v", err)
			return
		}

		// Verify the message format includes extended symbol information
		if !strings.Contains(string(body), "exchange") || !strings.Contains(string(body), "mic_code") {
			t.Errorf("Expected extended format in message: %s", body)
			return
		}

		// Send status response
		statusMsg := `{
			"event":"subscribe-status",
			"status":"ok","success":[
				{
					"symbol":"AAPL",
					"exchange":"NASDAQ",
					"country":"United States",
					"type":"Common Stock"
				}
			],"fails":null
		}`
		_ = conn.WriteMessage(websocket.TextMessage, []byte(statusMsg))

		time.Sleep(50 * time.Millisecond)
	}))
	defer server.Close()

	// Create WebSocket client
	logger := zerolog.New(os.Stdout)
	wsURL := strings.Replace(server.URL, "http://", "ws://", 1)
	cfg := &Conf{
		BaseWSURL: strings.TrimPrefix(wsURL, "ws://"),
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/quotes/price"},
	}

	ws := NewWS(cfg, &logger, nil)
	ws.url.Scheme = "ws"
	defer func() { _ = ws.Close() }()

	// Connect to server
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := ws.Connect(ctx)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Test extended subscription format
	extendedSymbols := []request.WSSymbolExtended{
		{
			Symbol:   "AAPL",
			Exchange: "NASDAQ",
			MicCode:  "XNAS",
			Type:     "Common Stock",
		},
	}

	err = ws.SubscribeExtended(extendedSymbols)
	if err != nil {
		t.Errorf("SubscribeExtended() failed: %v", err)
	}

	// Wait for status event
	select {
	case statusEvent := <-ws.ConsumeStatusEvents():
		if statusEvent.Status != "ok" {
			t.Errorf("Expected status 'ok', got %s", statusEvent.Status)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for status event")
	}
}

func TestWS_ResetAndHeartbeat(t *testing.T) {
	t.Parallel()

	messageCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("WebSocket upgrade failed: %v", err)
			return
		}
		defer func() { _ = conn.Close() }()

		for {
			_, body, err := conn.ReadMessage()
			if err != nil {
				return
			}
			messageCount++

			// Verify reset and heartbeat messages
			if messageCount == 1 && !strings.Contains(string(body), `"action":"reset"`) {
				t.Errorf("Expected reset message, got: %s", body)
			}
			if messageCount == 2 && !strings.Contains(string(body), `"action":"heartbeat"`) {
				t.Errorf("Expected heartbeat message, got: %s", body)
			}

			// Send acknowledgment for reset
			if messageCount == 1 {
				statusMsg := `{"event":"subscribe-status","status":"ok","success":[],"fails":null}`
				_ = conn.WriteMessage(websocket.TextMessage, []byte(statusMsg))
			}

			if messageCount >= 2 {
				return
			}
		}
	}))
	defer server.Close()

	// Create WebSocket client
	logger := zerolog.New(os.Stdout)
	wsURL := strings.Replace(server.URL, "http://", "ws://", 1)
	cfg := &Conf{
		BaseWSURL: strings.TrimPrefix(wsURL, "ws://"),
		APIKey:    "test-key",
		WebSocket: WebSocket{PriceURL: "/quotes/price"},
	}

	ws := NewWS(cfg, &logger, nil)
	ws.url.Scheme = "ws"
	defer func() { _ = ws.Close() }()

	// Connect to server
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := ws.Connect(ctx)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	// Test reset functionality
	err = ws.Reset()
	if err != nil {
		t.Errorf("Reset() failed: %v", err)
	}

	// Test heartbeat functionality
	err = ws.SendHeartbeat()
	if err != nil {
		t.Errorf("SendHeartbeat() failed: %v", err)
	}

	// Give time for messages to be processed
	time.Sleep(100 * time.Millisecond)

	if messageCount < 2 {
		t.Errorf("Expected at least 2 messages (reset + heartbeat), got %d", messageCount)
	}
}
