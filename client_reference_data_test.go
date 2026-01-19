package twelvedata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetExchanges(t *testing.T) {
	type args struct {
		req request.GetExchanges
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Exchanges
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetExchanges{
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
					      "name": "SSE",
					      "code": "XSHG",
					      "country": "China",
					      "timezone": "Asia/Shanghai",
					      "access": {
					        "global": "Level B",
					        "plan": "Pro"
					      }
					    }
					  ],
					  "status": "ok"
					}`,
					"/",
				),
			},
			want: response.Exchanges{
				Data: []response.Exchange{
					{
						Name:     "SSE",
						Code:     "XSHG",
						Country:  "China",
						Timezone: "Asia/Shanghai",
						Access: &response.Access{
							Global: "Level B",
							Plan:   "Pro",
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetExchanges{
					APIKey: request.APIKey{APIKey: ""},
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
			want:  response.Exchanges{},
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
						getExchanges: NewEndpoint[request.GetExchanges, response.Exchanges, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetExchanges) (response.Exchanges, response.Credits, error) {
					return cli.(client).GetExchanges(req)
				},
				"GetExchanges",
			)
		})
	}
}

func Test_client_GetMarketState(t *testing.T) {
	type args struct {
		req request.GetMarketState
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        []response.MarketState
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMarketState{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				// Mock uses only one object from the array
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`[
					  {
					    "name": "NYSE",
					    "code": "XNYS",
					    "country": "United States",
					    "is_market_open": true,
					    "time_after_open": "02:39:03",
					    "time_to_open": "00:00:00",
					    "time_to_close": "05:20:57"
					  }
					]`,
					"/",
				),
			},
			want: []response.MarketState{
				{
					Name:          "NYSE",
					Code:          "XNYS",
					Country:       "United States",
					IsMarketOpen:  true,
					TimeAfterOpen: "02:39:03",
					TimeToOpen:    "00:00:00",
					TimeToClose:   "05:20:57",
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetMarketState{
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
			want:        nil,
			want1:       response.NewCreditsImpl(100, 100),
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
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getMarketState: NewEndpoint[request.GetMarketState, []response.MarketState, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetMarketState) ([]response.MarketState, response.Credits, error) {
					return cli.(client).GetMarketState(req)
				},
				"GetMarketState",
			)
		})
	}
}

func Test_client_GetStocks(t *testing.T) {
	type args struct {
		req request.GetStock
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Stocks
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetStock{
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
					  "data": [
					    {
					      "symbol": "TCS",
					      "name": "Tata Consultancy Services Limited",
					      "currency": "INR",
					      "exchange": "NSE",
					      "mic_code": "XNSE",
					      "country": "India",
					      "type": "Common Stock",
					      "figi_code": "BBG000Q0WGC6",
					      "access": {
					        "global": "Level A",
					        "plan": "Grow"
					      }
					    }
					  ],
					  "status": "ok"
					}`,
					"/",
				),
			},
			want: response.Stocks{
				Data: []*response.Stock{
					{
						Symbol:   "TCS",
						Name:     "Tata Consultancy Services Limited",
						Currency: "INR",
						Exchange: "NSE",
						MicCode:  "XNSE",
						Country:  "India",
						Type:     "Common Stock",
						FigiCode: "BBG000Q0WGC6",
						Access: &response.StockAccess{
							Global: "Level A",
							Plan:   "Grow",
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "real api response format",
			args: args{
				req: request.GetStock{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Exchange: "NASDAQ",
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
					      "name": "Apple Inc",
					      "currency": "USD",
					      "exchange": "NASDAQ",
					      "mic_code": "XNGS",
					      "country": "United States",
					      "type": "Common Stock",
					      "figi_code": "BBG000B9XRY4",
					      "cfi_code": "ESVUFR",
					      "isin": "US0378331005",
					      "cusip": "037833100"
					    },
					    {
					      "symbol": "MSFT",
					      "name": "Microsoft Corp",
					      "currency": "USD",
					      "exchange": "NASDAQ",
					      "mic_code": "XNGS",
					      "country": "United States",
					      "type": "Common Stock",
					      "figi_code": "BBG000BPH459",
					      "cfi_code": "ESVUFR",
					      "isin": "US5949181045",
					      "cusip": "594918104"
					    }
					  ],
					  "count": 2,
					  "status": "ok"
					}`,
					"/?exchange=NASDAQ",
				),
			},
			want: response.Stocks{
				Data: []*response.Stock{
					{
						Symbol:   "AAPL",
						Name:     "Apple Inc",
						Currency: "USD",
						Exchange: "NASDAQ",
						MicCode:  "XNGS",
						Country:  "United States",
						Type:     "Common Stock",
						FigiCode: "BBG000B9XRY4",
						CfiCode:  "ESVUFR",
						Isin:     "US0378331005",
						Cusip:    "037833100",
					},
					{
						Symbol:   "MSFT",
						Name:     "Microsoft Corp",
						Currency: "USD",
						Exchange: "NASDAQ",
						MicCode:  "XNGS",
						Country:  "United States",
						Type:     "Common Stock",
						FigiCode: "BBG000BPH459",
						CfiCode:  "ESVUFR",
						Isin:     "US5949181045",
						Cusip:    "594918104",
					},
				},
				Count:  2,
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetStock{
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
			want: response.Stocks{
				Data:   nil,
				Status: "",
			},
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
						getStocks: NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetStock) (response.Stocks, response.Credits, error) {
					return cli.(client).GetStocks(req)
				},
				"GetStocks",
			)
		})
	}
}

func Test_client_GetETFs(t *testing.T) {
	type args struct {
		req request.GetETFs
		url string
	}

	missingAPIKeyURL := mockServerWithURL(
		t,
		http.StatusUnauthorized,
		100,
		1,
		`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
		"/?symbol=SPY",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFs
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFs{
					APIKey:          request.APIKey{APIKey: ""},
					Symbol:          "SPY",
					FIGI:            "BBG000BDTF76",
					ISIN:            "US0378331005",
					CUSIP:           "037833100",
					CIK:             "95953",
					Exchange:        "NYSE",
					MicCode:         "ARCX",
					Country:         "United States",
					Format:          "JSON",
					Delimiter:       ";",
					ShowPlan:        true,
					IncludeDelisted: true,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1,
					`{
						"data": [
							{
								"symbol": "SPY",
								"name": "SPDR S&P 500 ETF Trust",
								"currency": "USD",
								"exchange": "NYSE",
								"mic_code": "ARCX",
								"country": "United States",
								"figi_code": "BBG000BDTF76",
								"cfi_code": "CECILU",
								"isin": "US78462F1030",
								"cusip": "037833100",
								"access": {
									"global": "Basic",
									"plan": "Basic"
								}
							}
						],
						"status": "ok"
					}`,
					"/?cik=95953&country=United+States&cusip=037833100&delimiter=%3B&exchange=NYSE&figi=BBG000BDTF76&format=JSON&include_delisted=true&isin=US0378331005&mic_code=ARCX&show_plan=true&symbol=SPY",
				),
			},
			want: response.ETFs{
				Data: []*response.ETF{
					{
						Symbol:   "SPY",
						Name:     "SPDR S&P 500 ETF Trust",
						Currency: "USD",
						Exchange: "NYSE",
						MicCode:  "ARCX",
						Country:  "United States",
						FigiCode: "BBG000BDTF76",
						CfiCode:  "CECILU",
						Isin:     "US78462F1030",
						Cusip:    "037833100",
						Access: &response.ETFAccess{
							Global: "Basic",
							Plan:   "Basic",
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?cik=95953&country=United+States&cusip=037833100&delimiter=%3B&exchange=NYSE&figi=BBG000BDTF76&format=JSON&include_delisted=true&isin=US0378331005&mic_code=ARCX&show_plan=true&symbol=SPY",
		},
		{
			name: "missing api key",
			args: args{
				req: request.GetETFs{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "SPY",
				},
				url: missingAPIKeyURL,
			},
			want:  response.ETFs{},
			want1: response.NewCreditsImpl(100, 1),
			wantErr: fmt.Sprintf(
				"HTTP 401 Unauthorized: %s (URL: %s?symbol=SPY)",
				response.Error{
					Code:    401,
					Message: "**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer",
					Status:  "error",
				}.Error(),
				missingAPIKeyURL,
			),
			expectedURL: "/?symbol=SPY",
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
						getETFs: NewEndpoint[request.GetETFs, response.ETFs, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFs) (response.ETFs, response.Credits, error) {
					return cli.(client).GetETFs(req)
				},
				"GetETFs",
			)
		})
	}
}
