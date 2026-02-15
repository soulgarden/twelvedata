package response

import "github.com/guregu/null/v6"

// Dividends represents the response structure for dividends data.
type Dividends struct {
	Meta      DividendsMeta `json:"meta"`
	Dividends []Dividend    `json:"dividends"`
}

// DividendsMeta contains metadata for dividends data.
type DividendsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// Dividend represents a single dividend payment with ex-date and amount.
type Dividend struct {
	ExDate string     `json:"ex_date"`
	Amount null.Float `json:"amount"`
}
