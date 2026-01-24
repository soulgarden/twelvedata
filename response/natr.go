package response

import "github.com/guregu/null/v6"

// NATR represents the response structure for the Normalized Average True Range technical indicator endpoint.
type NATR struct {
	Meta   NATRMeta    `json:"meta"`
	Values []NATRValue `json:"values"`
	Status string      `json:"status"`
}

// NATRMeta contains metadata information about the Normalized Average True Range calculation.
type NATRMeta struct {
	Symbol           string        `json:"symbol"`
	Interval         string        `json:"interval"`
	Currency         string        `json:"currency"`
	ExchangeTimezone string        `json:"exchange_timezone"`
	Exchange         string        `json:"exchange"`
	MicCode          string        `json:"mic_code"`
	Type             string        `json:"type"`
	Indicator        NATRIndicator `json:"indicator"`
}

// NATRIndicator contains metadata about the Normalized Average True Range indicator configuration.
type NATRIndicator struct {
	Name       string   `json:"name"`
	TimePeriod null.Int `json:"time_period"`
}

// NATRValue represents individual Normalized Average True Range data points.
type NATRValue struct {
	Datetime string     `json:"datetime"`
	NATR     null.Float `json:"natr"`
}
