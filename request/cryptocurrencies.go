package request

// GetCryptocurrencies represents request parameters for cryptocurrencies data.
type GetCryptocurrencies struct {
	APIKey
	Symbol        string `schema:"symbol,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	CurrencyBase  string `schema:"currency_base,omitempty"`
	CurrencyQuote string `schema:"currency_quote,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
	// OutputSize limits the catalog page. Omit it to request all available records.
	OutputSize int `schema:"outputsize,omitempty"`
	// Page selects a result page; the provider defaults to 1.
	Page int `schema:"page,omitempty"`
}
