package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetEDGARFilings(t *testing.T) {
	type args struct {
		req request.GetEDGARFilings
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.EDGARFilings
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetEDGARFilings{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "AAPL",
					Figi:       "BBG01293F5X4",
					Isin:       "US0378331005",
					Cusip:      "594918104",
					Exchange:   "NASDAQ",
					MicCode:    "XNGS",
					Country:    "United States",
					FormType:   "8-K",
					FilledFrom: "2024-01-01",
					FilledTo:   "2024-02-01",
					Page:       2,
					PageSize:   25,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock"
					  },
					  "values": [
					    {
					      "cik": 1711463,
					      "filed_at": 1726617600,
					      "files": [
					        {
					          "name": "primary_doc.html",
					          "size": 2980,
					          "type": "144",
					          "url": "https://www.sec.gov/Archives/edgar/data/1711463/000197185724000581/primary_doc.xml"
					        }
					      ],
					      "filing_url": "https://www.sec.gov/Archives/edgar/data/1711463/0001971857-24-000581-index.htm",
					      "form_type": "144",
					      "ticker": [
					        "AAPL"
					      ]
					    }
					  ]
					}`,
					"/edgar_filings/archive?country=United+States&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&filled_from=2024-01-01&filled_to=2024-02-01&form_type=8-K&isin=US0378331005&mic_code=XNGS&page=2&page_size=25&symbol=AAPL",
				),
			},
			want: response.EDGARFilings{
				Meta: response.EDGARFilingsMeta{
					Symbol:   "AAPL",
					Exchange: "NASDAQ",
					MicCode:  "XNGS",
					Type:     "Common Stock",
				},
				Values: []response.EDGARFiling{
					{
						Cik:     null.IntFrom(1711463),
						FiledAt: null.IntFrom(1726617600),
						Files: []response.EDGARFilingFile{
							{
								Name: "primary_doc.html",
								Size: null.IntFrom(2980),
								Type: "144",
								URL:  "https://www.sec.gov/Archives/edgar/data/1711463/000197185724000581/primary_doc.xml",
							},
						},
						FilingURL: "https://www.sec.gov/Archives/edgar/data/1711463/0001971857-24-000581-index.htm",
						FormType:  "144",
						Ticker:    []string{"AAPL"},
					},
				},
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
						getEDGARFilings: NewEndpoint[request.GetEDGARFilings, response.EDGARFilings, response.Credits, error](httpCli, url+"/edgar_filings/archive"),
					}
				},
				func(cli interface{}, req request.GetEDGARFilings) (response.EDGARFilings, response.Credits, error) {
					return cli.(client).GetEDGARFilings(req)
				},
				"GetEDGARFilings",
			)
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "United States",
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
					  "insider_transactions": [
					    {
					      "full_name": "ADAMS KATHERINE L",
					      "position": "General Counsel",
					      "date_reported": "2021-05-03",
					      "is_direct": true,
					      "shares": 17000,
					      "value": 2257631,
					      "description": "Sale at price 132.57 - 133.93 per share."
					    }
					  ]
					}`,
					"/insider_transactions?country=United+States&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
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
						Shares:       null.IntFrom(17000),
						Value:        null.IntFrom(2257631),
						Description:  "Sale at price 132.57 - 133.93 per share.",
					},
				},
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
						getInsiderTransactions: NewEndpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error](httpCli, url+"/insider_transactions"),
					}
				},
				func(cli interface{}, req request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error) {
					return cli.(client).GetInsiderTransactions(req)
				},
				"GetInsiderTransactions",
			)
		})
	}
}

//nolint:dupl
func Test_client_GetInstitutionalHolders(t *testing.T) {
	type args struct {
		req request.GetInstitutionalHolders
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.InstitutionalHolders
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetInstitutionalHolders{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "United States",
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
					  "institutional_holders": [
					    {
					      "entity_name": "Vanguard Group Inc",
					      "date_reported": "2025-09-30",
					      "shares": 1399427162,
					      "value": 388536977757,
					      "percent_held": 0.0947
					    }
					  ]
					}`,
					"/institutional_holders?country=United+States&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want: response.InstitutionalHolders{
				Meta: response.InstitutionalHoldersMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
				},
				InstitutionalHolders: []response.InstitutionalHolder{
					{
						EntityName:   "Vanguard Group Inc",
						DateReported: "2025-09-30",
						Shares:       null.IntFrom(1399427162),
						Value:        null.IntFrom(388536977757),
						PercentHeld:  null.FloatFrom(0.0947),
					},
				},
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
						getInstitutionalHolders: NewEndpoint[request.GetInstitutionalHolders, response.InstitutionalHolders, response.Credits, error](httpCli, url+"/institutional_holders"),
					}
				},
				func(cli interface{}, req request.GetInstitutionalHolders) (response.InstitutionalHolders, response.Credits, error) {
					return cli.(client).GetInstitutionalHolders(req)
				},
				"GetInstitutionalHolders",
			)
		})
	}
}

//nolint:dupl
func Test_client_GetFundHolders(t *testing.T) {
	type args struct {
		req request.GetFundHolders
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.FundHolders
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetFundHolders{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "NASDAQ",
					MicCode:  "XNAS",
					Country:  "United States",
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
					  "fund_holders": [
					    {
					      "entity_name": "VANGUARD INDEX FUNDS-Vanguard Total Stock Market Index Fund",
					      "date_reported": "2025-09-30",
					      "shares": 467135722,
					      "value": 129695568698,
					      "percent_held": 0.031600002
					    }
					  ]
					}`,
					"/fund_holders?country=United+States&cusip=594918104&exchange=NASDAQ&figi=BBG01293F5X4&isin=US0378331005&mic_code=XNAS&symbol=AAPL",
				),
			},
			want: response.FundHolders{
				Meta: response.FundHoldersMeta{
					Symbol:           "AAPL",
					Name:             "Apple Inc",
					Currency:         "USD",
					Exchange:         "NASDAQ",
					MicCode:          "XNAS",
					ExchangeTimezone: "America/New_York",
				},
				FundHolders: []response.FundHolder{
					{
						EntityName:   "VANGUARD INDEX FUNDS-Vanguard Total Stock Market Index Fund",
						DateReported: "2025-09-30",
						Shares:       null.IntFrom(467135722),
						Value:        null.IntFrom(129695568698),
						PercentHeld:  null.FloatFrom(0.031600002),
					},
				},
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
						getFundHolders: NewEndpoint[request.GetFundHolders, response.FundHolders, response.Credits, error](httpCli, url+"/fund_holders"),
					}
				},
				func(cli interface{}, req request.GetFundHolders) (response.FundHolders, response.Credits, error) {
					return cli.(client).GetFundHolders(req)
				},
				"GetFundHolders",
			)
		})
	}
}

//nolint:dupl
func Test_client_GetDirectHolders(t *testing.T) {
	type args struct {
		req request.GetDirectHolders
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.DirectHolders
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetDirectHolders{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "7203",
					Figi:     "BBG01293F5X4",
					Isin:     "US0378331005",
					Cusip:    "594918104",
					Exchange: "Tadawul",
					MicCode:  "XSAU",
					Country:  "Saudi Arabia",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "7203",
					    "name": "Elm Co.",
					    "currency": "SAR",
					    "exchange": "Tadawul",
					    "mic_code": "XSAU",
					    "exchange_timezone": "Asia/Riyadh"
					  },
					  "direct_holders": [
					    {
					      "entity_name": "Public Investment Fund (Investment Company)",
					      "date_reported": "2025-03-13",
					      "shares": 53600000,
					      "value": 43148000000,
					      "percent_held": 0
					    }
					  ]
					}`,
					"/direct_holders?country=Saudi+Arabia&cusip=594918104&exchange=Tadawul&figi=BBG01293F5X4&isin=US0378331005&mic_code=XSAU&symbol=7203",
				),
			},
			want: response.DirectHolders{
				Meta: response.DirectHoldersMeta{
					Symbol:           "7203",
					Name:             "Elm Co.",
					Currency:         "SAR",
					Exchange:         "Tadawul",
					MicCode:          "XSAU",
					ExchangeTimezone: "Asia/Riyadh",
				},
				DirectHolders: []response.DirectHolder{
					{
						EntityName:   "Public Investment Fund (Investment Company)",
						DateReported: "2025-03-13",
						Shares:       null.IntFrom(53600000),
						Value:        null.IntFrom(43148000000),
						PercentHeld:  null.FloatFrom(0),
					},
				},
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
						getDirectHolders: NewEndpoint[request.GetDirectHolders, response.DirectHolders, response.Credits, error](httpCli, url+"/direct_holders"),
					}
				},
				func(cli interface{}, req request.GetDirectHolders) (response.DirectHolders, response.Credits, error) {
					return cli.(client).GetDirectHolders(req)
				},
				"GetDirectHolders",
			)
		})
	}
}

func Test_client_GetTaxInformation(t *testing.T) {
	type args struct {
		req request.GetTaxInformation
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.TaxInformation
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetTaxInformation{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "SKYQ",
					Figi:     "BBG019XJT9D6",
					Cusip:    "594918104",
					Isin:     "US5949181045",
					Exchange: "Nasdaq",
					MicCode:  "XNAS",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "meta": {
					    "symbol": "SKYQ",
					    "name": "Sky Quarry Inc.",
					    "exchange": "NASDAQ",
					    "mic_code": "XNCM",
					    "country": "United States"
					  },
					  "data": {
					    "tax_indicator": "us_1446f"
					  },
					  "status": "ok"
					}`,
					"/tax_info?cusip=594918104&exchange=Nasdaq&figi=BBG019XJT9D6&isin=US5949181045&mic_code=XNAS&symbol=SKYQ",
				),
			},
			want: response.TaxInformation{
				Meta: response.TaxInformationMeta{
					Symbol:   "SKYQ",
					Name:     "Sky Quarry Inc.",
					Exchange: "NASDAQ",
					MicCode:  "XNCM",
					Country:  "United States",
				},
				Data: response.TaxInformationData{
					TaxIndicator: null.StringFrom("us_1446f"),
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
						getTaxInformation: NewEndpoint[request.GetTaxInformation, response.TaxInformation, response.Credits, error](httpCli, url+"/tax_info"),
					}
				},
				func(cli interface{}, req request.GetTaxInformation) (response.TaxInformation, response.Credits, error) {
					return cli.(client).GetTaxInformation(req)
				},
				"GetTaxInformation",
			)
		})
	}
}

func Test_client_GetSanctionedEntities(t *testing.T) {
	type args struct {
		req request.GetSanctionedEntities
		url string
	}

	tests := []struct {
		name    string
		args    args
		want    response.SanctionedEntities
		want1   response.Credits
		wantErr string
	}{
		{
			name: "success",
			args: args{
				req: request.GetSanctionedEntities{
					APIKey: request.APIKey{APIKey: ""},
					Source: "ofac",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					  "sanctions": [
					    {
					      "symbol": "LOKESHMACH",
					      "name": "Lokesh Machines Ltd.",
					      "mic_code": "NSE",
					      "country": "India",
					      "sanction": {
					        "source": "ofac",
					        "program": "RUSSIA-EO14024",
					        "notes": "Block",
					        "lists": [
					          {
					            "name": "SDN List",
					            "published_at": "2024-10-30"
					          }
					        ]
					      }
					    }
					  ],
					  "count": 143,
					  "status": "ok"
					}`,
					"/sanctions/ofac",
				),
			},
			want: response.SanctionedEntities{
				Sanctions: []response.SanctionedEntity{
					{
						Symbol:  "LOKESHMACH",
						Name:    "Lokesh Machines Ltd.",
						MicCode: "NSE",
						Country: "India",
						Sanction: response.Sanction{
							Source:  "ofac",
							Program: "RUSSIA-EO14024",
							Notes:   "Block",
							Lists: []response.SanctionList{
								{
									Name:        "SDN List",
									PublishedAt: "2024-10-30",
								},
							},
						},
					},
				},
				Count:  null.IntFrom(143),
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
						getSanctionedEntities: NewEndpoint[request.GetSanctionedEntities, response.SanctionedEntities, response.Credits, error](httpCli, url+"/sanctions/{source}"),
					}
				},
				func(cli interface{}, req request.GetSanctionedEntities) (response.SanctionedEntities, response.Credits, error) {
					return cli.(client).GetSanctionedEntities(req)
				},
				"GetSanctionedEntities",
			)
		})
	}
}
