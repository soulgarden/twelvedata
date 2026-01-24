package twelvedata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetETFFamilies(t *testing.T) {
	type args struct {
		req request.GetETFFamilies
		url string
	}

	missingAPIKeyURL := mockServerWithURL(
		t,
		http.StatusUnauthorized,
		100,
		10,
		`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
		"/?country=United+States&fund_family=iShares",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFFamilies
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFFamilies{
					APIKey:     request.APIKey{APIKey: ""},
					Country:    "United States",
					FundFamily: "iShares",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
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
					"/?country=United+States&fund_family=iShares",
				),
			},
			want: response.ETFFamilies{
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
			want1:       response.NewCreditsImpl(100, 10),
			wantErr:     "",
			expectedURL: "/?country=United+States&fund_family=iShares",
		},
		{
			name: "missing_api_key",
			args: args{
				req: request.GetETFFamilies{
					APIKey:     request.APIKey{APIKey: ""},
					Country:    "United States",
					FundFamily: "iShares",
				},
				url: missingAPIKeyURL,
			},
			want:  response.ETFFamilies{},
			want1: response.NewCreditsImpl(100, 10),
			wantErr: fmt.Sprintf(
				"HTTP 401 Unauthorized: %s (URL: %s?country=United+States&fund_family=iShares)",
				response.Error{
					Code:    null.IntFrom(401),
					Message: "**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer",
					Status:  "error",
				}.Error(),
				missingAPIKeyURL,
			),
			expectedURL: "/?country=United+States&fund_family=iShares",
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
						getETFFamilies: NewEndpoint[request.GetETFFamilies, response.ETFFamilies, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFFamilies) (response.ETFFamilies, response.Credits, error) {
					return cli.(client).GetETFFamilies(req)
				},
				"GetETFFamilies",
			)
		})
	}
}

func Test_client_GetETFTypes(t *testing.T) {
	type args struct {
		req request.GetETFTypes
		url string
	}

	missingAPIKeyURL := mockServerWithURL(
		t,
		http.StatusUnauthorized,
		100,
		10,
		`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
		"/?country=United+States&fund_type=Large+Blend",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFTypes
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFTypes{
					APIKey:   request.APIKey{APIKey: ""},
					Country:  "United States",
					FundType: "Large Blend",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
						"result": {
							"Singapore": [
								"Property - Indirect Asia",
								"Sector Equity Water"
							],
							"United States": [
								"Asia-Pacific ex-Japan Equity",
								"EUR Flexible Allocation - Global"
							]
						},
						"status": "ok"
					}`,
					"/?country=United+States&fund_type=Large+Blend",
				),
			},
			want: response.ETFTypes{
				Result: map[string][]string{
					"Singapore": {
						"Property - Indirect Asia",
						"Sector Equity Water",
					},
					"United States": {
						"Asia-Pacific ex-Japan Equity",
						"EUR Flexible Allocation - Global",
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 10),
			wantErr:     "",
			expectedURL: "/?country=United+States&fund_type=Large+Blend",
		},
		{
			name: "missing_api_key",
			args: args{
				req: request.GetETFTypes{
					APIKey:   request.APIKey{APIKey: ""},
					Country:  "United States",
					FundType: "Large Blend",
				},
				url: missingAPIKeyURL,
			},
			want:  response.ETFTypes{},
			want1: response.NewCreditsImpl(100, 10),
			wantErr: fmt.Sprintf(
				"HTTP 401 Unauthorized: %s (URL: %s?country=United+States&fund_type=Large+Blend)",
				response.Error{
					Code:    null.IntFrom(401),
					Message: "**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer",
					Status:  "error",
				}.Error(),
				missingAPIKeyURL,
			),
			expectedURL: "/?country=United+States&fund_type=Large+Blend",
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
						getETFTypes: NewEndpoint[request.GetETFTypes, response.ETFTypes, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFTypes) (response.ETFTypes, response.Credits, error) {
					return cli.(client).GetETFTypes(req)
				},
				"GetETFTypes",
			)
		})
	}
}
