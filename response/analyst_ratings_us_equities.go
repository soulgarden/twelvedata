package response

import "github.com/guregu/null/v6"

// AnalystRatingsUSEquities represents analyst ratings data for US equities.
type AnalystRatingsUSEquities struct {
	Symbol      string                            `json:"symbol"`
	Exchange    string                            `json:"exchange"`
	Overview    AnalystRatingsUSEquitiesOverview  `json:"overview"`
	History     []AnalystRatingsUSEquitiesHistory `json:"history"`
	Firms       []AnalystRatingsUSEquitiesFirm    `json:"firms"`
	Consensus   AnalystRatingsUSEquitiesConsensus `json:"consensus"`
	LastUpdated string                            `json:"last_updated"`
}

// AnalystRatingsUSEquitiesOverview represents overview of US equity analyst ratings.
type AnalystRatingsUSEquitiesOverview struct {
	TotalRatings   int        `json:"total_ratings"`
	AverageRating  null.Float `json:"average_rating"`
	Recommendation string     `json:"recommendation"`
	LastRatingDate string     `json:"last_rating_date"`
	RatingScale    string     `json:"rating_scale"`
	Currency       string     `json:"currency"`
	MarketCap      null.Float `json:"market_cap"`
	Sector         string     `json:"sector"`
	Industry       string     `json:"industry"`
}

// AnalystRatingsUSEquitiesHistory represents historical analyst ratings data.
type AnalystRatingsUSEquitiesHistory struct {
	Date         string     `json:"date"`
	Firm         string     `json:"firm"`
	Analyst      string     `json:"analyst"`
	PriorRating  string     `json:"prior_rating"`
	NewRating    string     `json:"new_rating"`
	PriorTarget  null.Float `json:"prior_target"`
	NewTarget    null.Float `json:"new_target"`
	ActionType   string     `json:"action_type"`
	RatingChange string     `json:"rating_change"`
	TargetChange null.Float `json:"target_change"`
}

// AnalystRatingsUSEquitiesFirm represents analyst firm ratings data.
type AnalystRatingsUSEquitiesFirm struct {
	FirmName      string     `json:"firm_name"`
	AnalystCount  int        `json:"analyst_count"`
	CurrentRating string     `json:"current_rating"`
	PriceTarget   null.Float `json:"price_target"`
	LastUpdate    string     `json:"last_update"`
	AccuracyScore null.Float `json:"accuracy_score"`
	SuccessRate   null.Float `json:"success_rate"`
	AverageReturn null.Float `json:"average_return"`
}

// AnalystRatingsUSEquitiesConsensus represents consensus analyst ratings for US equities.
type AnalystRatingsUSEquitiesConsensus struct {
	StrongBuy    int        `json:"strong_buy"`
	Buy          int        `json:"buy"`
	Hold         int        `json:"hold"`
	Sell         int        `json:"sell"`
	StrongSell   int        `json:"strong_sell"`
	Mean         null.Float `json:"mean"`
	PriceTarget  null.Float `json:"price_target"`
	TargetHigh   null.Float `json:"target_high"`
	TargetLow    null.Float `json:"target_low"`
	TargetMedian null.Float `json:"target_median"`
	TargetMean   null.Float `json:"target_mean"`
	AnalystCount int        `json:"analyst_count"`
}
