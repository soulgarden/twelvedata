package twelvedata

import (
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

func TestErrImplError_Error(t *testing.T) {
	type testCase[Err error] struct {
		name string
		e    ErrImplError[Err]
		want string
	}

	tests := []testCase[error]{
		{
			name: "simple error",
			e: ErrImplError[error]{
				generic: nil,
				inner:   fasthttp.ErrTimeout,
			},
			want: "timeout",
		},
		{
			name: "api error",
			e: ErrImplError[error]{
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
		want ErrImplError[T]
	}

	tests := []testCase[error]{
		{
			name: "new error with nil generic",
			args: args[error]{
				err: fasthttp.ErrTimeout,
				t:   nil,
			},
			want: ErrImplError[error]{
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
			want: ErrImplError[error]{
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

// TestHTTPCli_makeRequest_TimeoutRetry tests the timeout retry logic in makeRequest function.
// This test verifies that:
// 1. ErrDialTimeout triggers a retry attempt
// 2. Successful retry returns correct response and credits
// 3. Failed retry returns proper error
// 4. Non-dial timeout errors do not trigger retry.
func TestHTTPCli_makeRequest_TimeoutRetry(t *testing.T) {
	logger := zerolog.Nop()

	tests := []struct {
		name            string
		setupServer     func() *httptest.Server
		dialError       error
		retryDialError  error
		wantCreditsLeft int64
		wantCreditsUsed int64
		wantErr         string
	}{
		{
			name: "successful retry after dial timeout",
			setupServer: func() *httptest.Server {
				callCount := 0
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					callCount++
					// The retry should succeed and return valid response
					w.Header().Set("Api-credits-left", "100")
					w.Header().Set("Api-credits-used", "5")
					w.WriteHeader(http.StatusOK)
					if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
						t.Errorf("Failed to write response: %v", err)
					}
				}))
			},
			dialError:       fasthttp.ErrDialTimeout,
			retryDialError:  nil, // Success on retry
			wantCreditsLeft: 100,
			wantCreditsUsed: 5,
			wantErr:         "",
		},
		{
			name: "retry fails after dial timeout",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					// This won't be reached due to dial errors
					w.WriteHeader(http.StatusOK)
				}))
			},
			dialError:       fasthttp.ErrDialTimeout,
			retryDialError:  fasthttp.ErrDialTimeout, // Also fails on retry
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         "http cli request",
		},
		{
			name: "non-dial timeout error should not retry",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					// This won't be reached due to timeout error
					w.WriteHeader(http.StatusOK)
				}))
			},
			dialError:       fasthttp.ErrTimeout, // Non-dial timeout
			retryDialError:  nil,
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         "http request",
		},
		{
			name: "successful retry but bad status code should return HTTP error",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Api-credits-left", "90")
					w.Header().Set("Api-credits-used", "15")
					w.WriteHeader(http.StatusBadRequest)
					if _, err := w.Write([]byte(`{"error":"bad request"}`)); err != nil {
						t.Errorf("Failed to write response: %v", err)
					}
				}))
			},
			dialError:       fasthttp.ErrDialTimeout,
			retryDialError:  nil, // Success on retry (but bad status)
			wantCreditsLeft: 0,
			wantCreditsUsed: 0,
			wantErr:         "HTTP 400",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.setupServer()
			defer server.Close()

			// Create a custom transport that can simulate dial errors
			callCount := 0
			transport := &fasthttp.Client{}

			// Override the Dial function to simulate errors
			transport.Dial = func(addr string) (net.Conn, error) {
				callCount++
				if callCount == 1 && tt.dialError != nil {
					return nil, tt.dialError
				}
				if callCount == 2 && tt.retryDialError != nil {
					return nil, tt.retryDialError
				}
				// For successful cases, use default dial
				return fasthttp.Dial(addr)
			}

			c := &HTTPCli{
				transport: transport,
				cfg:       &Conf{Timeout: 5},
				logger:    &logger,
			}

			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseResponse(resp)

			gotCreditsLeft, gotCreditsUsed, err := c.makeRequest(server.URL+"/test", resp)

			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("makeRequest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				if err == nil {
					t.Errorf("makeRequest() error = nil, wantErr containing %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("makeRequest() error = %v, wantErr containing %v", err.Error(), tt.wantErr)
					return
				}
			}

			if gotCreditsLeft != tt.wantCreditsLeft {
				t.Errorf("makeRequest() gotCreditsLeft = %v, want %v", gotCreditsLeft, tt.wantCreditsLeft)
			}

			if gotCreditsUsed != tt.wantCreditsUsed {
				t.Errorf("makeRequest() gotCreditsUsed = %v, want %v", gotCreditsUsed, tt.wantCreditsUsed)
			}

			// Verify that retry was attempted for dial timeout errors only
			expectedCalls := 1
			if errors.Is(tt.dialError, fasthttp.ErrDialTimeout) {
				expectedCalls = 2 // Original + retry
			}

			if callCount != expectedCalls {
				t.Errorf("Expected %d dial calls, got %d", expectedCalls, callCount)
			}
		})
	}
}
