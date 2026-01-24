package twelvedata

import (
	"encoding/json"
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
						APIKey: "demo",
					},
					Format:    "CSV",
					Delimiter: ";",
					TimeZone:  "America/New_York",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{"timestamp":"2025-05-07 11:10:12","current_usage":4003,"plan_limit":20000,"plan_category":"enterprise"}`,
					"/?apikey=demo&delimiter=%3B&format=CSV&timezone=America%2FNew_York",
				),
			},
			wantUsage: response.Usage{
				TimeStamp:    "2025-05-07 11:10:12",
				CurrentUsage: null.IntFrom(4003),
				PlanLimit:    null.IntFrom(20000),
				PlanCategory: "enterprise",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?apikey=demo&delimiter=%3B&format=CSV&timezone=America%2FNew_York",
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
				PlanCategory: "",
			},
			wantCredits: response.NewCreditsImpl(100, 100),
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

func Test_client_GetBatches(t *testing.T) {
	type args struct {
		req request.GetBatches
		url string
	}

	requests := map[string]request.BatchRequest{
		"req_1": {URL: "/time_series?symbol=AAPL&interval=1min&apikey=demo&outputsize=2"},
		"req_2": {URL: "/exchange_rate?symbol=USD/JPY&apikey=demo"},
	}

	responseBody := `{"code":200,"status":"success","data":{"req_1":{"status":"success","response":{"foo":"bar"}},"req_2":{"status":"success","response":{"rate":149.25999,"symbol":"USD/JPY","timestamp":1740160260}}}}`

	tests := []struct {
		name        string
		args        args
		want        response.Batches
		wantCredits response.Credits
		wantErr     string
	}{
		{
			name: "success",
			args: args{
				req: request.GetBatches{
					APIKey:   request.APIKey{APIKey: "demo"},
					Requests: requests,
				},
				url: mockServerWithRequest(
					t,
					http.StatusOK,
					100,
					100,
					responseBody,
					expectedRequest{
						Method: "POST",
						URL:    "/batch",
						Headers: map[string]string{
							"Authorization": "apikey demo",
							"Content-Type":  "application/json",
						},
						Body: requests,
					},
				),
			},
			want: response.Batches{
				Code:   null.IntFrom(200),
				Status: "success",
				Data: map[string]response.BatchResponse{
					"req_1": {
						Status:   "success",
						Response: json.RawMessage(`{"foo":"bar"}`),
					},
					"req_2": {
						Status:   "success",
						Response: json.RawMessage(`{"rate":149.25999,"symbol":"USD/JPY","timestamp":1740160260}`),
					},
				},
			},
			wantCredits: response.NewCreditsImpl(100, 100),
			wantErr:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				tt.args,
				tt.want,
				tt.wantCredits,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return client{
						getBatches: NewEndpoint[request.GetBatches, response.Batches, response.Credits, error](httpCli, url+"/batch"),
					}
				},
				func(cli interface{}, req request.GetBatches) (response.Batches, response.Credits, error) {
					return cli.(client).GetBatches(req)
				},
				"GetBatches",
			)
		})
	}
}
