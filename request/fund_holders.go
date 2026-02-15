package request

// GetFundHolders represents request parameters for fund holders data.
type GetFundHolders struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Figi     string `schema:"figi,omitempty"`
	Isin     string `schema:"isin,omitempty"`
	Cusip    string `schema:"cusip,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Country  string `schema:"country,omitempty"`
}
