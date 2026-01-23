package request

// GetMutualFundsDirectory represents request parameters for mutual funds directory data.
type GetMutualFundsDirectory struct {
	APIKey
	Symbol            string `schema:"symbol,omitempty"`
	FIGI              string `schema:"figi,omitempty"`
	ISIN              string `schema:"isin,omitempty"`
	CUSIP             string `schema:"cusip,omitempty"`
	CIK               string `schema:"cik,omitempty"`
	Country           string `schema:"country,omitempty"`
	FundFamily        string `schema:"fund_family,omitempty"`
	FundType          string `schema:"fund_type,omitempty"`
	PerformanceRating int    `schema:"performance_rating,omitempty"`
	RiskRating        int    `schema:"risk_rating,omitempty"`
	Page              int    `schema:"page,omitempty"`
	OutputSize        int    `schema:"outputsize,omitempty"`
}
