package response

// CurrencyConversion represents currency conversion data with amount and timestamp.
type CurrencyConversion struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}
