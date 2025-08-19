package request

// GetDividends represents request parameters for dividends data.
type GetDividends struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Exchange  string `schema:"exchange,omitempty"`
	MicCode   string `schema:"mic_code,omitempty"`
	Country   string `schema:"country,omitempty"`
	R         string `schema:"range,omitempty"`
	StartDate string `schema:"start_date,omitempty"`
	EndDate   string `schema:"end_date,omitempty"`
}
