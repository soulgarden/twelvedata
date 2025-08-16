package request

type GetCryptocurrencies struct {
	ApiKey
	Symbol        string `schema:"symbol,omitempty"`
	Exchange      string `schema:"exchange,omitempty"`
	CurrencyBase  string `schema:"currency_base,omitempty"`
	CurrencyQuote string `schema:"currency_quote,omitempty"`
	Format        string `schema:"format,omitempty"`
	Delimiter     string `schema:"delimiter,omitempty"`
}
