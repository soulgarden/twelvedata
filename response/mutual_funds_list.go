package response

import "github.com/guregu/null/v6"

// MutualFundsDirectory represents the response structure for the mutual funds directory endpoint.
type MutualFundsDirectory struct {
	Result MutualFundsDirectoryResult `json:"result"`
	Status string                     `json:"status"`
}

// MutualFundsDirectoryResult contains the mutual funds directory list and count.
type MutualFundsDirectoryResult struct {
	Count int                        `json:"count"`
	List  []MutualFundsDirectoryFund `json:"list"`
}

// MutualFundsDirectoryFund represents a single mutual fund in the directory list.
type MutualFundsDirectoryFund struct {
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	Country           string   `json:"country"`
	FundFamily        string   `json:"fund_family"`
	FundType          string   `json:"fund_type"`
	PerformanceRating null.Int `json:"performance_rating"`
	RiskRating        null.Int `json:"risk_rating"`
	Currency          string   `json:"currency"`
	Exchange          string   `json:"exchange"`
	MicCode           string   `json:"mic_code"`
}
