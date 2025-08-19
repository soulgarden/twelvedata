package twelvedata

import (
	"net/http"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

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
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetUsage{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"timestamp":"2025-02-02 18:02:32","current_usage":1,"plan_limit":610}`,
					"/",
				),
			},
			wantUsage: response.Usage{
				TimeStamp:    "2025-02-02 18:02:32",
				CurrentUsage: null.IntFrom(1),
				PlanLimit:    null.IntFrom(610),
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "real api response format",
			args: args{
				req: request.GetUsage{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"timestamp":"2025-08-19 19:22:24","current_usage":2,"plan_limit":8,"daily_usage":10,"plan_daily_limit":800,"plan_category":"basic"}`,
					"/",
				),
			},
			wantUsage: response.Usage{
				TimeStamp:      "2025-08-19 19:22:24",
				CurrentUsage:   null.IntFrom(2),
				PlanLimit:      null.IntFrom(8),
				DailyUsage:     null.IntFrom(10),
				PlanDailyLimit: null.IntFrom(800),
				PlanCategory:   "basic",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetUsage{
					APIKey: request.APIKey{
						APIKey: "",
					},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/",
				),
			},
			wantUsage: response.Usage{
				TimeStamp:    "",
				CurrentUsage: null.Int{},
				PlanLimit:    null.Int{},
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "error received: code: 401, message: **apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?end_date=2024-01-31&start_date=2024-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.wantUsage,
				tt.wantCredits,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getUsage: NewEndpoint[request.GetUsage, response.Usage, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetUsage) (response.Usage, response.Credits, error) {
					return cli.(client).GetUsage(req)
				},
				"GetUsage",
			)
		})
	}
}
