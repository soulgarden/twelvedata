package twelvedata

import (
	"net/http"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func TestEndpointQueryValues(t *testing.T) {
	const (
		symbol         = "AAPL"
		interval       = "1day"
		base           = "EUR"
		quote          = "USD"
		indicatorQuery = "interval=1day&sd=1.23456789&symbol=AAPL"
	)

	tests := []struct {
		path    string
		request any
		query   string
	}{
		{
			"/dividends",
			request.GetDividends{Symbol: symbol, Adjust: request.Bool(false)},
			"adjust=false&symbol=AAPL",
		},
		{
			"/time_series/cross",
			request.GetTimeSeriesCross{Base: base, Quote: quote, Interval: interval, Adjust: request.Bool(false)},
			"adjust=false&base=EUR&interval=1day&quote=USD",
		},
		{
			"/bbands",
			request.GetBBands{Symbol: symbol, Interval: interval, StandardDeviations: 1.23456789},
			indicatorQuery,
		},
		{
			"/percent_b",
			request.GetPercentB{Symbol: symbol, Interval: interval, StandardDeviations: 1.23456789},
			indicatorQuery,
		},
		{
			"/vwap",
			request.GetVWAP{Symbol: symbol, Interval: interval, StandardDeviations: 1.23456789},
			indicatorQuery,
		},
		{
			"/sar",
			request.GetSAR{Symbol: symbol, Interval: interval, Acceleration: 0.0000001, Maximum: 0.23456789},
			"acceleration=0.0000001&interval=1day&maximum=0.23456789&symbol=AAPL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			serverURL := mockServerWithURL(t, http.StatusOK, 100, 1, `{}`, tt.path+"?"+tt.query)
			endpoint := NewEndpoint[any, map[string]any, response.Credits, error](newTestHTTPCli(serverURL), serverURL+tt.path)
			if _, _, err := endpoint.Call(tt.request); err != nil {
				t.Fatal(err)
			}
		})
	}
}
