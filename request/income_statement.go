package request

// GetIncomeStatement represents request parameters for income statement data.
type GetIncomeStatement struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Exchange  string `schema:"exchange,omitempty"`
	MicCode   string `schema:"mic_code,omitempty"`
	Country   string `schema:"country,omitempty"`
	Period    string `schema:"period,omitempty"`
	StartDate string `schema:"start_date,omitempty"`
	EndDate   string `schema:"end_date,omitempty"`
}
