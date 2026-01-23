package response

import "github.com/guregu/null/v6"

// AnalystRatingsUSEquities represents analyst ratings data for US equities.
type AnalystRatingsUSEquities struct {
	Meta    AnalysisMeta                    `json:"meta"`
	Ratings []AnalystRatingsUSEquitiesEntry `json:"ratings"`
	Status  string                          `json:"status"`
}

// AnalystRatingsUSEquitiesEntry represents a single analyst rating record.
type AnalystRatingsUSEquitiesEntry struct {
	Date               string     `json:"date"`
	Firm               string     `json:"firm"`
	AnalystName        string     `json:"analyst_name"`
	RatingChange       string     `json:"rating_change"`
	RatingCurrent      string     `json:"rating_current"`
	RatingPrior        string     `json:"rating_prior"`
	Time               string     `json:"time"`
	ActionPriceTarget  string     `json:"action_price_target"`
	PriceTargetCurrent null.Float `json:"price_target_current"`
	PriceTargetPrior   null.Float `json:"price_target_prior"`
}
