package response

import "github.com/guregu/null/v6"

// AD represents the response structure for the Accumulation/Distribution technical indicator endpoint.
type AD struct {
	Meta   ADMeta    `json:"meta"`
	Values []ADValue `json:"values"`
	Status string    `json:"status"`
}

// ADMeta contains metadata information about the Accumulation/Distribution calculation.
type ADMeta struct {
	Symbol           string      `json:"symbol"`
	Interval         string      `json:"interval"`
	Currency         string      `json:"currency"`
	ExchangeTimezone string      `json:"exchange_timezone"`
	Exchange         string      `json:"exchange"`
	MicCode          string      `json:"mic_code"`
	Type             string      `json:"type"`
	Indicator        ADIndicator `json:"indicator"`
}

// ADIndicator contains metadata about the Accumulation/Distribution indicator configuration.
type ADIndicator struct {
	Name string `json:"name"`
}

// ADValue represents individual Accumulation/Distribution data points.
type ADValue struct {
	Datetime string     `json:"datetime"`
	AD       null.Float `json:"ad"`
}
