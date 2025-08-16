package response

type ForexPairs struct {
	Data   []*ForexPair `json:"data"`
	Status string       `json:"status"`
}

type ForexPair struct {
	Symbol        string `json:"symbol"`
	CurrencyGroup string `json:"currency_group"`
	CurrencyBase  string `json:"currency_base"`
	CurrencyQuote string `json:"currency_quote"`
}
