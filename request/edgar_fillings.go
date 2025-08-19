package request

// GetEDGARFillings represents request parameters for EDGAR filings data.
type GetEDGARFillings struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	Country    string `schema:"country,omitempty"`
	FormType   string `schema:"form_type,omitempty"`
	StartDate  string `schema:"start_date,omitempty"`
	EndDate    string `schema:"end_date,omitempty"`
	OutputSize string `schema:"outputsize,omitempty"`
	Format     string `schema:"format,omitempty"`
}
