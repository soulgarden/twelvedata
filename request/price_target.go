package request

// GetPriceTarget represents request parameters for analyst price targets.
type GetPriceTarget struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Figi     string `schema:"figi,omitempty"`
	Isin     string `schema:"isin,omitempty"`
	Cusip    string `schema:"cusip,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
}
