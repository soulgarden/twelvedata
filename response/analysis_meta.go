package response

// AnalysisMeta contains metadata for analysis endpoints.
type AnalysisMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	Type             string `json:"type"`
}
