package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetMutualFundsDirectory(t *testing.T) {
	type args struct {
		req request.GetMutualFundsDirectory
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		1,
		`{
			"result": {
				"count": 1,
				"list": [
					{
						"symbol": "0P0001LCQ3",
						"name": "JNL Small Cap Index Fund (I)",
						"country": "United States",
						"fund_family": "Jackson National",
						"fund_type": "Small Blend",
						"performance_rating": 2,
						"risk_rating": 4,
						"currency": "USD",
						"exchange": "OTC",
						"mic_code": "OTCM"
					}
				]
			},
			"status": "ok"
		}`,
		"/?cik=95953&country=United+States&cusip=120678230&figi=BBG00HMMLCH1&fund_family=Jackson+National&fund_type=Small+Blend&isin=LU1206782309&outputsize=100&page=2&performance_rating=4&risk_rating=2&symbol=0P0001LCQ3",
	)

	tests := []struct {
		name        string
		args        args
		want        response.MutualFundsDirectory
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMutualFundsDirectory{
					APIKey:            request.APIKey{APIKey: ""},
					Symbol:            "0P0001LCQ3",
					FIGI:              "BBG00HMMLCH1",
					ISIN:              "LU1206782309",
					CUSIP:             "120678230",
					CIK:               "95953",
					Country:           "United States",
					FundFamily:        "Jackson National",
					FundType:          "Small Blend",
					PerformanceRating: 4,
					RiskRating:        2,
					Page:              2,
					OutputSize:        100,
				},
				url: successURL,
			},
			want: response.MutualFundsDirectory{
				Result: response.MutualFundsDirectoryResult{
					Count: 1,
					List: []response.MutualFundsDirectoryFund{
						{
							Symbol:            "0P0001LCQ3",
							Name:              "JNL Small Cap Index Fund (I)",
							Country:           "United States",
							FundFamily:        "Jackson National",
							FundType:          "Small Blend",
							PerformanceRating: null.IntFrom(2),
							RiskRating:        null.IntFrom(4),
							Currency:          "USD",
							Exchange:          "OTC",
							MicCode:           "OTCM",
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?cik=95953&country=United+States&cusip=120678230&figi=BBG00HMMLCH1&fund_family=Jackson+National&fund_type=Small+Blend&isin=LU1206782309&outputsize=100&page=2&performance_rating=4&risk_rating=2&symbol=0P0001LCQ3",
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
						getMutualFundsDirectory: NewEndpoint[request.GetMutualFundsDirectory, response.MutualFundsDirectory, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetMutualFundsDirectory) (response.MutualFundsDirectory, response.Credits, error) {
					return cli.(client).GetMutualFundsDirectory(req)
				},
				"GetMutualFundsDirectory",
			)
		})
	}
}

func Test_client_GetMutualFundFamilies(t *testing.T) {
	type args struct {
		req request.GetMutualFundFamilies
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		1,
		`{
			"result": {
				"India": [
					"Aberdeen Standard Fund Managers Limited",
					"Aditya Birla Sun Life AMC Ltd"
				],
				"United States": [
					"Aegon Asset Management UK PLC",
					"Ampega Investment GmbH",
					"Aviva SpA"
				]
			},
			"status": "ok"
		}`,
		"/?country=United+States&fund_family=Jackson+National",
	)

	tests := []struct {
		name        string
		args        args
		want        response.MutualFundFamilies
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMutualFundFamilies{
					APIKey:     request.APIKey{APIKey: ""},
					Country:    "United States",
					FundFamily: "Jackson National",
				},
				url: successURL,
			},
			want: response.MutualFundFamilies{
				Result: map[string][]string{
					"India": {
						"Aberdeen Standard Fund Managers Limited",
						"Aditya Birla Sun Life AMC Ltd",
					},
					"United States": {
						"Aegon Asset Management UK PLC",
						"Ampega Investment GmbH",
						"Aviva SpA",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?country=United+States&fund_family=Jackson+National",
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
						getMutualFundFamilies: NewEndpoint[request.GetMutualFundFamilies, response.MutualFundFamilies, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetMutualFundFamilies) (response.MutualFundFamilies, response.Credits, error) {
					return cli.(client).GetMutualFundFamilies(req)
				},
				"GetMutualFundFamilies",
			)
		})
	}
}

func Test_client_GetMutualFundTypes(t *testing.T) {
	type args struct {
		req request.GetMutualFundTypes
		url string
	}

	successURL := mockServerWithURL(
		t,
		http.StatusOK,
		100,
		1,
		`{
			"result": {
				"India": [
					"Balanced",
					"Equity"
				],
				"United States": [
					"Large Blend",
					"Small Blend"
				]
			},
			"status": "ok"
		}`,
		"/?country=United+States&fund_type=Small+Blend",
	)

	tests := []struct {
		name        string
		args        args
		want        response.MutualFundTypes
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetMutualFundTypes{
					APIKey:   request.APIKey{APIKey: ""},
					Country:  "United States",
					FundType: "Small Blend",
				},
				url: successURL,
			},
			want: response.MutualFundTypes{
				Result: map[string][]string{
					"India": {
						"Balanced",
						"Equity",
					},
					"United States": {
						"Large Blend",
						"Small Blend",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?country=United+States&fund_type=Small+Blend",
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
						getMutualFundTypes: NewEndpoint[request.GetMutualFundTypes, response.MutualFundTypes, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetMutualFundTypes) (response.MutualFundTypes, response.Credits, error) {
					return cli.(client).GetMutualFundTypes(req)
				},
				"GetMutualFundTypes",
			)
		})
	}
}
