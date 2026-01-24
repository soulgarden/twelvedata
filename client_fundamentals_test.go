package twelvedata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetBalanceSheet(t *testing.T) {
	type args struct {
		req request.GetBalanceSheet
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.BalanceSheets
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetBalanceSheet{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
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
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "exchange_timezone": "America/New_York",
					    "period": "Quarterly"
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
					          "restricted_cash": null,
					          "assets_held_for_sale": null,
					          "hedging_assets": null,
					          "other_current_assets": 14111000000,
					          "total_current_assets": 134836000000
					        },
					        "non_current_assets": {
					          "properties": 0,
					          "land_and_improvements": 20041000000,
					          "machinery_furniture_equipment": 78659000000,
					          "construction_in_progress": null,
					          "leases": 11023000000,
					          "accumulated_depreciation": -70283000000,
					          "goodwill": null,
					          "investment_properties": null,
					          "financial_assets": null,
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
					          "tax_payable": null,
					          "pensions": null,
					          "other_current_liabilities": 47493000000,
					          "total_current_liabilities": 125481000000
					        },
					        "non_current_liabilities": {
					          "long_term_provisions": null,
					          "long_term_debt": 109106000000,
					          "provision_for_risks_and_charges": 24689000000,
					          "deferred_liabilities": null,
					          "derivative_product_liabilities": null,
					          "other_non_current_liabilities": 28636000000,
					          "total_non_current_liabilities": 162431000000
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
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.BalanceSheets{
				Meta: response.BalanceSheetsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Quarterly",
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
								PrepaidAssets:             null.Int{},
								RestrictedCash:            null.Int{},
								AssetsHeldForSale:         null.Int{},
								HedgingAssets:             null.Int{},
								OtherCurrentAssets:        null.IntFrom(14111000000),
								TotalCurrentAssets:        null.IntFrom(134836000000),
							},
							NonCurrentAssets: response.BalanceSheetNonCurrentAssets{
								Properties:                  null.IntFrom(0),
								LandAndImprovements:         null.IntFrom(20041000000),
								MachineryFurnitureEquipment: null.IntFrom(78659000000),
								ConstructionInProgress:      null.Int{},
								Leases:                      null.IntFrom(11023000000),
								AccumulatedDepreciation:     null.IntFrom(-70283000000),
								Goodwill:                    null.Int{},
								InvestmentProperties:        null.Int{},
								FinancialAssets:             null.Int{},
								IntangibleAssets:            null.Int{},
								InvestmentsAndAdvances:      null.IntFrom(127877000000),
								OtherNonCurrentAssets:       null.IntFrom(48849000000),
								TotalNonCurrentAssets:       null.IntFrom(216166000000),
							},
							TotalAssets: null.IntFrom(351002000000),
						},
						Liabilities: response.BalanceSheetLiabilities{
							CurrentLiabilities: response.BalanceSheetCurrentLiabilities{
								AccountsPayable:         null.IntFrom(54763000000),
								AccruedExpenses:         null.Int{},
								ShortTermDebt:           null.IntFrom(15613000000),
								DeferredRevenue:         null.IntFrom(7612000000),
								TaxPayable:              null.Int{},
								Pensions:                null.Int{},
								OtherCurrentLiabilities: null.IntFrom(47493000000),
								TotalCurrentLiabilities: null.IntFrom(125481000000),
							},
							NonCurrentLiabilities: response.BalanceSheetNonCurrentLiabilities{
								LongTermProvisions:           null.Int{},
								LongTermDebt:                 null.IntFrom(109106000000),
								ProvisionForRisksAndCharges:  null.IntFrom(24689000000),
								DeferredLiabilities:          null.Int{},
								DerivativeProductLiabilities: null.Int{},
								OtherNonCurrentLiabilities:   null.IntFrom(28636000000),
								TotalNonCurrentLiabilities:   null.IntFrom(162431000000),
							},
							TotalLiabilities: null.IntFrom(287912000000),
						},
						ShareholdersEquity: response.BalanceSheetShareholdersEquity{
							CommonStock:             null.IntFrom(57365000000),
							RetainedEarnings:        null.IntFrom(5562000000),
							OtherShareholdersEquity: null.IntFrom(163000000),
							TotalShareholdersEquity: null.IntFrom(63090000000),
							AdditionalPaidInCapital: null.Int{},
							TreasuryStock:           null.Int{},
							MinorityInterest:        null.Int{},
						},
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetBalanceSheet{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want:  response.BalanceSheets{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
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
						getBalanceSheet: NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error) {
					return cli.(client).GetBalanceSheet(req)
				},
				"GetBalanceSheet",
			)
		})
	}
}

func Test_client_GetCashFlow(t *testing.T) {
	type args struct {
		req request.GetCashFlow
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.CashFlows
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCashFlow{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
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
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "exchange_timezone": "America/New_York",
					    "period": "Quarterly"
					  },
					  "cash_flow": [
					    {
					      "fiscal_date": "2021-12-31",
					      "quarter": 1,
					      "operating_activities": {
					        "net_income": 34630000000,
					        "depreciation": 2697000000,
					        "deferred_taxes": 682000000,
					        "stock_based_compensation": 2265000000,
					        "other_non_cash_items": 167000000,
					        "accounts_receivable": -13746000000,
					        "accounts_payable": 19813000000,
					        "other_assets_liabilities": 458000000,
					        "operating_cash_flow": 46966000000
					      },
					      "investing_activities": {
					        "capital_expenditures": -2803000000,
					        "net_intangibles": null,
					        "net_acquisitions": null,
					        "purchase_of_investments": -34913000000,
					        "sale_of_investments": 21984000000,
					        "other_investing_activity": -374000000,
					        "investing_cash_flow": -16106000000
					      },
					      "financing_activities": {
					        "long_term_debt_issuance": null,
					        "long_term_debt_payments": 0,
					        "short_term_debt_issuance": -1000000000,
					        "common_stock_issuance": null,
					        "common_stock_repurchase": -20478000000,
					        "common_dividends": -3732000000,
					        "other_financing_charges": -2949000000,
					        "financing_cash_flow": -28159000000
					      },
					      "end_cash_position": 38630000000,
					      "income_tax_paid": 5235000000,
					      "interest_paid": 531000000,
					      "free_cash_flow": 49769000000
					    }
					  ]
					}`,
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.CashFlows{
				Meta: response.CashFlowsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Quarterly",
				},
				CashFlow: []response.CashFlow{
					{
						FiscalDate: "2021-12-31",
						Quarter:    null.IntFrom(1),
						OperatingActivities: response.CashFlowOperatingActivities{
							NetIncome:              null.IntFrom(34630000000),
							Depreciation:           null.IntFrom(2697000000),
							DeferredTaxes:          null.IntFrom(682000000),
							StockBasedCompensation: null.IntFrom(2265000000),
							OtherNonCashItems:      null.IntFrom(167000000),
							AccountsReceivable:     null.IntFrom(-13746000000),
							AccountsPayable:        null.IntFrom(19813000000),
							OtherAssetsLiabilities: null.IntFrom(458000000),
							OperatingCashFlow:      null.IntFrom(46966000000),
						},
						InvestingActivities: response.CashFlowInvestingActivities{
							CapitalExpenditures:    null.IntFrom(-2803000000),
							NetIntangibles:         null.Int{},
							NetAcquisitions:        null.Int{},
							PurchaseOfInvestments:  null.IntFrom(-34913000000),
							SaleOfInvestments:      null.IntFrom(21984000000),
							OtherInvestingActivity: null.IntFrom(-374000000),
							InvestingCashFlow:      null.IntFrom(-16106000000),
						},
						FinancingActivities: response.CashFlowFinancingActivities{
							LongTermDebtIssuance:  null.Int{},
							LongTermDebtPayments:  null.IntFrom(0),
							ShortTermDebtIssuance: null.IntFrom(-1000000000),
							CommonStockIssuance:   null.Int{},
							CommonStockRepurchase: null.IntFrom(-20478000000),
							CommonDividends:       null.IntFrom(-3732000000),
							OtherFinancingCharges: null.IntFrom(-2949000000),
							FinancingCashFlow:     null.IntFrom(-28159000000),
						},
						EndCashPosition: null.IntFrom(38630000000),
						IncomeTaxPaid:   null.IntFrom(5235000000),
						InterestPaid:    null.IntFrom(531000000),
						FreeCashFlow:    null.IntFrom(49769000000),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetCashFlow{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want:  response.CashFlows{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
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
						getCashFlow: NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetCashFlow) (response.CashFlows, response.Credits, error) {
					return cli.(client).GetCashFlow(req)
				},
				"GetCashFlow",
			)
		})
	}
}

func Test_client_GetDividends(t *testing.T) {
	type args struct {
		req request.GetDividends
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Dividends
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetDividends{
					APIKey:    request.APIKey{APIKey: ""},
					Symbol:    "AAPL",
					Figi:      "BBG000B9XRY4",
					Isin:      "US0378331005",
					Cusip:     "037833100",
					Exchange:  "NASDAQ",
					MicCode:   "XNAS",
					Country:   "US",
					Range:     "1y",
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
					Adjust:    true,
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
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "exchange_timezone": "America/New_York"
					  },
					  "dividends": [
					    {
					      "ex_date": "2021-08-06",
					      "amount": 0.22
					    }
					  ]
					}`,
					"/?adjust=true&country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&range=1y&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.Dividends{
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
						ExDate: "2021-08-06",
						Amount: null.FloatFrom(0.22),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?adjust=true&country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&range=1y&start_date=2024-01-01&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetDividends{
					APIKey:    request.APIKey{APIKey: ""},
					Symbol:    "AAPL",
					Figi:      "BBG000B9XRY4",
					Isin:      "US0378331005",
					Cusip:     "037833100",
					Exchange:  "NASDAQ",
					MicCode:   "XNAS",
					Country:   "US",
					Range:     "1y",
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
					Adjust:    true,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?adjust=true&country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&range=1y&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want:        response.Dividends{},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?adjust=true&country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&range=1y&start_date=2024-01-01&symbol=AAPL",
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
						getDividends: NewEndpoint[request.GetDividends, response.Dividends, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetDividends) (response.Dividends, response.Credits, error) {
					return cli.(client).GetDividends(req)
				},
				"GetDividends",
			)
		})
	}
}

func Test_client_GetSplits(t *testing.T) {
	type args struct {
		req request.GetSplits
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        response.Splits
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with demo API response",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "AAPL",
						"name": "Apple Inc.",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "splits": [
						{
						  "date": "2020-08-31",
						  "description": "4-for-1 split",
						  "ratio": 0.25,
						  "from_factor": 4,
						  "to_factor": 1
						}
					  ]
					}`,
					"/?symbol=AAPL",
				),
			},
			want: response.Splits{
				Meta: response.SplitsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				Splits: []response.SplitEvent{
					{
						Date:        "2020-08-31",
						Description: "4-for-1 split",
						Ratio:       null.FloatFrom(0.25),
						FromFactor:  null.IntFrom(4),
						ToFactor:    null.IntFrom(1),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?symbol=AAPL",
		},
		{
			name: "success with range parameter",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "AAPL",
					Range:  "full",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "AAPL",
						"name": "Apple Inc.",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "splits": [
						{
						  "date": "2020-08-31",
						  "description": "4-for-1 split",
						  "ratio": 0.25,
						  "from_factor": 4,
						  "to_factor": 1
						},
						{
						  "date": "2014-06-09",
						  "description": "7-for-1 split",
						  "ratio": 0.14286,
						  "from_factor": 7,
						  "to_factor": 1
						}
					  ]
					}`,
					"/?range=full&symbol=AAPL",
				),
			},
			want: response.Splits{
				Meta: response.SplitsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				Splits: []response.SplitEvent{
					{
						Date:        "2020-08-31",
						Description: "4-for-1 split",
						Ratio:       null.FloatFrom(0.25),
						FromFactor:  null.IntFrom(4),
						ToFactor:    null.IntFrom(1),
					},
					{
						Date:        "2014-06-09",
						Description: "7-for-1 split",
						Ratio:       null.FloatFrom(0.14286),
						FromFactor:  null.IntFrom(7),
						ToFactor:    null.IntFrom(1),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?range=full&symbol=AAPL",
		},
		{
			name: "success with date range",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:    "MSFT",
					Exchange:  "NASDAQ",
					StartDate: "2003-01-01",
					EndDate:   "2003-12-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "MSFT",
						"name": "Microsoft Corporation",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "splits": [
						{
						  "date": "2003-02-18",
						  "description": "2-for-1 split",
						  "ratio": 0.5,
						  "from_factor": 2,
						  "to_factor": 1
						},
						{
						  "date": "1999-03-29",
						  "description": "2-for-1 split",
						  "ratio": 0.5,
						  "from_factor": 2,
						  "to_factor": 1
						}
					  ]
					}`,
					"/?end_date=2003-12-31&exchange=NASDAQ&start_date=2003-01-01&symbol=MSFT",
				),
			},
			want: response.Splits{
				Meta: response.SplitsMeta{
					Symbol:           "MSFT",
					Name:             "Microsoft Corporation",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				Splits: []response.SplitEvent{
					{
						Date:        "2003-02-18",
						Description: "2-for-1 split",
						Ratio:       null.FloatFrom(0.5),
						FromFactor:  null.IntFrom(2),
						ToFactor:    null.IntFrom(1),
					},
					{
						Date:        "1999-03-29",
						Description: "2-for-1 split",
						Ratio:       null.FloatFrom(0.5),
						FromFactor:  null.IntFrom(2),
						ToFactor:    null.IntFrom(1),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?end_date=2003-12-31&exchange=NASDAQ&start_date=2003-01-01&symbol=MSFT",
		},
		{
			name: "no splits found",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "NEWCO",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "NEWCO",
						"name": "New Company Inc.",
						"currency": "USD",
						"exchange": "NYSE",
						"mic_code": "XNYS",
						"exchange_timezone": "America/New_York"
					  },
					  "splits": []
					}`,
					"/?symbol=NEWCO",
				),
			},
			want: response.Splits{
				Meta: response.SplitsMeta{
					Symbol:           "NEWCO",
					Name:             "New Company Inc.",
					Currency:         "USD",
					Exchange:         "NYSE",
					MicCode:          "XNYS",
					ExchangeTimezone: "America/New_York",
				},
				Splits: []response.SplitEvent{},
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?symbol=NEWCO",
		},
		{
			name: "symbol not found",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "INVALID",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":400,"message":"**INVALID** not found: symbol may be delisted","status":"error"}`,
					"/?symbol=INVALID",
				),
			},
			want:        response.Splits{},
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Symbol Not Found: INVALID - **INVALID** not found: symbol may be delisted",
			expectedURL: "/?symbol=INVALID",
		},
		{
			name: "plan limitation error",
			args: args{
				req: request.GetSplits{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":403,"message":"/splits is available exclusively with grow or pro or ultra or enterprise plans","status":"error"}`,
					"/?symbol=AAPL",
				),
			},
			want:        response.Splits{},
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Plan Limitation: /splits is available exclusively with grow or pro or ultra or enterprise plans",
			expectedURL: "/?symbol=AAPL",
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
						getSplits: NewEndpoint[request.GetSplits, response.Splits, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetSplits) (response.Splits, response.Credits, error) {
					return cli.(client).GetSplits(req)
				},
				"GetSplits",
			)
		})
	}
}

func Test_client_GetSplitsCalendar(t *testing.T) {
	type args struct {
		req request.GetSplitsCalendar
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.SplitsCalendar
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetSplitsCalendar{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 1,
					Page:       1,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`[
					  {
					    "date": "2024-01-15",
					    "symbol": "AAPL",
					    "mic_code": "XNAS",
					    "exchange": "NASDAQ",
					    "description": "2-for-1 split",
					    "ratio": 2,
					    "from_factor": 1,
					    "to_factor": 2
					  }
					]`,
					"/?end_date=2024-01-31&outputsize=1&page=1&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.SplitsCalendar{
				{
					Date:        "2024-01-15",
					Symbol:      "AAPL",
					MicCode:     "XNAS",
					Exchange:    "NASDAQ",
					Description: "2-for-1 split",
					Ratio:       null.FloatFrom(2),
					FromFactor:  null.IntFrom(1),
					ToFactor:    null.IntFrom(2),
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&outputsize=1&page=1&start_date=2024-01-01&symbol=AAPL",
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
						getSplitsCalendar: NewEndpoint[request.GetSplitsCalendar, response.SplitsCalendar, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetSplitsCalendar) (response.SplitsCalendar, response.Credits, error) {
					return cli.(client).GetSplitsCalendar(req)
				},
				"GetSplitsCalendar",
			)
		})
	}
}

func Test_client_GetDividendsCalendar(t *testing.T) {
	type args struct {
		req request.GetDividendsCalendar
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        response.DividendsCalendar
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetDividendsCalendar{
					APIKey: request.APIKey{
						APIKey: "",
					},
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`[
						{
						  "symbol": "AAPL",
						  "mic_code": "XNAS",
						  "exchange": "NASDAQ",
						  "ex_date": "2024-01-15",
						  "amount": 0.24
						},
						{
						  "symbol": "MSFT",
						  "mic_code": "XNGS",
						  "exchange": "NASDAQ",
						  "ex_date": "2024-01-20",
						  "amount": 0.75
						}
					]`,
					"/?end_date=2024-01-31&start_date=2024-01-01",
				),
			},
			want: response.DividendsCalendar{
				{
					Symbol:   "AAPL",
					MicCode:  "XNAS",
					Exchange: "NASDAQ",
					ExDate:   "2024-01-15",
					Amount:   null.FloatFrom(0.24),
				},
				{
					Symbol:   "MSFT",
					MicCode:  "XNGS",
					Exchange: "NASDAQ",
					ExDate:   "2024-01-20",
					Amount:   null.FloatFrom(0.75),
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "plan limitation error",
			args: args{
				req: request.GetDividendsCalendar{
					APIKey: request.APIKey{
						APIKey: "",
					},
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":403,"message":"/dividends_calendar is available exclusively with grow or pro or ultra or enterprise plans. Consider upgrading your API Key now at https://twelvedata.com/pricing","status":"error"}`,
					"/?end_date=2024-01-31&start_date=2024-01-01",
				),
			},
			want:  response.DividendsCalendar(nil),
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "Plan Limitation: /dividends_calendar is available exclusively with grow or pro or ultra or enterprise plans. " +
				"Consider upgrading your API Key now at https://twelvedata.com/pricing",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "demo API key limitation",
			args: args{
				req: request.GetDividendsCalendar{
					APIKey: request.APIKey{
						APIKey: "demo",
					},
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":403,"message":"The 'demo' API key is only used for initial familiarity. To become a full user, you can request your own API key at https://twelvedata.com/pricing. It is absolutely free, and it's yours for a lifetime. It only takes 10 seconds to obtain your own API key!","status":"error"}`,
					"/?apikey=demo&end_date=2024-01-31&start_date=2024-01-01",
				),
			},
			want:  response.DividendsCalendar(nil),
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "Plan Limitation: The 'demo' API key is only used for initial familiarity. To become a full user, " +
				"you can request your own API key at https://twelvedata.com/pricing. It is absolutely free, and it's yours for a lifetime. " +
				"It only takes 10 seconds to obtain your own API key!",
			expectedURL: "/?apikey=demo&end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetDividendsCalendar{
					APIKey: request.APIKey{
						APIKey: "",
					},
					StartDate: "2024-01-01",
					EndDate:   "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?end_date=2024-01-31&start_date=2024-01-01",
				),
			},
			want:  response.DividendsCalendar(nil),
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
						getDividendsCalendar: NewEndpoint[request.GetDividendsCalendar, response.DividendsCalendar, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetDividendsCalendar) (response.DividendsCalendar, response.Credits, error) {
					return cli.(client).GetDividendsCalendar(req)
				},
				"GetDividendsCalendar",
			)
		})
	}
}

func Test_client_GetEarningsCalendar(t *testing.T) {
	type args struct {
		req request.GetEarningsCalendar
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.EarningsCalendar
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEarningsCalendar{
					APIKey:        request.APIKey{APIKey: ""},
					Exchange:      "NASDAQ",
					MicCode:       "XNAS",
					Country:       "US",
					Format:        "json",
					Delimiter:     ",",
					DecimalPlaces: 2,
					StartDate:     "2024-01-01",
					EndDate:       "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "earnings": {
					    "2020-05-08": [
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
					  "status": "ok"
					}`,
					"/?country=US&delimiter=%2C&dp=2&end_date=2024-01-31&exchange=NASDAQ&format=json&mic_code=XNAS&start_date=2024-01-01",
				),
			},
			want: response.EarningsCalendar{
				Earnings: map[string][]*response.EarningsCalendarItem{
					"2020-05-08": {
						{
							Symbol:      "BR",
							Name:        "Broadridge Financial Solutions Inc",
							Currency:    "USD",
							Exchange:    "NYSE",
							MicCode:     "XNYS",
							Country:     "United States",
							Time:        "Time Not Supplied",
							EPSEstimate: null.FloatFrom(1.72),
							EPSActual:   null.FloatFrom(1.67),
							Difference:  null.FloatFrom(-0.05),
							SurprisePrc: null.FloatFrom(-2.9),
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&delimiter=%2C&dp=2&end_date=2024-01-31&exchange=NASDAQ&format=json&mic_code=XNAS&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEarningsCalendar{
					APIKey:        request.APIKey{APIKey: ""},
					Exchange:      "NASDAQ",
					MicCode:       "XNAS",
					Country:       "US",
					Format:        "json",
					Delimiter:     ",",
					DecimalPlaces: 2,
					StartDate:     "2024-01-01",
					EndDate:       "2024-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&delimiter=%2C&dp=2&end_date=2024-01-31&exchange=NASDAQ&format=json&mic_code=XNAS&start_date=2024-01-01",
				),
			},
			want:  response.EarningsCalendar{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&delimiter=%2C&dp=2&end_date=2024-01-31&exchange=NASDAQ&format=json&mic_code=XNAS&start_date=2024-01-01",
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
						getEarningsCalendar: NewEndpoint[request.GetEarningsCalendar, response.EarningsCalendar, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetEarningsCalendar) (response.EarningsCalendar, response.Credits, error) {
					return cli.(client).GetEarningsCalendar(req)
				},
				"GetEarningsCalendar",
			)
		})
	}
}

func Test_client_GetIncomeStatement(t *testing.T) {
	type args struct {
		req request.GetIncomeStatement
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.IncomeStatements
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetIncomeStatement{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
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
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "exchange_timezone": "America/New_York",
					    "period": "Quarterly"
					  },
					  "income_statement": [
					    {
					      "fiscal_date": "2021-12-31",
					      "quarter": 1,
					      "year": 2022,
					      "sales": 123945000000,
					      "cost_of_goods": 69702000000,
					      "gross_profit": 54243000000,
					      "operating_expense": {
					        "research_and_development": 6306000000,
					        "selling_general_and_administrative": 6449000000,
					        "other_operating_expenses": null
					      },
					      "operating_income": 41488000000,
					      "non_operating_interest": {
					        "income": 650000000,
					        "expense": 694000000
					      },
					      "other_income_expense": -203000000,
					      "pretax_income": 41241000000,
					      "income_tax": 6611000000,
					      "net_income": 34630000000,
					      "eps_basic": 2.11,
					      "eps_diluted": 2.1,
					      "basic_shares_outstanding": 16391724000,
					      "diluted_shares_outstanding": 16391724000,
					      "ebitda": 44632000000,
					      "net_income_continuous_operations": null,
					      "minority_interests": null,
					      "preferred_stock_dividends": null
					    }
					  ]
					}`,
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.IncomeStatements{
				Meta: response.IncomeStatementsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
					Period:           "Quarterly",
				},
				IncomeStatement: []response.IncomeStatement{
					{
						FiscalDate:  "2021-12-31",
						Quarter:     null.IntFrom(1),
						Year:        null.IntFrom(2022),
						Sales:       null.IntFrom(123945000000),
						CostOfGoods: null.IntFrom(69702000000),
						GrossProfit: null.IntFrom(54243000000),
						OperatingExpense: response.IncomeStatementOperatingExpense{
							ResearchAndDevelopment:          null.IntFrom(6306000000),
							SellingGeneralAndAdministrative: null.IntFrom(6449000000),
							OtherOperatingExpenses:          null.Int{},
						},
						OperatingIncome: null.IntFrom(41488000000),
						NonOperatingInterest: response.IncomeStatementNonOperatingInterest{
							Income:  null.IntFrom(650000000),
							Expense: null.IntFrom(694000000),
						},
						OtherIncomeExpense:            null.IntFrom(-203000000),
						PretaxIncome:                  null.IntFrom(41241000000),
						IncomeTax:                     null.IntFrom(6611000000),
						NetIncome:                     null.IntFrom(34630000000),
						EPSBasic:                      null.FloatFrom(2.11),
						EPSDiluted:                    null.FloatFrom(2.1),
						BasicSharesOutstanding:        null.IntFrom(16391724000),
						DilutedSharesOutstanding:      null.IntFrom(16391724000),
						EBITDA:                        null.IntFrom(44632000000),
						NetIncomeContinuousOperations: null.Int{},
						MinorityInterests:             null.Int{},
						PreferredStockDividends:       null.Int{},
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetIncomeStatement{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Period:     "quarterly",
					StartDate:  "2024-01-01",
					EndDate:    "2024-01-31",
					OutputSize: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want:  response.IncomeStatements{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&cusip=037833100&end_date=2024-01-31&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL",
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
						getIncomeStatement: NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error) {
					return cli.(client).GetIncomeStatement(req)
				},
				"GetIncomeStatement",
			)
		})
	}
}

func Test_client_GetLogo(t *testing.T) {
	type args struct {
		req request.GetLogo
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        response.Logo
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetLogo{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
						"symbol": "AAPL",
						"exchange": "NASDAQ"
					  },
					  "url": "https://api.twelvedata.com/logo/apple.com"
					}`,
					"/?symbol=AAPL",
				),
			},
			want: response.Logo{
				Meta: response.LogoMeta{
					Symbol:   "AAPL",
					Exchange: "NASDAQ",
				},
				URL: "https://api.twelvedata.com/logo/apple.com",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?symbol=AAPL",
		},
		{
			name: "crypto success with base and quote logos",
			args: args{
				req: request.GetLogo{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "BTC/USD",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
						"symbol": "BTC/USD",
						"exchange": "Coinbase Pro"
					  },
					  "url": "https://api.twelvedata.com/logo/coinbase.com",
					  "logo_base": "https://logo.twelvedata.com/crypto/btc.png",
					  "logo_quote": "https://logo.twelvedata.com/crypto/usd.png"
					}`,
					"/?symbol=BTC%2FUSD",
				),
			},
			want: response.Logo{
				Meta: response.LogoMeta{
					Symbol:   "BTC/USD",
					Exchange: "Coinbase Pro",
				},
				URL:       "https://api.twelvedata.com/logo/coinbase.com",
				LogoBase:  "https://logo.twelvedata.com/crypto/btc.png",
				LogoQuote: "https://logo.twelvedata.com/crypto/usd.png",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?symbol=BTC%2FUSD",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetLogo{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?symbol=AAPL",
				),
			},
			want:  response.Logo{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?symbol=AAPL",
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
						getLogo: NewEndpoint[request.GetLogo, response.Logo, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetLogo) (response.Logo, response.Credits, error) {
					return cli.(client).GetLogo(req)
				},
				"GetLogo",
			)
		})
	}
}

func runConsolidatedStatementTest[Req any, Resp any](
	t *testing.T,
	req Req,
	responseBody string,
	want Resp,
	expectedURL string,
	createClient func(*HTTPCli, string) interface{},
	callMethod func(interface{}, Req) (Resp, response.Credits, error),
	methodName string,
) {
	t.Helper()

	args := struct {
		req Req
		url string
	}{
		req: req,
		url: mockServerWithURL(t, http.StatusOK, 100, 100, responseBody, expectedURL),
	}

	testEndpointCall(
		t,
		methodName,
		args,
		want,
		response.NewCreditsImpl(100, 100),
		"",
		createClient,
		callMethod,
		methodName,
	)
}

func Test_client_GetIncomeStatementConsolidated(t *testing.T) {
	runConsolidatedStatementTest(
		t,
		request.GetIncomeStatement{
			APIKey: request.APIKey{APIKey: ""},
			Symbol: "AAPL",
		},
		`{
		  "meta": {
		    "symbol": "AAPL",
		    "name": "Apple Inc",
		    "currency": "USD",
		    "exchange": "NASDAQ",
		    "mic_code": "XNAS",
		    "exchange_timezone": "America/New_York",
		    "period": "Quarterly"
		  },
		  "income_statement": []
		}`,
		response.IncomeStatements{
			Meta: response.IncomeStatementsMeta{
				Symbol:           "AAPL",
				Name:             "Apple Inc",
				Currency:         "USD",
				Exchange:         "NASDAQ",
				MicCode:          "XNAS",
				ExchangeTimezone: "America/New_York",
				Period:           "Quarterly",
			},
			IncomeStatement: []response.IncomeStatement{},
		},
		"/income_statement/consolidated?symbol=AAPL",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getIncomeStatementConsolidated: NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](httpCli, url+"/income_statement/consolidated"),
			}
		},
		func(cli interface{}, req request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error) {
			return cli.(client).GetIncomeStatementConsolidated(req)
		},
		"GetIncomeStatementConsolidated",
	)
}

func Test_client_GetBalanceSheetConsolidated(t *testing.T) {
	runConsolidatedStatementTest(
		t,
		request.GetBalanceSheet{
			APIKey: request.APIKey{APIKey: ""},
			Symbol: "AAPL",
		},
		`{
		  "meta": {
		    "symbol": "AAPL",
		    "name": "Apple Inc",
		    "currency": "USD",
		    "exchange": "NASDAQ",
		    "mic_code": "XNAS",
		    "exchange_timezone": "America/New_York",
		    "period": "Quarterly"
		  },
		  "balance_sheet": []
		}`,
		response.BalanceSheets{
			Meta: response.BalanceSheetsMeta{
				Symbol:           "AAPL",
				Name:             "Apple Inc",
				Currency:         "USD",
				Exchange:         "NASDAQ",
				MicCode:          "XNAS",
				ExchangeTimezone: "America/New_York",
				Period:           "Quarterly",
			},
			BalanceSheet: []response.BalanceSheet{},
		},
		"/balance_sheet/consolidated?symbol=AAPL",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getBalanceSheetConsolidated: NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](httpCli, url+"/balance_sheet/consolidated"),
			}
		},
		func(cli interface{}, req request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error) {
			return cli.(client).GetBalanceSheetConsolidated(req)
		},
		"GetBalanceSheetConsolidated",
	)
}

func Test_client_GetCashFlowConsolidated(t *testing.T) {
	runConsolidatedStatementTest(
		t,
		request.GetCashFlow{
			APIKey: request.APIKey{APIKey: ""},
			Symbol: "AAPL",
		},
		`{
		  "meta": {
		    "symbol": "AAPL",
		    "name": "Apple Inc",
		    "currency": "USD",
		    "exchange": "NASDAQ",
		    "mic_code": "XNAS",
		    "exchange_timezone": "America/New_York",
		    "period": "Quarterly"
		  },
		  "cash_flow": []
		}`,
		response.CashFlows{
			Meta: response.CashFlowsMeta{
				Symbol:           "AAPL",
				Name:             "Apple Inc",
				Currency:         "USD",
				Exchange:         "NASDAQ",
				MicCode:          "XNAS",
				ExchangeTimezone: "America/New_York",
				Period:           "Quarterly",
			},
			CashFlow: []response.CashFlow{},
		},
		"/cash_flow/consolidated?symbol=AAPL",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getCashFlowConsolidated: NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](httpCli, url+"/cash_flow/consolidated"),
			}
		},
		func(cli interface{}, req request.GetCashFlow) (response.CashFlows, response.Credits, error) {
			return cli.(client).GetCashFlowConsolidated(req)
		},
		"GetCashFlowConsolidated",
	)
}

func Test_client_GetProfile(t *testing.T) {
	type args struct {
		req request.GetProfile
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Profile
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetProfile{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:   "AAPL",
					Figi:     "BBG000B9XRY4",
					Isin:     "US0378331005",
					Cusip:    "037833100",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "US",
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
					  "sector": "Technology",
					  "industry": "Consumer Electronics",
					  "employees": 147000,
					  "website": "http://www.apple.com",
					  "description": "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. It also sells various related services. The company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, HomePod, iPod touch, and other Apple-branded and third-party accessories. It also provides AppleCare support services; cloud services store services; and operates various platforms, including the App Store, that allow customers to discover and download applications and digital content, such as books, music, video, games, and podcasts. In addition, the company offers various services, such as Apple Arcade, a game subscription service; Apple Music, which offers users a curated listening experience with on-demand radio stations; Apple News+, a subscription news and magazine service; Apple TV+, which offers exclusive original content; Apple Card, a co-branded credit card; and Apple Pay, a cashless payment service, as well as licenses its intellectual property. The company serves consumers, and small and mid-sized businesses; and the education, enterprise, and government markets. It sells and delivers third-party applications for its products through the App Store. The company also sells its products through its retail and online stores, and direct sales force; and third-party cellular network carriers, wholesalers, retailers, and resellers. Apple Inc. was founded in 1977 and is headquartered in Cupertino, California.",
					  "type": "Common Stock",
					  "CEO": "Mr. Timothy D. Cook",
					  "address": "One Apple Park Way",
					  "address2": "Cupertino, CA 95014",
					  "city": "Cupertino",
					  "zip": "95014",
					  "state": "CA",
					  "country": "US",
					  "phone": "408-996-1010"
					}`,
					"/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want: response.Profile{
				Symbol:      "AAPL",
				Name:        "Apple Inc",
				Exchange:    "NASDAQ",
				MicCode:     "XNAS",
				Sector:      "Technology",
				Industry:    "Consumer Electronics",
				Employees:   null.IntFrom(147000),
				Website:     "http://www.apple.com",
				Description: "Apple Inc. designs, manufactures, and markets smartphones, personal computers, tablets, wearables, and accessories worldwide. It also sells various related services. The company offers iPhone, a line of smartphones; Mac, a line of personal computers; iPad, a line of multi-purpose tablets; and wearables, home, and accessories comprising AirPods, Apple TV, Apple Watch, Beats products, HomePod, iPod touch, and other Apple-branded and third-party accessories. It also provides AppleCare support services; cloud services store services; and operates various platforms, including the App Store, that allow customers to discover and download applications and digital content, such as books, music, video, games, and podcasts. In addition, the company offers various services, such as Apple Arcade, a game subscription service; Apple Music, which offers users a curated listening experience with on-demand radio stations; Apple News+, a subscription news and magazine service; Apple TV+, which offers exclusive original content; Apple Card, a co-branded credit card; and Apple Pay, a cashless payment service, as well as licenses its intellectual property. The company serves consumers, and small and mid-sized businesses; and the education, enterprise, and government markets. It sells and delivers third-party applications for its products through the App Store. The company also sells its products through its retail and online stores, and direct sales force; and third-party cellular network carriers, wholesalers, retailers, and resellers. Apple Inc. was founded in 1977 and is headquartered in Cupertino, California.",
				Type:        "Common Stock",
				CEO:         "Mr. Timothy D. Cook",
				Address:     "One Apple Park Way",
				Address2:    "Cupertino, CA 95014",
				City:        "Cupertino",
				Zip:         "95014",
				State:       "CA",
				Country:     "US",
				Phone:       "408-996-1010",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetProfile{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:   "AAPL",
					Figi:     "BBG000B9XRY4",
					Isin:     "US0378331005",
					Cusip:    "037833100",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want:  response.Profile{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
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
						getProfile: NewEndpoint[request.GetProfile, response.Profile, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetProfile) (response.Profile, response.Credits, error) {
					return cli.(client).GetProfile(req)
				},
				"GetProfile",
			)
		})
	}
}

func Test_client_GetStatistics(t *testing.T) {
	type args struct {
		req request.GetStatistics
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.Statistics
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetStatistics{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:   "AAPL",
					Figi:     "BBG000B9XRY4",
					Isin:     "US0378331005",
					Cusip:    "037833100",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
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
					    "exchange": "NASDAQ",
					    "mic_code": "XNAS",
					    "exchange_timezone": "America/New_York"
					  },
					  "statistics": {
					    "valuations_metrics": {
					      "market_capitalization": 2546807865344,
					      "enterprise_value": 2620597731328,
					      "trailing_pe": 30.162493,
					      "forward_pe": 26.982489,
					      "peg_ratio": 1.4,
					      "price_to_sales_ttm": 7.336227,
					      "price_to_book_mrq": 39.68831,
					      "enterprise_to_revenue": 7.549,
					      "enterprise_to_ebitda": 23.623
					    },
					    "financials": {
					      "fiscal_year_ends": "2020-09-26",
					      "most_recent_quarter": "2021-06-26",
					      "gross_margin": 46.57807,
					      "profit_margin": 0.25004,
					      "operating_margin": 0.28788,
					      "return_on_assets_ttm": 0.19302,
					      "return_on_equity_ttm": 1.27125,
					      "income_statement": {
					        "revenue_ttm": 347155005440,
					        "revenue_per_share_ttm": 20.61,
					        "quarterly_revenue_growth": 0.364,
					        "gross_profit_ttm": 104956000000,
					        "ebitda": 110934999040,
					        "net_income_to_common_ttm": 86801997824,
					        "diluted_eps_ttm": 5.108,
					        "quarterly_earnings_growth_yoy": 0.932
					      },
					      "balance_sheet": {
					        "revenue_ttm": 347155005440,
					        "total_cash_mrq": 61696000000,
					        "total_cash_per_share_mrq": 3.732,
					        "total_debt_mrq": 135491002368,
					        "total_debt_to_equity_mrq": 210.782,
					        "current_ratio_mrq": 1.062,
					        "book_value_per_share_mrq": 3.882
					      },
					      "cash_flow": {
					        "operating_cash_flow_ttm": 104414003200,
					        "levered_free_cash_flow_ttm": 80625876992
					      }
					    },
					    "stock_statistics": {
					      "shares_outstanding": 16530199552,
					      "float_shares": 16513305231,
					      "avg_10_volume": 72804757,
					      "avg_90_volume": 77013078,
					      "shares_short": 93105968,
					      "short_ratio": 1.19,
					      "short_percent_of_shares_outstanding": 0.0056,
					      "percent_held_by_insiders": 0.00071000005,
					      "percent_held_by_institutions": 0.58474
					    },
					    "stock_price_summary": {
					      "fifty_two_week_low": 103.1,
					      "fifty_two_week_high": 157.26,
					      "fifty_two_week_change": 0.375625,
					      "beta": 1.201965,
					      "day_50_ma": 148.96686,
					      "day_200_ma": 134.42506
					    },
					    "dividends_and_splits": {
					      "forward_annual_dividend_rate": 0.88,
					      "forward_annual_dividend_yield": 0.0057,
					      "trailing_annual_dividend_rate": 0.835,
					      "trailing_annual_dividend_yield": 0.0053832764,
					      "5_year_average_dividend_yield": 1.27,
					      "payout_ratio": 0.16309999,
					      "dividend_frequency": "Quarterly",
					      "dividend_date": "2021-08-12",
					      "ex_dividend_date": "2021-08-06",
					      "last_split_factor": "4-for-1 split",
					      "last_split_date": "2020-08-31"
					    }
					  }
					}`,
					"/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want: response.Statistics{
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
						MarketCapitalization: null.IntFrom(2546807865344),
						EnterpriseValue:      null.IntFrom(2620597731328),
						TrailingPE:           null.FloatFrom(30.162493),
						ForwardPE:            null.FloatFrom(26.982489),
						PEGRatio:             null.FloatFrom(1.4),
						PriceToSalesTTM:      null.FloatFrom(7.336227),
						PriceToBookMRQ:       null.FloatFrom(39.68831),
						EnterpriseToRevenue:  null.FloatFrom(7.549),
						EnterpriseToEBITDA:   null.FloatFrom(23.623),
					},
					Financials: response.StatisticsFinancials{
						FiscalYearEnds:    "2020-09-26",
						MostRecentQuarter: "2021-06-26",
						GrossMargin:       null.FloatFrom(46.57807),
						ProfitMargin:      null.FloatFrom(0.25004),
						OperatingMargin:   null.FloatFrom(0.28788),
						ReturnOnAssetsTTM: null.FloatFrom(0.19302),
						ReturnOnEquityTTM: null.FloatFrom(1.27125),
						IncomeStatement: response.StatisticsIncomeStatement{
							RevenueTTM:                 null.IntFrom(347155005440),
							RevenuePerShareTTM:         null.FloatFrom(20.61),
							QuarterlyRevenueGrowth:     null.FloatFrom(0.364),
							GrossProfitTTM:             null.IntFrom(104956000000),
							EBITDA:                     null.IntFrom(110934999040),
							NetIncomeToCommonTTM:       null.IntFrom(86801997824),
							DilutedEPSTTM:              null.FloatFrom(5.108),
							QuarterlyEarningsGrowthYoY: null.FloatFrom(0.932),
						},
						BalanceSheet: response.StatisticsBalanceSheet{
							RevenueTTM:           null.IntFrom(347155005440),
							TotalCashMRQ:         null.IntFrom(61696000000),
							TotalCashPerShareMRQ: null.FloatFrom(3.732),
							TotalDebtMRQ:         null.IntFrom(135491002368),
							TotalDebtToEquityMRQ: null.FloatFrom(210.782),
							CurrentRatioMRQ:      null.FloatFrom(1.062),
							BookValuePerShareMRQ: null.FloatFrom(3.882),
						},
						CashFlow: response.StatisticsCashFlow{
							OperatingCashFlowTTM:   null.IntFrom(104414003200),
							LeveredFreeCashFlowTTM: null.IntFrom(80625876992),
						},
					},
					StockStatistics: response.StockStatistics{
						SharesOutstanding:               null.IntFrom(16530199552),
						FloatShares:                     null.IntFrom(16513305231),
						Avg10Volume:                     null.IntFrom(72804757),
						Avg90Volume:                     null.IntFrom(77013078),
						SharesShort:                     null.IntFrom(93105968),
						ShortRatio:                      null.FloatFrom(1.19),
						ShortPercentOfSharesOutstanding: null.FloatFrom(0.0056),
						PercentHeldByInsiders:           null.FloatFrom(0.00071000005),
						PercentHeldByInstitutions:       null.FloatFrom(0.58474),
					},
					StockPriceSummary: response.StockPriceSummary{
						FiftyTwoWeekLow:    null.FloatFrom(103.1),
						FiftyTwoWeekHigh:   null.FloatFrom(157.26),
						FiftyTwoWeekChange: null.FloatFrom(0.375625),
						Beta:               null.FloatFrom(1.201965),
						Day50MA:            null.FloatFrom(148.96686),
						Day200MA:           null.FloatFrom(134.42506),
					},
					DividendsAndSplits: response.DividendsAndSplits{
						ForwardAnnualDividendRate:    null.FloatFrom(0.88),
						ForwardAnnualDividendYield:   null.FloatFrom(0.0057),
						TrailingAnnualDividendRate:   null.FloatFrom(0.835),
						TrailingAnnualDividendYield:  null.FloatFrom(0.0053832764),
						FiveYearAverageDividendYield: null.FloatFrom(1.27),
						PayoutRatio:                  null.FloatFrom(0.16309999),
						DividendFrequency:            "Quarterly",
						DividendDate:                 "2021-08-12",
						ExDividendDate:               "2021-08-06",
						LastSplitFactor:              "4-for-1 split",
						LastSplitDate:                "2020-08-31",
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetStatistics{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:   "AAPL",
					Figi:     "BBG000B9XRY4",
					Isin:     "US0378331005",
					Cusip:    "037833100",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want:  response.Statistics{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?country=US&cusip=037833100&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
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
						getStatistics: NewEndpoint[request.GetStatistics, response.Statistics, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetStatistics) (response.Statistics, response.Credits, error) {
					return cli.(client).GetStatistics(req)
				},
				"GetStatistics",
			)
		})
	}
}

func Test_client_GetEarnings(t *testing.T) {
	type args struct {
		req request.GetEarnings
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        response.Earnings
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with demo API response",
			args: args{
				req: request.GetEarnings{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:        "AAPL",
					Figi:          "BBG000B9XRY4",
					Isin:          "US0378331005",
					Cusip:         "037833100",
					Exchange:      "NASDAQ",
					MicCode:       "XNAS",
					Country:       "US",
					Type:          "actual",
					Period:        "quarterly",
					OutputSize:    2,
					Format:        "json",
					Delimiter:     ",",
					DecimalPlaces: 2,
					StartDate:     "2024-01-01",
					EndDate:       "2024-12-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "AAPL",
						"name": "Apple Inc.",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "earnings": [
						{
						  "date": "2025-05-02",
						  "time": "After Hours",
						  "eps_estimate": 1.63,
						  "eps_actual": 1.65,
						  "difference": 0.02,
						  "surprise_prc": 1.23
						},
						{
						  "date": "2025-01-31",
						  "time": "After Hours",
						  "eps_estimate": 2.35,
						  "eps_actual": 2.4,
						  "difference": 0.05,
						  "surprise_prc": 2.13
						}
					  ],
					  "status": "ok"
					}`,
					"/?country=US&cusip=037833100&delimiter=%2C&dp=2&end_date=2024-12-31&exchange=NASDAQ&figi=BBG000B9XRY4&format=json&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL&type=actual",
				),
			},
			want: response.Earnings{
				Meta: response.EarningsMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				Earnings: []response.EarningsItem{
					{
						Date:        "2025-05-02",
						Time:        "After Hours",
						EPSEstimate: null.FloatFrom(1.63),
						EPSActual:   null.FloatFrom(1.65),
						Difference:  null.FloatFrom(0.02),
						SurprisePrc: null.FloatFrom(1.23),
					},
					{
						Date:        "2025-01-31",
						Time:        "After Hours",
						EPSEstimate: null.FloatFrom(2.35),
						EPSActual:   null.FloatFrom(2.4),
						Difference:  null.FloatFrom(0.05),
						SurprisePrc: null.FloatFrom(2.13),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&delimiter=%2C&dp=2&end_date=2024-12-31&exchange=NASDAQ&figi=BBG000B9XRY4&format=json&isin=US0378331005&mic_code=XNAS&outputsize=2&period=quarterly&start_date=2024-01-01&symbol=AAPL&type=actual",
		},
		{
			name: "success with multiple parameters",
			args: args{
				req: request.GetEarnings{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:     "MSFT",
					Exchange:   "NASDAQ",
					StartDate:  "2024-01-01",
					EndDate:    "2024-12-31",
					OutputSize: 2,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					20,
					`{
					  "meta": {
						"symbol": "MSFT",
						"name": "Microsoft Corporation",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "earnings": [
						{
						  "date": "2024-10-23",
						  "time": "After Hours",
						  "eps_estimate": 2.95,
						  "eps_actual": 3.30,
						  "difference": 0.35,
						  "surprise_prc": 11.86
						}
					  ],
					  "status": "ok"
					}`,
					"/?end_date=2024-12-31&exchange=NASDAQ&outputsize=2&start_date=2024-01-01&symbol=MSFT",
				),
			},
			want: response.Earnings{
				Meta: response.EarningsMeta{
					Symbol:           "MSFT",
					Name:             "Microsoft Corporation",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				Earnings: []response.EarningsItem{
					{
						Date:        "2024-10-23",
						Time:        "After Hours",
						EPSEstimate: null.FloatFrom(2.95),
						EPSActual:   null.FloatFrom(3.30),
						Difference:  null.FloatFrom(0.35),
						SurprisePrc: null.FloatFrom(11.86),
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 20),
			wantErr:     "",
			expectedURL: "/?end_date=2024-12-31&exchange=NASDAQ&outputsize=2&start_date=2024-01-01&symbol=MSFT",
		},
		{
			name: "symbol not found",
			args: args{
				req: request.GetEarnings{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol: "INVALID",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":400,"message":"**INVALID** not found: symbol may be delisted","status":"error"}`,
					"/?symbol=INVALID",
				),
			},
			want:        response.Earnings{},
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Symbol Not Found: INVALID - **INVALID** not found: symbol may be delisted",
			expectedURL: "/?symbol=INVALID",
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
						getEarnings: NewEndpoint[request.GetEarnings, response.Earnings, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetEarnings) (response.Earnings, response.Credits, error) {
					return cli.(client).GetEarnings(req)
				},
				"GetEarnings",
			)
		})
	}
}

func Test_client_GetIPOCalendar(t *testing.T) {
	type args struct {
		req request.GetIPOCalendar
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.IPOCalendar
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with date range",
			args: args{
				req: request.GetIPOCalendar{
					APIKey:    request.APIKey{APIKey: ""},
					Exchange:  "NASDAQ",
					MicCode:   "XNAS",
					Country:   "US",
					StartDate: "2023-01-01",
					EndDate:   "2023-01-31",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					40,
					`[
					  {
					    "symbol": "EXAMPLE",
					    "name": "Example Corporation",
					    "exchange": "NASDAQ",
					    "date": "2023-01-15"
					  },
					  {
					    "symbol": "TEST",
					    "name": "Test Inc",
					    "exchange": "NYSE",
					    "date": "2023-01-20"
					  }
					]`,
					"/?country=US&end_date=2023-01-31&exchange=NASDAQ&mic_code=XNAS&start_date=2023-01-01",
				),
			},
			want: response.IPOCalendar{
				{
					Symbol:   "EXAMPLE",
					Name:     "Example Corporation",
					Exchange: "NASDAQ",
					Date:     "2023-01-15",
				},
				{
					Symbol:   "TEST",
					Name:     "Test Inc",
					Exchange: "NYSE",
					Date:     "2023-01-20",
				},
			},
			want1:       response.NewCreditsImpl(100, 40),
			wantErr:     "",
			expectedURL: "/?country=US&end_date=2023-01-31&exchange=NASDAQ&mic_code=XNAS&start_date=2023-01-01",
		},
		{
			name: "success with minimal parameters",
			args: args{
				req: request.GetIPOCalendar{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					40,
					`[
					  {
					    "symbol": "NEWCO",
					    "name": "New Company Ltd",
					    "exchange": "NYSE",
					    "date": "2023-02-15"
					  }
					]`,
					"/",
				),
			},
			want: response.IPOCalendar{
				{
					Symbol:   "NEWCO",
					Name:     "New Company Ltd",
					Exchange: "NYSE",
					Date:     "2023-02-15",
				},
			},
			want1:       response.NewCreditsImpl(100, 40),
			wantErr:     "",
			expectedURL: "/",
		},
		{
			name: "plan limitation error",
			args: args{
				req: request.GetIPOCalendar{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":403,"message":"IPO Calendar is not available with your plan. Please upgrade to a Grow plan or higher to access this endpoint: https://twelvedata.com/pricing","status":"error"}`,
					"/",
				),
			},
			want:        nil,
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Plan Limitation: IPO Calendar is not available with your plan. Please upgrade to a Grow plan or higher to access this endpoint: https://twelvedata.com/pricing",
			expectedURL: "/",
		},
		{
			name: "invalid api key",
			args: args{
				req: request.GetIPOCalendar{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/",
				),
			},
			want:        nil,
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/",
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
						getIPOCalendar: NewEndpoint[request.GetIPOCalendar, response.IPOCalendar, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetIPOCalendar) (response.IPOCalendar, response.Credits, error) {
					return cli.(client).GetIPOCalendar(req)
				},
				"GetIPOCalendar",
			)
		})
	}
}

func Test_client_GetKeyExecutives(t *testing.T) {
	type args struct {
		req request.GetKeyExecutives
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        response.KeyExecutives
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with demo API response",
			args: args{
				req: request.GetKeyExecutives{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1000,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc.",
					    "currency": "USD",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "exchange_timezone": "America/New_York"
					  },
					  "key_executives": [
					    {
					      "name": "Mr. Timothy D. Cook",
					      "title": "CEO & Director",
					      "age": 63,
					      "year_born": 1961,
					      "pay": 16520856
					    },
					    {
					      "name": "Mr. Chris  Kondo",
					      "title": "Senior Director of Corporate Accounting",
					      "age": 0,
					      "year_born": 0,
					      "pay": null
					    },
					    {
					      "name": "Ms. Katherine L. Adams",
					      "title": "Senior VP, General Counsel & Secretary",
					      "age": 60,
					      "year_born": 1964,
					      "pay": 5022182
					    }
					  ]
					}`,
					"/?symbol=AAPL",
				),
			},
			want: response.KeyExecutives{
				Meta: response.KeyExecutivesMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				KeyExecutives: []response.KeyExecutive{
					{
						Name:     "Mr. Timothy D. Cook",
						Title:    "CEO & Director",
						Age:      null.IntFrom(63),
						YearBorn: null.IntFrom(1961),
						Pay:      null.IntFrom(16520856),
					},
					{
						Name:     "Mr. Chris  Kondo",
						Title:    "Senior Director of Corporate Accounting",
						Age:      null.IntFrom(0),
						YearBorn: null.IntFrom(0),
						Pay:      null.Int{},
					},
					{
						Name:     "Ms. Katherine L. Adams",
						Title:    "Senior VP, General Counsel & Secretary",
						Age:      null.IntFrom(60),
						YearBorn: null.IntFrom(1964),
						Pay:      null.IntFrom(5022182),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 1000),
			wantErr:     "",
			expectedURL: "/?symbol=AAPL",
		},
		{
			name: "success with additional parameters",
			args: args{
				req: request.GetKeyExecutives{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Exchange: "NASDAQ",
					Country:  "US",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1000,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "name": "Apple Inc.",
					    "currency": "USD",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "exchange_timezone": "America/New_York"
					  },
					  "key_executives": [
					    {
					      "name": "Mr. Timothy D. Cook",
					      "title": "CEO & Director",
					      "age": 63,
					      "year_born": 1961,
					      "pay": 16520856
					    }
					  ]
					}`,
					"/?country=US&exchange=NASDAQ&symbol=AAPL",
				),
			},
			want: response.KeyExecutives{
				Meta: response.KeyExecutivesMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				KeyExecutives: []response.KeyExecutive{
					{
						Name:     "Mr. Timothy D. Cook",
						Title:    "CEO & Director",
						Age:      null.IntFrom(63),
						YearBorn: null.IntFrom(1961),
						Pay:      null.IntFrom(16520856),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 1000),
			wantErr:     "",
			expectedURL: "/?country=US&exchange=NASDAQ&symbol=AAPL",
		},
		{
			name: "symbol not found error",
			args: args{
				req: request.GetKeyExecutives{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "INVALID",
				},
				url: mockServerWithURL(
					t,
					http.StatusBadRequest,
					100,
					0,
					`{
					  "code": 400,
					  "message": "**INVALID** not found: symbol not found",
					  "status": "error"
					}`,
					"/?symbol=INVALID",
				),
			},
			want:        response.KeyExecutives{},
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Symbol Not Found: INVALID - **INVALID** not found: symbol not found",
			expectedURL: "/?symbol=INVALID",
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
						getKeyExecutives: NewEndpoint[request.GetKeyExecutives, response.KeyExecutives, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetKeyExecutives) (response.KeyExecutives, response.Credits, error) {
					return cli.(client).GetKeyExecutives(req)
				},
				"GetKeyExecutives",
			)
		})
	}
}

func Test_client_GetMarketCap(t *testing.T) {
	type args struct {
		req request.GetMarketCap
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.MarketCap
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with demo API response",
			args: args{
				req: request.GetMarketCap{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:     "AAPL",
					Figi:       "BBG000B9XRY4",
					Isin:       "US0378331005",
					Cusip:      "037833100",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					StartDate:  "2025-01-01",
					EndDate:    "2025-01-10",
					Page:       2,
					OutputSize: 5,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					5,
					`{
					  "meta": {
						"symbol": "AAPL",
						"name": "Apple Inc.",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "market_cap": [
						{
						  "date": "2025-08-20",
						  "value": 3354078693448
						},
						{
						  "date": "2025-08-19",
						  "value": 3421602594488
						},
						{
						  "date": "2025-08-18",
						  "value": 3426499926446
						}
					  ]
					}`,
					"/?country=US&cusip=037833100&end_date=2025-01-10&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=5&page=2&start_date=2025-01-01&symbol=AAPL",
				),
			},
			want: response.MarketCap{
				Meta: response.MarketCapMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				MarketCap: []response.MarketCapData{
					{
						Date:  "2025-08-20",
						Value: null.IntFrom(3354078693448),
					},
					{
						Date:  "2025-08-19",
						Value: null.IntFrom(3421602594488),
					},
					{
						Date:  "2025-08-18",
						Value: null.IntFrom(3426499926446),
					},
				},
			},
			want1:       response.NewCreditsImpl(100, 5),
			wantErr:     "",
			expectedURL: "/?country=US&cusip=037833100&end_date=2025-01-10&exchange=NASDAQ&figi=BBG000B9XRY4&isin=US0378331005&mic_code=XNAS&outputsize=5&page=2&start_date=2025-01-01&symbol=AAPL",
		},
		{
			name: "success with additional parameters",
			args: args{
				req: request.GetMarketCap{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Symbol:     "AAPL",
					Exchange:   "NASDAQ",
					OutputSize: 5,
					StartDate:  "2025-01-01",
					EndDate:    "2025-01-10",
					Page:       1,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					5,
					`{
					  "meta": {
						"symbol": "AAPL",
						"name": "Apple Inc.",
						"currency": "USD",
						"exchange": "NASDAQ",
						"mic_code": "XNGS",
						"exchange_timezone": "America/New_York"
					  },
					  "market_cap": []
					}`,
					"/?end_date=2025-01-10&exchange=NASDAQ&outputsize=5&page=1&start_date=2025-01-01&symbol=AAPL",
				),
			},
			want: response.MarketCap{
				Meta: response.MarketCapMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc.",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					ExchangeTimezone: "America/New_York",
				},
				MarketCap: []response.MarketCapData{},
			},
			want1:       response.NewCreditsImpl(100, 5),
			wantErr:     "",
			expectedURL: "/?end_date=2025-01-10&exchange=NASDAQ&outputsize=5&page=1&start_date=2025-01-01&symbol=AAPL",
		},
		{
			name: "symbol not found error",
			args: args{
				req: request.GetMarketCap{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "INVALID",
				},
				url: mockServerWithURL(
					t,
					http.StatusBadRequest,
					100,
					0,
					`{
					  "code": 400,
					  "message": "**INVALID** not found: symbol not found",
					  "status": "error"
					}`,
					"/?symbol=INVALID",
				),
			},
			want:        response.MarketCap{},
			want1:       response.NewCreditsImpl(100, 0),
			wantErr:     "Symbol Not Found: INVALID - **INVALID** not found: symbol not found",
			expectedURL: "/?symbol=INVALID",
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
						getMarketCap: NewEndpoint[request.GetMarketCap, response.MarketCap, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetMarketCap) (response.MarketCap, response.Credits, error) {
					return cli.(client).GetMarketCap(req)
				},
				"GetMarketCap",
			)
		})
	}
}

func Test_client_GetLastChange(t *testing.T) {
	type args struct {
		req request.GetLastChange
		url string
	}

	invalidEndpointURL := mockServerWithURL(
		t,
		http.StatusBadRequest,
		100,
		50,
		`{"code":400,"message":"Invalid endpoint parameter: invalid_endpoint","status":"error"}`,
		"/last_change/invalid_endpoint?symbol=AAPL",
	)

	tests := []struct {
		name        string
		args        args
		want        response.LastChange
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with real API response format",
			args: args{
				req: request.GetLastChange{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Endpoint:   "time_series",
					Symbol:     "AAPL",
					MicCode:    "XNAS",
					StartDate:  "2024-01-01",
					Page:       1,
					OutputSize: 30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					50,
					`{"pagination":{"current_page":1,"per_page":30},"data":[]}`,
					"/last_change/time_series?mic_code=XNAS&outputsize=30&page=1&start_date=2024-01-01&symbol=AAPL",
				),
			},
			want: response.LastChange{
				Pagination: response.LastChangePagination{
					CurrentPage: null.IntFrom(1),
					PerPage:     null.IntFrom(30),
				},
				Data: []response.LastChangeData{},
			},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     "",
			expectedURL: "/last_change/time_series?mic_code=XNAS&outputsize=30&page=1&start_date=2024-01-01&symbol=AAPL",
		},
		{
			name: "success with quote endpoint",
			args: args{
				req: request.GetLastChange{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Endpoint:   "quote",
					Symbol:     "MSFT",
					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					Country:    "US",
					Page:       1,
					OutputSize: 30,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					90,
					50,
					`{"pagination":{"current_page":1,"per_page":30},"data":[{"symbol":"MSFT","exchange":"NASDAQ","mic_code":"XNAS","country":"United States","endpoint":"quote","last_change":"2025-08-21T10:30:00Z","change_type":"update","description":"Quote data updated","timestamp":"2025-08-21T10:30:00Z"}]}`,
					"/last_change/quote?country=US&exchange=NASDAQ&mic_code=XNAS&outputsize=30&page=1&symbol=MSFT",
				),
			},
			want: response.LastChange{
				Pagination: response.LastChangePagination{
					CurrentPage: null.IntFrom(1),
					PerPage:     null.IntFrom(30),
				},
				Data: []response.LastChangeData{
					{
						Symbol:      "MSFT",
						Exchange:    "NASDAQ",
						MicCode:     "XNAS",
						Country:     "United States",
						Endpoint:    "quote",
						LastChange:  "2025-08-21T10:30:00Z",
						ChangeType:  "update",
						Description: "Quote data updated",
						Timestamp:   "2025-08-21T10:30:00Z",
					},
				},
			},
			want1:       response.NewCreditsImpl(90, 50),
			wantErr:     "",
			expectedURL: "/last_change/quote?country=US&exchange=NASDAQ&mic_code=XNAS&outputsize=30&page=1&symbol=MSFT",
		},
		{
			name: "error - invalid endpoint parameter",
			args: args{
				req: request.GetLastChange{
					APIKey: request.APIKey{
						APIKey: "",
					},
					Endpoint: "invalid_endpoint",
					Symbol:   "AAPL",
				},
				url: invalidEndpointURL,
			},
			want:        response.LastChange{},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     fmt.Sprintf("HTTP 400 Bad Request: code: 400, message: Invalid endpoint parameter: invalid_endpoint, status: error (URL: %s/last_change/invalid_endpoint?symbol=AAPL)", invalidEndpointURL),
			expectedURL: "/last_change/invalid_endpoint?symbol=AAPL",
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
						getLastChange: NewEndpoint[request.GetLastChange, response.LastChange, response.Credits, error](httpCli, url+"/last_change/{endpoint}"),
					}
				},
				func(cli interface{}, req request.GetLastChange) (response.LastChange, response.Credits, error) {
					return cli.(client).GetLastChange(req)
				},
				"GetLastChange",
			)
		})
	}
}
