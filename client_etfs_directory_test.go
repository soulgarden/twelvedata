package twelvedata

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetETFsDirectory(t *testing.T) {
	type args struct {
		req request.GetETFsDirectory
		url string
	}

	missingAPIKeyURL := mockServerWithURL(
		t,
		http.StatusUnauthorized,
		100,
		1,
		`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
		"/?outputsize=1",
	)

	tests := []struct {
		name        string
		args        args
		want        response.ETFsDirectory
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetETFsDirectory{
					APIKey:     request.APIKey{APIKey: ""},
					Symbol:     "IVV",
					FundFamily: "iShares",
					Page:       1,
					OutputSize: 50,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1,
					`{
						"result": {
							"count": 1000,
							"list": [
								{
									"symbol": "IVV",
									"name": "iShares Core S&P 500 ETF",
									"country": "United States",
									"mic_code": "XNAS",
									"fund_family": "iShares",
									"fund_type": "Large Blend"
								}
							]
						},
						"status": "ok"
					}`,
					"/?fund_family=iShares&outputsize=50&page=1&symbol=IVV",
				),
			},
			want: response.ETFsDirectory{
				Result: response.ETFsDirectoryResult{
					Count: 1000,
					List: []response.ETFsDirectoryETF{
						{
							Symbol:     "IVV",
							Name:       "iShares Core S&P 500 ETF",
							Country:    "United States",
							MicCode:    "XNAS",
							FundFamily: "iShares",
							FundType:   "Large Blend",
						},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?fund_family=iShares&outputsize=50&page=1&symbol=IVV",
		},
		{
			name: "missing_api_key",
			args: args{
				req: request.GetETFsDirectory{
					APIKey:     request.APIKey{APIKey: ""},
					OutputSize: 1,
				},
				url: missingAPIKeyURL,
			},
			want:  response.ETFsDirectory{},
			want1: response.NewCreditsImpl(100, 1),
			wantErr: fmt.Sprintf(
				"HTTP 401 Unauthorized: %s (URL: %s?outputsize=1)",
				response.Error{
					Code:    401,
					Message: "**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer",
					Status:  "error",
				}.Error(),
				missingAPIKeyURL,
			),
			expectedURL: "/?outputsize=1",
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
						getETFsDirectory: NewEndpoint[request.GetETFsDirectory, response.ETFsDirectory, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetETFsDirectory) (response.ETFsDirectory, response.Credits, error) {
					return cli.(client).GetETFsDirectory(req)
				},
				"GetETFsDirectory",
			)
		})
	}
}
