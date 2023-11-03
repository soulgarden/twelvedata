package twelvedata //nolint: testpackage

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/fasthttp/websocket"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/response"
)

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
		name          string
		fields        fields
		args          args
		wantErr       bool
		expectedPrice float64
	}{
		{
			name: "1",
			fields: fields{

				url: &url.URL{
					Scheme: "ws",
					Host:   "127.0.0.1",
					Path:   "/quotes/price",
				},
				eventsCh: make(chan response.PriceEvent),
				dialer:   websocket.DefaultDialer,
				logger:   zerolog.New(os.Stdout),
			},
			args: args{
				ctx:     context.Background(),
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(cw http.ResponseWriter, sr *http.Request) {
				upgrader := websocket.Upgrader{
					Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
						http.Error(w, reason.Error(), status)
					},
				}

				ws, err := upgrader.Upgrade(cw, sr, nil)
				if err != nil {
					t.Error(err)
					_ = ws.Close()
				}

				go func() {
					for {
						_, body, err := ws.ReadMessage()
						if err != nil {
							t.Error(err)
						}

						tt.fields.logger.Debug().Bytes("body", body).Msg("read message")
					}
				}()

				err = ws.WriteMessage(websocket.TextMessage, []byte(tt.args.subscribeStatusMsg))
				if err != nil {
					t.Error(err)
				}

				err = ws.WriteMessage(websocket.TextMessage, []byte(tt.args.priceEventMsg))
				if err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()

			ws := &WS{
				url: &url.URL{
					Scheme: "ws",
					Host:   strings.Replace(server.URL, "http://", "", 1),
					Path:   "/quotes/price",
				},
				eventsCh: tt.fields.eventsCh,
				dialer:   tt.fields.dialer,
				logger:   &tt.fields.logger,
			}

			go func() {
				if err := ws.Subscribe(tt.args.ctx, tt.args.symbols); (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()

			resp := <-ws.Consume()

			if resp.Price != tt.expectedPrice {
				t.Error("price not equal")
			}
		})
	}
}
