package response

type ExchangeRate struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}
