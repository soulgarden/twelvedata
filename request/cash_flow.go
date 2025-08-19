package request

// GetCashFlow represents request parameters for cash flow data.
type GetCashFlow struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Exchange  string `schema:"exchange,omitempty"`
	MicCode   string `schema:"mic_code,omitempty"`
	Country   string `schema:"country,omitempty"`
	StartDate string `schema:"start_date,omitempty"`
	EndDate   string `schema:"end_date,omitempty"`
	Period    string `schema:"period,omitempty"`
}
