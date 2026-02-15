package response

import "github.com/guregu/null/v6"

// MutualFundRatings represents the response structure for mutual fund ratings data.
type MutualFundRatings struct {
	MutualFund MutualFundRatingsData `json:"mutual_fund"`
	Status     string                `json:"status"`
}

// MutualFundRatingsData contains the ratings information for a mutual fund.
type MutualFundRatingsData struct {
	Ratings MutualFundRatingsInfo `json:"ratings"`
}

// MutualFundRatingsInfo contains ratings information for a mutual fund.
type MutualFundRatingsInfo struct {
	PerformanceRating null.Int `json:"performance_rating"`
	RiskRating        null.Int `json:"risk_rating"`
	ReturnRating      null.Int `json:"return_rating"`
}
