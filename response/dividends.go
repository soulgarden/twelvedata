package response

// nolint: tagliatelle
type Dividends struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	Dividends []struct {
		PaymentDate string  `json:"payment_date"`
		Amount      float64 `json:"amount"`
	} `json:"dividends"`
}
