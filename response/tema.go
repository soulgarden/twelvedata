package response

import "github.com/guregu/null/v6"

// TEMA represents the Triple Exponential Moving Average technical indicator response data.
type TEMA struct {
	Meta   TEMAMeta    `json:"meta"`
	Values []TEMAValue `json:"values"`
	Status string      `json:"status"`
}

// TEMAMeta represents the metadata for the Triple Exponential Moving Average technical indicator response.
type TEMAMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        TEMAIndicator `json:"indicator"`
}

// TEMAIndicator contains metadata about the Triple Exponential Moving Average indicator configuration.
type TEMAIndicator struct {
	Name       string   `json:"name"`
	SeriesType string   `json:"series_type"`
	TimePeriod null.Int `json:"time_period"`
}

// TEMAValue represents a single data point in the Triple Exponential Moving Average technical indicator response.
type TEMAValue struct {
	Datetime string `json:"datetime"`
	TEMA     string `json:"tema"`
}
