package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func expectedMutualFundSummaryInfo() response.MutualFundSummaryInfo {
	return response.MutualFundSummaryInfo{
		Symbol:                  "0P0001LCQ3",
		Name:                    "JNL Small Cap Index Fund (I)",
		FundFamily:              "Jackson National",
		FundType:                "Small Blend",
		Currency:                "USD",
		ShareClassInceptionDate: "2021-04-26",
		YTDReturn:               null.FloatFrom(-0.02986),
		ExpenseRatioNet:         null.FloatFrom(0.001),
		Yield:                   null.FloatFrom(0),
		NAV:                     null.FloatFrom(10.09),
		MinInvestment:           null.IntFrom(0),
		TurnoverRate:            null.FloatFrom(0.32),
		NetAssets:               null.IntFrom(2400762112),
		Overview:                "The fund invests, normally, at least 80% of its assets in the stocks...",
		People: []response.MutualFundManager{
			{
				Name:        "John Doe",
				TenureSince: "2018-01-01",
			},
		},
	}
}

func expectedMutualFundPerformanceInfo() response.MutualFundPerformanceInfo {
	return response.MutualFundPerformanceInfo{
		TrailingReturns: []response.MutualFundTrailingReturn{
			{
				Period:           "ytd",
				ShareClassReturn: null.FloatFrom(-0.02986),
				CategoryReturn:   null.FloatFrom(0.2019),
				RankInCategory:   null.IntFrom(76),
			},
		},
		AnnualTotalReturns: []response.MutualFundAnnualTotalReturn{
			{
				Year:             null.IntFrom(2024),
				ShareClassReturn: null.FloatFrom(0.08546),
				CategoryReturn:   null.FloatFrom(0.1119),
			},
		},
		QuarterlyTotalReturns: []response.MutualFundQuarterlyTotalReturn{
			{
				Year: null.IntFrom(2024),
				Q1:   null.FloatFrom(0.02358),
				Q2:   null.FloatFrom(-0.03071),
				Q3:   null.FloatFrom(0.10099),
				Q4:   null.FloatFrom(-0.00629),
			},
		},
		LoadAdjustedReturn: []response.MutualFundLoadAdjustedReturn{
			{
				Period: "1_year",
				Return: null.FloatFrom(0.06139),
			},
		},
	}
}

func expectedMutualFundRiskInfo() response.MutualFundRiskInfo {
	return response.MutualFundRiskInfo{
		VolatilityMeasures: []response.MutualFundVolatilityMeasure{
			{
				Period:                   "3_year",
				Alpha:                    null.FloatFrom(-9.12),
				AlphaCategory:            null.FloatFrom(-0.0939),
				Beta:                     null.FloatFrom(1),
				BetaCategory:             null.FloatFrom(0.0126),
				MeanAnnualReturn:         null.FloatFrom(0.45),
				MeanAnnualReturnCategory: null.FloatFrom(0.0117),
				RSquared:                 null.FloatFrom(69),
				RSquaredCategory:         null.FloatFrom(0.8309),
				Std:                      null.FloatFrom(23.15),
				StdCategory:              null.FloatFrom(0.2554),
				SharpeRatio:              null.FloatFrom(0.04),
				SharpeRatioCategory:      null.FloatFrom(0.005),
				TreynorRatio:             null.FloatFrom(-1.41),
				TreynorRatioCategory:     null.FloatFrom(0.0806),
			},
		},
		ValuationMetrics: response.MutualFundValuationMetrics{
			PriceToEarnings:                    null.FloatFrom(0.05695),
			PriceToEarningsCategory:            null.FloatFrom(20.63),
			PriceToBook:                        null.FloatFrom(0.55626),
			PriceToBookCategory:                null.FloatFrom(2.87),
			PriceToSales:                       null.FloatFrom(0.97803),
			PriceToSalesCategory:               null.FloatFrom(1.34),
			PriceToCashflow:                    null.FloatFrom(0.10564),
			PriceToCashflowCategory:            null.FloatFrom(11.81),
			MedianMarketCapitalization:         null.IntFrom(2965),
			MedianMarketCapitalizationCategory: null.IntFrom(4925),
			ThreeYearEarningsGrowth:            null.FloatFrom(16.32),
			ThreeYearEarningsGrowthCategory:    null.FloatFrom(10.55),
		},
	}
}

func expectedMutualFundRatingsInfo() response.MutualFundRatingsInfo {
	return response.MutualFundRatingsInfo{
		PerformanceRating: null.IntFrom(2),
		RiskRating:        null.IntFrom(4),
		ReturnRating:      null.IntFrom(0),
	}
}

func expectedMutualFundCompositionInfo() response.MutualFundCompositionInfo {
	return response.MutualFundCompositionInfo{
		MajorMarketSectors: []response.MutualFundSectorWeight{
			{
				Sector: "Industrials",
				Weight: null.FloatFrom(0.1742),
			},
		},
		AssetAllocation: response.MutualFundAssetAllocation{
			Cash:            null.FloatFrom(0.0043),
			Stocks:          null.FloatFrom(0.9956),
			PreferredStocks: null.FloatFrom(0),
			Convertibles:    null.FloatFrom(0),
			Bonds:           null.FloatFrom(0),
			Others:          null.FloatFrom(0),
		},
		TopHoldings: []response.MutualFundTopHolding{
			{
				Symbol:   "BBWI",
				Name:     "Bath & Body Works Inc",
				Exchange: "NASDAQ",
				MicCode:  "XNAS",
				Weight:   null.FloatFrom(0.00624),
			},
		},
		BondBreakdown: response.MutualFundBondBreakdown{
			AverageMaturity: response.MutualFundBondMetric{
				Fund:     null.Float{},
				Category: null.FloatFrom(1.97),
			},
			AverageDuration: response.MutualFundBondMetric{
				Fund:     null.Float{},
				Category: null.FloatFrom(1.64),
			},
			CreditQuality: []response.MutualFundCreditQuality{
				{
					Grade:  "U.S. Government",
					Weight: null.FloatFrom(0),
				},
			},
		},
	}
}

func expectedMutualFundPurchaseInfoDetails() response.MutualFundPurchaseInfoDetails {
	return response.MutualFundPurchaseInfoDetails{
		Expenses: response.MutualFundPurchaseExpenses{
			ExpenseRatioGross: null.FloatFrom(0.0022),
			ExpenseRatioNet:   null.FloatFrom(0.001),
		},
		Minimums: response.MutualFundPurchaseMinimums{
			InitialInvestment:       null.IntFrom(0),
			AdditionalInvestment:    null.IntFrom(0),
			InitialIRAInvestment:    null.String{},
			AdditionalIRAInvestment: null.String{},
		},
		Pricing: response.MutualFundPurchasePricing{
			NAV:             null.FloatFrom(10.09),
			TwelveMonthLow:  null.FloatFrom(9.630000114441),
			TwelveMonthHigh: null.FloatFrom(12.10000038147),
			LastMonth:       null.FloatFrom(11.050000190735),
		},
		Brokerages: []string{},
	}
}

func expectedMutualFundSustainabilityDetails() response.MutualFundSustainabilityDetails {
	return response.MutualFundSustainabilityDetails{
		Score: null.IntFrom(22),
		CorporateESGPillars: response.MutualFundESGPillars{
			Environmental: null.FloatFrom(3.73),
			Social:        null.FloatFrom(10.44),
			Governance:    null.FloatFrom(7.86),
		},
		SustainableInvestment: null.BoolFrom(false),
		CorporateAUM:          null.FloatFrom(0.99486),
	}
}

func runMutualFundWorldTest[Req any, Resp any](
	t *testing.T,
	methodName string,
	req Req,
	want Resp,
	responseBody string,
	expectedURL string,
	createEndpoint func(*HTTPCli, string) interface{},
	callEndpoint func(interface{}, Req) (Resp, response.Credits, error),
) {
	t.Helper()

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		1,
		responseBody,
		expectedURL,
	)

	testEndpointCall(
		t,
		"success",
		struct {
			req Req
			url string
		}{
			req: req,
			url: successURL,
		},
		want,
		response.NewCreditsImpl(100, 1),
		"",
		createEndpoint,
		callEndpoint,
		methodName,
	)
}

func Test_client_GetMutualFundSummary(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundSummary",
		request.GetMutualFundSummary{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundSummary{
			MutualFund: response.MutualFundSummaryData{
				Summary: expectedMutualFundSummaryInfo(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"summary": {
					"symbol": "0P0001LCQ3",
					"name": "JNL Small Cap Index Fund (I)",
					"fund_family": "Jackson National",
					"fund_type": "Small Blend",
					"currency": "USD",
					"share_class_inception_date": "2021-04-26",
					"ytd_return": -0.02986,
					"expense_ratio_net": 0.001,
					"yield": 0,
					"nav": 10.09,
					"min_investment": 0,
					"turnover_rate": 0.32,
					"net_assets": 2400762112,
					"overview": "The fund invests, normally, at least 80% of its assets in the stocks...",
					"people": [
						{
							"name": "John Doe",
							"tenure_since": "2018-01-01"
						}
					]
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundSummary: NewEndpoint[request.GetMutualFundSummary, response.MutualFundSummary, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundSummary) (response.MutualFundSummary, response.Credits, error) {
			return cli.(client).GetMutualFundSummary(req)
		},
	)
}

func Test_client_GetMutualFundPerformance(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundPerformance",
		request.GetMutualFundPerformance{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundPerformance{
			MutualFund: response.MutualFundPerformanceData{
				Performance: expectedMutualFundPerformanceInfo(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"performance": {
					"trailing_returns": [
						{
							"period": "ytd",
							"share_class_return": -0.02986,
							"category_return": 0.2019,
							"rank_in_category": 76
						}
					],
					"annual_total_returns": [
						{
							"year": 2024,
							"share_class_return": 0.08546,
							"category_return": 0.1119
						}
					],
					"quarterly_total_returns": [
						{
							"year": 2024,
							"q1": 0.02358,
							"q2": -0.03071,
							"q3": 0.10099,
							"q4": -0.00629
						}
					],
					"load_adjusted_return": [
						{
							"period": "1_year",
							"return": 0.06139
						}
					]
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundPerformance: NewEndpoint[request.GetMutualFundPerformance, response.MutualFundPerformance, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundPerformance) (response.MutualFundPerformance, response.Credits, error) {
			return cli.(client).GetMutualFundPerformance(req)
		},
	)
}

func Test_client_GetMutualFundRisk(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundRisk",
		request.GetMutualFundRisk{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundRisk{
			MutualFund: response.MutualFundRiskData{
				Risk: expectedMutualFundRiskInfo(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"risk": {
					"volatility_measures": [
						{
							"period": "3_year",
							"alpha": -9.12,
							"alpha_category": -0.0939,
							"beta": 1,
							"beta_category": 0.0126,
							"mean_annual_return": 0.45,
							"mean_annual_return_category": 0.0117,
							"r_squared": 69,
							"r_squared_category": 0.8309,
							"std": 23.15,
							"std_category": 0.2554,
							"sharpe_ratio": 0.04,
							"sharpe_ratio_category": 0.005,
							"treynor_ratio": -1.41,
							"treynor_ratio_category": 0.0806
						}
					],
					"valuation_metrics": {
						"price_to_earnings": 0.05695,
						"price_to_earnings_category": 20.63,
						"price_to_book": 0.55626,
						"price_to_book_category": 2.87,
						"price_to_sales": 0.97803,
						"price_to_sales_category": 1.34,
						"price_to_cashflow": 0.10564,
						"price_to_cashflow_category": 11.81,
						"median_market_capitalization": 2965,
						"median_market_capitalization_category": 4925,
						"3_year_earnings_growth": 16.32,
						"3_year_earnings_growths_category": 10.55
					}
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundRisk: NewEndpoint[request.GetMutualFundRisk, response.MutualFundRisk, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundRisk) (response.MutualFundRisk, response.Credits, error) {
			return cli.(client).GetMutualFundRisk(req)
		},
	)
}

func Test_client_GetMutualFundRatings(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundRatings",
		request.GetMutualFundRatings{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundRatings{
			MutualFund: response.MutualFundRatingsData{
				Ratings: expectedMutualFundRatingsInfo(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"ratings": {
					"performance_rating": 2,
					"risk_rating": 4,
					"return_rating": 0
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundRatings: NewEndpoint[request.GetMutualFundRatings, response.MutualFundRatings, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundRatings) (response.MutualFundRatings, response.Credits, error) {
			return cli.(client).GetMutualFundRatings(req)
		},
	)
}

func Test_client_GetMutualFundComposition(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundComposition",
		request.GetMutualFundComposition{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundComposition{
			MutualFund: response.MutualFundCompositionData{
				Composition: expectedMutualFundCompositionInfo(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"composition": {
					"major_market_sectors": [
						{
							"sector": "Industrials",
							"weight": 0.1742
						}
					],
					"asset_allocation": {
						"cash": 0.0043,
						"stocks": 0.9956,
						"preferred_stocks": 0,
						"convertibles": 0,
						"bonds": 0,
						"others": 0
					},
					"top_holdings": [
						{
							"symbol": "BBWI",
							"name": "Bath & Body Works Inc",
							"exchange": "NASDAQ",
							"mic_code": "XNAS",
							"weight": 0.00624
						}
					],
					"bond_breakdown": {
						"average_maturity": {
							"fund": null,
							"category": 1.97
						},
						"average_duration": {
							"fund": null,
							"category": 1.64
						},
						"credit_quality": [
							{
								"grade": "U.S. Government",
								"weight": 0
							}
						]
					}
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundComposition: NewEndpoint[request.GetMutualFundComposition, response.MutualFundComposition, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundComposition) (response.MutualFundComposition, response.Credits, error) {
			return cli.(client).GetMutualFundComposition(req)
		},
	)
}

func Test_client_GetMutualFundPurchaseInfo(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundPurchaseInfo",
		request.GetMutualFundPurchaseInfo{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundPurchaseInfo{
			MutualFund: response.MutualFundPurchaseInfoData{
				PurchaseInfo: expectedMutualFundPurchaseInfoDetails(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"purchase_info": {
					"expenses": {
						"expense_ratio_gross": 0.0022,
						"expense_ratio_net": 0.001
					},
					"minimums": {
						"initial_investment": 0,
						"additional_investment": 0,
						"initial_ira_investment": null,
						"additional_ira_investment": null
					},
					"pricing": {
						"nav": 10.09,
						"12_month_low": 9.630000114441,
						"12_month_high": 12.10000038147,
						"last_month": 11.050000190735
					},
					"brokerages": []
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundPurchaseInfo: NewEndpoint[request.GetMutualFundPurchaseInfo, response.MutualFundPurchaseInfo, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundPurchaseInfo) (response.MutualFundPurchaseInfo, response.Credits, error) {
			return cli.(client).GetMutualFundPurchaseInfo(req)
		},
	)
}

func Test_client_GetMutualFundSustainability(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundSustainability",
		request.GetMutualFundSustainability{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			DecimalPlaces: 5,
		},
		response.MutualFundSustainability{
			MutualFund: response.MutualFundSustainabilityData{
				Sustainability: expectedMutualFundSustainabilityDetails(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"sustainability": {
					"score": 22,
					"corporate_esg_pillars": {
						"environmental": 3.73,
						"social": 10.44,
						"governance": 7.86
					},
					"sustainable_investment": false,
					"corporate_aum": 0.99486
				}
			},
			"status": "ok"
		}`,
		"/?dp=5&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundSustainability: NewEndpoint[request.GetMutualFundSustainability, response.MutualFundSustainability, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundSustainability) (response.MutualFundSustainability, response.Credits, error) {
			return cli.(client).GetMutualFundSustainability(req)
		},
	)
}

func Test_client_GetMutualFundFullData(t *testing.T) {
	runMutualFundWorldTest(
		t,
		"GetMutualFundFullData",
		request.GetMutualFundFullData{
			APIKey:        request.APIKey{APIKey: ""},
			Symbol:        "0P0001LCQ3",
			FIGI:          "BBG00HMMLCH1",
			ISIN:          "LU1206782309",
			CUSIP:         "120678230",
			Country:       "United States",
			DecimalPlaces: 5,
		},
		response.MutualFundFullData{
			MutualFund: response.MutualFundFullDataData{
				Summary:        expectedMutualFundSummaryInfo(),
				Performance:    expectedMutualFundPerformanceInfo(),
				Risk:           expectedMutualFundRiskInfo(),
				Ratings:        expectedMutualFundRatingsInfo(),
				Composition:    expectedMutualFundCompositionInfo(),
				PurchaseInfo:   expectedMutualFundPurchaseInfoDetails(),
				Sustainability: expectedMutualFundSustainabilityDetails(),
			},
			Status: "ok",
		},
		`{
			"mutual_fund": {
				"summary": {
					"symbol": "0P0001LCQ3",
					"name": "JNL Small Cap Index Fund (I)",
					"fund_family": "Jackson National",
					"fund_type": "Small Blend",
					"currency": "USD",
					"share_class_inception_date": "2021-04-26",
					"ytd_return": -0.02986,
					"expense_ratio_net": 0.001,
					"yield": 0,
					"nav": 10.09,
					"min_investment": 0,
					"turnover_rate": 0.32,
					"net_assets": 2400762112,
					"overview": "The fund invests, normally, at least 80% of its assets in the stocks...",
					"people": [
						{
							"name": "John Doe",
							"tenure_since": "2018-01-01"
						}
					]
				},
				"performance": {
					"trailing_returns": [
						{
							"period": "ytd",
							"share_class_return": -0.02986,
							"category_return": 0.2019,
							"rank_in_category": 76
						}
					],
					"annual_total_returns": [
						{
							"year": 2024,
							"share_class_return": 0.08546,
							"category_return": 0.1119
						}
					],
					"quarterly_total_returns": [
						{
							"year": 2024,
							"q1": 0.02358,
							"q2": -0.03071,
							"q3": 0.10099,
							"q4": -0.00629
						}
					],
					"load_adjusted_return": [
						{
							"period": "1_year",
							"return": 0.06139
						}
					]
				},
				"risk": {
					"volatility_measures": [
						{
							"period": "3_year",
							"alpha": -9.12,
							"alpha_category": -0.0939,
							"beta": 1,
							"beta_category": 0.0126,
							"mean_annual_return": 0.45,
							"mean_annual_return_category": 0.0117,
							"r_squared": 69,
							"r_squared_category": 0.8309,
							"std": 23.15,
							"std_category": 0.2554,
							"sharpe_ratio": 0.04,
							"sharpe_ratio_category": 0.005,
							"treynor_ratio": -1.41,
							"treynor_ratio_category": 0.0806
						}
					],
					"valuation_metrics": {
						"price_to_earnings": 0.05695,
						"price_to_earnings_category": 20.63,
						"price_to_book": 0.55626,
						"price_to_book_category": 2.87,
						"price_to_sales": 0.97803,
						"price_to_sales_category": 1.34,
						"price_to_cashflow": 0.10564,
						"price_to_cashflow_category": 11.81,
						"median_market_capitalization": 2965,
						"median_market_capitalization_category": 4925,
						"3_year_earnings_growth": 16.32,
						"3_year_earnings_growths_category": 10.55
					}
				},
				"ratings": {
					"performance_rating": 2,
					"risk_rating": 4,
					"return_rating": 0
				},
				"composition": {
					"major_market_sectors": [
						{
							"sector": "Industrials",
							"weight": 0.1742
						}
					],
					"asset_allocation": {
						"cash": 0.0043,
						"stocks": 0.9956,
						"preferred_stocks": 0,
						"convertibles": 0,
						"bonds": 0,
						"others": 0
					},
					"top_holdings": [
						{
							"symbol": "BBWI",
							"name": "Bath & Body Works Inc",
							"exchange": "NASDAQ",
							"mic_code": "XNAS",
							"weight": 0.00624
						}
					],
					"bond_breakdown": {
						"average_maturity": {
							"fund": null,
							"category": 1.97
						},
						"average_duration": {
							"fund": null,
							"category": 1.64
						},
						"credit_quality": [
							{
								"grade": "U.S. Government",
								"weight": 0
							}
						]
					}
				},
				"purchase_info": {
					"expenses": {
						"expense_ratio_gross": 0.0022,
						"expense_ratio_net": 0.001
					},
					"minimums": {
						"initial_investment": 0,
						"additional_investment": 0,
						"initial_ira_investment": null,
						"additional_ira_investment": null
					},
					"pricing": {
						"nav": 10.09,
						"12_month_low": 9.630000114441,
						"12_month_high": 12.10000038147,
						"last_month": 11.050000190735
					},
					"brokerages": []
				},
				"sustainability": {
					"score": 22,
					"corporate_esg_pillars": {
						"environmental": 3.73,
						"social": 10.44,
						"governance": 7.86
					},
					"sustainable_investment": false,
					"corporate_aum": 0.99486
				}
			},
			"status": "ok"
		}`,
		"/?country=United+States&cusip=120678230&dp=5&figi=BBG00HMMLCH1&isin=LU1206782309&symbol=0P0001LCQ3",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMutualFundFullData: NewEndpoint[request.GetMutualFundFullData, response.MutualFundFullData, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMutualFundFullData) (response.MutualFundFullData, response.Credits, error) {
			return cli.(client).GetMutualFundFullData(req)
		},
	)
}
