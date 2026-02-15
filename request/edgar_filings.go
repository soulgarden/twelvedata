package request

// GetEDGARFilings represents request parameters for EDGAR filings data.
type GetEDGARFilings struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	Figi       string `schema:"figi,omitempty"`
	Isin       string `schema:"isin,omitempty"`
	Cusip      string `schema:"cusip,omitempty"`
	Exchange   string `schema:"exchange,omitempty"`
	MicCode    string `schema:"mic_code,omitempty"`
	Country    string `schema:"country,omitempty"`
	FormType   string `schema:"form_type,omitempty"`
	FilledFrom string `schema:"filled_from,omitempty"`
	FilledTo   string `schema:"filled_to,omitempty"`
	Page       int    `schema:"page,omitempty"`
	PageSize   int    `schema:"page_size,omitempty"`
}
