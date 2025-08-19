package response

// Cryptocurrencies represents the response structure for cryptocurrency data.
type Cryptocurrencies struct {
	Data   []*Cryptocurrency `json:"data"`
	Status string            `json:"status"`
}

// Cryptocurrency represents a single cryptocurrency with exchange availability information.
type Cryptocurrency struct {
	Symbol             string   `json:"symbol"`
	AvailableExchanges []string `json:"available_exchanges"`
	CurrencyBase       string   `json:"currency_base"`
	CurrencyQuote      string   `json:"currency_quote"`
}
