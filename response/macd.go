package response

import "github.com/guregu/null/v6"

// MACD represents the response structure for Moving Average Convergence Divergence technical indicator.
type MACD struct {
	Meta   MACDMeta   `json:"meta"`
	Values []MACDData `json:"values"`
	Status string     `json:"status"`
}

// MACDMeta contains metadata for MACD response.
type MACDMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        MACDIndicator `json:"indicator"`
}

// MACDIndicator contains MACD indicator configuration.
type MACDIndicator struct {
	Name         string   `json:"name"`
	SeriesType   string   `json:"series_type"`
	FastPeriod   null.Int `json:"fast_period"`
	SlowPeriod   null.Int `json:"slow_period"`
	SignalPeriod null.Int `json:"signal_period"`
}

// MACDData represents individual MACD data points.
type MACDData struct {
	Datetime   string `json:"datetime"`
	MACD       string `json:"macd"`
	MACDSignal string `json:"macd_signal"`
	MACDHist   string `json:"macd_hist"`
}
