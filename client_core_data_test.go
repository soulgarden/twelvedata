package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetMarketMovers(t *testing.T) {
	type args struct {
		req request.GetMarketMovers
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.MarketMovers
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMarketMovers{
					APIKey:           request.APIKey{APIKey: ""},
					Market:           "stocks",
					Direction:        "gainers",
					PriceGreaterThan: "10",
				},
				// Mock uses only the first object from the array
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "values": [
					    {
					      "symbol": "BSET",
					      "name": "Bassett Furniture Industries Inc",
					      "exchange": "NASDAQ",
					      "mic_code": "XNAS",
					      "datetime": "2023-10-01T12:00:00Z",
					      "last": 17.25,
					      "high": 18,
					      "low": 16.5,
					      "volume": 108297,
					      "change": 3.31,
					      "percent_change": 23.74462
					    }
					  ],
					  "status": "ok"
					}`,
					"/market_movers/stocks?direction=gainers&price_greater_than=10",
				),
			},
			want: response.MarketMovers{
				Values: []response.MarketMover{
					{
						Symbol:        "BSET",
						Name:          "Bassett Furniture Industries Inc",
						Exchange:      "NASDAQ",
						MicCode:       "XNAS",
						Datetime:      "2023-10-01T12:00:00Z",
						Last:          17.25,
						High:          18,
						Low:           16.5,
						Volume:        108297,
						Change:        3.31,
						PercentChange: 23.74462,
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/market_movers/stocks?direction=gainers&price_greater_than=10",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetMarketMovers{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/market_movers/",
				),
			},
			want:  response.MarketMovers{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getMarketMovers: NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](httpCli, url+"/market_movers/{market}"),
					}
				},
				func(cli interface{}, req request.GetMarketMovers) (response.MarketMovers, response.Credits, error) {
					return cli.(client).GetMarketMovers(req)
				},
				"GetMarketMovers",
			)
		})
	}
}

func Test_client_GetQuote(t *testing.T) {
	type args struct {
		req request.GetQuote
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Quote
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetQuote{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "symbol": "AAPL",
					  "name": "Apple Inc",
					  "exchange": "NASDAQ",
					  "mic_code": "XNAS",
					  "currency": "USD",
					  "datetime": "2021-09-16",
					  "timestamp": 1631772000,
					  "last_quote_at": 1631772000,
					  "open": "148.44000",
					  "high": "148.96840",
					  "low": "147.22099",
					  "close": "148.85001",
					  "volume": "67903927",
					  "previous_close": "149.09000",
					  "change": "-0.23999",
					  "percent_change": "-0.16097",
					  "average_volume": "83571571",
					  "rolling_1d_change": "123.123",
					  "rolling_7d_change": "123.123",
					  "rolling_change": "123.123",
					  "is_market_open": false,
					  "fifty_two_week": {
					    "low": "103.10000",
					    "high": "157.25999",
					    "low_change": "45.75001",
					    "high_change": "-8.40999",
					    "low_change_percent": "44.37440",
					    "high_change_percent": "-5.34782",
					    "range": "103.099998 - 157.259995"
					  },
					  "extended_change": "0.09",
					  "extended_percent_change": "0.05",
					  "extended_price": "125.22",
					  "extended_timestamp": "1649845281"
					}`,
					"/quote",
				),
			},
			want: response.Quote{
				Symbol:          "AAPL",
				Name:            "Apple Inc",
				Exchange:        "NASDAQ",
				MicCode:         "XNAS",
				Currency:        "USD",
				Datetime:        "2021-09-16",
				Timestamp:       1631772000,
				LastQuoteAt:     1631772000,
				Open:            "148.44000",
				High:            "148.96840",
				Low:             "147.22099",
				Close:           "148.85001",
				Volume:          "67903927",
				PreviousClose:   "149.09000",
				Change:          "-0.23999",
				PercentChange:   "-0.16097",
				AverageVolume:   "83571571",
				Rolling1DChange: "123.123",
				Rolling7DChange: "123.123",
				RollingChange:   "123.123",
				IsMarketOpen:    false,
				FiftyTwoWeek: &response.QuoteFiftyTwoWeek{
					Low:               "103.10000",
					High:              "157.25999",
					LowChange:         "45.75001",
					HighChange:        "-8.40999",
					LowChangePercent:  "44.37440",
					HighChangePercent: "-5.34782",
					Range:             "103.099998 - 157.259995",
				},
				ExtendedChange:        "0.09",
				ExtendedPercentChange: "0.05",
				ExtendedPrice:         "125.22",
				ExtendedTimestamp:     null.StringFrom("1649845281"),
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetQuote{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/quote",
				),
			},
			want:  response.Quote{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getQuote: NewEndpoint[request.GetQuote, response.Quote, response.Credits, error](httpCli, url+"/quote"),
					}
				},
				func(cli interface{}, req request.GetQuote) (response.Quote, response.Credits, error) {
					return cli.(client).GetQuote(req)
				},
				"GetQuote",
			)
		})
	}
}

func Test_client_GetTimeSeries(t *testing.T) {
	type args struct {
		req request.GetTimeSeries
		url string
	}

	tests := []struct {
		name           string
		args           args
		wantTimeSeries response.TimeSeries
		wantCredits    response.Credits
		wantErr        string
		expectedURL    string
	}{
		{
			name: "success",
			args: args{
				req: request.GetTimeSeries{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1min",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "type": "Common Stock"
					  },
					  "values": [
					    {
					      "datetime": "2021-09-16 15:59:00",
					      "open": "148.73500",
					      "high": "148.86000",
					      "low": "148.73000",
					      "close": "148.85001",
					      "volume": "624277"
					    }
					  ],
					  "status": "ok"
					}`,
					"/",
				),
			},
			wantTimeSeries: response.TimeSeries{
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
						Datetime: "2021-09-16 15:59:00",
						Open:     "148.73500",
						High:     "148.86000",
						Low:      "148.73000",
						Close:    "148.85001",
						Volume:   "624277",
					},
				},
				Status: "ok",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetTimeSeries{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/",
				),
			},
			wantTimeSeries: response.TimeSeries{
				Meta:   response.TimeSeriesMeta{},
				Values: nil,
				Status: "",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.wantTimeSeries,
				tt.wantCredits,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getTimeSeries: NewEndpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetTimeSeries) (response.TimeSeries, response.Credits, error) {
					return cli.(client).GetTimeSeries(req)
				},
				"GetTimeSeries",
			)
		})
	}
}

func Test_client_GetPrice(t *testing.T) {
	type args struct {
		req request.GetPrice
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Price
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetPrice{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"price": "200.99001"}`,
					"/price?symbol=AAPL",
				),
			},
			want:        response.Price{Price: "200.99001"},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/price?symbol=AAPL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getPrice: NewEndpoint[request.GetPrice, response.Price, response.Credits, error](httpCli, url+"/price"),
					}
				},
				func(cli interface{}, req request.GetPrice) (response.Price, response.Credits, error) {
					return cli.(client).GetPrice(req)
				},
				"GetPrice",
			)
		})
	}
}

func Test_client_GetEOD(t *testing.T) {
	type args struct {
		req request.GetEOD
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.EOD
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEOD{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "AAPL",
					Date:   "2021-09-16",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
						"symbol": "AAPL",
						"exchange": "NASDAQ",
						"mic_code": "XNAS",
						"currency": "USD",
						"datetime": "2021-09-16",
						"close": "148.79"
					}`,
					"/eod?date=2021-09-16&symbol=AAPL",
				),
			},
			want: response.EOD{
				Symbol:   "AAPL",
				Exchange: "NASDAQ",
				MicCode:  "XNAS",
				Currency: "USD",
				Datetime: "2021-09-16",
				Close:    "148.79",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/eod?date=2021-09-16&symbol=AAPL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getEOD: NewEndpoint[request.GetEOD, response.EOD, response.Credits, error](httpCli, url+"/eod"),
					}
				},
				func(cli interface{}, req request.GetEOD) (response.EOD, response.Credits, error) {
					return cli.(client).GetEOD(req)
				},
				"GetEOD",
			)
		})
	}
}

func Test_client_GetTimeSeriesCross(t *testing.T) {
	type args struct {
		req request.GetTimeSeriesCross
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.TimeSeriesCross
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetTimeSeriesCross{
					APIKey:   request.APIKey{APIKey: ""},
					Base:     "JPY",
					Quote:    "BTC",
					Interval: "1day",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
						"meta": {
							"base_instrument": "JPY/USD",
							"base_currency": "",
							"base_exchange": "PHYSICAL CURRENCY",
							"interval": "1day",
							"quote_instrument": "BTC/USD",
							"quote_currency": "",
							"quote_exchange": "Coinbase Pro"
						},
						"values": [
							{
								"datetime": "2025-02-28 14:30:00",
								"open": "0.0000081115665",
								"high": "0.0000081273069",
								"low": "0.0000081088287",
								"close": "0.0000081268066"
							}
						]
					}`,
					"/time_series/cross?base=JPY&interval=1day&quote=BTC",
				),
			},
			want: response.TimeSeriesCross{
				Meta: response.TimeSeriesCrossMeta{
					BaseInstrument:  "JPY/USD",
					BaseCurrency:    "",
					BaseExchange:    "PHYSICAL CURRENCY",
					Interval:        "1day",
					QuoteInstrument: "BTC/USD",
					QuoteCurrency:   "",
					QuoteExchange:   "Coinbase Pro",
				},
				Values: []response.TimeSeriesCrossValue{
					{
						Datetime: "2025-02-28 14:30:00",
						Open:     "0.0000081115665",
						High:     "0.0000081273069",
						Low:      "0.0000081088287",
						Close:    "0.0000081268066",
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/time_series/cross?base=JPY&interval=1day&quote=BTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getTimeSeriesCross: NewEndpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error](httpCli, url+"/time_series/cross"),
					}
				},
				func(cli interface{}, req request.GetTimeSeriesCross) (response.TimeSeriesCross, response.Credits, error) {
					return cli.(client).GetTimeSeriesCross(req)
				},
				"GetTimeSeriesCross",
			)
		})
	}
}
