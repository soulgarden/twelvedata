package request

type GetCryptocurrencyExchanges struct {
	ApiKey
	Format    string `schema:"format,omitempty"`
	Delimiter string `schema:"delimiter,omitempty"`
}
