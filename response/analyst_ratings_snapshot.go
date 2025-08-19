package response

import "github.com/guregu/null/v6"

// AnalystRatingsSnapshot represents analyst ratings snapshot data.
type AnalystRatingsSnapshot struct {
	Symbol      string                          `json:"symbol"`
	Exchange    string                          `json:"exchange"`
	Current     AnalystRatingsSnapshotCurrent   `json:"current"`
	Summary     AnalystRatingsSnapshotSummary   `json:"summary"`
	Breakdown   AnalystRatingsSnapshotBreakdown `json:"breakdown"`
	Changes     AnalystRatingsSnapshotChanges   `json:"changes"`
	LastUpdated string                          `json:"last_updated"`
}

// AnalystRatingsSnapshotCurrent represents current analyst ratings snapshot.
type AnalystRatingsSnapshotCurrent struct {
	AverageScore   null.Float `json:"average_score"`
	Recommendation string     `json:"recommendation"`
	TotalAnalysts  int        `json:"total_analysts"`
	RatingScale    string     `json:"rating_scale"`
	Currency       string     `json:"currency"`
	LastRatingDate string     `json:"last_rating_date"`
}

// AnalystRatingsSnapshotSummary represents analyst ratings summary.
type AnalystRatingsSnapshotSummary struct {
	Mean           null.Float `json:"mean"`
	Median         null.Float `json:"median"`
	Mode           null.Float `json:"mode"`
	StandardDev    null.Float `json:"standard_deviation"`
	BullishPercent null.Float `json:"bullish_percent"`
	BearishPercent null.Float `json:"bearish_percent"`
	NeutralPercent null.Float `json:"neutral_percent"`
}

// AnalystRatingsSnapshotBreakdown represents detailed breakdown of analyst ratings.
type AnalystRatingsSnapshotBreakdown struct {
	StrongBuy     int        `json:"strong_buy"`
	Buy           int        `json:"buy"`
	Hold          int        `json:"hold"`
	Sell          int        `json:"sell"`
	StrongSell    int        `json:"strong_sell"`
	StrongBuyPct  null.Float `json:"strong_buy_pct"`
	BuyPct        null.Float `json:"buy_pct"`
	HoldPct       null.Float `json:"hold_pct"`
	SellPct       null.Float `json:"sell_pct"`
	StrongSellPct null.Float `json:"strong_sell_pct"`
}

// AnalystRatingsSnapshotChanges represents recent changes in analyst ratings.
type AnalystRatingsSnapshotChanges struct {
	Last7Days  AnalystRatingsChangesPeriod `json:"last_7_days"`
	Last30Days AnalystRatingsChangesPeriod `json:"last_30_days"`
	Last90Days AnalystRatingsChangesPeriod `json:"last_90_days"`
	LastYear   AnalystRatingsChangesPeriod `json:"last_year"`
}

// AnalystRatingsChangesPeriod represents analyst rating changes for a specific period.
type AnalystRatingsChangesPeriod struct {
	Upgrades     int        `json:"upgrades"`
	Downgrades   int        `json:"downgrades"`
	Initiations  int        `json:"initiations"`
	Reiterations int        `json:"reiterations"`
	NetChange    null.Float `json:"net_change"`
	TotalChanges int        `json:"total_changes"`
}
