package twelvedata

import (
	"net/http"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetExchangeRate(t *testing.T) {
	type args struct {
		req request.GetExchangeRate
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.ExchangeRate
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetExchangeRate{
					APIKey: request.APIKey{APIKey: ""},
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					100,
					`{
					    "symbol": "USD/JPY",
					    "rate": 105.12,
					    "timestamp": 1602714051
					}`,
					"/",
				),
			},
			want: response.ExchangeRate{
				Symbol:    "USD/JPY",
				Rate:      105.12,
				Timestamp: 1602714051,
			},
			want1:       response.NewCreditsImpl(100, 100),
			wantErr:     "",
			expectedURL: "/?amount=100&symbol=EUR%2FUSD",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetExchangeRate{
					APIKey: request.APIKey{APIKey: ""},
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
			want:  response.ExchangeRate{},
			want1: response.NewCreditsImpl(100, 100),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?amount=100&symbol=EUR%2FUSD",
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
						getExchangeRate: NewEndpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetExchangeRate) (response.ExchangeRate, response.Credits, error) {
					return cli.(client).GetExchangeRate(req)
				},
				"GetExchangeRate",
			)
		})
	}
}

func Test_client_GetCurrencyConversion(t *testing.T) {
	type args struct {
		req request.GetCurrencyConversion
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.CurrencyConversion
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success",
			args: args{
				req: request.GetCurrencyConversion{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "EUR/USD",
					Amount: "100",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1,
					`{
						"symbol": "EUR/USD",
						"rate": 1.16009,
						"amount": 116.009,
						"timestamp": 1755861240
					}`,
					"/?amount=100&symbol=EUR%2FUSD",
				),
			},
			want: response.CurrencyConversion{
				Symbol:    "EUR/USD",
				Rate:      1.16009,
				Amount:    116.009,
				Timestamp: 1755861240,
			},
			want1:       response.NewCreditsImpl(100, 1),
			wantErr:     "",
			expectedURL: "/?amount=100&symbol=EUR%2FUSD",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetCurrencyConversion{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "EUR/USD",
					Amount: "100",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					1,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?amount=100&symbol=EUR%2FUSD",
				),
			},
			want:  response.CurrencyConversion{},
			want1: response.NewCreditsImpl(100, 1),
			wantErr: "error received: code: 401, message: **apikey** parameter is incorrect or not specified. " +
				"You can get your free API key instantly following this link: https://twelvedata.com/pricing. " +
				"If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer, status: error",
			expectedURL: "/?amount=100&symbol=EUR%2FUSD",
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
						getCurrencyConversion: NewEndpoint[request.GetCurrencyConversion, response.CurrencyConversion, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetCurrencyConversion) (response.CurrencyConversion, response.Credits, error) {
					return cli.(client).GetCurrencyConversion(req)
				},
				"GetCurrencyConversion",
			)
		})
	}
}
