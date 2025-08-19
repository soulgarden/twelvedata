package request

// GetMutualFundSustainability represents request parameters for mutual fund sustainability endpoint.
type GetMutualFundSustainability struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	FIGI     string `schema:"figi,omitempty"`
	ISIN     string `schema:"isin,omitempty"`
	CUSIP    string `schema:"cusip,omitempty"`
	Exchange string `schema:"exchange,omitempty"`
	Country  string `schema:"country,omitempty"`
	Format   string `schema:"format,omitempty"`
}
