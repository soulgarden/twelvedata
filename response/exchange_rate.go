package response

import "github.com/guregu/null/v6"

// ExchangeRate represents currency exchange rate data with timestamp.
type ExchangeRate struct {
	Symbol    string     `json:"symbol"`
	Rate      null.Float `json:"rate"`
	Timestamp null.Int   `json:"timestamp"`
}
