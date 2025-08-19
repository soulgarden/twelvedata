package response

import "github.com/guregu/null/v6"

// TR represents the response structure for the True Range technical indicator endpoint.
type TR struct {
	Meta   TRMeta    `json:"meta"`
	Values []TRValue `json:"values"`
	Status string    `json:"status"`
}

// TRMeta contains metadata information about the True Range calculation.
type TRMeta struct {
	Symbol           string      `json:"symbol"`
	Interval         string      `json:"interval"`
	Currency         string      `json:"currency"`
	ExchangeTimezone string      `json:"exchange_timezone"`
	Exchange         string      `json:"exchange"`
	MicCode          string      `json:"mic_code"`
	Type             string      `json:"type"`
	Indicator        TRIndicator `json:"indicator"`
}

// TRIndicator contains metadata about the True Range indicator configuration.
type TRIndicator struct {
	Name string `json:"name"`
}

// TRValue represents individual True Range data points.
type TRValue struct {
	Datetime string     `json:"datetime"`
	TR       null.Float `json:"tr"`
}
