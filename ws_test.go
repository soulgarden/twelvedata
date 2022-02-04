package twelvedata

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

// nolint: funlen, gocognit
func TestWS_Subscribe(t *testing.T) {
	t.Parallel()

	type fields struct {
		url      *url.URL
		eventsCh chan *response.PriceEvent
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
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				// nolint: exhaustivestruct
				url: &url.URL{
					Scheme: "ws",
					Host:   "127.0.0.1",
					Path:   "/quotes/price",
				},
				eventsCh: make(chan *response.PriceEvent),
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
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := httptest.NewServer(http.HandlerFunc(func(cw http.ResponseWriter, sr *http.Request) {
				// nolint: exhaustivestruct
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
			defer s.Close()

			// nolint: exhaustivestruct
			w := &WS{
				url: &url.URL{
					Scheme: "ws",
					Host:   strings.Replace(s.URL, "http://", "", 1),
					Path:   "/quotes/price",
				},
				eventsCh: tt.fields.eventsCh,
				dialer:   tt.fields.dialer,
				logger:   &tt.fields.logger,
			}

			go func() {
				if err := w.Subscribe(tt.args.ctx, tt.args.symbols); (err != nil) != tt.wantErr {
					t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()

			resp := <-w.Consume()

			if resp.Price != 172.8700 {
				t.Error("not equal")
			}
		})
	}
}
