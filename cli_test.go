package twelvedata //nolint: testpackage

import (
	"database/sql"
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
	"gopkg.in/guregu/null.v4"
)

type fields struct {
	cfg     *Conf
	httpCli *fasthttp.Client
	logger  *zerolog.Logger
}

func startServer(t *testing.T, responseCode int, wantCreditsLeft, wantCreditsUsed int64, responseBody string) string {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(cw http.ResponseWriter, sr *http.Request) {
		if responseCode == http.StatusInternalServerError {
			cw.WriteHeader(responseCode)
		}

		cw.Header().Add("api-credits-left", strconv.FormatInt(wantCreditsLeft, 10))
		cw.Header().Add("api-credits-used", strconv.FormatInt(wantCreditsUsed, 10))

		_, err := cw.Write([]byte(responseBody))
		if err != nil {
			t.Error(err)
		}
	}))

	t.Cleanup(func() {
		server.Close()
	})

	return server.URL
}

func runAssertions(
	t *testing.T,
	gotCreditsLeft, gotCreditsUsed, wantCreditsLeft, wantCreditsUsed int64,
	gotErr, wantErr error,
	gotResp, wantResp interface{},
) {
	t.Helper()

	if (gotErr != nil && wantErr == nil) || (!errors.Is(gotErr, wantErr)) {
		t.Errorf("error = %v, wantErr %v", gotErr, wantErr)

		return
	}

	if !reflect.DeepEqual(gotResp, wantResp) {
		t.Errorf("gotResp = %v, want %v", gotResp, wantResp)
	}

	if gotCreditsLeft != wantCreditsLeft {
		t.Errorf("gotCreditsLeft = %v, want %v", gotResp, wantCreditsLeft)
	}

	if gotCreditsUsed != wantCreditsUsed {
		t.Errorf("gotCreditsUsed = %v, want %v", gotResp, wantCreditsUsed)
	}
}

func TestCli_GetStocks(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol          string
		exchange        string
		micCode         string
		country         string
		instrumentType  string
		showPlan        bool
		includeDelisted bool
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Stocks
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				micCode:        "",
				country:        "",
				instrumentType: "",
				showPlan:       true,
			},
			responseCode: http.StatusOK,

			responseBody: `{
				"data":[
					{
						"symbol":"AAPL",
						"name":"Apple Inc",
						"currency":"USD",
						"exchange":"NASDAQ",
						"mic_code":"XNGS",
						"country":"United States",
						"type":"Common Stock",
						"access": {
							"global": "Basic",
							"plan": "Basic"
						}
					},
					{
						"symbol":"AAPL",
						"name":"1X AAPL",
						"currency":"EUR",
						"exchange":"Euronext",
						"mic_code":"XAMS",
						"country":"Netherlands",
						"type":"Common Stock",
						"access": {
							"global": "Basic",
							"plan": "Basic"
						}
					}
				],
				"status":"ok"
			}`,
			wantResp: &response.Stocks{
				Data: []*response.Stock{
					{
						Symbol:   "AAPL",
						Name:     "Apple Inc",
						Currency: "USD",
						Exchange: "NASDAQ",
						MicCode:  "XNGS",
						Country:  "United States",
						Type:     "Common Stock",
						Access: &response.Access{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
					{
						Symbol:   "AAPL",
						Name:     "1X AAPL",
						Currency: "EUR",
						Exchange: "Euronext",
						MicCode:  "XAMS",
						Country:  "Netherlands",
						Type:     "Common Stock",
						Access: &response.Access{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				micCode:        "",
				country:        "",
				instrumentType: "",
				showPlan:       true,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "NOTFOUND",
				exchange:       "",
				micCode:        "",
				country:        "",
				instrumentType: "",
				showPlan:       true,
			},
			responseCode: http.StatusOK,

			responseBody: `{
				"data":[],
				"status":"ok"
			}`,
			wantResp: &response.Stocks{
				Data: []*response.Stock{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				exchange:       "",
				micCode:        "",
				country:        "",
				instrumentType: "",
				showPlan:       true,
			},
			responseCode: http.StatusInternalServerError,

			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetStocks(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.instrumentType,
				tt.args.showPlan,
				tt.args.includeDelisted,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetExchanges(t *testing.T) {
	t.Parallel()

	type args struct {
		instrumentType string
		name           string
		code           string
		country        string
		showPlan       bool
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Exchanges
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Exchanges{
				Data: []response.Exchange{
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

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Exchanges{
				Data: []response.Exchange{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				instrumentType: "etf",
				name:           "",
				code:           "",
				country:        "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetExchanges(
				tt.args.instrumentType,
				tt.args.name,
				tt.args.code,
				tt.args.country,
				tt.args.showPlan,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetEtfs(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol          string
		exchange        string
		micCode         string
		country         string
		showPlan        bool
		includeDelisted bool
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Etfs
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "QQQ",
				showPlan: true,
			},
			responseCode: http.StatusOK,

			responseBody: `
			{
				"data":[
					{
						"symbol":"QQQ",
						"name":"Invesco QQQ Trust, Series 1",
						"currency":"MXN",
						"exchange":"BMV",
						"mic_code":"XMEX",
						"country":"Mexico",
						"access": {
							"global": "Basic",
							"plan": "Basic"
						}
					},
					{
						"symbol":"QQQ",
						"name":"Invesco QQQ Trust",
						"currency":"USD",
						"exchange":"NASDAQ",
						"mic_code":"XNMS",
						"country":"United States",
						"access": {
							"global": "Basic",
							"plan": "Basic"
						}
					}
				],
				"status":"ok"
			}`,
			wantResp: &response.Etfs{
				Data: []response.Etf{
					{
						Symbol:   "QQQ",
						Name:     "Invesco QQQ Trust, Series 1",
						Currency: "MXN",
						Exchange: "BMV",
						MicCode:  "XMEX",
						Country:  "Mexico",
						Access: &response.Access{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
					{
						Symbol:   "QQQ",
						Name:     "Invesco QQQ Trust",
						Currency: "USD",
						Exchange: "NASDAQ",
						MicCode:  "XNMS",
						Country:  "United States",
						Access: &response.Access{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "QQQ",
				showPlan: true,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "QQQ",
				showPlan: true,
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"data":[],
				"status":"ok"
			}`,
			wantResp: &response.Etfs{
				Data: []response.Etf{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol: "QQQ",
			},
			responseCode: http.StatusInternalServerError,

			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetEtfs(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.showPlan,
				tt.args.includeDelisted,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetIndices(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol          string
		country         string
		showPlan        bool
		includeDelisted bool
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Indices
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Indices{
				Data: []response.Index{
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

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Indices{
				Data: []response.Index{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:  "IXIC",
				country: "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetIndices(
				tt.args.symbol,
				tt.args.country,
				tt.args.showPlan,
				tt.args.includeDelisted,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetTimeSeries(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol         string
		interval       string
		exchange       string
		micCode        string
		country        string
		instrumentType string
		outputSize     int
		prePost        string
		db             int
		order          string
		timezone       string
		date           string
		startDate      string
		endDate        string
		previousClose  bool
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.TimeSeries
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				interval:       "1min",
				exchange:       "",
				country:        "",
				instrumentType: "",
				outputSize:     3,
				prePost:        "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `
			{
				"meta":{
					"symbol":"AAPL",
					"interval":"1min",
					"currency":"USD",
					"exchange_timezone":"America/New_York",
					"exchange":"NASDAQ",
					"mic_code": "XNAS",
					"type":"Common Stock"
				},
				"values":[
					{"datetime":"2022-02-07 15:59:00","open":"171.42000","high":"171.75999","low":"171.41000","close":"171.71001","volume":"863231"},
					{"datetime":"2022-02-07 15:58:00","open":"171.27499","high":"171.45000","low":"171.27000","close":"171.41000","volume":"374529"},
					{"datetime":"2022-02-07 15:57:00","open":"171.12000","high":"171.36000","low":"171.12000","close":"171.27000","volume":"337196"}
				],
				"status":"ok"
			}`,
			wantResp: &response.TimeSeries{
				Meta: response.TimeSeriesMeta{
					Symbol:           "AAPL",
					Interval:         "1min",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					Type:             "Common Stock",
				},
				Values: []response.TimeSeriesValue{
					{
						Datetime: "2022-02-07 15:59:00",
						Open:     "171.42000",
						High:     "171.75999",
						Low:      "171.41000",
						Close:    "171.71001",
						Volume:   "863231",
					},
					{
						Datetime: "2022-02-07 15:58:00",
						Open:     "171.27499",
						High:     "171.45000",
						Low:      "171.27000",
						Close:    "171.41000",
						Volume:   "374529",
					},
					{
						Datetime: "2022-02-07 15:57:00",
						Open:     "171.12000",
						High:     "171.36000",
						Low:      "171.12000",
						Close:    "171.27000",
						Volume:   "337196",
					},
				},
				Status: "ok",
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				interval:       "1min",
				exchange:       "",
				country:        "",
				instrumentType: "",
				outputSize:     1000,
				prePost:        "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "symbol is not available with your plan",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				interval:       "1min",
				exchange:       "",
				country:        "",
				instrumentType: "",
				outputSize:     1000,
				prePost:        "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code": 400,
				"message": "**symbol** ALD is not available with your plan. You may select the appropriate plan at https://twelvedata.com/pricing","status":"error","meta":{"symbol":"ALD:TASE","interval":"1day","exchange":""}}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrIsNotAvailableWithYourPlan,
		},
		{
			name: "not found symbol",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "NOTFOUND",
				interval:       "1min",
				exchange:       "",
				country:        "",
				instrumentType: "",
				outputSize:     3,
				prePost:        "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"code":400,
				"message":"**symbol** not found: NOTFOUND. Please specify it correctly according to API Documentation.",
				"status":"error",
				"meta":{
					"symbol":"NOTFOUND",
					"interval":"1min",
					"exchange":""
				}
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:         "AAPL",
				interval:       "1min",
				exchange:       "",
				country:        "",
				instrumentType: "",
				outputSize:     3,
				prePost:        "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetTimeSeries(
				tt.args.symbol,
				tt.args.interval,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.instrumentType,
				tt.args.outputSize,
				tt.args.prePost,
				tt.args.db,
				tt.args.order,
				tt.args.timezone,
				tt.args.date,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.previousClose,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetExchangeRate(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol   string
		date     string
		timeZone string
		dp       int
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.ExchangeRate
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "USD/JPY",
				timeZone: "",
				dp:       2,
			},
			responseCode: http.StatusOK,

			responseBody: `{"symbol":"USD/JPY","rate":115.58,"timestamp":1644344940}`,
			wantResp: &response.ExchangeRate{
				Symbol:    "USD/JPY",
				Rate:      115.58,
				Timestamp: 1644344940,
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "USD/JPY",
				timeZone: "",
				dp:       2,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol 1",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "NOT/FOUND",
				timeZone: "",
				dp:       2,
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"code":400,
				"message":"**symbol** not found: NOT/FOUND. Please specify it correctly according to API Documentation.",
				"status":"error",
				"meta":{}
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "not found symbol 2",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "ZAC/USD",
				timeZone: "",
				dp:       2,
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"code":400,
				"message":"**symbol** not found: NOT/FOUND. Please specify it correctly according to API Documentation.",
				"status":"error",
				"meta":{}
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "USD/JPY",
				timeZone: "",
				dp:       2,
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetExchangeRate(
				tt.args.symbol,
				tt.args.date,
				tt.args.timeZone,
				tt.args.dp,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetQuote(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol           string
		interval         string
		exchange         string
		micCode          string
		country          string
		volumeTimePeriod string
		instrumentType   string
		prepost          string
		eod              bool
		rollingPeriod    int
		dp               int
		timezone         string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Quotes
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:           "AAPL",
				interval:         "1day",
				exchange:         "",
				micCode:          "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prepost:          "",
				eod:              false,
				rollingPeriod:    24,
				dp:               5,
				timezone:         "",
			},
			responseCode: http.StatusOK,

			responseBody: `
			{
				"symbol":"AAPL",
				"name":"Apple Inc",
				"exchange":"NASDAQ",
				"mic_code": "XNAS",
				"currency":"USD",
				"datetime":"2022-02-08",
				"open":"171.73000",
				"high":"175.35001",
				"low":"171.42999",
				"close":"174.83000",
				"volume":"74706900",
				"previous_close":"171.66000",
				"change":"3.17000",
				"percent_change":"1.84667",
				"average_volume":"102060300",
				"fifty_two_week":{
					"low":"116.21000",
					"high":"182.94000",
					"low_change":"58.62000",
					"high_change":"-8.11000",
					"low_change_percent":"50.44317",
					"high_change_percent":"-4.43315",
					"range":"116.209999 - 182.940002"
				}
			}`,
			wantResp: &response.Quotes{
				Data: []response.Quote{{
					Symbol:        "AAPL",
					Name:          "Apple Inc",
					Exchange:      "NASDAQ",
					MicCode:       "XNAS",
					Currency:      "USD",
					Datetime:      "2022-02-08",
					Open:          "171.73000",
					High:          "175.35001",
					Low:           "171.42999",
					Close:         "174.83000",
					Volume:        "74706900",
					PreviousClose: "171.66000",
					Change:        "3.17000",
					PercentChange: "1.84667",
					AverageVolume: "102060300",
					FiftyTwoWeek: &response.FiftyTwoWeek{
						Low:               "116.21000",
						High:              "182.94000",
						LowChange:         "58.62000",
						HighChange:        "-8.11000",
						LowChangePercent:  "50.44317",
						HighChangePercent: "-4.43315",
						Range:             "116.209999 - 182.940002",
					},
				}},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:           "AAPL",
				interval:         "1day",
				exchange:         "",
				micCode:          "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prepost:          "",
				eod:              false,
				rollingPeriod:    24,
				dp:               5,
				timezone:         "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:           "NOTFOUND1,NOTFOUND2",
				interval:         "1day",
				exchange:         "",
				micCode:          "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				eod:              false,
				rollingPeriod:    24,
				dp:               5,
				timezone:         "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"NOTFOUND1": {
					"code": 400,
					"message": "**symbol** not found: NOTFOUND1. Please specify it correctly according to API Documentation.",
					"status": "error",
					"meta": {
						"symbol": "NOTFOUND1,NOTFOUND2",
						"interval": "1day",
						"exchange": ""
					}
				},
				"NOTFOUND2": {
					"code": 400,
					"message": "**symbol** not found: NOTFOUND2. Please specify it correctly according to API Documentation.",
					"status": "error",
					"meta": {
						"symbol": "NOTFOUND1,NOTFOUND2",
						"interval": "1day",
						"exchange": ""
					}
				}
			}`,
			wantResp: &response.Quotes{
				Errors: []response.QuoteError{
					{
						Code:    400,
						Message: "**symbol** not found: NOTFOUND1. Please specify it correctly according to API Documentation.",
						Status:  "error",
						Meta: &response.QuoteErrorMeta{
							Symbol:   "NOTFOUND1,NOTFOUND2",
							Interval: "1day",
							Exchange: "",
						},
					},
					{
						Code:    400,
						Message: "**symbol** not found: NOTFOUND2. Please specify it correctly according to API Documentation.",
						Status:  "error",
						Meta: &response.QuoteErrorMeta{
							Symbol:   "NOTFOUND1,NOTFOUND2",
							Interval: "1day",
							Exchange: "",
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:           "AAPL",
				interval:         "1day",
				exchange:         "",
				micCode:          "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prepost:          "",
				eod:              false,
				rollingPeriod:    24,
				dp:               5,
				timezone:         "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetQuote(
				tt.args.symbol,
				tt.args.interval,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.volumeTimePeriod,
				tt.args.instrumentType,
				tt.args.prepost,
				tt.args.eod,
				tt.args.rollingPeriod,
				tt.args.dp,
				tt.args.timezone,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetProfile(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol   string
		exchange string
		micCode  string
		country  string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Profile
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `
			{
				"symbol":"AAPL",
				"name":"Apple Inc",
				"exchange":"NASDAQ",
				"mic_code": "XNAS",
				"sector":"Technology",
				"industry":"Consumer Electronics",
				"employees":154000,
				"website":"https://www.apple.com",
				"description":"Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. It also sells various related services. In addition, the company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; AirPods Max, an over-ear wireless headphone; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, HomePod, and iPod touch. Further, it provides AppleCare support services; cloud services store services; and operates various platforms, including the App Store that allow customers to discover and download applications and digital content, such as books, music, video, games, and podcasts. Additionally, the company offers various services, such as Apple Arcade, a game subscription service; Apple Music, which offers users a curated listening experience with on-demand radio stations; Apple News+, a subscription news and magazine service; Apple TV+, which offers exclusive original content; Apple Card, a co-branded credit card; and Apple Pay, a cashless payment service, as well as licenses its intellectual property. The company serves consumers, and small and mid-sized businesses; and the education, enterprise, and government markets. It distributes third-party applications for its products through the App Store. The company also sells its products through its retail and online stores, and direct sales force; and third-party cellular network carriers, wholesalers, retailers, and resellers. Apple Inc. was incorporated in 1977 and is headquartered in Cupertino, California.",
				"type":"Common Stock",
				"CEO":"Mr. Timothy D. Cook",
				"address":"One Apple Park Way",
				"city":"Cupertino",
				"zip":"95014",
				"state":"CA",
				"country":"US",
				"phone":"408 996 1010"
			}`,
			wantResp: &response.Profile{
				Symbol:   "AAPL",
				Name:     "Apple Inc",
				Exchange: "NASDAQ",
				MicCode:  "XNAS",
				Sector:   "Technology",
				Industry: "Consumer Electronics",
				Employees: null.Int{
					NullInt64: sql.NullInt64{
						Int64: 154000,
						Valid: true,
					},
				},
				Website: "https://www.apple.com",
				//nolint: lll
				Description: "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. It also sells various related services. In addition, the company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; AirPods Max, an over-ear wireless headphone; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, HomePod, and iPod touch. Further, it provides AppleCare support services; cloud services store services; and operates various platforms, including the App Store that allow customers to discover and download applications and digital content, such as books, music, video, games, and podcasts. Additionally, the company offers various services, such as Apple Arcade, a game subscription service; Apple Music, which offers users a curated listening experience with on-demand radio stations; Apple News+, a subscription news and magazine service; Apple TV+, which offers exclusive original content; Apple Card, a co-branded credit card; and Apple Pay, a cashless payment service, as well as licenses its intellectual property. The company serves consumers, and small and mid-sized businesses; and the education, enterprise, and government markets. It distributes third-party applications for its products through the App Store. The company also sells its products through its retail and online stores, and direct sales force; and third-party cellular network carriers, wholesalers, retailers, and resellers. Apple Inc. was incorporated in 1977 and is headquartered in Cupertino, California.",
				Type:        "Common Stock",
				CEO:         "Mr. Timothy D. Cook",
				Address:     "One Apple Park Way",
				City:        "Cupertino",
				Zip:         "95014",
				State:       "CA",
				Country:     "US",
				Phone:       "408 996 1010",
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 10,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 10,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "NOTFOUND",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 10,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetProfile(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetDividends(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		exchange  string
		micCode   string
		country   string
		r         string
		startDate string
		endDate   string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Dividends
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				r:         "last",
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,

			responseBody: `
			{
				"meta":{
					"symbol":"AAPL",
					"name":"Apple Inc",
					"currency":"USD",
					"exchange":"NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone":"America/New_York"
				},
				"dividends":[
					{"payment_date":"2022-02-04","amount":0.22}
				]
			}`,
			wantResp: &response.Dividends{
				Meta: response.DividendsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
				},
				Dividends: []response.Dividend{
					{
						PaymentDate: "2022-02-04",
						Amount:      0.22,
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				r:         "last",
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOTFOUND",
				exchange:  "",
				micCode:   "",
				country:   "",
				r:         "last",
				startDate: "",
				endDate:   "",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				r:         "last",
				startDate: "",
				endDate:   "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetDividends(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.r,
				tt.args.startDate,
				tt.args.endDate,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetEarning(t *testing.T) {
	t.Parallel()

	type args struct {
		time string
		date string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Earnings
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				date: "",
				time: "",
			},
			responseCode: http.StatusOK,
			responseBody: `
				{
					 "earnings":{
						:2023-01-02:[
								{
									"date": "2020-04-30",
									"time": "After Hours",
									"eps_estimate": 2.09,
									"eps_actual": 2.55,
									"difference": 0.46,
									"surprise_prc": 22.01
								}
							]
						},
						"status":"ok"
					}`,
			wantResp: &response.Earnings{
				Earnings: map[string][]*response.Earning{
					"2020-05-08": {
						{
							Symbol:   "BR",
							Name:     "Broadridge Financial Solutions Inc",
							Currency: "USD",
							Exchange: "NYSE",
							MicCode:  "XNYS",
							Country:  "United States",
							Time:     "Time Not Supplied",
							EpsEstimate: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: 1.72,
									Valid:   true,
								},
							},
							EpsActual: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: 1.67,
									Valid:   true,
								},
							},
							Difference: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: -0.05,
									Valid:   true,
								},
							},
							SurprisePrc: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: -2.9,
									Valid:   true,
								},
							},
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 40,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				date: "",
				time: "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 40,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				date: "",
				time: "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetEarning(
				tt.args.date,
				tt.args.time,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetEarningsCalendar(t *testing.T) {
	t.Parallel()

	type args struct {
		exchange  string
		micCode   string
		country   string
		dp        int
		startDate string
		endDate   string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Earnings
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange:  "",
				micCode:   "",
				country:   "",
				dp:        2,
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			responseBody: `
				{
					"earnings":{
						"2020-05-08":[
							{
								"symbol": "BR",
								"name": "Broadridge Financial Solutions Inc",
								"currency": "USD",
								"exchange": "NYSE",
								"mic_code": "XNYS",
								"country": "United States",
								"time": "Time Not Supplied",
								"eps_estimate": 1.72,
								"eps_actual": 1.67,
								"difference": -0.05,
								"surprise_prc": -2.9
							}
						]
					},
					"status":"ok"
				}`,
			wantResp: &response.Earnings{
				Earnings: map[string][]*response.Earning{
					"2020-05-08": {
						{
							Symbol:   "BR",
							Name:     "Broadridge Financial Solutions Inc",
							Currency: "USD",
							Exchange: "NYSE",
							MicCode:  "XNYS",
							Country:  "United States",
							Time:     "Time Not Supplied",
							EpsEstimate: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: 1.72,
									Valid:   true,
								},
							},
							EpsActual: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: 1.67,
									Valid:   true,
								},
							},
							Difference: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: -0.05,
									Valid:   true,
								},
							},
							SurprisePrc: null.Float{
								NullFloat64: sql.NullFloat64{
									Float64: -2.9,
									Valid:   true,
								},
							},
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 40,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange:  "",
				micCode:   "",
				country:   "",
				dp:        2,
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 40,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange:  "",
				micCode:   "",
				country:   "",
				dp:        2,
				startDate: "",
				endDate:   "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetEarningsCalendar(
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.dp,
				tt.args.startDate,
				tt.args.endDate,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetStatistics(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol   string
		exchange string
		micCode  string
		country  string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Statistics
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"meta": {
					"symbol": "AAPL",
					"name": "Apple Inc",
					"currency": "USD",
					"exchange": "NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone": "America/New_York"
				},
				"statistics": {
					"valuations_metrics": {
						"market_capitalization": 2880798195712,
						"enterprise_value": 3022112423936,
						"trailing_pe": 31.299448,
						"forward_pe": 28.412607,
						"peg_ratio": 2,
						"price_to_sales_ttm": 7.874971,
						"price_to_book_mrq": 45.71463,
						"enterprise_to_revenue": 8.261,
						"enterprise_to_ebitda": 25.135
					},
					"financials": {
						"fiscal_year_ends": "2021-09-25",
						"most_recent_quarter": "2021-09-25",
						"profit_margin": 0.25882,
						"operating_margin": 0.29782,
						"return_on_assets_ttm": 0.20179,
						"return_on_equity_ttm": 1.47443,
						"income_statement": {
							"revenue_ttm": 365817004032,
							"revenue_per_share_ttm": 21.904,
							"quarterly_revenue_growth": 0.288,
							"gross_profit_ttm": 152836000000,
							"ebitda": 120233000960,
							"net_income_to_common_ttm": 94679998464,
							"diluted_eps_ttm": 5.61,
							"quarterly_earnings_growth_yoy": 0.622
						},
						"balance_sheet": {
							"revenue_ttm": 365817004032,
							"total_cash_mrq": 62639001600,
							"total_cash_per_share_mrq": 3.818,
							"total_debt_mrq": 136521998336,
							"total_debt_to_equity_mrq": 216.392,
							"current_ratio_mrq": 1.075,
							"book_value_per_share_mrq": 3.841
						},
						"cash_flow": {
							"operating_cash_flow_ttm": 104037998592,
							"levered_free_cash_flow_ttm": 73295003648
						}
					},
					"stock_statistics": {
						"shares_outstanding": 16406400000,
						"float_shares": 16389662475,
						"avg_10_volume": 94468150,
						"avg_30_volume": 94056423,
						"shares_short": 113277024,
						"short_ratio": 1,
						"short_percent_of_shares_outstanding": 0.0069,
						"percent_held_by_insiders": 0.0007,
						"percent_held_by_institutions": 0.58707
					},
					"stock_price_summary": {
						"fifty_two_week_low": 116.21,
						"fifty_two_week_high": 182.94,
						"fifty_two_week_change": null,
						"beta": 1.203116,
						"day_50_ma": 171.6632,
						"day_200_ma": 149.5189
					},
					"dividends_and_splits": {
						"forward_annual_dividend_rate": 0.88,
						"forward_annual_dividend_yield": 0.0049,
						"trailing_annual_dividend_rate": 0.85,
						"trailing_annual_dividend_yield": 0.004861866,
						"5_year_average_dividend_yield": 1.17,
						"payout_ratio": 0.1515,
						"dividend_date": "2022-02-10",
						"ex_dividend_date": "2021-11-05",
						"last_split_factor": "4-for-1 split",
						"last_split_date": "2020-08-31"
					}
				}
			}`,
			wantResp: &response.Statistics{
				Meta: response.StatisticsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
				},
				Statistics: response.StatisticsValues{
					ValuationsMetrics: response.StatisticsValuationsMetrics{
						MarketCapitalization: null.IntFrom(2880798195712),
						EnterpriseValue:      null.IntFrom(3022112423936),
						TrailingPe:           null.FloatFrom(31.299448),
						ForwardPe:            null.FloatFrom(28.412607),
						PegRatio:             null.FloatFrom(2),
						PriceToSalesTtm:      null.FloatFrom(7.874971),
						PriceToBookMrq:       null.FloatFrom(45.71463),
						EnterpriseToRevenue:  null.FloatFrom(8.261),
						EnterpriseToEbitda:   null.FloatFrom(25.135),
					},
					Financials: response.StatisticsFinancials{
						FiscalYearEnds:    "2021-09-25",
						MostRecentQuarter: "2021-09-25",
						ProfitMargin:      null.FloatFrom(0.25882),
						OperatingMargin:   null.FloatFrom(0.29782),
						ReturnOnAssetsTtm: null.FloatFrom(0.20179),
						ReturnOnEquityTtm: null.FloatFrom(1.47443),
						IncomeStatement: response.StatisticsIncomeStatement{
							RevenueTtm:                 null.IntFrom(365817004032),
							RevenuePerShareTtm:         null.FloatFrom(21.904),
							QuarterlyRevenueGrowth:     null.FloatFrom(0.288),
							GrossProfitTtm:             null.IntFrom(152836000000),
							Ebitda:                     null.IntFrom(120233000960),
							NetIncomeToCommonTtm:       null.IntFrom(94679998464),
							DilutedEpsTtm:              null.FloatFrom(5.61),
							QuarterlyEarningsGrowthYoy: null.FloatFrom(0.622),
						},
						BalanceSheet: response.StatisticsBalanceSheet{
							RevenueTtm:           null.IntFrom(365817004032),
							TotalCashMrq:         null.IntFrom(62639001600),
							TotalCashPerShareMrq: null.FloatFrom(3.818),
							TotalDebtMrq:         null.IntFrom(136521998336),
							TotalDebtToEquityMrq: null.FloatFrom(216.392),
							CurrentRatioMrq:      null.FloatFrom(1.075),
							BookValuePerShareMrq: null.FloatFrom(3.841),
						},
						CashFlow: response.StatisticsCashFlow{
							OperatingCashFlowTtm:   null.IntFrom(104037998592),
							LeveredFreeCashFlowTtm: null.IntFrom(73295003648),
						},
					},
					StockStatistics: response.StockStatistics{
						SharesOutstanding:               null.IntFrom(16406400000),
						FloatShares:                     null.IntFrom(16389662475),
						Avg10Volume:                     null.IntFrom(94468150),
						Avg30Volume:                     null.IntFrom(94056423),
						SharesShort:                     null.IntFrom(113277024),
						ShortRatio:                      null.FloatFrom(1),
						ShortPercentOfSharesOutstanding: null.FloatFrom(0.0069),
						PercentHeldByInsiders:           null.FloatFrom(0.0007),
						PercentHeldByInstitutions:       null.FloatFrom(0.58707),
					},
					StockPriceSummary: response.StockPriceSummary{
						FiftyTwoWeekLow:    null.FloatFrom(116.21),
						FiftyTwoWeekHigh:   null.FloatFrom(182.94),
						FiftyTwoWeekChange: null.Float{NullFloat64: sql.NullFloat64{}},
						Beta:               null.FloatFrom(1.203116),
						Day50Ma:            null.FloatFrom(171.6632),
						Day200Ma:           null.FloatFrom(149.5189),
					},
					DividendsAndSplits: response.DividendsAndSplits{
						ForwardAnnualDividendRate:   null.FloatFrom(0.88),
						ForwardAnnualDividendYield:  null.FloatFrom(0.0049),
						TrailingAnnualDividendRate:  null.FloatFrom(0.85),
						TrailingAnnualDividendYield: null.FloatFrom(0.004861866),
						YearAverageDividendYield:    null.FloatFrom(1.17),
						PayoutRatio:                 null.FloatFrom(0.1515),
						DividendDate:                "2022-02-10",
						ExDividendDate:              "2021-11-05",
						LastSplitFactor:             "4-for-1 split",
						LastSplitDate:               "2020-08-31",
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "NOTFOUND",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "forbidden",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":403,
				"message":"The 'demo' API key is only used for initial familiarity. To become a full user, you can request your own API key at https://twelvedata.com/pricing. It is absolutely free, and it’s yours for a lifetime. It only takes 10 seconds to obtain your own API key!",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrForbidden,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetStatistics(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetBalanceSheet(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		exchange  string
		country   string
		micCode   string
		startDate string
		endDate   string
		period    string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.BalanceSheets
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				period:    "annual",
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"meta": {
					"symbol": "AAPL",
					"name": "Apple Inc",
					"currency": "USD",
					"exchange": "NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone": "America/New_York",
					"period": "Annual"
				},
				"balance_sheet": [
					{
						"fiscal_date": "2021-09-30",
						"assets": {
							"current_assets": {
								"cash": 17305000000,
								"cash_equivalents": 17635000000,
								"cash_and_cash_equivalents": 34940000000,
								"other_short_term_investments": 27699000000,
								"accounts_receivable": 26278000000,
								"other_receivables": 25228000000,
								"inventory": 6580000000,
								"prepaid_assets": null,
								"other_current_assets": 14111000000,
								"total_current_assets": 134836000000
							},
							"non_current_assets": {
								"properties": 0,
								"land_and_improvements": 20041000000,
								"machinery_furniture_equipment": 78659000000,
								"leases": 11023000000,
								"accumulated_depreciation": -70283000000,
								"goodwill": null,
								"intangible_assets": null,
								"investments_and_advances": 127877000000,
								"other_non_current_assets": 48849000000,
								"total_non_current_assets": 216166000000
							},
							"total_assets": 351002000000
						},
						"liabilities": {
							"current_liabilities": {
								"accounts_payable": 54763000000,
								"accrued_expenses": null,
								"short_term_debt": 15613000000,
								"deferred_revenue": 7612000000,
								"other_current_liabilities": 47493000000,
								"total_current_liabilities": 125481000000,
								"tax_payable": null
							},
							"non_current_liabilities": {
								"long_term_debt": 109106000000,
								"provision_for_risks_and_charges": 24689000000,
								"deferred_liabilities": null,
								"other_non_current_liabilities": 28636000000,
								"total_non_current_liabilities": 162431000000,
								"long_term_provisions": null
							},
							"total_liabilities": 287912000000
						},
						"shareholders_equity": {
							"common_stock": 57365000000,
							"retained_earnings": 5562000000,
							"other_shareholders_equity": 163000000,
							"total_shareholders_equity": 63090000000,
							"additional_paid_in_capital": null,
							"treasury_stock": null,
							"minority_interest": null
						}
					}
				]
			}`,
			wantResp: &response.BalanceSheets{
				Meta: response.BalanceSheetsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Annual",
				},
				BalanceSheet: []response.BalanceSheet{
					{
						FiscalDate: "2021-09-30",
						Assets: response.BalanceSheetAssets{
							CurrentAssets: response.BalanceSheetCurrentAssets{
								Cash:                      null.IntFrom(17305000000),
								CashEquivalents:           null.IntFrom(17635000000),
								CashAndCashEquivalents:    null.IntFrom(34940000000),
								OtherShortTermInvestments: null.IntFrom(27699000000),
								AccountsReceivable:        null.IntFrom(26278000000),
								OtherReceivables:          null.IntFrom(25228000000),
								Inventory:                 null.IntFrom(6580000000),
								PrepaidAssets:             null.Int{NullInt64: sql.NullInt64{}},
								OtherCurrentAssets:        null.IntFrom(14111000000),
								TotalCurrentAssets:        null.IntFrom(134836000000),
							},
							NonCurrentAssets: response.BalanceSheetNonCurrentAssets{
								Properties:                  null.IntFrom(0),
								LandAndImprovements:         null.IntFrom(20041000000),
								MachineryFurnitureEquipment: null.IntFrom(78659000000),
								Leases:                      null.IntFrom(11023000000),
								AccumulatedDepreciation:     null.IntFrom(-70283000000),
								Goodwill:                    null.Int{NullInt64: sql.NullInt64{}},
								IntangibleAssets:            null.Int{NullInt64: sql.NullInt64{}},
								InvestmentsAndAdvances:      null.IntFrom(127877000000),
								OtherNonCurrentAssets:       null.IntFrom(48849000000),
								TotalNonCurrentAssets:       null.IntFrom(216166000000),
							},
							TotalAssets: null.IntFrom(351002000000),
						},
						Liabilities: response.BalanceSheetLiabilities{
							CurrentLiabilities: response.BalanceSheetCurrentLiabilities{
								AccountsPayable:         null.IntFrom(54763000000),
								AccruedExpenses:         null.Int{NullInt64: sql.NullInt64{}},
								ShortTermDebt:           null.IntFrom(15613000000),
								DeferredRevenue:         null.IntFrom(7612000000),
								OtherCurrentLiabilities: null.IntFrom(47493000000),
								TotalCurrentLiabilities: null.IntFrom(125481000000),
								TaxPayable:              null.Int{NullInt64: sql.NullInt64{}},
							},
							NonCurrentLiabilities: response.BalanceSheetNonCurrentLiabilities{
								LongTermDebt:                null.IntFrom(109106000000),
								ProvisionForRisksAndCharges: null.IntFrom(24689000000),
								DeferredLiabilities:         null.Int{NullInt64: sql.NullInt64{}},
								OtherNonCurrentLiabilities:  null.IntFrom(28636000000),
								TotalNonCurrentLiabilities:  null.IntFrom(162431000000),
								LongTermProvisions:          null.Int{NullInt64: sql.NullInt64{}},
							},
							TotalLiabilities: null.IntFrom(287912000000),
						},
						ShareholdersEquity: response.BalanceSheetShareholdersEquity{
							CommonStock:             null.IntFrom(57365000000),
							RetainedEarnings:        null.IntFrom(5562000000),
							OtherShareholdersEquity: null.IntFrom(163000000),
							TotalShareholdersEquity: null.IntFrom(63090000000),
							AdditionalPaidInCapital: null.Int{NullInt64: sql.NullInt64{}},
							TreasuryStock:           null.Int{NullInt64: sql.NullInt64{}},
							MinorityInterest:        null.Int{NullInt64: sql.NullInt64{}},
						},
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOTFOUND",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetBalanceSheet(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.country,
				tt.args.micCode,
				tt.args.period,
				tt.args.startDate,
				tt.args.endDate,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetCashFlow(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		exchange  string
		country   string
		micCode   string
		startDate string
		endDate   string
		period    string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.CashFlows
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 0, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"meta": {
					"symbol": "AAPL",
					"name": "Apple Inc",
					"currency": "USD",
					"exchange": "NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone": "America/New_York",
					"period": "Annual"
				},
				"cash_flow": [
					{
						"fiscal_date": "2021-09-30",
						"operating_activities": {
							"net_income": 94680000000,
							"depreciation": 11284000000,
							"deferred_taxes": -4774000000,
							"stock_based_compensation": 7906000000,
							"other_non_cash_items": -147000000,
							"accounts_receivable": -14028000000,
							"accounts_payable": 12326000000,
							"other_assets_liabilities": -3209000000,
							"operating_cash_flow": 104038000000
						},
						"investing_activities": {
							"capital_expenditures": -11085000000,
							"net_intangibles": null,
							"net_acquisitions": -33000000,
							"purchase_of_investments": -109689000000,
							"sale_of_investments": 106870000000,
							"other_investing_activity": -608000000,
							"investing_cash_flow": -14545000000
						},
						"financing_activities": {
							"long_term_debt_issuance": 20393000000,
							"long_term_debt_payments": -8750000000,
							"short_term_debt_issuance": 1022000000,
							"common_stock_issuance": 1105000000,
							"common_stock_repurchase": -85971000000,
							"common_dividends": -14467000000,
							"other_financing_charges": -6685000000,
							"financing_cash_flow": -93353000000
						},
						"end_cash_position": 35929000000,
						"income_tax_paid": 25385000000,
						"interest_paid": 2687000000,
						"free_cash_flow": 115123000000
					}
				]
			}`,
			wantResp: &response.CashFlows{
				Meta: response.CashFlowsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Annual",
				},
				CashFlow: []response.CashFlow{
					{
						FiscalDate: "2021-09-30",
						OperatingActivities: response.CashFlowOperatingActivities{
							NetIncome:              null.IntFrom(94680000000),
							Depreciation:           null.IntFrom(11284000000),
							DeferredTaxes:          null.IntFrom(-4774000000),
							StockBasedCompensation: null.IntFrom(7906000000),
							OtherNonCashItems:      null.IntFrom(-147000000),
							AccountsReceivable:     null.IntFrom(-14028000000),
							AccountsPayable:        null.IntFrom(12326000000),
							OtherAssetsLiabilities: null.IntFrom(-3209000000),
							OperatingCashFlow:      null.IntFrom(104038000000),
						},
						InvestingActivities: response.CashFlowInvestingActivities{
							CapitalExpenditures:    null.IntFrom(-11085000000),
							NetIntangibles:         null.Int{NullInt64: sql.NullInt64{}},
							NetAcquisitions:        null.IntFrom(-33000000),
							PurchaseOfInvestments:  null.IntFrom(-109689000000),
							SaleOfInvestments:      null.IntFrom(106870000000),
							OtherInvestingActivity: null.IntFrom(-608000000),
							InvestingCashFlow:      null.IntFrom(-14545000000),
						},
						FinancingActivities: response.CashFlowFinancingActivities{
							LongTermDebtIssuance:  null.IntFrom(20393000000),
							LongTermDebtPayments:  null.IntFrom(-8750000000),
							ShortTermDebtIssuance: null.IntFrom(1022000000),
							CommonStockIssuance:   null.IntFrom(1105000000),
							CommonStockRepurchase: null.IntFrom(-85971000000),
							CommonDividends:       null.IntFrom(-14467000000),
							OtherFinancingCharges: null.IntFrom(-6685000000),
							FinancingCashFlow:     null.IntFrom(-93353000000),
						},
						EndCashPosition: null.IntFrom(35929000000),
						IncomeTaxPaid:   null.IntFrom(25385000000),
						InterestPaid:    null.IntFrom(2687000000),
						FreeCashFlow:    null.IntFrom(115123000000),
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOTFOUND",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				micCode:   "",
				startDate: "",
				endDate:   "",
				period:    "annual",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetCashFlow(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.country,
				tt.args.micCode,
				tt.args.startDate,
				tt.args.endDate,
				tt.args.period,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetIncomeStatement(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		exchange  string
		micCode   string
		country   string
		period    string
		startDate string
		endDate   string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.IncomeStatements
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				period:    "annual",
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"meta": {
					"symbol": "AAPL",
					"name": "Apple Inc",
					"currency": "USD",
					"exchange": "NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone": "America/New_York",
					"period": "Annual"
				},
				"income_statement": [
					{
						"fiscal_date": "2021-09-30",
						"sales": 365817000000,
						"cost_of_goods": 212981000000,
						"gross_profit": 152836000000,
						"operating_expense": {
							"research_and_development": 21914000000,
							"selling_general_and_administrative": 21973000000,
							"other_operating_expenses": null
						},
						"operating_income": 108949000000,
						"non_operating_interest": {
							"income": 2843000000,
							"expense": 2645000000
						},
						"other_income_expense": 60000000,
						"pretax_income": 109207000000,
						"income_tax": 14527000000,
						"net_income": 94680000000,
						"eps_basic": 5.67,
						"eps_diluted": 5.61,
						"basic_shares_outstanding": 16701272000,
						"diluted_shares_outstanding": 16701272000,
						"ebitda": 123136000000
					}
				]
			}`,
			wantResp: &response.IncomeStatements{
				Meta: response.IncomeStatementsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Annual",
				},
				IncomeStatement: []response.IncomeStatement{{
					FiscalDate:  "2021-09-30",
					Sales:       null.IntFrom(365817000000),
					CostOfGoods: null.IntFrom(212981000000),
					GrossProfit: null.IntFrom(152836000000),
					OperatingExpense: response.IncomeStatementOperatingExpense{
						ResearchAndDevelopment:          null.IntFrom(21914000000),
						SellingGeneralAndAdministrative: null.IntFrom(21973000000),
						OtherOperatingExpenses:          null.Int{NullInt64: sql.NullInt64{}},
					},
					OperatingIncome: null.IntFrom(108949000000),
					NonOperatingInterest: response.IncomeStatementNonOperatingInterest{
						Income:  null.IntFrom(2843000000),
						Expense: null.IntFrom(2645000000),
					},
					OtherIncomeExpense:       null.IntFrom(60000000),
					PretaxIncome:             null.IntFrom(109207000000),
					IncomeTax:                null.IntFrom(14527000000),
					NetIncome:                null.IntFrom(94680000000),
					EpsBasic:                 null.FloatFrom(5.67),
					EpsDiluted:               null.FloatFrom(5.61),
					BasicSharesOutstanding:   null.IntFrom(16701272000),
					DilutedSharesOutstanding: null.IntFrom(16701272000),
					Ebitda:                   null.IntFrom(123136000000),
				}},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				period:    "annual",
				startDate: "",
				endDate:   "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOTFOUND",
				exchange:  "",
				micCode:   "",
				country:   "",
				period:    "annual",
				startDate: "",
				endDate:   "",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				micCode:   "",
				country:   "",
				period:    "annual",
				startDate: "",
				endDate:   "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetIncomeStatement(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
				tt.args.period,
				tt.args.startDate,
				tt.args.endDate,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetInsiderTransactions(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol   string
		exchange string
		micCode  string
		country  string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.InsiderTransactions
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"meta": {
					"symbol": "AAPL",
					"name": "Apple Inc",
					"currency": "USD",
					"exchange": "NASDAQ",
					"mic_code": "XNAS",
					"exchange_timezone": "America/New_York"
				},
				"insider_transactions": [
					{
						"full_name": "COOK TIMOTHY D",
						"position": "Chief Executive Officer",
						"date_reported": "2021-08-25",
						"is_direct": true,
						"shares": 2386440,
						"value": 354568479,
						"description": "Sale at price 148.30 - 149.97 per share."
					}
				]
			}`,
			wantResp: &response.InsiderTransactions{
				Meta: response.InsiderTransactionsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
				},
				InsiderTransactions: []response.InsiderTransaction{
					{
						FullName:     "COOK TIMOTHY D",
						Position:     "Chief Executive Officer",
						DateReported: "2021-08-25",
						IsDirect:     true,
						Shares:       2386440,
						Value:        354568479,
						Description:  "Sale at price 148.30 - 149.97 per share.",
					},
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "NOTFOUND",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusOK,
			responseBody:    `{"code":404,"message":"Data not found","status":"error"}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				micCode:  "",
				country:  "",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetInsiderTransactions(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.micCode,
				tt.args.country,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetUsage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		fields          fields
		responseCode    int
		responseBody    string
		wantResp        *response.Usage
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"timestamp":"2022-02-11 13:05:55",
				"current_usage":312,
				"plan_limit":610
			}`,
			wantResp: &response.Usage{
				TimeStamp:      "2022-02-11 13:05:55",
				CurrentUsage:   null.IntFrom(312),
				PlanLimit:      null.IntFrom(610),
				DailyUsage:     null.Int{},
				PlanDailyLimit: null.Int{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetUsage()

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetMarketMovers(t *testing.T) {
	t.Parallel()

	type args struct {
		instrument string
		direction  string
		outputSize int
		country    string
		dp         int
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.MarketMovers
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				instrument: "stocks",
				direction:  "gainers",
				outputSize: 30,
				country:    "india",
				dp:         5,
			},
			responseCode: http.StatusOK,
			responseBody: `
			{
				"values": [
					{
						"symbol": "FINKURVE",
						"name": "Finkurve Financial Services Limited",
						"exchange": "BSE",
						"mic_code": "XNAS",
						"datetime": "2022-02-18 14:17:00",
						"last": 50.3,
						"high": 50.3,
						"low": 41.5,
						"volume": 430798,
						"change": 8.8,
						"percent_change": 21.2048
					}
				]
			}
			`,
			wantResp: &response.MarketMovers{
				Values: []response.MarketMover{{
					Symbol:        "FINKURVE",
					Name:          "Finkurve Financial Services Limited",
					Exchange:      "BSE",
					MicCode:       "XNAS",
					Datetime:      "2022-02-18 14:17:00",
					Last:          50.3,
					High:          50.3,
					Low:           41.5,
					Volume:        430798,
					Change:        8.8,
					PercentChange: 21.2048,
				}},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         nil,
		},
		{
			name: "too many requests",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				instrument: "stocks",
				direction:  "gainers",
				outputSize: 30,
				country:    "",
				dp:         5,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				instrument: "notfound",
				direction:  "gainers",
				outputSize: 30,
				country:    "",
				dp:         5,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":404,
				"message":"There is an error in the query. Please check your query and try again. If you're unable to resolve it contact our support at https://twelvedata.com/contact/customer","status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 100,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				instrument: "stocks",
				direction:  "gainers",
				outputSize: 30,
				country:    "",
				dp:         5,
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)
			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetMarketMovers(
				tt.args.instrument,
				tt.args.direction,
				tt.args.outputSize,
				tt.args.country,
				tt.args.dp,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}

func TestCli_GetMarketState(t *testing.T) {
	t.Parallel()

	type args struct {
		exchange string
		code     string
		country  string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        []response.MarketState
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         error
	}{
		{
			name: "success",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange: "NYSE",
				code:     "XNYS",
				country:  "United States",
			},
			responseCode: http.StatusOK,
			responseBody: `
			[
				{
					"name": "NYSE",
					"code": "XNYS",
					"country": "United States",
					"is_market_open": true,
					"time_to_open": "00:00:00",
					"time_to_close": "05:20:57"
				}
			]
			`,
			wantResp: []response.MarketState{
				{
					Name:         "NYSE",
					Code:         "XNYS",
					Country:      "United States",
					IsMarketOpen: true,
					TimeToOpen:   "00:00:00",
					TimeToClose:  "05:20:57",
				},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange: "NYSE",
				code:     "XNYS",
				country:  "United States",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 1000 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange: "NYSE",
				code:     "XNYS",
				country:  "United States",
			},
			responseCode: http.StatusOK,

			responseBody:    `[]`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrNotFound,
		},
		{
			name: "500 internal server error",

			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				exchange: "NYSE",
				code:     "XNYS",
				country:  "United States",
			},
			responseCode:    http.StatusInternalServerError,
			responseBody:    ``,
			wantResp:        nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         dictionary.ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.fields.cfg.BaseURL = startServer(t, tt.responseCode, tt.wantCreditsLeft, tt.wantCreditsUsed, tt.responseBody)

			c := NewCli(tt.fields.cfg, NewHTTPCli(tt.fields.httpCli, tt.fields.cfg, tt.fields.logger), tt.fields.logger)
			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetMarketState(
				tt.args.exchange,
				tt.args.code,
				tt.args.country,
			)

			runAssertions(
				t,
				gotCreditsLeft,
				gotCreditsUsed,
				tt.wantCreditsLeft,
				tt.wantCreditsUsed,
				gotErr,
				tt.wantErr,
				gotResp,
				tt.wantResp,
			)
		})
	}
}
