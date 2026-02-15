package response

import "github.com/guregu/null/v6"

// MutualFundSustainability represents the response structure for mutual fund sustainability endpoint.
type MutualFundSustainability struct {
	MutualFund MutualFundSustainabilityData `json:"mutual_fund"`
	Status     string                       `json:"status"`
}

// MutualFundSustainabilityData contains the sustainability information for a mutual fund.
type MutualFundSustainabilityData struct {
	Sustainability MutualFundSustainabilityDetails `json:"sustainability"`
}

// MutualFundSustainabilityDetails represents sustainability metrics for a mutual fund.
type MutualFundSustainabilityDetails struct {
	Score                 null.Int             `json:"score"`
	CorporateESGPillars   MutualFundESGPillars `json:"corporate_esg_pillars"`
	SustainableInvestment null.Bool            `json:"sustainable_investment"`
	CorporateAUM          null.Float           `json:"corporate_aum"`
}

// MutualFundESGPillars represents ESG pillar scores.
type MutualFundESGPillars struct {
	Environmental null.Float `json:"environmental"`
	Social        null.Float `json:"social"`
	Governance    null.Float `json:"governance"`
}
