package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

const analysisWrongAPIKeyResponse = `{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`

const analysisWrongAPIKeyError = "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
	"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
	"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error"

var analysisMeta = response.AnalysisMeta{
	Symbol:           "AAPL",
	Name:             "Apple Inc",
	Currency:         "USD",
	ExchangeTimezone: "America/New_York",
	Exchange:         "NASDAQ",
	MicCode:          "XNGS",
	Type:             "Common Stock",
}

//nolint:dupl
func Test_client_GetEarningsEstimate(t *testing.T) {
	type args struct {
		req request.GetEarningsEstimate
		url string
	}

	expectedURL := "/earnings_estimate?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.EarningsEstimate
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEarningsEstimate{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "earnings_estimate": [
					    {
					      "date": "2022-09-30",
					      "period": "current_quarter",
					      "number_of_analysts": 27,
					      "avg_estimate": 1.26,
					      "low_estimate": 1.13,
					      "high_estimate": 1.35,
					      "year_ago_eps": 1.24
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.EarningsEstimate{
				Meta: analysisMeta,
				EarningsEstimate: []response.EarningsEstimateEntry{
					{
						Date:             "2022-09-30",
						Period:           "current_quarter",
						NumberOfAnalysts: null.IntFrom(27),
						AvgEstimate:      null.FloatFrom(1.26),
						LowEstimate:      null.FloatFrom(1.13),
						HighEstimate:     null.FloatFrom(1.35),
						YearAgoEPS:       null.FloatFrom(1.24),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEarningsEstimate{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.EarningsEstimate{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getEarningsEstimate: NewEndpoint[request.GetEarningsEstimate, response.EarningsEstimate, response.Credits, error](httpCli, url+"/earnings_estimate"),
					}
				},
				func(cli interface{}, req request.GetEarningsEstimate) (response.EarningsEstimate, response.Credits, error) {
					return cli.(client).GetEarningsEstimate(req)
				},
				"GetEarningsEstimate",
			)
		})
	}
}

func Test_client_GetRevenueEstimate(t *testing.T) {
	type args struct {
		req request.GetRevenueEstimate
		url string
	}

	expectedURL := "/revenue_estimate?country=US&cusip=594918104&dp=2&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.RevenueEstimate
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetRevenueEstimate{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					Figi:          "BBG01293F5X4",
					Isin:          "US0378331005",
					Cusip:         "594918104",
					Exchange:      "NASDAQ",
					Country:       "US",
					DecimalPlaces: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "revenue_estimate": [
					    {
					      "date": "2022-09-30",
					      "period": "current_quarter",
					      "number_of_analysts": 24,
					      "avg_estimate": 88631500000,
					      "low_estimate": 85144300000,
					      "high_estimate": 92794900000,
					      "year_ago_sales": 83360000000,
					      "sales_growth": 0.06
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.RevenueEstimate{
				Meta: analysisMeta,
				RevenueEstimate: []response.RevenueEstimateEntry{
					{
						Date:             "2022-09-30",
						Period:           "current_quarter",
						NumberOfAnalysts: null.IntFrom(24),
						AvgEstimate:      null.FloatFrom(88631500000),
						LowEstimate:      null.FloatFrom(85144300000),
						HighEstimate:     null.FloatFrom(92794900000),
						YearAgoSales:     null.FloatFrom(83360000000),
						SalesGrowth:      null.FloatFrom(0.06),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetRevenueEstimate{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					Figi:          "BBG01293F5X4",
					Isin:          "US0378331005",
					Cusip:         "594918104",
					Exchange:      "NASDAQ",
					Country:       "US",
					DecimalPlaces: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.RevenueEstimate{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getRevenueEstimate: NewEndpoint[request.GetRevenueEstimate, response.RevenueEstimate, response.Credits, error](httpCli, url+"/revenue_estimate"),
					}
				},
				func(cli interface{}, req request.GetRevenueEstimate) (response.RevenueEstimate, response.Credits, error) {
					return cli.(client).GetRevenueEstimate(req)
				},
				"GetRevenueEstimate",
			)
		})
	}
}

//nolint:dupl
func Test_client_GetEPSTrend(t *testing.T) {
	type args struct {
		req request.GetEPSTrend
		url string
	}

	expectedURL := "/eps_trend?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.EPSTrend
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEPSTrend{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "eps_trend": [
					    {
					      "date": "2022-09-30",
					      "period": "current_quarter",
					      "current_estimate": 1.26,
					      "7_days_ago": 1.26,
					      "30_days_ago": 1.31,
					      "60_days_ago": 1.32,
					      "90_days_ago": 1.33
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.EPSTrend{
				Meta: analysisMeta,
				EPSTrend: []response.EPSTrendEntry{
					{
						Date:            "2022-09-30",
						Period:          "current_quarter",
						CurrentEstimate: null.FloatFrom(1.26),
						SevenDaysAgo:    null.FloatFrom(1.26),
						ThirtyDaysAgo:   null.FloatFrom(1.31),
						SixtyDaysAgo:    null.FloatFrom(1.32),
						NinetyDaysAgo:   null.FloatFrom(1.33),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEPSTrend{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.EPSTrend{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getEPSTrend: NewEndpoint[request.GetEPSTrend, response.EPSTrend, response.Credits, error](httpCli, url+"/eps_trend"),
					}
				},
				func(cli interface{}, req request.GetEPSTrend) (response.EPSTrend, response.Credits, error) {
					return cli.(client).GetEPSTrend(req)
				},
				"GetEPSTrend",
			)
		})
	}
}

func Test_client_GetEPSRevisions(t *testing.T) {
	type args struct {
		req request.GetEPSRevisions
		url string
	}

	expectedURL := "/eps_revisions?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.EPSRevisions
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEPSRevisions{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "eps_revision": [
					    {
					      "date": "2022-09-30",
					      "period": "current_quarter",
					      "up_last_week": 1,
					      "up_last_month": 5,
					      "down_last_week": 0,
					      "down_last_month": 0
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.EPSRevisions{
				Meta: analysisMeta,
				EPSRevision: []response.EPSRevisionEntry{
					{
						Date:          "2022-09-30",
						Period:        "current_quarter",
						UpLastWeek:    null.IntFrom(1),
						UpLastMonth:   null.IntFrom(5),
						DownLastWeek:  null.IntFrom(0),
						DownLastMonth: null.IntFrom(0),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEPSRevisions{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.EPSRevisions{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getEPSRevisions: NewEndpoint[request.GetEPSRevisions, response.EPSRevisions, response.Credits, error](httpCli, url+"/eps_revisions"),
					}
				},
				func(cli interface{}, req request.GetEPSRevisions) (response.EPSRevisions, response.Credits, error) {
					return cli.(client).GetEPSRevisions(req)
				},
				"GetEPSRevisions",
			)
		})
	}
}

func Test_client_GetGrowthEstimates(t *testing.T) {
	type args struct {
		req request.GetGrowthEstimates
		url string
	}

	expectedURL := "/growth_estimates?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.GrowthEstimates
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetGrowthEstimates{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "growth_estimates": {
					    "current_quarter": 0.016,
					    "next_quarter": 0.01,
					    "current_year": 0.087,
					    "next_year": 0.055999998,
					    "next_5_years_pa": 0.094799995,
					    "past_5_years_pa": 0.23867
					  },
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.GrowthEstimates{
				Meta: analysisMeta,
				GrowthEstimates: response.GrowthEstimatesData{
					CurrentQuarter: null.FloatFrom(0.016),
					NextQuarter:    null.FloatFrom(0.01),
					CurrentYear:    null.FloatFrom(0.087),
					NextYear:       null.FloatFrom(0.055999998),
					Next5YearsPA:   null.FloatFrom(0.094799995),
					Past5YearsPA:   null.FloatFrom(0.23867),
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetGrowthEstimates{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.GrowthEstimates{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getGrowthEstimates: NewEndpoint[request.GetGrowthEstimates, response.GrowthEstimates, response.Credits, error](httpCli, url+"/growth_estimates"),
					}
				},
				func(cli interface{}, req request.GetGrowthEstimates) (response.GrowthEstimates, response.Credits, error) {
					return cli.(client).GetGrowthEstimates(req)
				},
				"GetGrowthEstimates",
			)
		})
	}
}

func Test_client_GetRecommendations(t *testing.T) {
	type args struct {
		req request.GetRecommendations
		url string
	}

	expectedURL := "/recommendations?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.Recommendations
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetRecommendations{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "trends": {
					    "current_month": {
					      "strong_buy": 13,
					      "buy": 20,
					      "hold": 8,
					      "sell": 0,
					      "strong_sell": 0
					    },
					    "previous_month": {
					      "strong_buy": 13,
					      "buy": 20,
					      "hold": 8,
					      "sell": 0,
					      "strong_sell": 0
					    },
					    "2_months_ago": {
					      "strong_buy": 13,
					      "buy": 20,
					      "hold": 8,
					      "sell": 0,
					      "strong_sell": 0
					    },
					    "3_months_ago": {
					      "strong_buy": 13,
					      "buy": 20,
					      "hold": 8,
					      "sell": 0,
					      "strong_sell": 0
					    }
					  },
					  "rating": 8.2,
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.Recommendations{
				Meta: analysisMeta,
				Trends: response.RecommendationTrends{
					CurrentMonth: response.RecommendationTrend{
						StrongBuy:  null.IntFrom(13),
						Buy:        null.IntFrom(20),
						Hold:       null.IntFrom(8),
						Sell:       null.IntFrom(0),
						StrongSell: null.IntFrom(0),
					},
					PreviousMonth: response.RecommendationTrend{
						StrongBuy:  null.IntFrom(13),
						Buy:        null.IntFrom(20),
						Hold:       null.IntFrom(8),
						Sell:       null.IntFrom(0),
						StrongSell: null.IntFrom(0),
					},
					TwoMonthsAgo: response.RecommendationTrend{
						StrongBuy:  null.IntFrom(13),
						Buy:        null.IntFrom(20),
						Hold:       null.IntFrom(8),
						Sell:       null.IntFrom(0),
						StrongSell: null.IntFrom(0),
					},
					ThreeMonthsAgo: response.RecommendationTrend{
						StrongBuy:  null.IntFrom(13),
						Buy:        null.IntFrom(20),
						Hold:       null.IntFrom(8),
						Sell:       null.IntFrom(0),
						StrongSell: null.IntFrom(0),
					},
				},
				Rating: null.FloatFrom(8.2),
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetRecommendations{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.Recommendations{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getRecommendations: NewEndpoint[request.GetRecommendations, response.Recommendations, response.Credits, error](httpCli, url+"/recommendations"),
					}
				},
				func(cli interface{}, req request.GetRecommendations) (response.Recommendations, response.Credits, error) {
					return cli.(client).GetRecommendations(req)
				},
				"GetRecommendations",
			)
		})
	}
}

func Test_client_GetPriceTarget(t *testing.T) {
	type args struct {
		req request.GetPriceTarget
		url string
	}

	expectedURL := "/price_target?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.PriceTarget
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetPriceTarget{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "price_target": {
					    "high": 220,
					    "median": 185,
					    "low": 136,
					    "average": 184.01,
					    "current": 169.5672,
					    "currency": "USD"
					  },
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.PriceTarget{
				Meta: analysisMeta,
				PriceTarget: response.PriceTargetData{
					High:     null.FloatFrom(220),
					Median:   null.FloatFrom(185),
					Low:      null.FloatFrom(136),
					Average:  null.FloatFrom(184.01),
					Current:  null.FloatFrom(169.5672),
					Currency: "USD",
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetPriceTarget{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.PriceTarget{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getPriceTarget: NewEndpoint[request.GetPriceTarget, response.PriceTarget, response.Credits, error](httpCli, url+"/price_target"),
					}
				},
				func(cli interface{}, req request.GetPriceTarget) (response.PriceTarget, response.Credits, error) {
					return cli.(client).GetPriceTarget(req)
				},
				"GetPriceTarget",
			)
		})
	}
}

func Test_client_GetAnalystRatingsSnapshot(t *testing.T) {
	type args struct {
		req request.GetAnalystRatingsSnapshot
		url string
	}

	expectedURL := "/analyst_ratings/light?country=US&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&outputsize=30&rating_change=Maintains&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.AnalystRatingsSnapshot
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetAnalystRatingsSnapshot{
					APIKey:       request.APIKey{APIKey: ""},
					Symbol:       "AAPL",
					Figi:         "BBG01293F5X4",
					Isin:         "US0378331005",
					Cusip:        "594918104",
					Exchange:     "NASDAQ",
					Country:      "US",
					RatingChange: "Maintains",
					OutputSize:   30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "ratings": [
					    {
					      "date": "2022-08-19",
					      "firm": "Keybanc",
					      "rating_change": "Maintains",
					      "rating_current": "Overweight",
					      "rating_prior": "Overweight"
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.AnalystRatingsSnapshot{
				Meta: analysisMeta,
				Ratings: []response.AnalystRatingsSnapshotEntry{
					{
						Date:          "2022-08-19",
						Firm:          "Keybanc",
						RatingChange:  "Maintains",
						RatingCurrent: "Overweight",
						RatingPrior:   "Overweight",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetAnalystRatingsSnapshot{
					APIKey:       request.APIKey{APIKey: ""},
					Symbol:       "AAPL",
					Figi:         "BBG01293F5X4",
					Isin:         "US0378331005",
					Cusip:        "594918104",
					Exchange:     "NASDAQ",
					Country:      "US",
					RatingChange: "Maintains",
					OutputSize:   30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.AnalystRatingsSnapshot{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getAnalystRatingsSnapshot: NewEndpoint[request.GetAnalystRatingsSnapshot, response.AnalystRatingsSnapshot, response.Credits, error](httpCli, url+"/analyst_ratings/light"),
					}
				},
				func(cli interface{}, req request.GetAnalystRatingsSnapshot) (response.AnalystRatingsSnapshot, response.Credits, error) {
					return cli.(client).GetAnalystRatingsSnapshot(req)
				},
				"GetAnalystRatingsSnapshot",
			)
		})
	}
}

func Test_client_GetAnalystRatingsUSEquities(t *testing.T) {
	type args struct {
		req request.GetAnalystRatingsUSEquities
		url string
	}

	expectedURL := "/analyst_ratings/us_equities?cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&outputsize=30&rating_change=Maintains&symbol=AAPL"

	tests := []struct {
		name        string
		args        args
		want        response.AnalystRatingsUSEquities
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetAnalystRatingsUSEquities{
					APIKey:       request.APIKey{APIKey: ""},
					Symbol:       "AAPL",
					Figi:         "BBG01293F5X4",
					Isin:         "US0378331005",
					Cusip:        "594918104",
					Exchange:     "NASDAQ",
					RatingChange: "Maintains",
					OutputSize:   30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "ratings": [
					    {
					      "date": "2022-08-19",
					      "firm": "Keybanc",
					      "analyst_name": "Brandon Nispel",
					      "rating_change": "Maintains",
					      "rating_current": "Overweight",
					      "rating_prior": "Overweight",
					      "time": "08:29:48",
					      "action_price_target": "Raises",
					      "price_target_current": 185.14,
					      "price_target_prior": 177.01
					    }
					  ],
					  "status": "ok"
					}`,
					expectedURL,
				),
			},
			want: response.AnalystRatingsUSEquities{
				Meta: analysisMeta,
				Ratings: []response.AnalystRatingsUSEquitiesEntry{
					{
						Date:               "2022-08-19",
						Firm:               "Keybanc",
						AnalystName:        "Brandon Nispel",
						RatingChange:       "Maintains",
						RatingCurrent:      "Overweight",
						RatingPrior:        "Overweight",
						Time:               "08:29:48",
						ActionPriceTarget:  "Raises",
						PriceTargetCurrent: null.FloatFrom(185.14),
						PriceTargetPrior:   null.FloatFrom(177.01),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: expectedURL,
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetAnalystRatingsUSEquities{
					APIKey:       request.APIKey{APIKey: ""},
					Symbol:       "AAPL",
					Figi:         "BBG01293F5X4",
					Isin:         "US0378331005",
					Cusip:        "594918104",
					Exchange:     "NASDAQ",
					RatingChange: "Maintains",
					OutputSize:   30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					analysisWrongAPIKeyResponse,
					expectedURL,
				),
			},
			want:        response.AnalystRatingsUSEquities{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     analysisWrongAPIKeyError,
			expectedURL: expectedURL,
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
						getAnalystRatingsUSEquities: NewEndpoint[request.GetAnalystRatingsUSEquities, response.AnalystRatingsUSEquities, response.Credits, error](httpCli, url+"/analyst_ratings/us_equities"),
					}
				},
				func(cli interface{}, req request.GetAnalystRatingsUSEquities) (response.AnalystRatingsUSEquities, response.Credits, error) {
					return cli.(client).GetAnalystRatingsUSEquities(req)
				},
				"GetAnalystRatingsUSEquities",
			)
		})
	}
}
