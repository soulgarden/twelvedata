package twelvedata

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func TestEndpointResponseFormat(t *testing.T) {
	for _, format := range []string{"", "JSON", "json", "CSV", "csv", "xml"} {
		t.Run(format, func(t *testing.T) {
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				if r.URL.Query().Get("format") != format {
					t.Errorf("format query = %q, want %q", r.URL.Query().Get("format"), format)
				}
				if _, err := w.Write([]byte(`{"price":"12.5"}`)); err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()
			endpoint := NewEndpoint[request.GetPrice, response.Price, response.Credits, error](newTestHTTPCli(server.URL), server.URL)
			_, credits, err := endpoint.Call(request.GetPrice{Symbol: "AAPL", Format: format})
			if format == "" || strings.EqualFold(format, "json") {
				if err != nil || calls.Load() != 1 {
					t.Fatalf("JSON request: calls=%d, error=%v", calls.Load(), err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), "only JSON") {
				t.Errorf("expected unsupported format error, got %v", err)
			}
			if calls.Load() != 0 || credits != nil {
				t.Errorf("unsupported format sent an HTTP request: calls=%d, credits=%v", calls.Load(), credits)
			}
		})
	}
}
