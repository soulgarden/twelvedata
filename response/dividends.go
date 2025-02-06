package response

import "gopkg.in/guregu/null.v4"

type Dividends struct {
	Meta      DividendsMeta `json:"meta"`
	Dividends []Dividend    `json:"dividends"`
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
	ExDate string     `json:"ex_date"`
	Amount null.Float `json:"amount"`
}
