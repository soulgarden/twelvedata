package response

import "github.com/guregu/null/v6"

// MarketMovers represents the response structure for market movers data.
type MarketMovers struct {
	Values []MarketMover `json:"values"`
	Status string        `json:"status"`
}

// MarketMover represents a single market mover with price and volume information.
type MarketMover struct {
	Symbol        string     `json:"symbol"`
	Name          string     `json:"name"`
	Exchange      string     `json:"exchange"`
	MicCode       string     `json:"mic_code"`
	Datetime      string     `json:"datetime"`
	Last          null.Float `json:"last"`
	High          null.Float `json:"high"`
	Low           null.Float `json:"low"`
	Volume        null.Int   `json:"volume"`
	Change        null.Float `json:"change"`
	PercentChange null.Float `json:"percent_change"`
}
