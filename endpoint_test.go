package twelvedata

import (
	"net/http"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

type testResponse struct {
	Status string `json:"status"`
}

type headerRequest struct{}

func (headerRequest) Headers() map[string]string {
	return map[string]string{"X-Test": "value"}
}

type methodRequest struct{}

func (methodRequest) Method() string {
	return http.MethodDelete
}

type rawBodyRequest struct{}

func (rawBodyRequest) RawBody() ([]byte, string, error) {
	return []byte(`{"foo":"bar"}`), "application/custom+json", nil
}

type pathRequest struct {
	ID string `schema:"-"`
}

func (req pathRequest) PathParams() map[string]string {
	return map[string]string{"id": req.ID}
}

type getWithBodyRequest struct{}

func (getWithBodyRequest) Method() string {
	return http.MethodGet
}

func (getWithBodyRequest) Body() (any, string, error) {
	return map[string]string{"foo": "bar"}, "application/json", nil
}

func newTestHTTPCli(baseURL string) *HTTPCli {
	return &HTTPCli{
		transport: &fasthttp.Client{},
		cfg: &Conf{
			Timeout: 1,
			BaseURL: baseURL,
		},
		logger: &zerolog.Logger{},
	}
}

func TestEndpoint_Call_UsesHeaders(t *testing.T) {
	serverURL := mockServerWithRequest(
		t,
		http.StatusOK,
		100,
		1,
		`{"status":"ok"}`,
		expectedRequest{
			Method:  http.MethodGet,
			URL:     "/",
			Headers: map[string]string{"X-Test": "value"},
		},
	)

	endpoint := NewEndpoint[headerRequest, testResponse, response.Credits, error](newTestHTTPCli(serverURL), serverURL)
	resp, _, err := endpoint.Call(headerRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "ok" {
		t.Fatalf("unexpected status: %s", resp.Status)
	}
}

func TestEndpoint_Call_UsesMethoder(t *testing.T) {
	serverURL := mockServerWithRequest(
		t,
		http.StatusOK,
		100,
		1,
		`{"status":"ok"}`,
		expectedRequest{
			Method: http.MethodDelete,
			URL:    "/",
		},
	)

	endpoint := NewEndpoint[methodRequest, testResponse, response.Credits, error](newTestHTTPCli(serverURL), serverURL)
	resp, _, err := endpoint.Call(methodRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "ok" {
		t.Fatalf("unexpected status: %s", resp.Status)
	}
}

func TestEndpoint_Call_UsesRawBodyer(t *testing.T) {
	serverURL := mockServerWithRequest(
		t,
		http.StatusOK,
		100,
		1,
		`{"status":"ok"}`,
		expectedRequest{
			Method: http.MethodPost,
			URL:    "/",
			Headers: map[string]string{
				"Content-Type": "application/custom+json",
			},
			Body: map[string]string{"foo": "bar"},
		},
	)

	endpoint := NewEndpoint[rawBodyRequest, testResponse, response.Credits, error](newTestHTTPCli(serverURL), serverURL)
	resp, _, err := endpoint.Call(rawBodyRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "ok" {
		t.Fatalf("unexpected status: %s", resp.Status)
	}
}

func TestEndpoint_Call_RejectsGetWithBody(t *testing.T) {
	endpoint := NewEndpoint[getWithBodyRequest, testResponse, response.Credits, error](newTestHTTPCli("http://example.com"), "http://example.com")
	_, _, err := endpoint.Call(getWithBodyRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "GET request cannot include a body") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEndpoint_Call_UsesPathParams(t *testing.T) {
	serverURL := mockServerWithRequest(
		t,
		http.StatusOK,
		100,
		1,
		`{"status":"ok"}`,
		expectedRequest{
			Method: http.MethodGet,
			URL:    "/path/abc",
		},
	)

	endpoint := NewEndpoint[pathRequest, testResponse, response.Credits, error](newTestHTTPCli(serverURL), serverURL+"/path/{id}")
	resp, _, err := endpoint.Call(pathRequest{ID: "abc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "ok" {
		t.Fatalf("unexpected status: %s", resp.Status)
	}
}
