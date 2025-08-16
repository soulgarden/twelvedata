package response

type Cryptocurrencies struct {
	Data   []*Cryptocurrency `json:"data"`
	Status string            `json:"status"`
}

type Cryptocurrency struct {
	Symbol             string   `json:"symbol"`
	AvailableExchanges []string `json:"available_exchanges"`
	CurrencyBase       string   `json:"currency_base"`
	CurrencyQuote      string   `json:"currency_quote"`
}
