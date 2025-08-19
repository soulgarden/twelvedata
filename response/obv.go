package response

import "github.com/guregu/null/v6"

// OBV represents the response structure for the On Balance Volume technical indicator endpoint.
type OBV struct {
	Meta   OBVMeta    `json:"meta"`
	Values []OBVValue `json:"values"`
	Status string     `json:"status"`
}

// OBVMeta contains metadata information about the On Balance Volume calculation.
type OBVMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        OBVIndicator `json:"indicator"`
}

// OBVIndicator contains metadata about the On Balance Volume indicator configuration.
type OBVIndicator struct {
	Name string `json:"name"`
}

// OBVValue represents individual On Balance Volume data points.
type OBVValue struct {
	Datetime string     `json:"datetime"`
	OBV      null.Float `json:"obv"`
}
