package response

import "github.com/guregu/null/v6"

// PriceTarget represents analyst price target data.
type PriceTarget struct {
	Symbol      string               `json:"symbol"`
	Exchange    string               `json:"exchange"`
	Current     PriceTargetCurrent   `json:"current"`
	History     []PriceTargetHistory `json:"history"`
	LastUpdated string               `json:"last_updated"`
}

// PriceTargetCurrent represents current price target information.
type PriceTargetCurrent struct {
	Mean         null.Float `json:"mean"`
	Median       null.Float `json:"median"`
	High         null.Float `json:"high"`
	Low          null.Float `json:"low"`
	StandardDev  null.Float `json:"standard_deviation"`
	AnalystCount int        `json:"analyst_count"`
	Currency     string     `json:"currency"`
}

// PriceTargetHistory represents historical price target data.
type PriceTargetHistory struct {
	Date         string     `json:"date"`
	Mean         null.Float `json:"mean"`
	High         null.Float `json:"high"`
	Low          null.Float `json:"low"`
	AnalystCount int        `json:"analyst_count"`
}
