package response

import "github.com/guregu/null/v6"

// ATR represents the Average True Range technical indicator response data.
type ATR struct {
	Meta   ATRMeta    `json:"meta"`
	Values []ATRValue `json:"values"`
	Status string     `json:"status"`
}

// ATRMeta represents the metadata for the Average True Range technical indicator response.
type ATRMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        ATRIndicator `json:"indicator"`
}

// ATRIndicator contains metadata about the Average True Range indicator configuration.
type ATRIndicator struct {
	Name       string   `json:"name"`
	TimePeriod null.Int `json:"time_period"`
}

// ATRValue represents a single data point in the Average True Range technical indicator response.
type ATRValue struct {
	Datetime string `json:"datetime"`
	ATR      string `json:"atr"`
}
