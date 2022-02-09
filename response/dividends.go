package response

type Dividends struct {
	Meta      *DividendsMeta `json:"meta"`
	Dividends []*Dividend    `json:"dividends"`
}

// nolint: tagliatelle
type DividendsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// nolint: tagliatelle
type Dividend struct {
	PaymentDate string  `json:"payment_date"`
	Amount      float64 `json:"amount"`
}
