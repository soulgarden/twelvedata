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

// nolint: funlen,dupl
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

// nolint: funlen,dupl
func TestCli_GetExchanges(t *testing.T) {
	t.Parallel()

	type fields struct {
		cfg     *Conf
		httpCli *fasthttp.Client
		logger  *zerolog.Logger
	}

	type args struct {
		instrumentType string
		name           string
		code           string
		country        string
	}

	tests := []struct {
		name              string
		fields            fields
		args              args
		responseCode      int
		responseBody      string
		wantExchangesResp *response.Exchanges
		wantCreditsLeft   int
		wantCreditsUsed   int
		wantErr           error
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
				instrumentType: "etf",
				name:           "",
				code:           "",
				country:        "",
			},
			responseCode: http.StatusOK,

			responseBody: `{
				"data":[
					{"name":"ASX","code":"XASX","country":"Australia","timezone":"Australia/Sydney"},
					{"name":"VSE","code":"XWBO","country":"Austria","timezone":"Europe/Vienna"}
				],
				"status":"ok"
			}`,
			wantExchangesResp: &response.Exchanges{
				Data: []*response.Exchange{
					{
						Name:     "ASX",
						Code:     "XASX",
						Country:  "Australia",
						Timezone: "Australia/Sydney",
					},
					{
						Name:     "VSE",
						Code:     "XWBO",
						Country:  "Austria",
						Timezone: "Europe/Vienna",
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
				instrumentType: "etf",
				name:           "",
				code:           "",
				country:        "",
			},
			responseCode: http.StatusOK,
			// nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantExchangesResp: nil,
			wantCreditsLeft:   10,
			wantCreditsUsed:   1,
			wantErr:           dictionary.ErrTooManyRequests,
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
				instrumentType: "",
				name:           "NOTFOUND",
				code:           "",
				country:        "",
			},
			responseCode: http.StatusOK,

			responseBody: `{
				"data":[],
				"status":"ok"
			}`,
			wantExchangesResp: &response.Exchanges{
				Data: []*response.Exchange{},
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
				instrumentType: "etf",
				name:           "",
				code:           "",
				country:        "",
			},
			responseCode: http.StatusInternalServerError,

			responseBody:      ``,
			wantExchangesResp: nil,
			wantCreditsLeft:   0,
			wantCreditsUsed:   0,
			wantErr:           dictionary.ErrBadStatusCode,
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

			gotExchangesResp, gotCreditsLeft, gotCreditsUsed, err := c.GetExchanges(
				tt.args.instrumentType,
				tt.args.name,
				tt.args.code,
				tt.args.country,
			)
			if (err != nil && tt.wantErr == nil) || (!errors.Is(err, tt.wantErr)) {
				t.Errorf("GetStocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotExchangesResp, tt.wantExchangesResp) {
				t.Errorf("GetExchanges() gotExchangesResp = %v, want %v", gotExchangesResp, tt.wantExchangesResp)
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("GetExchanges() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("GetExchanges() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}
		})
	}
}

// nolint:funlen
func TestCli_GetEtfs(t *testing.T) {
	t.Parallel()

	type fields struct {
		cfg     *Conf
		httpCli *fasthttp.Client
		logger  *zerolog.Logger
	}

	type args struct {
		symbol string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantEtfsResp    *response.Etfs
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
				symbol: "QQQ",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"data":[
					{"symbol":"QQQ","name":"Invesco QQQ Trust, Series 1","currency":"MXN","exchange":"BMV"},
					{"symbol":"QQQ","name":"Invesco QQQ Trust","currency":"USD","exchange":"NASDAQ"}
				],
				"status":"ok"
			}`,
			wantEtfsResp: &response.Etfs{
				Data: []*response.Etf{
					{
						Symbol:   "QQQ",
						Name:     "Invesco QQQ Trust, Series 1",
						Currency: "MXN",
						Exchange: "BMV",
					},
					{
						Symbol:   "QQQ",
						Name:     "Invesco QQQ Trust",
						Currency: "USD",
						Exchange: "NASDAQ",
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
				symbol: "QQQ",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantEtfsResp:    nil,
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
				symbol: "QQQ",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"data":[],
				"status":"ok"
			}`,
			wantEtfsResp: &response.Etfs{
				Data: []*response.Etf{},
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
				symbol: "QQQ",
			},
			responseCode: http.StatusInternalServerError,

			responseBody:    ``,
			wantEtfsResp:    nil,
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

			gotEtfsResp, gotCreditsLeft, gotCreditsUsed, err := c.GetEtfs(tt.args.symbol)
			if (err != nil && tt.wantErr == nil) || (!errors.Is(err, tt.wantErr)) {
				t.Errorf("GetStocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotEtfsResp, tt.wantEtfsResp) {
				t.Errorf("GetEtfs() gotEtfsResp = %v, want %v", gotEtfsResp, tt.wantEtfsResp)
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("GetEtfs() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("GetEtfs() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}
		})
	}
}

// nolint:funlen
func TestCli_GetIndices(t *testing.T) {
	t.Parallel()

	type fields struct {
		cfg     *Conf
		httpCli *fasthttp.Client
		logger  *zerolog.Logger
	}

	type args struct {
		symbol  string
		country string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantIndicesResp *response.Indices
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
				symbol:  "IXIC",
				country: "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"data":[
					{"symbol":"IXIC","name":"NASDAQ Composite","country":"United States","currency":"USD"}
				],
				"status":"ok"
			}`,
			wantIndicesResp: &response.Indices{
				Data: []*response.Index{
					{
						Symbol:   "IXIC",
						Name:     "NASDAQ Composite",
						Country:  "United States",
						Currency: "USD",
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
				symbol:  "IXIC",
				country: "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantIndicesResp: nil,
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
				symbol:  "NOTFOUND",
				country: "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"data":[],
				"status":"ok"
			}`,
			wantIndicesResp: &response.Indices{
				Data: []*response.Index{},
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
				symbol:  "IXIC",
				country: "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantIndicesResp: nil,
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

			gotIndicesResp, gotCreditsLeft, gotCreditsUsed, err := c.GetIndices(tt.args.symbol, tt.args.country)
			if (err != nil && tt.wantErr == nil) || (!errors.Is(err, tt.wantErr)) {
				t.Errorf("GetStocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotIndicesResp, tt.wantIndicesResp) {
				t.Errorf("GetIndices() gotIndicesResp = %v, want %v", gotIndicesResp, tt.wantIndicesResp)
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("GetIndices() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("GetIndices() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}
		})
	}
}
