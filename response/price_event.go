package response

import "github.com/guregu/null/v6"

// PriceEvent represents a price event with timestamp and volume information.
type PriceEvent struct {
	Event     string     `json:"event"`
	Symbol    string     `json:"symbol"`
	Currency  string     `json:"currency"`
	Exchange  string     `json:"exchange"`
	MicCode   string     `json:"mic_code"`
	Type      string     `json:"type"`
	Timestamp null.Int   `json:"timestamp"`
	Price     null.Float `json:"price"`
	DayVolume null.Int   `json:"day_volume"`
}
