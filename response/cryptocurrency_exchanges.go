package response

// CryptocurrencyExchanges represents the response structure for cryptocurrency exchanges data.
type CryptocurrencyExchanges struct {
	Data   []*CryptocurrencyExchange `json:"data"`
	Status string                    `json:"status"`
}

// CryptocurrencyExchange represents a single cryptocurrency exchange.
type CryptocurrencyExchange struct {
	Name string `json:"name"`
}
