package response

import "github.com/guregu/null/v6"

// MutualFundRatingsResponse represents the response structure for Mutual Fund Ratings data.
type MutualFundRatingsResponse struct {
	MutualFund MutualFundRatingsData `json:"mutual_fund"`
	Status     string                `json:"status"`
}

// MutualFundRatingsData contains the ratings information for a mutual fund.
type MutualFundRatingsData struct {
	Ratings MutualFundRatingsInfo `json:"ratings"`
}

// MutualFundRatingsInfo contains detailed ratings information for a mutual fund.
type MutualFundRatingsInfo struct {
	OverallRating        null.String                `json:"overall_rating"`
	MorningstarRating    null.String                `json:"morningstar_rating"`
	LipperRating         null.String                `json:"lipper_rating"`
	TwelveDataRating     null.String                `json:"twelve_data_rating"`
	PerformanceRating    null.Float                 `json:"performance_rating"`
	RiskRating           null.Float                 `json:"risk_rating"`
	ExpenseRating        null.String                `json:"expense_rating"`
	SustainabilityRating null.String                `json:"sustainability_rating"`
	CategoryRankings     MutualFundCategoryRankings `json:"category_rankings"`
	RatingMethodologies  MutualFundRatingMethods    `json:"rating_methodologies"`
	LastUpdated          null.String                `json:"last_updated"`
}

// MutualFundCategoryRankings represents category-specific rankings.
type MutualFundCategoryRankings struct {
	CategoryName        null.String `json:"category_name"`
	CategoryRank        null.Int    `json:"category_rank"`
	CategorySize        null.Int    `json:"category_size"`
	Percentile          null.Float  `json:"percentile"`
	QuartileRank        null.Int    `json:"quartile_rank"`
	MorningstarCategory null.String `json:"morningstar_category"`
	LipperCategory      null.String `json:"lipper_category"`
}

// MutualFundRatingMethods represents rating methodologies used by different agencies.
type MutualFundRatingMethods struct {
	MorningstarMethod null.String `json:"morningstar_method"`
	LipperMethod      null.String `json:"lipper_method"`
	TwelveDataMethod  null.String `json:"twelve_data_method"`
	RatingScale       null.String `json:"rating_scale"`
	LastReviewDate    null.String `json:"last_review_date"`
}

// MutualFundRatings represents ratings for a mutual fund (preserved for compatibility).
type MutualFundRatings struct {
	OverallRating        null.String `json:"overall_rating"`
	MorningstarRating    null.String `json:"morningstar_rating"`
	LipperRating         null.String `json:"lipper_rating"`
	TwelveDataRating     null.String `json:"twelve_data_rating"`
	PerformanceRating    null.Float  `json:"performance_rating"`
	RiskRating           null.Float  `json:"risk_rating"`
	ExpenseRating        null.String `json:"expense_rating"`
	SustainabilityRating null.String `json:"sustainability_rating"`
	LastUpdated          null.String `json:"last_updated"`
}
