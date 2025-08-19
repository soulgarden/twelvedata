package response

// MarketMovers represents the response structure for market movers data.
type MarketMovers struct {
	Values []MarketMover `json:"values"`
	Status string        `json:"status"`
}

// MarketMover represents a single market mover with price and volume information.
type MarketMover struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Exchange      string  `json:"exchange"`
	MicCode       string  `json:"mic_code"`
	Datetime      string  `json:"datetime"`
	Last          float64 `json:"last"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Volume        int64   `json:"volume"`
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percent_change"`
}
