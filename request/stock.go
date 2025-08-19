package request

// GetStock represents request parameters for stock data.
type GetStock struct {
	APIKey
	Symbol          string `schema:"symbol,omitempty"`
	Exchange        string `schema:"exchange,omitempty"`
	MicCode         string `schema:"mic_code,omitempty"`
	Country         string `schema:"country,omitempty"`
	InstrumentType  string `schema:"type,omitempty"`
	ShowPlan        bool   `schema:"show_plan,omitempty"`
	IncludeDelisted bool   `schema:"include_delisted,omitempty"`
}
