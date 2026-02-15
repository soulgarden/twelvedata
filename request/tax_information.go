package request

// GetTaxInformation represents request parameters for tax information data.
type GetTaxInformation struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Figi     string `schema:"figi,omitempty"`
	Cusip    string `schema:"cusip,omitempty"`
	Isin     string `schema:"isin,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
}
