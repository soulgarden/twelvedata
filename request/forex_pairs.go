package request

// GetForexPairs represents request parameters for forex pairs data.
type GetForexPairs struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	CurrencyBase  string `schema:"currency_base,omitempty"`
	CurrencyQuote string `schema:"currency_quote,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
}
