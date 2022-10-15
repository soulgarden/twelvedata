package response

type Dividends struct {
	Meta      *DividendsMeta `json:"meta"`
	Dividends []*Dividend    `json:"dividends"`
}

type DividendsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

type Dividend struct {
	PaymentDate string  `json:"payment_date"`
	Amount      float64 `json:"amount"`
}
