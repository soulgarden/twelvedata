package twelvedata // nolint: testpackage

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

// nolint: funlen
func TestCli_GetStocks(t *testing.T) {
	t.Parallel()

	type fields struct {
		cfg     *Conf
		httpCli *fasthttp.Client
		logger  *zerolog.Logger
	}

	type args struct {
		symbol         string
		exchange       string
		country        string
		instrumentType string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantStocksResp  *response.Stocks
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				country:        "",
				instrumentType: "",
			},
			responseCode: http.StatusOK,
			// nolint: lll
			responseBody: `{
				"data":[
					{"symbol":"AAPL","name":"Apple Inc","currency":"USD","exchange":"NASDAQ","country":"United States","type":"Common Stock"},
					{"symbol":"AAPL","name":"1X AAPL","currency":"EUR","exchange":"Euronext","country":"Netherlands","type":"Common Stock"}
				],
				"status":"ok"
			}`,
			wantStocksResp: &response.Stocks{
				Data: []*response.Stock{
					{
						Symbol:   "AAPL",
						Name:     "Apple Inc",
						Currency: "USD",
						Exchange: "NASDAQ",
						Country:  "United States",
						Type:     "Common Stock",
					},
					{
						Symbol:   "AAPL",
						Name:     "1X AAPL",
						Currency: "EUR",
						Exchange: "Euronext",
						Country:  "Netherlands",
						Type:     "Common Stock",
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				country:        "",
				instrumentType: "",
			},
			responseCode: http.StatusOK,
			// nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantStocksResp:  nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "NOTFOUND",
				exchange:       "",
				country:        "",
				instrumentType: "",
			},
			responseCode: http.StatusOK,

			responseBody: `{
				"data":[],
				"status":"ok"
			}`,
			wantStocksResp: &response.Stocks{
				Data: []*response.Stock{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				country:        "",
				instrumentType: "",
			},
			responseCode: http.StatusInternalServerError,

			responseBody:    ``,
			wantStocksResp:  nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := httptest.NewServer(http.HandlerFunc(func(cw http.ResponseWriter, sr *http.Request) {
				if tt.responseCode == http.StatusInternalServerError {
					cw.WriteHeader(tt.responseCode)
				}

				cw.Header().Add("api-credits-left", strconv.Itoa(tt.wantCreditsLeft))
				cw.Header().Add("api-credits-used", strconv.Itoa(tt.wantCreditsUsed))

				_, err := cw.Write([]byte(tt.responseBody))
				if err != nil {
					t.Error(err)
				}
			}))

			defer s.Close()

			tt.fields.cfg.BaseURL = s.URL

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotStocksResp, gotCreditsLeft, gotCreditsUsed, err := c.GetStocks(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.country,
				tt.args.instrumentType,
			)
			if (err != nil && tt.wantErr == nil) || (!errors.Is(err, tt.wantErr)) {
				t.Errorf("GetStocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotStocksResp, tt.wantStocksResp) {
				t.Errorf("GetStocks() gotStocksResp = %v, want %v", gotStocksResp, tt.wantStocksResp)
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("GetStocks() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("GetStocks() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}
		})
	}
}
