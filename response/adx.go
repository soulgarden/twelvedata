package response

import "github.com/guregu/null/v6"

// ADX represents the response structure for the Average Directional Index technical indicator endpoint.
type ADX struct {
	Meta   ADXMeta   `json:"meta"`
	Values []ADXData `json:"values"`
	Status string    `json:"status"`
}

// ADXMeta contains metadata information about the Average Directional Index calculation.
type ADXMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        ADXIndicator `json:"indicator"`
}

// ADXIndicator contains metadata about the Average Directional Index indicator configuration.
type ADXIndicator struct {
	Name       string   `json:"name"`
	TimePeriod null.Int `json:"time_period"`
}

// ADXData represents individual Average Directional Index data points.
type ADXData struct {
	Datetime string `json:"datetime"`
	ADX      string `json:"adx"`
}
