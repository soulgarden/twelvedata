package response

import "github.com/guregu/null/v6"

// Stoch represents the response structure for Stochastic Oscillator technical indicator.
type Stoch struct {
	Meta   StochMeta   `json:"meta"`
	Values []StochData `json:"values"`
	Status string      `json:"status"`
}

// StochMeta contains metadata for Stochastic response.
type StochMeta struct {
	Symbol           string         `json:"symbol"`
	Interval         string         `json:"interval"`
	Currency         string         `json:"currency"`
	ExchangeTimezone string         `json:"exchange_timezone"`
	Exchange         string         `json:"exchange"`
	MicCode          string         `json:"mic_code"`
	Type             string         `json:"type"`
	Indicator        StochIndicator `json:"indicator"`
}

// StochIndicator contains Stochastic indicator configuration.
type StochIndicator struct {
	Name        string   `json:"name"`
	FastKPeriod null.Int `json:"fast_k_period"`
	SlowKPeriod null.Int `json:"slow_k_period"`
	SlowDPeriod null.Int `json:"slow_d_period"`
	SlowKMAType string   `json:"slow_kma_type"`
	SlowDMAType string   `json:"slow_dma_type"`
}

// StochData represents individual Stochastic data points.
type StochData struct {
	Datetime string `json:"datetime"`
	SlowK    string `json:"slow_k"`
	SlowD    string `json:"slow_d"`
}
