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

type fields struct {
	cfg     *Conf
	httpCli *fasthttp.Client
	logger  *zerolog.Logger
}

func startServer(t *testing.T, responseCode, wantCreditsLeft, wantCreditsUsed int, responseBody string) string {
	t.Helper()

	s := httptest.NewServer(http.HandlerFunc(func(cw http.ResponseWriter, sr *http.Request) {
		if responseCode == http.StatusInternalServerError {
			cw.WriteHeader(responseCode)
		}

		cw.Header().Add("api-credits-left", strconv.Itoa(wantCreditsLeft))
		cw.Header().Add("api-credits-used", strconv.Itoa(wantCreditsUsed))

		_, err := cw.Write([]byte(responseBody))
		if err != nil {
			t.Error(err)
		}
	}))

	t.Cleanup(func() {
		s.Close()
	})

	return s.URL
}

func runAssertions(
	t *testing.T,
	gotCreditsLeft, gotCreditsUsed, wantCreditsLeft, wantCreditsUsed int,
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

// nolint: funlen
func TestCli_GetStocks(t *testing.T) {
	t.Parallel()

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
		wantResp        *response.Stocks
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Stocks{
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
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Stocks{
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
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetStocks(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.country,
				tt.args.instrumentType,
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

// nolint: funlen
func TestCli_GetExchanges(t *testing.T) {
	t.Parallel()

	type args struct {
		instrumentType string
		name           string
		code           string
		country        string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Exchanges
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
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
			// nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetExchanges(
				tt.args.instrumentType,
				tt.args.name,
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

// nolint:funlen
func TestCli_GetEtfs(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Etfs
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Etfs{
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
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
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
			wantResp: &response.Etfs{
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetEtfs(tt.args.symbol)

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

// nolint:funlen
func TestCli_GetIndices(t *testing.T) {
	t.Parallel()

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
		wantResp        *response.Indices
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
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
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetIndices(tt.args.symbol, tt.args.country)

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

// nolint: funlen
func TestCli_GetTimeSeries(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol         string
		interval       string
		exchange       string
		country        string
		instrumentType string
		outputSize     int
		prePost        string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.TimeSeries
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
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
				Meta: &response.TimeSeriesMeta{
					Symbol:           "AAPL",
					Interval:         "1min",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					Type:             "Common Stock",
				},
				Values: []*response.TimeSeriesValue{
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
			// nolint: exhaustivestruct
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
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
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
			wantErr:         dictionary.ErrInvalidTwelveDataResponse,
		},
		{
			name: "500 internal server error",
			// nolint: exhaustivestruct
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetTimeSeries(
				tt.args.symbol,
				tt.args.interval,
				tt.args.exchange,
				tt.args.country,
				tt.args.instrumentType,
				tt.args.outputSize,
				tt.args.prePost,
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

// nolint: funlen
func TestCli_GetExchangeRate(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		timeZone  string
		precision int
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.ExchangeRate
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "USD/JPY",
				timeZone:  "",
				precision: 0,
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "USD/JPY",
				timeZone:  "",
				precision: 0,
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbol",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOT/FOUND",
				timeZone:  "",
				precision: 0,
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
			wantErr:         dictionary.ErrInvalidTwelveDataResponse,
		},
		{
			name: "500 internal server error",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "USD/JPY",
				timeZone:  "",
				precision: 0,
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetExchangeRate(
				tt.args.symbol,
				tt.args.timeZone,
				tt.args.precision,
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

// nolint:funlen
func TestCli_GetQuotes(t *testing.T) {
	t.Parallel()

	type args struct {
		interval         string
		exchange         string
		country          string
		volumeTimePeriod string
		instrumentType   string
		prePost          string
		timezone         string
		decimalPlaces    int
		symbols          []string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Quotes
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				interval:         "1day",
				exchange:         "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prePost:          "",
				timezone:         "",
				decimalPlaces:    5,
				symbols:          []string{"AAPL"},
			},
			responseCode: http.StatusOK,

			responseBody: `
			{
				"symbol":"AAPL",
				"name":"Apple Inc",
				"exchange":"NASDAQ",
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
				Data: []*response.Quote{{
					Symbol:        "AAPL",
					Name:          "Apple Inc",
					Exchange:      "NASDAQ",
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
				Errors: []*response.QuoteError{},
			},
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         nil,
		},
		{
			name: "too many requests",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				interval:         "1day",
				exchange:         "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prePost:          "",
				timezone:         "",
				decimalPlaces:    5,
				symbols:          []string{"AAPL"},
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 1,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				interval:         "1day",
				exchange:         "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prePost:          "",
				timezone:         "",
				decimalPlaces:    5,
				symbols:          []string{"NOTFOUND1", "NOTFOUND2"},
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
				Data: []*response.Quote{},
				Errors: []*response.QuoteError{
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				interval:         "1day",
				exchange:         "",
				country:          "",
				volumeTimePeriod: "",
				instrumentType:   "",
				prePost:          "",
				timezone:         "",
				decimalPlaces:    5,
				symbols:          []string{"AAPL"},
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetQuotes(
				tt.args.interval,
				tt.args.exchange,
				tt.args.country,
				tt.args.volumeTimePeriod,
				tt.args.instrumentType,
				tt.args.prePost,
				tt.args.timezone,
				tt.args.decimalPlaces,
				tt.args.symbols,
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

// nolint: funlen
func TestCli_GetProfile(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol   string
		exchange string
		country  string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Profile
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				country:  "",
			},
			responseCode: http.StatusOK,
			// nolint: lll
			responseBody: `
			{
				"symbol":"AAPL",
				"name":"Apple Inc",
				"exchange":"NASDAQ",
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
				Symbol:    "AAPL",
				Name:      "Apple Inc",
				Exchange:  "NASDAQ",
				Sector:    "Technology",
				Industry:  "Consumer Electronics",
				Employees: 154000,
				Website:   "https://www.apple.com",
				// nolint: lll
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
				country:  "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 10,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "NOTFOUND",
				exchange: "",
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:   "AAPL",
				exchange: "",
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetProfile(
				tt.args.symbol,
				tt.args.exchange,
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

// nolint: funlen
func TestCli_GetDividends(t *testing.T) {
	t.Parallel()

	type args struct {
		symbol    string
		exchange  string
		country   string
		r         string
		startTime string
		endTime   string
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		responseCode    int
		responseBody    string
		wantResp        *response.Dividends
		wantCreditsLeft int
		wantCreditsUsed int
		wantErr         error
	}{
		{
			name: "success",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				r:         "last",
				startTime: "",
				endTime:   "",
			},
			responseCode: http.StatusOK,

			responseBody: `
			{
				"meta":{
					"symbol":"AAPL",
					"name":"Apple Inc",
					"currency":"USD",
					"exchange":"NASDAQ",
					"exchange_timezone":"America/New_York"
				},
				"dividends":[
					{"payment_date":"2022-02-04","amount":0.22}
				]
			}`,
			wantResp: &response.Dividends{
				Meta: &response.DividendsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					ExchangeTimezone: "America/New_York",
				},
				Dividends: []*response.Dividend{
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				r:         "last",
				startTime: "",
				endTime:   "",
			},
			responseCode: http.StatusOK,
			//nolint: lll
			responseBody: `{
				"code":429,
				"message":"You have run out of API credits for the current minute. 10 API credits were used, with the current limit being 987. Wait for the next minute or consider switching to a higher tier plan at https://twelvedata.com/pricing",
				"status":"error"
			}`,
			wantResp:        nil,
			wantCreditsLeft: 10,
			wantCreditsUsed: 20,
			wantErr:         dictionary.ErrTooManyRequests,
		},
		{
			name: "not found symbols",
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "NOTFOUND",
				exchange:  "",
				country:   "",
				r:         "last",
				startTime: "",
				endTime:   "",
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
			// nolint: exhaustivestruct
			fields: fields{
				cfg:     &Conf{Timeout: 10, APIKey: "demo"},
				httpCli: &fasthttp.Client{},
				logger:  &zerolog.Logger{},
			},
			args: args{
				symbol:    "AAPL",
				exchange:  "",
				country:   "",
				r:         "last",
				startTime: "",
				endTime:   "",
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

			c := NewCli(tt.fields.cfg, tt.fields.httpCli, tt.fields.logger)

			gotResp, gotCreditsLeft, gotCreditsUsed, gotErr := c.GetDividends(
				tt.args.symbol,
				tt.args.exchange,
				tt.args.country,
				tt.args.r,
				tt.args.startTime,
				tt.args.endTime,
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
