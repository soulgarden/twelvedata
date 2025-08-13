package twelvedata

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
	"gopkg.in/guregu/null.v4"
)

func TestErrImpl_Error(t *testing.T) {
	type testCase[Err error] struct {
		name string
		e    ErrImpl[Err]
		want string
	}

	tests := []testCase[error]{
		{
			name: "simple error",
			e: ErrImpl[error]{
				generic: nil,
				inner:   fasthttp.ErrTimeout,
			},
			want: "timeout",
		},
		{
			name: "api error",
			e: ErrImpl[error]{
				generic: nil,
				inner:   response.Error{Code: 401, Message: "Invalid API key", Status: "error"},
			},
			want: "code: 401, message: Invalid API key, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPCli_getCredits(t *testing.T) {
	type fields struct {
		transport *fasthttp.Client
		cfg       *Conf
		logger    *zerolog.Logger
	}

	type args struct {
		resp *fasthttp.Response
	}

	tests := []struct {
		name            string
		fields          fields
		args            args
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         bool
	}{
		{
			name: "valid credits headers",
			fields: fields{
				transport: &fasthttp.Client{},
				cfg:       &Conf{},
				logger:    &zerolog.Logger{},
			},
			args: args{
				resp: func() *fasthttp.Response {
					resp := fasthttp.AcquireResponse()
					resp.Header.Set("Api-credits-left", "500")
					resp.Header.Set("Api-credits-used", "10")

					return resp
				}(),
			},
			wantCreditsLeft: 500,
			wantCreditsUsed: 10,
			wantErr:         false,
		},
		{
			name: "invalid credits left header",
			fields: fields{
				transport: &fasthttp.Client{},
				cfg:       &Conf{},
				logger:    &zerolog.Logger{},
			},
			args: args{
				resp: func() *fasthttp.Response {
					resp := fasthttp.AcquireResponse()
					resp.Header.Set("Api-credits-left", "invalid")
					resp.Header.Set("Api-credits-used", "5")

					return resp
				}(),
			},
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         true,
		},
		{
			name: "missing headers",
			fields: fields{
				transport: &fasthttp.Client{},
				cfg:       &Conf{},
				logger:    &zerolog.Logger{},
			},
			args: args{
				resp: fasthttp.AcquireResponse(),
			},
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPCli{
				transport: tt.fields.transport,
				cfg:       tt.fields.cfg,
				logger:    tt.fields.logger,
			}

			gotCreditsLeft, gotCreditsUsed, err := c.getCredits(tt.args.resp)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCredits() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("getCredits() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("getCredits() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		httpCli *HTTPCli
		cfg     *Conf
	}

	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			name: "new client with valid config",
			args: args{
				httpCli: &HTTPCli{
					transport: &fasthttp.Client{},
					cfg:       &Conf{BaseURL: "https://api.twelvedata.com"},
					logger:    &zerolog.Logger{},
				},
				cfg: &Conf{
					BaseURL:       "https://api.twelvedata.com",
					ReferenceData: ReferenceData{StocksURL: "/stocks"},
					CoreData:      CoreData{TimeSeriesURL: "/time_series"},
				},
			},
			want: client{
				getStocks:     NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](&HTTPCli{transport: &fasthttp.Client{}, cfg: &Conf{BaseURL: "https://api.twelvedata.com"}, logger: &zerolog.Logger{}}, "https://api.twelvedata.com/stocks"),
				getTimeSeries: NewEndpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error](&HTTPCli{transport: &fasthttp.Client{}, cfg: &Conf{BaseURL: "https://api.twelvedata.com"}, logger: &zerolog.Logger{}}, "https://api.twelvedata.com/time_series"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.httpCli, tt.args.cfg)
			if got == nil {
				t.Errorf("NewClient() returned nil")
			}
		})
	}
}

func TestNewEndpoint(t *testing.T) {
	type args struct {
		httpCli *HTTPCli
		URI     string
	}

	type testCase[Request any, Response any, Credits response.Credits, Error error] struct {
		name string
		args args
		want *Endpoint[Request, Response, Credits, Error]
	}

	tests := []testCase[request.GetStock, response.Stocks, response.Credits, error]{
		{
			name: "new endpoint for stocks",
			args: args{
				httpCli: &HTTPCli{
					transport: &fasthttp.Client{},
					cfg:       &Conf{},
					logger:    &zerolog.Logger{},
				},
				URI: "https://api.twelvedata.com/stocks",
			},
			want: &Endpoint[request.GetStock, response.Stocks, response.Credits, error]{
				httpCli: &HTTPCli{
					transport: &fasthttp.Client{},
					cfg:       &Conf{},
					logger:    &zerolog.Logger{},
				},
				URL: "https://api.twelvedata.com/stocks",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](tt.args.httpCli, tt.args.URI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewError(t *testing.T) {
	type args[T error] struct {
		err error
		t   T
	}

	type testCase[T error] struct {
		name string
		args args[T]
		want ErrImpl[T]
	}

	tests := []testCase[error]{
		{
			name: "new error with nil generic",
			args: args[error]{
				err: fasthttp.ErrTimeout,
				t:   nil,
			},
			want: ErrImpl[error]{
				generic: nil,
				inner:   fasthttp.ErrTimeout,
			},
		},
		{
			name: "new error with api error",
			args: args[error]{
				err: fasthttp.ErrDialTimeout,
				t:   response.Error{Code: 500, Message: "Internal server error", Status: "error"},
			},
			want: ErrImpl[error]{
				generic: response.Error{Code: 500, Message: "Internal server error", Status: "error"},
				inner:   fasthttp.ErrDialTimeout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.err, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHTTPCli(t *testing.T) {
	type args struct {
		transport *fasthttp.Client
		cfg       *Conf
		logger    *zerolog.Logger
	}

	tests := []struct {
		name string
		args args
		want *HTTPCli
	}{
		{
			name: "new http client",
			args: args{
				transport: &fasthttp.Client{},
				cfg: &Conf{
					APIKey:  "test-key",
					Timeout: 15,
				},
				logger: &zerolog.Logger{},
			},
			want: &HTTPCli{
				transport: &fasthttp.Client{},
				cfg: &Conf{
					APIKey:  "test-key",
					Timeout: 15,
				},
				logger: &zerolog.Logger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPCli(tt.args.transport, tt.args.cfg, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPCli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetBalanceSheet(t *testing.T) {
	type args struct {
		req request.GetBalanceSheet
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.BalanceSheets
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetBalanceSheet{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetBalanceSheet{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.BalanceSheets{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getBalanceSheet: NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetBalanceSheet(tt.args.req)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("GetBalanceSheet() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr) {
				t.Errorf("GetBalanceSheet() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalanceSheet() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetBalanceSheet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetCashFlow(t *testing.T) {
	type args struct {
		req request.GetCashFlow
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.CashFlows
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCashFlow{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetCashFlow{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.CashFlows{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getCashFlow: NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetCashFlow(tt.args.req)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("GetCashFlow() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr) {
				t.Errorf("GetCashFlow() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCashFlow() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetCashFlow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetDividends(t *testing.T) {
	type args struct {
		req request.GetDividends
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Dividends
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetDividends{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetDividends{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:    response.Dividends{},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getDividends: NewEndpoint[request.GetDividends, response.Dividends, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetDividends(tt.args.req)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("GetDividends() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr) {
				t.Errorf("GetDividends() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDividends() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetDividends() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetEarningsCalendar(t *testing.T) {
	type args struct {
		req request.GetEarningsCalendar
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Earnings
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEarningsCalendar{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
				),
			},
			want: response.Earnings{
				Earnings: map[string][]*response.Earning{
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEarningsCalendar{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Earnings{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getEarningsCalendar: NewEndpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetEarningsCalendar(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetEarningsCalendar() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEarningsCalendar() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetEarningsCalendar() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetEtfs(t *testing.T) {
	type args struct {
		req request.GetEtfs
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Etfs
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEtfs{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "SPY",
					      "name": "SPDR S&P 500 ETF Trust",
					      "currency": "MXN",
					      "exchange": "BMV",
					      "mic_code": "XMEX",
					      "country": "Mexico",
					      "figi_code": "",
					      "access": {
					        "global": "Level B",
					        "plan": "Pro"
					      }
					    }
					  ],
					  "status": "ok"
					}`,
				),
			},
			want: response.Etfs{
				Data: []response.Etf{
					{
						Symbol:   "SPY",
						Name:     "SPDR S&P 500 ETF Trust",
						Currency: "MXN",
						Exchange: "BMV",
						MicCode:  "XMEX",
						Country:  "Mexico",
						FigiCode: "",
						Access: &response.Access{
							Global: "Level B",
							Plan:   "Pro",
						},
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetEtfs{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Etfs{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getEtfs: NewEndpoint[request.GetEtfs, response.Etfs, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetEtfs(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetEtfs() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEtfs() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetEtfs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetExchangeRate(t *testing.T) {
	type args struct {
		req request.GetExchangeRate
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.ExchangeRate
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetExchangeRate{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{
					    "symbol": "USD/JPY",
					    "rate": 105.12,
					    "timestamp": 1602714051
					}`,
				),
			},
			want: response.ExchangeRate{
				Symbol:    "USD/JPY",
				Rate:      105.12,
				Timestamp: 1602714051,
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetExchangeRate{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.ExchangeRate{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getExchangeRate: NewEndpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetExchangeRate(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetExchangeRate() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExchangeRate() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetExchangeRate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetExchanges(t *testing.T) {
	type args struct {
		req request.GetExchanges
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Exchanges
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetExchanges{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetExchanges{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Exchanges{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getExchanges: NewEndpoint[request.GetExchanges, response.Exchanges, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetExchanges(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetExchanges() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExchanges() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetExchanges() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetIncomeStatement(t *testing.T) {
	type args struct {
		req request.GetIncomeStatement
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.IncomeStatements
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetIncomeStatement{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetIncomeStatement{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.IncomeStatements{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getIncomeStatement: NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetIncomeStatement(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetIncomeStatement() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIncomeStatement() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetIncomeStatement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetIndices(t *testing.T) {
	type args struct {
		req request.GetIndices
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Indices
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetIndices{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "data": [
					    {
					      "symbol": "00000020",
					      "name": "BOLSA BILBAO INDEX",
					      "country": "Spain",
					      "currency": "EUR",
					      "exchange": "BME",
					      "mic_code": "XBIL"
					    }
					  ],
					  "count": 1209,
					  "status": "ok"
					}`,
				),
			},
			want: response.Indices{
				Data: []response.Index{
					{
						Symbol:   "00000020",
						Name:     "BOLSA BILBAO INDEX",
						Country:  "Spain",
						Currency: "EUR",
						Exchange: "BME",
						MicCode:  "XBIL",
					},
				},
				Count:  1209,
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetIndices{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Indices{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getIndices: NewEndpoint[request.GetIndices, response.Indices, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetIndices(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetIndices() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndices() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetIndices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetInsiderTransactions(t *testing.T) {
	type args struct {
		req request.GetInsiderTransactions
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.InsiderTransactions
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetInsiderTransactions{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
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
					  "insider_transactions": [
					    {
					      "full_name": "ADAMS KATHERINE L",
					      "position": "General Counsel",
					      "date_reported": "2021-05-03",
					      "is_direct": true,
					      "shares": 17000,
					      "value": 2257631,
					      "description": "Sale at price 132.57 - 133.93 per share."
					    },
					    {
					      "full_name": "MAESTRI LUCA",
					      "position": "Chief Financial Officer",
					      "date_reported": "2021-05-03",
					      "is_direct": true,
					      "shares": 121072,
					      "value": 16079338,
					      "description": "Sale at price 132.58 - 133.93 per share."
					    },
					    {
					      "full_name": "O'BRIEN DEIRDRE",
					      "position": "Officer",
					      "date_reported": "2021-04-16",
					      "is_direct": true,
					      "shares": 18216,
					      "value": 2441126,
					      "description": "Sale at price 134.01 per share."
					    },
					    {
					      "full_name": "KONDO CHRISTOPHER",
					      "position": "Officer",
					      "date_reported": "2021-04-15",
					      "is_direct": true,
					      "shares": 17603,
					      "value": 0,
					      "description": ""
					    }
					  ]
					}`,
				),
			},
			want: response.InsiderTransactions{
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
						FullName:     "ADAMS KATHERINE L",
						Position:     "General Counsel",
						DateReported: "2021-05-03",
						IsDirect:     true,
						Shares:       17000,
						Value:        2257631,
						Description:  "Sale at price 132.57 - 133.93 per share.",
					},
					{
						FullName:     "MAESTRI LUCA",
						Position:     "Chief Financial Officer",
						DateReported: "2021-05-03",
						IsDirect:     true,
						Shares:       121072,
						Value:        16079338,
						Description:  "Sale at price 132.58 - 133.93 per share.",
					},
					{
						FullName:     "O'BRIEN DEIRDRE",
						Position:     "Officer",
						DateReported: "2021-04-16",
						IsDirect:     true,
						Shares:       18216,
						Value:        2441126,
						Description:  "Sale at price 134.01 per share.",
					},
					{
						FullName:     "KONDO CHRISTOPHER",
						Position:     "Officer",
						DateReported: "2021-04-15",
						IsDirect:     true,
						Shares:       17603,
						Value:        0,
						Description:  "",
					},
				},
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetInsiderTransactions{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.InsiderTransactions{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getInsiderTransactions: NewEndpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetInsiderTransactions(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetInsiderTransactions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInsiderTransactions() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetInsiderTransactions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
					ApiKey:           request.ApiKey{ApiKey: ""},
					Market:           "stocks",
					Direction:        "gainers",
					PriceGreaterThan: 10.0,
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
					"/market_movers/stocks?direction=gainers&price_greater_than=10.000000",
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
			expectedURL: "/market_movers/stocks?direction=gainers&price_greater_than=10.000000",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetMarketMovers{
					ApiKey: request.ApiKey{ApiKey: ""},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.MarketMovers{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getMarketMovers: NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url+"/market_movers/{market}",
				),
			}

			got, got1, err := cli.GetMarketMovers(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetMarketMovers() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarketMovers() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetMarketMovers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetMarketState(t *testing.T) {
	type args struct {
		req request.GetMarketState
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    []response.MarketState
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMarketState{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				// Mock uses only one object from the array
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetMarketState{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:    nil,
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getMarketState: NewEndpoint[request.GetMarketState, []response.MarketState, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetMarketState(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetMarketState() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarketState() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetMarketState() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetProfile(t *testing.T) {
	type args struct {
		req request.GetProfile
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Profile
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetProfile{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetProfile{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Profile{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getProfile: NewEndpoint[request.GetProfile, response.Profile, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetProfile(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProfile() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetProfile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetQuote(t *testing.T) {
	type args struct {
		req request.GetQuote
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Quote
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetQuote{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
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
					  "rolling_period_change": "123.123",
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
					  "extended_timestamp": 1649845281
					}`,
				),
			},
			want: response.Quote{
				Symbol:              "AAPL",
				Name:                "Apple Inc",
				Exchange:            "NASDAQ",
				MicCode:             "XNAS",
				Currency:            "USD",
				Datetime:            "2021-09-16",
				Timestamp:           1631772000,
				Open:                "148.44000",
				High:                "148.96840",
				Low:                 "147.22099",
				Close:               "148.85001",
				Volume:              "67903927",
				PreviousClose:       "149.09000",
				Change:              "-0.23999",
				PercentChange:       "-0.16097",
				AverageVolume:       "83571571",
				Rolling1DChange:     "123.123",
				Rolling7DChange:     "123.123",
				RollingPeriodChange: "123.123",
				IsMarketOpen:        false,
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
				ExtendedTimestamp:     null.IntFrom(1649845281),
			},
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetQuote{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Quote{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getQuote: NewEndpoint[request.GetQuote, response.Quote, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url+"/quote",
				),
			}

			got, got1, err := cli.GetQuote(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetQuote() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuote() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetQuote() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetStatistics(t *testing.T) {
	type args struct {
		req request.GetStatistics
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Statistics
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetStatistics{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetStatistics{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			want:  response.Statistics{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getStatistics: NewEndpoint[request.GetStatistics, response.Statistics, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetStatistics(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetStatistics() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatistics() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetStatistics() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_GetStocks(t *testing.T) {
	type args struct {
		req request.GetStock
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.Stocks
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetStock{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
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
			want1:   response.NewCreditsImpl(100, 100),
			wantErr: "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetStock{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getStocks: NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			got, got1, err := cli.GetStocks(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetStocks() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStocks() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetStocks() got1 = %v, want %v", got1, tt.want1)
			}
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
	}{
		{
			name: "success",
			args: args{
				req: request.GetTimeSeries{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
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
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetTimeSeries{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			wantTimeSeries: response.TimeSeries{
				Meta:   response.TimeSeriesMeta{},
				Values: nil,
				Status: "",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getTimeSeries: NewEndpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			gotTimeSeries, gotCredits, err := cli.GetTimeSeries(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetTimeSeries() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotTimeSeries, tt.wantTimeSeries) {
				t.Errorf("GetTimeSeries() gotTimeSeries = %v, wantTimeSeries %v", gotTimeSeries, tt.wantTimeSeries)
			}

			if !reflect.DeepEqual(gotCredits, tt.wantCredits) {
				t.Errorf("GetTimeSeries() gotCredits = %v, wantCredits %v", gotCredits, tt.wantCredits)
			}
		})
	}
}

func Test_client_GetUsage(t *testing.T) {
	type args struct {
		req request.GetUsage
		url string
	}

	tests := []struct {
		name        string
		args        args
		wantUsage   response.Usage
		wantCredits response.Credits
		wantErr     string
	}{
		{
			name: "success",
			args: args{
				req: request.GetUsage{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"timestamp":"2025-02-02 18:02:32","current_usage":1,"plan_limit":610}`,
				),
			},
			wantUsage: response.Usage{
				TimeStamp:    "2025-02-02 18:02:32",
				CurrentUsage: null.IntFrom(1),
				PlanLimit:    null.IntFrom(610),
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetUsage{
					ApiKey: request.ApiKey{
						ApiKey: "",
					},
				},
				url: mockServer(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
				),
			},
			wantUsage: response.Usage{
				TimeStamp:    "",
				CurrentUsage: null.Int{},
				PlanLimit:    null.Int{},
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getUsage: NewEndpoint[request.GetUsage, response.Usage, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url,
				),
			}

			gotUsage, gotCredits, err := cli.GetUsage(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetUsage() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(gotUsage, tt.wantUsage) {
				t.Errorf("GetUsage() gotUsage = %v, wantUsage %v", gotUsage, tt.wantUsage)
			}

			if !reflect.DeepEqual(gotCredits, tt.wantCredits) {
				t.Errorf("GetUsage() gotCredits = %v, wantCredits %v", gotCredits, tt.wantCredits)
			}
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
					ApiKey: request.ApiKey{ApiKey: ""},
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
			cli := client{
				getPrice: NewEndpoint[request.GetPrice, response.Price, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url+"/price",
				),
			}

			got, got1, err := cli.GetPrice(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrice() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPrice() got1 = %v, want %v", got1, tt.want1)
			}
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
					ApiKey: request.ApiKey{ApiKey: ""},
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
			cli := client{
				getEOD: NewEndpoint[request.GetEOD, response.EOD, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url+"/eod",
				),
			}

			got, got1, err := cli.GetEOD(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetEOD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEOD() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetEOD() got1 = %v, want %v", got1, tt.want1)
			}
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
					ApiKey:   request.ApiKey{ApiKey: ""},
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
							"symbol": "JPY/BTC",
							"interval": "1day",
							"currency": "BTC",
							"exchange_timezone": "UTC",
							"exchange": "Cross Rate",
							"mic_code": "",
							"type": "Cross Rate"
						},
						"values": [
							{
								"datetime": "2023-01-15 00:00:00",
								"open": "0.00007541",
								"high": "0.00007612",
								"low": "0.00007398",
								"close": "0.00007456",
								"volume": "0"
							}
						],
						"status": "ok"
					}`,
					"/time_series/cross?base=JPY&interval=1day&quote=BTC",
				),
			},
			want: response.TimeSeriesCross{
				Meta: response.TimeSeriesMeta{
					Symbol:           "JPY/BTC",
					Interval:         "1day",
					Currency:         "BTC",
					ExchangeTimezone: "UTC",
					Exchange:         "Cross Rate",
					MicCode:          "",
					Type:             "Cross Rate",
				},
				Values: []response.TimeSeriesValue{
					{
						Datetime: "2023-01-15 00:00:00",
						Open:     "0.00007541",
						High:     "0.00007612",
						Low:      "0.00007398",
						Close:    "0.00007456",
						Volume:   "0",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/time_series/cross?base=JPY&interval=1day&quote=BTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := client{
				getTimeSeriesCross: NewEndpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error](
					&HTTPCli{
						transport: &fasthttp.Client{},
						cfg: &Conf{
							Timeout: 1,
							BaseURL: tt.args.url,
						},
						logger: &zerolog.Logger{},
					},
					tt.args.url+"/time_series/cross",
				),
			}

			got, got1, err := cli.GetTimeSeriesCross(tt.args.req)
			if (err != nil) != (tt.wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), tt.wantErr)) {
				t.Errorf("GetTimeSeriesCross() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeSeriesCross() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetTimeSeriesCross() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func mockServer(t *testing.T, responseCode int, wantCreditsLeft, wantCreditsUsed int64, responseBody string) string {
	t.Helper()

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(cw http.ResponseWriter, _ *http.Request) {
		if responseCode == http.StatusInternalServerError {
			cw.WriteHeader(responseCode)
		}

		cw.Header().Add("Api-credits-left", strconv.FormatInt(wantCreditsLeft, 10))
		cw.Header().Add("Api-credits-used", strconv.FormatInt(wantCreditsUsed, 10))

		_, err := cw.Write([]byte(responseBody))
		if err != nil {
			t.Error(err)
		}
	}))

	server.Start()

	t.Cleanup(func() {
		server.Close()
	})

	return server.URL
}

func mockServerWithURL(t *testing.T, responseCode int, wantCreditsLeft, wantCreditsUsed int64, responseBody string, expectedURL string) string {
	t.Helper()

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(cw http.ResponseWriter, r *http.Request) {
		// Check request URL
		if expectedURL != "" && r.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
		}

		if responseCode == http.StatusInternalServerError {
			cw.WriteHeader(responseCode)
		}

		cw.Header().Add("Api-credits-left", strconv.FormatInt(wantCreditsLeft, 10))
		cw.Header().Add("Api-credits-used", strconv.FormatInt(wantCreditsUsed, 10))

		_, err := cw.Write([]byte(responseBody))
		if err != nil {
			t.Error(err)
		}
	}))

	server.Start()

	t.Cleanup(func() {
		server.Close()
	})

	return server.URL
}
