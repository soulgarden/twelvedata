package response

import "github.com/guregu/null/v6"

// WillR represents the response structure for the Williams %R technical indicator endpoint.
type WillR struct {
	Meta   WillRMeta    `json:"meta"`
	Values []WillRValue `json:"values"`
	Status string       `json:"status"`
}

// WillRMeta contains metadata information about the Williams %R calculation.
type WillRMeta struct {
	Symbol           string         `json:"symbol"`
	Interval         string         `json:"interval"`
	Currency         string         `json:"currency"`
	ExchangeTimezone string         `json:"exchange_timezone"`
	Exchange         string         `json:"exchange"`
	MicCode          string         `json:"mic_code"`
	Type             string         `json:"type"`
	Indicator        WillRIndicator `json:"indicator"`
}

// WillRIndicator contains metadata about the Williams %R indicator configuration.
type WillRIndicator struct {
	Name       string `json:"name"`
	TimePeriod int    `json:"time_period"`
}

// WillRValue represents individual Williams %R data points.
type WillRValue struct {
	Datetime  string     `json:"datetime"`
	WilliamsR null.Float `json:"williams_r"`
}
