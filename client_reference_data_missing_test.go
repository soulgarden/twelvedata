package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetForexPairs(t *testing.T) {
	type args struct {
		req request.GetForexPairs
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.ForexPairs
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetForexPairs{
					APIKey:        request.APIKey{APIKey: ""},
					CurrencyBase:  "EUR",
					CurrencyQuote: "USD",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "EUR/USD",
					      "currency_group": "Major",
					      "currency_base": "EUR",
					      "currency_quote": "USD"
					    }
					  ],
					  "status": "ok"
					}`,
					"/forex_pairs?currency_base=EUR&currency_quote=USD",
				),
			},
			want: response.ForexPairs{
				Data: []*response.ForexPair{
					{
						Symbol:        "EUR/USD",
						CurrencyGroup: "Major",
						CurrencyBase:  "EUR",
						CurrencyQuote: "USD",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getForexPairs: NewEndpoint[request.GetForexPairs, response.ForexPairs, response.Credits, error](httpCli, url+"/forex_pairs"),
					}
				},
				func(cli interface{}, req request.GetForexPairs) (response.ForexPairs, response.Credits, error) {
					return cli.(client).GetForexPairs(req)
				},
				"GetForexPairs",
			)
		})
	}
}

func Test_client_GetCryptocurrencies(t *testing.T) {
	type args struct {
		req request.GetCryptocurrencies
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Cryptocurrencies
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCryptocurrencies{
					APIKey:        request.APIKey{APIKey: ""},
					CurrencyBase:  "BTC",
					CurrencyQuote: "USD",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "BTC/USD",
					      "available_exchanges": ["Binance"],
					      "currency_base": "BTC",
					      "currency_quote": "USD"
					    }
					  ],
					  "status": "ok"
					}`,
					"/cryptocurrencies?currency_base=BTC&currency_quote=USD",
				),
			},
			want: response.Cryptocurrencies{
				Data: []*response.Cryptocurrency{
					{
						Symbol:             "BTC/USD",
						AvailableExchanges: []string{"Binance"},
						CurrencyBase:       "BTC",
						CurrencyQuote:      "USD",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getCryptocurrencies: NewEndpoint[request.GetCryptocurrencies, response.Cryptocurrencies, response.Credits, error](httpCli, url+"/cryptocurrencies"),
					}
				},
				func(cli interface{}, req request.GetCryptocurrencies) (response.Cryptocurrencies, response.Credits, error) {
					return cli.(client).GetCryptocurrencies(req)
				},
				"GetCryptocurrencies",
			)
		})
	}
}

func Test_client_GetCommodities(t *testing.T) {
	type args struct {
		req request.GetCommodities
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Commodities
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCommodities{
					APIKey:   request.APIKey{APIKey: ""},
					Category: "Energy",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "WTI",
					      "name": "Crude Oil WTI",
					      "category": "Energy",
					      "description": "West Texas Intermediate"
					    }
					  ],
					  "status": "ok"
					}`,
					"/commodities?category=Energy",
				),
			},
			want: response.Commodities{
				Data: []*response.Commodity{
					{
						Symbol:      "WTI",
						Name:        "Crude Oil WTI",
						Category:    "Energy",
						Description: "West Texas Intermediate",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getCommodities: NewEndpoint[request.GetCommodities, response.Commodities, response.Credits, error](httpCli, url+"/commodities"),
					}
				},
				func(cli interface{}, req request.GetCommodities) (response.Commodities, response.Credits, error) {
					return cli.(client).GetCommodities(req)
				},
				"GetCommodities",
			)
		})
	}
}

func Test_client_GetSymbolSearch(t *testing.T) {
	type args struct {
		req request.GetSymbolSearch
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.SymbolSearch
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetSymbolSearch{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					OutputSize: 2,
					ShowPlan:   true,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "AAPL",
					      "instrument_name": "Apple Inc",
					      "exchange": "NASDAQ",
					      "mic_code": "XNAS",
					      "exchange_timezone": "America/New_York",
					      "instrument_type": "Common Stock",
					      "country": "United States",
					      "currency": "USD",
					      "access": {
					        "global": "Basic",
					        "plan": "Basic"
					      }
					    }
					  ],
					  "status": "ok"
					}`,
					"/symbol_search?outputsize=2&show_plan=true&symbol=AAPL",
				),
			},
			want: response.SymbolSearch{
				Data: []*response.SymbolSearchResult{
					{
						Symbol:           "AAPL",
						InstrumentName:   "Apple Inc",
						Exchange:         "NASDAQ",
						MicCode:          "XNAS",
						ExchangeTimezone: "America/New_York",
						InstrumentType:   "Common Stock",
						Country:          "United States",
						Currency:         "USD",
						Access: &response.SymbolSearchResultAccess{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getSymbolSearch: NewEndpoint[request.GetSymbolSearch, response.SymbolSearch, response.Credits, error](httpCli, url+"/symbol_search"),
					}
				},
				func(cli interface{}, req request.GetSymbolSearch) (response.SymbolSearch, response.Credits, error) {
					return cli.(client).GetSymbolSearch(req)
				},
				"GetSymbolSearch",
			)
		})
	}
}

func Test_client_GetEarliestTimestamp(t *testing.T) {
	type args struct {
		req request.GetEarliestTimestamp
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.EarliestTimestamp
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEarliestTimestamp{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "datetime": "2023-01-01 00:00:00",
					  "unix_time": 1672531200
					}`,
					"/earliest_timestamp?interval=1day&symbol=AAPL",
				),
			},
			want: response.EarliestTimestamp{
				Datetime: "2023-01-01 00:00:00",
				UnixTime: null.IntFrom(1672531200),
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getEarliestTimestamp: NewEndpoint[request.GetEarliestTimestamp, response.EarliestTimestamp, response.Credits, error](httpCli, url+"/earliest_timestamp"),
					}
				},
				func(cli interface{}, req request.GetEarliestTimestamp) (response.EarliestTimestamp, response.Credits, error) {
					return cli.(client).GetEarliestTimestamp(req)
				},
				"GetEarliestTimestamp",
			)
		})
	}
}

func Test_client_GetExchangeSchedule(t *testing.T) {
	type args struct {
		req request.GetExchangeSchedule
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.ExchangeSchedule
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetExchangeSchedule{
					APIKey:  request.APIKey{APIKey: ""},
					Date:    "2024-01-01",
					MicCode: "XNAS",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "title": "NASDAQ",
					      "name": "NASDAQ",
					      "code": "XNAS",
					      "country": "United States",
					      "time_zone": "America/New_York",
					      "sessions": [
					        {
					          "open_time": "09:30",
					          "close_time": "16:00",
					          "session_name": "Regular",
					          "session_type": "regular"
					        }
					      ]
					    }
					  ],
					  "status": "ok"
					}`,
					"/exchange_schedule?date=2024-01-01&mic_code=XNAS",
				),
			},
			want: response.ExchangeSchedule{
				Data: []*response.ExchangeScheduleItem{
					{
						Title:    "NASDAQ",
						Name:     "NASDAQ",
						Code:     "XNAS",
						Country:  "United States",
						TimeZone: "America/New_York",
						Sessions: []*response.ExchangeScheduleSession{
							{
								OpenTime:    "09:30",
								CloseTime:   "16:00",
								SessionName: "Regular",
								SessionType: "regular",
							},
						},
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getExchangeSchedule: NewEndpoint[request.GetExchangeSchedule, response.ExchangeSchedule, response.Credits, error](httpCli, url+"/exchange_schedule"),
					}
				},
				func(cli interface{}, req request.GetExchangeSchedule) (response.ExchangeSchedule, response.Credits, error) {
					return cli.(client).GetExchangeSchedule(req)
				},
				"GetExchangeSchedule",
			)
		})
	}
}

func Test_client_GetCryptocurrencyExchanges(t *testing.T) {
	type args struct {
		req request.GetCryptocurrencyExchanges
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.CryptocurrencyExchanges
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCryptocurrencyExchanges{
					APIKey: request.APIKey{APIKey: ""},
					Format: "JSON",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "name": "Binance"
					    }
					  ],
					  "status": "ok"
					}`,
					"/cryptocurrency_exchanges?format=JSON",
				),
			},
			want: response.CryptocurrencyExchanges{
				Data: []*response.CryptocurrencyExchange{
					{
						Name: "Binance",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getCryptocurrencyExchanges: NewEndpoint[request.GetCryptocurrencyExchanges, response.CryptocurrencyExchanges, response.Credits, error](httpCli, url+"/cryptocurrency_exchanges"),
					}
				},
				func(cli interface{}, req request.GetCryptocurrencyExchanges) (response.CryptocurrencyExchanges, response.Credits, error) {
					return cli.(client).GetCryptocurrencyExchanges(req)
				},
				"GetCryptocurrencyExchanges",
			)
		})
	}
}

func Test_client_GetCountries(t *testing.T) {
	type args struct {
		req request.GetCountries
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Countries
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCountries{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "iso2": "US",
					      "iso3": "USA",
					      "numeric": "840",
					      "name": "United States",
					      "official_name": "United States of America",
					      "capital": "Washington",
					      "currency": "USD"
					    }
					  ],
					  "status": "ok"
					}`,
					"/countries",
				),
			},
			want: response.Countries{
				Data: []*response.Country{
					{
						Iso2:         "US",
						Iso3:         "USA",
						Numeric:      "840",
						Name:         "United States",
						OfficialName: "United States of America",
						Capital:      "Washington",
						Currency:     "USD",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getCountries: NewEndpoint[request.GetCountries, response.Countries, response.Credits, error](httpCli, url+"/countries"),
					}
				},
				func(cli interface{}, req request.GetCountries) (response.Countries, response.Credits, error) {
					return cli.(client).GetCountries(req)
				},
				"GetCountries",
			)
		})
	}
}

func Test_client_GetInstrumentType(t *testing.T) {
	type args struct {
		req request.GetInstrumentType
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.InstrumentType
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetInstrumentType{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "result": ["Stock", "ETF"],
					  "status": "ok"
					}`,
					"/instrument_type",
				),
			},
			want: response.InstrumentType{
				Result: []string{"Stock", "ETF"},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
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
						getInstrumentType: NewEndpoint[request.GetInstrumentType, response.InstrumentType, response.Credits, error](httpCli, url+"/instrument_type"),
					}
				},
				func(cli interface{}, req request.GetInstrumentType) (response.InstrumentType, response.Credits, error) {
					return cli.(client).GetInstrumentType(req)
				},
				"GetInstrumentType",
			)
		})
	}
}
