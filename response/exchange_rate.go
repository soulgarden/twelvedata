package response

// ExchangeRate represents currency exchange rate data with timestamp.
type ExchangeRate struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}
