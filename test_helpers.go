package twelvedata

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

// mockServerWithURL creates a test HTTP server with the specified response parameters
// and validates that the request URL matches the expected URL pattern.
func mockServerWithURL(t *testing.T, responseCode int, wantCreditsLeft, wantCreditsUsed int64, responseBody string, expectedURL string) string {
	t.Helper()

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(cw http.ResponseWriter, r *http.Request) {
		// Check request URL
		if expectedURL != "" && r.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
		}

		cw.Header().Add("Api-credits-left", strconv.FormatInt(wantCreditsLeft, 10))
		cw.Header().Add("Api-credits-used", strconv.FormatInt(wantCreditsUsed, 10))

		if responseCode != http.StatusOK {
			cw.WriteHeader(responseCode)
		}

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

// testEndpointCall is a generic helper function that eliminates duplicate test execution logic
// across different endpoint tests. It creates a client with the specified endpoint, calls it,
// and validates the results according to the test expectations.
func testEndpointCall[Req any, Resp any](
	t *testing.T,
	_ string,
	args struct {
		req Req
		url string
	},
	want Resp,
	wantCredits response.Credits,
	wantErr string,
	createEndpoint func(httpCli *HTTPCli, url string) interface{},
	callEndpoint func(client interface{}, req Req) (Resp, response.Credits, error),
	methodName string,
) {
	t.Helper()

	endpoint := createEndpoint(&HTTPCli{
		transport: &fasthttp.Client{},
		cfg: &Conf{
			Timeout: 1,
			BaseURL: args.url,
		},
		logger: &zerolog.Logger{},
	}, args.url)

	got, gotCredits, err := callEndpoint(endpoint, args.req)

	// Error checking - supports both standard and compressed patterns
	if (err != nil) != (wantErr != "") || (err != nil && !reflect.DeepEqual(err.Error(), wantErr)) {
		t.Errorf("%s() error = %v, wantErr %v", methodName, err, wantErr)
		return
	}

	// Response validation
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s() got = %v, want %v", methodName, got, want)
	}

	if !reflect.DeepEqual(gotCredits, wantCredits) {
		t.Errorf("%s() gotCredits = %v, want %v", methodName, gotCredits, wantCredits)
	}
}

// TechnicalIndicatorTestArgs represents common test arguments for technical indicators.
type TechnicalIndicatorTestArgs[T any] struct {
	req T
	url string
}

// TechnicalIndicatorTestCase represents a common test case structure for technical indicators.
type TechnicalIndicatorTestCase[T any, R any] struct {
	name    string
	args    TechnicalIndicatorTestArgs[T]
	want    R
	want1   response.Credits
	wantErr string
}

// runTechnicalIndicatorTest runs a generic test for technical indicators, eliminating code duplication
// across different indicator tests by providing a common test execution pattern.
func runTechnicalIndicatorTest[T any, R any](
	t *testing.T,
	testCases []TechnicalIndicatorTestCase[T, R],
	endpointPath string,
	methodName string,
	createClient func(*HTTPCli, string) interface{},
	callMethod func(interface{}, T) (R, response.Credits, error),
) {
	t.Helper()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testEndpointCall(
				t,
				tt.name,
				struct {
					req T
					url string
				}{tt.args.req, tt.args.url},
				tt.want,
				tt.want1,
				tt.wantErr,
				func(httpCli *HTTPCli, url string) interface{} {
					return createClient(httpCli, url+endpointPath)
				},
				func(clientInterface interface{}, req T) (R, response.Credits, error) {
					return callMethod(clientInterface, req)
				},
				methodName,
			)
		})
	}
}
