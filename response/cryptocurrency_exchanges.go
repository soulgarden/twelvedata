package response

type CryptocurrencyExchanges struct {
	Data   []*CryptocurrencyExchange `json:"data"`
	Status string                    `json:"status"`
}

type CryptocurrencyExchange struct {
	Name string `json:"name"`
}
