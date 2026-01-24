package twelvedata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

const (
	etfWorldSummaryPlanRestrictionMessage     = "/etfs/world/summary is available exclusively with ultra or enterprise plans. Consider upgrading your API Key now at https://twelvedata.com/pricing"
	etfWorldPerformancePlanRestrictionMessage = "/etfs/world/performance is available exclusively with ultra or enterprise plans. Consider upgrading your API Key now at https://twelvedata.com/pricing"
	etfWorldRiskPlanRestrictionMessage        = "/etfs/world/risk is available exclusively with ultra or enterprise plans. Consider upgrading your API Key now at https://twelvedata.com/pricing"
	etfWorldCompositionPlanRestrictionMessage = "/etfs/world/composition is available exclusively with ultra or enterprise plans. Consider upgrading your API Key now at https://twelvedata.com/pricing"
)

func expectedETFWorldComposition() response.ETFWorldComposition {
	return response.ETFWorldComposition{
		MajorMarketSectors: []response.ETFSectorWeight{
			{
				Sector: "Technology",
				Weight: null.FloatFrom(0.2424),
			},
		},
		CountryAllocation: []response.ETFCountryAllocation{
			{
				Country:    "United Kingdom",
				Allocation: null.FloatFrom(0.9855),
			},
		},
		AssetAllocation: response.ETFAssetAllocation{
			Cash:            null.FloatFrom(0.0004),
			Stocks:          null.FloatFrom(0.9996),
			PreferredStocks: null.FloatFrom(0),
			Convertibles:    null.FloatFrom(0),
			Bonds:           null.FloatFrom(0),
			Others:          null.FloatFrom(0),
		},
		TopHoldings: []response.ETFTopHolding{
			{
				Symbol:   "AAPL",
				Name:     "Apple Inc",
				Exchange: "NASDAQ",
				MicCode:  "XNAS",
				Weight:   null.FloatFrom(0.0592),
			},
		},
		BondBreakdown: response.ETFBondBreakdown{
			AverageMaturity: response.ETFBondMetric{
				Fund:     null.FloatFrom(6.65),
				Category: null.FloatFrom(7.81),
			},
			AverageDuration: response.ETFBondMetric{
				Fund:     null.FloatFrom(5.72),
				Category: null.FloatFrom(5.64),
			},
			CreditQuality: []response.ETFCreditQuality{
				{
					Grade:  "AAA",
					Weight: null.FloatFrom(0),
				},
			},
		},
	}
}

func expectedETFWorldRisk() response.ETFWorldRisk {
	return response.ETFWorldRisk{
		VolatilityMeasures: []response.ETFVolatilityMeasure{
			{
				Period:                   "3_year",
				Alpha:                    null.FloatFrom(-0.03),
				AlphaCategory:            null.FloatFrom(-0.02),
				Beta:                     null.FloatFrom(1),
				BetaCategory:             null.FloatFrom(0.01),
				MeanAnnualReturn:         null.FloatFrom(1.58),
				MeanAnnualReturnCategory: null.FloatFrom(0.01),
				RSquared:                 null.FloatFrom(100),
				RSquaredCategory:         null.FloatFrom(0.95),
				Std:                      null.FloatFrom(18.52),
				StdCategory:              null.FloatFrom(0.19),
				SharpeRatio:              null.FloatFrom(0.95),
				SharpeRatioCategory:      null.FloatFrom(0.01),
				TreynorRatio:             null.FloatFrom(17.41),
				TreynorRatioCategory:     null.FloatFrom(0.16),
			},
		},
		ValuationMetrics: response.ETFValuationMetrics{
			PriceToEarnings: null.FloatFrom(26.46),
			PriceToBook:     null.FloatFrom(4.42),
			PriceToSales:    null.FloatFrom(2.96),
			PriceToCashflow: null.FloatFrom(17.57),
		},
	}
}

func Test_client_GetETFSummary(t *testing.T) {
	type args struct {
		req request.GetETFSummary
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		50,
		`{
			"etf": {
				"summary": {
					"symbol": "IVV",
					"name": "iShares Core S&P 500 ETF",
					"fund_family": "iShares",
					"fund_type": "Large Blend",
					"currency": "USD",
					"share_class_inception_date": "2000-11-13",
					"ytd_return": -0.0537,
					"expense_ratio_net": -0.004,
					"yield": 0.0133,
					"nav": 413.24,
					"last_price": 413.24,
					"turnover_rate": 0.04,
					"net_assets": 753409982464,
					"overview": "The investment seeks to track the performance of the Standard & Poor's 500..."
				}
			},
			"status": "ok"
		}`,
		"/?symbol=IVV",
	)

	planRestrictedURL := mockServerWithURL(
		t,
		http.StatusForbidden,
		100,
		50,
		fmt.Sprintf(`{"code":403,"message":"%s","status":"error"}`, etfWorldSummaryPlanRestrictionMessage),
		"/?symbol=IVV",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFWorldSummary
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFSummary{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "IVV",
				},
				url: successURL,
			},
			want: response.ETFWorldSummary{
				ETF: response.ETFWorldSummaryData{
					Summary: response.ETFWorldSummaryInfo{
						Symbol:                  "IVV",
						Name:                    "iShares Core S&P 500 ETF",
						FundFamily:              "iShares",
						FundType:                "Large Blend",
						Currency:                "USD",
						ShareClassInceptionDate: "2000-11-13",
						YTDReturn:               null.FloatFrom(-0.0537),
						ExpenseRatioNet:         null.FloatFrom(-0.004),
						Yield:                   null.FloatFrom(0.0133),
						NAV:                     null.FloatFrom(413.24),
						LastPrice:               null.FloatFrom(413.24),
						TurnoverRate:            null.FloatFrom(0.04),
						NetAssets:               null.IntFrom(753409982464),
						Overview:                "The investment seeks to track the performance of the Standard & Poor's 500...",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     "",
			expectedURL: "/?symbol=IVV",
		},
		{
			name: "plan_restricted",
			args: args{
				req: request.GetETFSummary{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "IVV",
				},
				url: planRestrictedURL,
			},
			want:        response.ETFWorldSummary{},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     fmt.Sprintf("Plan Limitation: %s", etfWorldSummaryPlanRestrictionMessage),
			expectedURL: "/?symbol=IVV",
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
						getETFSummary: NewEndpoint[request.GetETFSummary, response.ETFWorldSummary, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFSummary) (response.ETFWorldSummary, response.Credits, error) {
					return cli.(client).GetETFSummary(req)
				},
				"GetETFSummary",
			)
		})
	}
}

func Test_client_GetETFRisk(t *testing.T) {
	type args struct {
		req request.GetETFRisk
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		50,
		`{
			"etf": {
				"risk": {
					"volatility_measures": [
						{
							"period": "3_year",
							"alpha": -0.03,
							"alpha_category": -0.02,
							"beta": 1,
							"beta_category": 0.01,
							"mean_annual_return": 1.58,
							"mean_annual_return_category": 0.01,
							"r_squared": 100,
							"r_squared_category": 0.95,
							"std": 18.52,
							"std_category": 0.19,
							"sharpe_ratio": 0.95,
							"sharpe_ratio_category": 0.01,
							"treynor_ratio": 17.41,
							"treynor_ratio_category": 0.16
						}
					],
					"valuation_metrics": {
						"price_to_earnings": 26.46,
						"price_to_book": 4.42,
						"price_to_sales": 2.96,
						"price_to_cashflow": 17.57
					}
				}
			},
			"status": "ok"
		}`,
		"/?symbol=IVV",
	)

	planRestrictedURL := mockServerWithURL(
		t,
		http.StatusForbidden,
		100,
		50,
		fmt.Sprintf(`{"code":403,"message":"%s","status":"error"}`, etfWorldRiskPlanRestrictionMessage),
		"/?symbol=IVV",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFRisk
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFRisk{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "IVV",
				},
				url: successURL,
			},
			want: response.ETFRisk{
				ETF: response.ETFRiskData{
					Risk: expectedETFWorldRisk(),
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     "",
			expectedURL: "/?symbol=IVV",
		},
		{
			name: "plan_restricted",
			args: args{
				req: request.GetETFRisk{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "IVV",
				},
				url: planRestrictedURL,
			},
			want:        response.ETFRisk{},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     fmt.Sprintf("Plan Limitation: %s", etfWorldRiskPlanRestrictionMessage),
			expectedURL: "/?symbol=IVV",
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
						getETFRisk: NewEndpoint[request.GetETFRisk, response.ETFRisk, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFRisk) (response.ETFRisk, response.Credits, error) {
					return cli.(client).GetETFRisk(req)
				},
				"GetETFRisk",
			)
		})
	}
}

func Test_client_GetETFFullData(t *testing.T) {
	type args struct {
		req request.GetETFFullData
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		800,
		`{
			"etf": {
				"summary": {
					"symbol": "VTI",
					"name": "Vanguard Total Stock Market ETF",
					"fund_family": "Vanguard",
					"fund_type": "Large Blend",
					"currency": "USD",
					"share_class_inception_date": "2001-05-24",
					"ytd_return": -0.0537,
					"expense_ratio_net": 0.0003,
					"yield": 0.0133,
					"nav": 413.24,
					"last_price": 413.24,
					"turnover_rate": 0.04,
					"net_assets": 753409982464,
					"overview": "The investment seeks to track the performance of the Standard & Poor's 500..."
				},
				"performance": {
					"trailing_returns": [
						{
							"period": "ytd",
							"share_class_return": -0.0751,
							"category_return": 0.1484
						}
					],
					"annual_total_returns": [
						{
							"year": 2021,
							"share_class_return": 0.2866,
							"category_return": 0
						}
					]
				},
				"risk": {
					"volatility_measures": [
						{
							"period": "3_year",
							"alpha": -0.03,
							"alpha_category": -0.02,
							"beta": 1,
							"beta_category": 0.01,
							"mean_annual_return": 1.58,
							"mean_annual_return_category": 0.01,
							"r_squared": 100,
							"r_squared_category": 0.95,
							"std": 18.52,
							"std_category": 0.19,
							"sharpe_ratio": 0.95,
							"sharpe_ratio_category": 0.01,
							"treynor_ratio": 17.41,
							"treynor_ratio_category": 0.16
						}
					],
					"valuation_metrics": {
						"price_to_earnings": 26.46,
						"price_to_book": 4.42,
						"price_to_sales": 2.96,
						"price_to_cashflow": 17.57
					}
				},
				"composition": {
					"major_market_sectors": [
						{
							"sector": "Technology",
							"weight": 0.2424
						}
					],
					"country_allocation": [
						{
							"country": "United Kingdom",
							"allocation": 0.9855
						}
					],
					"asset_allocation": {
						"cash": 0.0004,
						"stocks": 0.9996,
						"preferred_stocks": 0,
						"convertibles": 0,
						"bonds": 0,
						"others": 0
					},
					"top_holdings": [
						{
							"symbol": "AAPL",
							"name": "Apple Inc",
							"exchange": "NASDAQ",
							"mic_code": "XNAS",
							"weight": 0.0592
						}
					],
					"bond_breakdown": {
						"average_maturity": {
							"fund": 6.65,
							"category": 7.81
						},
						"average_duration": {
							"fund": 5.72,
							"category": 5.64
						},
						"credit_quality": [
							{
								"grade": "AAA",
								"weight": 0
							}
						]
					}
				}
			},
			"status": "ok"
		}`,
		"/?symbol=VTI",
	)
	wrongAPIKeyURL := mockServerWithURL(
		t,
		http.StatusUnauthorized,
		100,
		800,
		`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
		"/?symbol=VTI",
	)

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
				url: successURL,
			},
			want: response.ETFFullData{
				ETF: response.ETFFullDataData{
					Summary: response.ETFWorldSummaryInfo{
						Symbol:                  "VTI",
						Name:                    "Vanguard Total Stock Market ETF",
						FundFamily:              "Vanguard",
						FundType:                "Large Blend",
						Currency:                "USD",
						ShareClassInceptionDate: "2001-05-24",
						YTDReturn:               null.FloatFrom(-0.0537),
						ExpenseRatioNet:         null.FloatFrom(0.0003),
						Yield:                   null.FloatFrom(0.0133),
						NAV:                     null.FloatFrom(413.24),
						LastPrice:               null.FloatFrom(413.24),
						TurnoverRate:            null.FloatFrom(0.04),
						NetAssets:               null.IntFrom(753409982464),
						Overview:                "The investment seeks to track the performance of the Standard & Poor's 500...",
					},
					Performance: response.ETFWorldPerformance{
						TrailingReturns: []response.ETFTrailingReturn{
							{
								Period:           "ytd",
								ShareClassReturn: null.FloatFrom(-0.0751),
								CategoryReturn:   null.FloatFrom(0.1484),
							},
						},
						AnnualTotalReturns: []response.ETFAnnualTotalReturn{
							{
								Year:             null.IntFrom(2021),
								ShareClassReturn: null.FloatFrom(0.2866),
								CategoryReturn:   null.FloatFrom(0),
							},
						},
					},
					Risk:        expectedETFWorldRisk(),
					Composition: expectedETFWorldComposition(),
				},
				Status: "ok",
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
				url: wrongAPIKeyURL,
			},
			want:  response.ETFFullData{},
			want1: response.NewCreditsImpl(100, 800),
			wantErr: fmt.Sprintf(
				"HTTP 401 Unauthorized: code: 401, message: **apikey** parameter is incorrect or not specified. "+
					"You can get your free API key instantly following this link: https://twelvedata.com/pricing. "+
					"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error (URL: %s?symbol=VTI)",
				wrongAPIKeyURL,
			),
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

	planRestrictedURL := mockServerWithURL(
		t,
		http.StatusForbidden,
		100,
		200,
		fmt.Sprintf(`{"code":403,"message":"%s","status":"error"}`, etfWorldPerformancePlanRestrictionMessage),
		"/?symbol=VTI",
	)

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
						"etf": {
							"performance": {
								"trailing_returns": [
									{
										"period": "ytd",
										"share_class_return": -0.0751,
										"category_return": 0.1484
									}
								],
								"annual_total_returns": [
									{
										"year": 2021,
										"share_class_return": 0.2866,
										"category_return": 0
									}
								]
							}
						},
						"status": "ok"
					}`,
					"/?symbol=VTI",
				),
			},
			want: response.ETFPerformance{
				ETF: response.ETFPerformanceData{
					Performance: response.ETFWorldPerformance{
						TrailingReturns: []response.ETFTrailingReturn{
							{
								Period:           "ytd",
								ShareClassReturn: null.FloatFrom(-0.0751),
								CategoryReturn:   null.FloatFrom(0.1484),
							},
						},
						AnnualTotalReturns: []response.ETFAnnualTotalReturn{
							{
								Year:             null.IntFrom(2021),
								ShareClassReturn: null.FloatFrom(0.2866),
								CategoryReturn:   null.FloatFrom(0),
							},
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 200),
			wantErr:     "",
			expectedURL: "/?symbol=VTI",
		},
		{
			name: "plan_restricted",
			args: args{
				req: request.GetETFPerformance{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: planRestrictedURL,
			},
			want:        response.ETFPerformance{},
			want1:       response.NewCreditsImpl(100, 200),
			wantErr:     fmt.Sprintf("Plan Limitation: %s", etfWorldPerformancePlanRestrictionMessage),
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

	planRestrictedURL := mockServerWithURL(
		t,
		http.StatusForbidden,
		100,
		100,
		fmt.Sprintf(`{"code":403,"message":"%s","status":"error"}`, etfWorldCompositionPlanRestrictionMessage),
		"/?symbol=VTI",
	)

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
						"etf": {
							"composition": {
								"major_market_sectors": [
									{
										"sector": "Technology",
										"weight": 0.2424
									}
								],
								"country_allocation": [
									{
										"country": "United Kingdom",
										"allocation": 0.9855
									}
								],
								"asset_allocation": {
									"cash": 0.0004,
									"stocks": 0.9996,
									"preferred_stocks": 0,
									"convertibles": 0,
									"bonds": 0,
									"others": 0
								},
								"top_holdings": [
									{
										"symbol": "AAPL",
										"name": "Apple Inc",
										"exchange": "NASDAQ",
										"mic_code": "XNAS",
										"weight": 0.0592
									}
								],
								"bond_breakdown": {
									"average_maturity": {
										"fund": 6.65,
										"category": 7.81
									},
									"average_duration": {
										"fund": 5.72,
										"category": 5.64
									},
									"credit_quality": [
										{
											"grade": "AAA",
											"weight": 0
										}
									]
								}
							}
						},
						"status": "ok"
					}`,
					"/?symbol=VTI",
				),
			},
			want: response.ETFComposition{
				ETF: response.ETFCompositionData{
					Composition: expectedETFWorldComposition(),
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?symbol=VTI",
		},
		{
			name: "plan_restricted",
			args: args{
				req: request.GetETFComposition{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "VTI",
				},
				url: planRestrictedURL,
			},
			want:        response.ETFComposition{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     fmt.Sprintf("Plan Limitation: %s", etfWorldCompositionPlanRestrictionMessage),
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
