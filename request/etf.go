package request

// GetEtfs represents request parameters for ETF data.
type GetEtfs struct {
	APIKey
	Symbol          string `schema:"symbol,omitempty"`
	Exchange        string `schema:"exchange,omitempty"`
	MicCode         string `schema:"mic_code,omitempty"`
	Country         string `schema:"country,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
}
