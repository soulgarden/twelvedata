package response

import "github.com/guregu/null/v6"

// ROC represents the response structure for the Rate of Change technical indicator endpoint.
type ROC struct {
	Meta   ROCMeta    `json:"meta"`
	Values []ROCValue `json:"values"`
	Status string     `json:"status"`
}

// ROCMeta contains metadata information about the Rate of Change calculation.
type ROCMeta struct {
	Symbol           string       `json:"symbol"`
	Interval         string       `json:"interval"`
	Currency         string       `json:"currency"`
	ExchangeTimezone string       `json:"exchange_timezone"`
	Exchange         string       `json:"exchange"`
	MicCode          string       `json:"mic_code"`
	Type             string       `json:"type"`
	Indicator        ROCIndicator `json:"indicator"`
}

// ROCIndicator contains metadata about the Rate of Change indicator configuration.
type ROCIndicator struct {
	Name       string `json:"name"`
	SeriesType string `json:"series_type"`
	TimePeriod int    `json:"time_period"`
}

// ROCValue represents individual Rate of Change data points.
type ROCValue struct {
	Datetime string     `json:"datetime"`
	ROC      null.Float `json:"roc"`
}
