package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetETFFullData(t *testing.T) {
	type args struct {
		req request.GetETFFullData
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.ETFFullData
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFFullData{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					800,
					`{
						"summary": {
							"symbol": "VTI",
							"name": "Vanguard Total Stock Market ETF",
							"currency": "USD",
							"exchange": "ARCX",
							"country": "United States",
							"asset_class": "Equity",
							"net_assets": 1500000000000.0,
							"expense_ratio": 0.03,
							"inception_date": "2001-05-24",
							"last_updated": "2024-01-15"
						},
						"performance": {
							"trailing_returns": {
								"1d": -0.5,
								"5d": 1.2,
								"1m": 3.4,
								"3m": 8.7,
								"6m": 12.3,
								"1y": 15.6,
								"3y": 9.8,
								"5y": 11.2,
								"10y": 13.4
							},
							"annual_returns": [
								{"year": 2023, "return": 26.29},
								{"year": 2022, "return": -19.47}
							]
						},
						"risk": {
							"beta": 1.0,
							"volatility": 15.2,
							"standard_deviation": 14.8,
							"sharpe_ratio": 0.85,
							"alpha": 0.12,
							"r_squared": 0.98
						},
						"composition": {
							"top_holdings": [
								{
									"symbol": "AAPL",
									"name": "Apple Inc.",
									"percentage": 7.1,
									"shares": 123456789,
									"market_value": 106500000000.0
								}
							],
							"sector_allocation": [
								{
									"sector": "Technology",
									"percentage": 28.5
								}
							],
							"country_allocation": [
								{
									"country": "United States",
									"percentage": 100.0
								}
							],
							"asset_allocation": {
								"stocks": 99.8,
								"bonds": 0.0,
								"cash": 0.2,
								"other": 0.0
							},
							"last_updated": "2024-01-15"
						}
					}`,
					"/?symbol=VTI",
				),
			},
			want: response.ETFFullData{
				Summary: response.ETFSummary{
					Symbol:        "VTI",
					Name:          "Vanguard Total Stock Market ETF",
					Currency:      "USD",
					Exchange:      "ARCX",
					Country:       "United States",
					AssetClass:    "Equity",
					NetAssets:     null.FloatFrom(1.5e+12),
					ExpenseRatio:  null.FloatFrom(0.03),
					InceptionDate: null.StringFrom("2001-05-24"),
					LastUpdated:   null.StringFrom("2024-01-15"),
				},
				Performance: response.ETFPerformance{
					TrailingReturns: response.TrailingReturns{
						OneDay:      null.FloatFrom(-0.5),
						FiveDays:    null.FloatFrom(1.2),
						OneMonth:    null.FloatFrom(3.4),
						ThreeMonths: null.FloatFrom(8.7),
						SixMonths:   null.FloatFrom(12.3),
						OneYear:     null.FloatFrom(15.6),
						ThreeYears:  null.FloatFrom(9.8),
						FiveYears:   null.FloatFrom(11.2),
						TenYears:    null.FloatFrom(13.4),
					},
					AnnualReturns: []response.AnnualReturn{
						{Year: 2023, Return: null.FloatFrom(26.29)},
						{Year: 2022, Return: null.FloatFrom(-19.47)},
					},
				},
				Risk: response.ETFFullDataRisk{
					Beta:              null.FloatFrom(1.0),
					Alpha:             null.FloatFrom(0.12),
					StandardDeviation: null.FloatFrom(14.8),
					SharpeRatio:       null.FloatFrom(0.85),
					Volatility:        null.FloatFrom(15.2),
					RSquared:          null.FloatFrom(0.98),
				},
				Composition: response.ETFComposition{
					TopHoldings: []response.Holding{
						{
							Symbol:      "AAPL",
							Name:        "Apple Inc.",
							Percentage:  null.FloatFrom(7.1),
							Shares:      null.IntFrom(123456789),
							MarketValue: null.FloatFrom(1.065e+11),
						},
					},
					SectorAllocation: []response.SectorAllocation{
						{
							Sector:     "Technology",
							Percentage: null.FloatFrom(28.5),
						},
					},
					CountryAllocation: []response.CountryAllocation{
						{
							Country:    "United States",
							Percentage: null.FloatFrom(100),
						},
					},
					AssetAllocation: response.AssetAllocation{
						Stocks: null.FloatFrom(99.8),
						Bonds:  null.FloatFrom(0),
						Cash:   null.FloatFrom(0.2),
						Other:  null.FloatFrom(0),
					},
					LastUpdated: "2024-01-15",
				},
			},
			want1:       response.NewCreditsImpl(100, 800),
			wantErr:     "",
			expectedURL: "/?symbol=VTI",
		},
		{
			name: "wrong_api_key",
			args: args{
				req: request.GetETFFullData{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: mockServerWithURL(
					t,
					http.StatusUnauthorized,
					100,
					800,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?symbol=VTI",
				),
			},
			want:  response.ETFFullData{},
			want1: response.NewCreditsImpl(100, 800),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?symbol=VTI",
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
						getETFFullData: NewEndpoint[request.GetETFFullData, response.ETFFullData, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFFullData) (response.ETFFullData, response.Credits, error) {
					return cli.(client).GetETFFullData(req)
				},
				"GetETFFullData",
			)
		})
	}
}

func Test_client_GetETFPerformance(t *testing.T) {
	type args struct {
		req request.GetETFPerformance
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.ETFPerformance
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFPerformance{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					200,
					`{
						"trailing_returns": {
							"1d": -0.5,
							"5d": 1.2,
							"1m": 3.4,
							"3m": 8.7,
							"6m": 12.3,
							"1y": 15.6,
							"3y": 9.8,
							"5y": 11.2,
							"10y": 13.4
						},
						"annual_returns": [
							{"year": 2023, "return": 26.29},
							{"year": 2022, "return": -19.47}
						]
					}`,
					"/?symbol=VTI",
				),
			},
			want: response.ETFPerformance{
				TrailingReturns: response.TrailingReturns{
					OneDay:      null.FloatFrom(-0.5),
					FiveDays:    null.FloatFrom(1.2),
					OneMonth:    null.FloatFrom(3.4),
					ThreeMonths: null.FloatFrom(8.7),
					SixMonths:   null.FloatFrom(12.3),
					OneYear:     null.FloatFrom(15.6),
					ThreeYears:  null.FloatFrom(9.8),
					FiveYears:   null.FloatFrom(11.2),
					TenYears:    null.FloatFrom(13.4),
				},
				AnnualReturns: []response.AnnualReturn{
					{Year: 2023, Return: null.FloatFrom(26.29)},
					{Year: 2022, Return: null.FloatFrom(-19.47)},
				},
			},
			want1:       response.NewCreditsImpl(100, 200),
			wantErr:     "",
			expectedURL: "/?symbol=VTI",
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
						getETFPerformance: NewEndpoint[request.GetETFPerformance, response.ETFPerformance, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFPerformance) (response.ETFPerformance, response.Credits, error) {
					return cli.(client).GetETFPerformance(req)
				},
				"GetETFPerformance",
			)
		})
	}
}

func Test_client_GetETFComposition(t *testing.T) {
	type args struct {
		req request.GetETFComposition
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.ETFComposition
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFComposition{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
						"top_holdings": [
							{
								"symbol": "AAPL",
								"name": "Apple Inc.",
								"percentage": 7.1,
								"shares": 123456789,
								"market_value": 106500000000.0
							}
						],
						"sector_allocation": [
							{
								"sector": "Technology",
								"percentage": 28.5
							}
						],
						"country_allocation": [
							{
								"country": "United States",
								"percentage": 100.0
							}
						],
						"asset_allocation": {
							"stocks": 99.8,
							"bonds": 0.0,
							"cash": 0.2,
							"other": 0.0
						},
						"last_updated": "2024-01-15"
					}`,
					"/?symbol=VTI",
				),
			},
			want: response.ETFComposition{
				TopHoldings: []response.Holding{
					{
						Symbol:      "AAPL",
						Name:        "Apple Inc.",
						Percentage:  null.FloatFrom(7.1),
						Shares:      null.IntFrom(123456789),
						MarketValue: null.FloatFrom(106500000000.0),
					},
				},
				SectorAllocation: []response.SectorAllocation{
					{Sector: "Technology", Percentage: null.FloatFrom(28.5)},
				},
				CountryAllocation: []response.CountryAllocation{
					{Country: "United States", Percentage: null.FloatFrom(100.0)},
				},
				AssetAllocation: response.AssetAllocation{
					Stocks: null.FloatFrom(99.8),
					Bonds:  null.FloatFrom(0.0),
					Cash:   null.FloatFrom(0.2),
					Other:  null.FloatFrom(0.0),
				},
				LastUpdated: "2024-01-15",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?symbol=VTI",
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
						getETFComposition: NewEndpoint[request.GetETFComposition, response.ETFComposition, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFComposition) (response.ETFComposition, response.Credits, error) {
					return cli.(client).GetETFComposition(req)
				},
				"GetETFComposition",
			)
		})
	}
}
