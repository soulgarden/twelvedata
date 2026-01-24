package response

import "github.com/guregu/null/v6"

// CurrencyConversion represents currency conversion data with amount and timestamp.
type CurrencyConversion struct {
	Symbol    string     `json:"symbol"`
	Rate      null.Float `json:"rate"`
	Amount    null.Float `json:"amount"`
	Timestamp null.Int   `json:"timestamp"`
}
