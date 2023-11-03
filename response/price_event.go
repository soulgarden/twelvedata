package response

type PriceEvent struct {
	Event     string  `json:"event"`
	Symbol    string  `json:"symbol"`
	Currency  string  `json:"currency"`
	Exchange  string  `json:"exchange"`
	MicCode   string  `json:"mic_code"`
	Type      string  `json:"type"`
	Timestamp int     `json:"timestamp"`
	Price     float64 `json:"price"`
	DayVolume int     `json:"day_volume"`
}
