package request

// GetCryptocurrencyExchanges represents request parameters for cryptocurrency exchanges data.
type GetCryptocurrencyExchanges struct {
	APIKey
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
}
