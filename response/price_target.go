package response

import "github.com/guregu/null/v6"

// PriceTarget represents analyst price target data.
type PriceTarget struct {
	Meta        AnalysisMeta    `json:"meta"`
	PriceTarget PriceTargetData `json:"price_target"`
	Status      string          `json:"status"`
}

// PriceTargetData represents price target information.
type PriceTargetData struct {
	High     null.Float `json:"high"`
	Median   null.Float `json:"median"`
	Low      null.Float `json:"low"`
	Average  null.Float `json:"average"`
	Current  null.Float `json:"current"`
	Currency string     `json:"currency"`
}
