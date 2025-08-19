package response

// ForexPairs represents the response structure for forex pairs data.
type ForexPairs struct {
	Data   []*ForexPair `json:"data"`
	Status string       `json:"status"`
}

// ForexPair represents a single forex currency pair with base and quote currencies.
type ForexPair struct {
	Symbol        string `json:"symbol"`
	CurrencyGroup string `json:"currency_group"`
	CurrencyBase  string `json:"currency_base"`
	CurrencyQuote string `json:"currency_quote"`
}
