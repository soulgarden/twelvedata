package response

import "github.com/guregu/null/v6"

// RSI represents the response structure for Relative Strength Index technical indicator.
type RSI struct {
	Meta   RSIMeta   `json:"meta"`
	Values []RSIData `json:"values"`
	Status string    `json:"status"`
}

// RSIMeta contains metadata for RSI response.
type RSIMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        RSIIndicator `json:"indicator"`
}

// RSIIndicator contains RSI indicator configuration.
type RSIIndicator struct {
	Name       string   `json:"name"`
	SeriesType string   `json:"series_type"`
	TimePeriod null.Int `json:"time_period"`
}

// RSIData represents individual RSI data points.
type RSIData struct {
	Datetime string `json:"datetime"`
	RSI      string `json:"rsi"`
}
