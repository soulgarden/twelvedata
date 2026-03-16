package twelvedata

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

func Test_client_GetPressReleases(t *testing.T) {
	type args struct {
		req request.GetPressReleases
		url string
	}

	tests := []struct {
		name        string
		args        args
		want        response.PressReleases
		want1       response.Credits
		wantErr     string
		expectedURL string
	}{
		{
			name: "success with all filters",
			args: args{
				req: request.GetPressReleases{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "AAPL",
					Figi:   "BBG000B9Y5X2",
					Isin:   "US0378331005",
					Cusip:  "037833100",

					Exchange:   "NASDAQ",
					MicCode:    "XNAS",
					StartDate:  "2025-12-01T00:00:00",
					EndDate:    "2025-12-31T23:59:00",
					Language:   "en,en-US",
					TimeZone:   "America/New_York",
					OutputSize: 3,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					50,
					`{
					  "press_releases": [
					    {
					      "id": "20251201SF35699",
					      "datetime": "2025-12-01T11:21:00+01:00",
					      "title": "NVIDIA and Synopsys Announce Strategic Partnership",
					      "body": "<b>Key Highlights</b><ul><li>Example</li></ul>",
					      "style": "/* Style Definitions */",
					      "language": ["en", "en-US"]
					    }
					  ],
					  "status": "ok"
					}`,
					"/?cusip=037833100&end_date=2025-12-31T23%3A59%3A00&exchange=NASDAQ&figi=BBG000B9Y5X2&isin=US0378331005&language=en%2Cen-US&mic_code=XNAS&outputsize=3&start_date=2025-12-01T00%3A00%3A00&symbol=AAPL&timezone=America%2FNew_York",
				),
			},
			want: response.PressReleases{
				PressReleases: []response.PressRelease{
					{
						ID:       "20251201SF35699",
						Datetime: "2025-12-01T11:21:00+01:00",
						Title:    "NVIDIA and Synopsys Announce Strategic Partnership",
						Body:     "<b>Key Highlights</b><ul><li>Example</li></ul>",
						Style:    "/* Style Definitions */",
						Language: []string{"en", "en-US"},
					},
				},
				Status: "ok",
			},
			want1:       response.NewCreditsImpl(100, 50),
			wantErr:     "",
			expectedURL: "/?cusip=037833100&end_date=2025-12-31T23%3A59%3A00&exchange=NASDAQ&figi=BBG000B9Y5X2&isin=US0378331005&language=en%2Cen-US&mic_code=XNAS&outputsize=3&start_date=2025-12-01T00%3A00%3A00&symbol=AAPL&timezone=America%2FNew_York",
		},
		{
			name: "success with figi only and empty list",
			args: args{
				req: request.GetPressReleases{
					APIKey:     request.APIKey{APIKey: ""},
					Figi:       "BBG000B9Y5X2",
					OutputSize: 1,
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					80,
					50,
					`{
					  "press_releases": [],
					  "status": "ok"
					}`,
					"/?figi=BBG000B9Y5X2&outputsize=1",
				),
			},
			want: response.PressReleases{
				PressReleases: []response.PressRelease{},
				Status:        "ok",
			},
			want1:       response.NewCreditsImpl(80, 50),
			wantErr:     "",
			expectedURL: "/?figi=BBG000B9Y5X2&outputsize=1",
		},
		{
			name: "wrong api key",
			args: args{
				req: request.GetPressReleases{
					APIKey: request.APIKey{APIKey: ""},
					Symbol: "AAPL",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					0,
					`{"code":401,"message":"**apikey** parameter is incorrect or not specified. You can get your free API key instantly following this link: https://twelvedata.com/pricing. If you believe that everything is correct, you can contact us at https://twelvedata.com/contact/customer","status":"error"}`,
					"/?symbol=AAPL",
				),
			},
			want:  response.PressReleases{},
			want1: response.NewCreditsImpl(100, 0),
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
						getPressReleases: NewEndpoint[request.GetPressReleases, response.PressReleases, response.Credits, error](httpCli, url),
					}
				},
				func(cli interface{}, req request.GetPressReleases) (response.PressReleases, response.Credits, error) {
					return cli.(client).GetPressReleases(req)
				},
				"GetPressReleases",
			)
		})
	}
}

func TestNewClient_GetPressReleases(t *testing.T) {
	serverURL := mockServerWithURL(
		t,
		http.StatusOK,
		42,
		50,
		`{
		  "press_releases": [
		    {
		      "id": "20260311C7828",
		      "datetime": "2026-03-11T14:45:00Z",
		      "title": "Example title",
		      "body": "<p>Example body</p>",
		      "style": "/* style */",
		      "language": ["en"]
		    }
		  ],
		  "status": "ok"
		}`,
		"/press_releases?symbol=AAPL",
	)

	cfg := &Conf{
		BaseURL: serverURL,
		Timeout: 1,
		Fundamentals: Fundamentals{
			PressReleasesURL: "/press_releases",
		},
	}
	logger := zerolog.Nop()
	httpCli := NewHTTPCli(&fasthttp.Client{}, cfg, &logger)
	cli := NewClient(httpCli, cfg)

	got, credits, err := cli.GetPressReleases(request.GetPressReleases{
		APIKey: request.APIKey{APIKey: ""},
		Symbol: "AAPL",
	})
	if err != nil {
		t.Fatalf("GetPressReleases() error = %v", err)
	}

	want := response.PressReleases{
		PressReleases: []response.PressRelease{
			{
				ID:       "20260311C7828",
				Datetime: "2026-03-11T14:45:00Z",
				Title:    "Example title",
				Body:     "<p>Example body</p>",
				Style:    "/* style */",
				Language: []string{"en"},
			},
		},
		Status: "ok",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetPressReleases() got = %#v, want %#v", got, want)
	}

	wantCredits := response.NewCreditsImpl(42, 50)
	if credits.GetCreditsLeft() != wantCredits.GetCreditsLeft() || credits.GetCreditsUsed() != wantCredits.GetCreditsUsed() {
		t.Fatalf(
			"GetPressReleases() gotCredits = left:%d used:%d, want left:%d used:%d",
			credits.GetCreditsLeft(),
			credits.GetCreditsUsed(),
			wantCredits.GetCreditsLeft(),
			wantCredits.GetCreditsUsed(),
		)
	}
}
