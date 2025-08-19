package response

import "github.com/guregu/null/v6"

// MutualFundSustainability represents the response structure for mutual fund sustainability endpoint.
type MutualFundSustainability struct {
	Meta   MutualFundSustainabilityMeta `json:"meta"`
	Data   MutualFundSustainabilityData `json:"data"`
	Status string                       `json:"status"`
}

// MutualFundSustainabilityMeta contains metadata for mutual fund sustainability.
type MutualFundSustainabilityMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// MutualFundSustainabilityData represents sustainability metrics for a mutual fund.
type MutualFundSustainabilityData struct {
	SustainabilityRank    null.Int    `json:"sustainability_rank"`
	SustainabilityScore   null.Float  `json:"sustainability_score"`
	EnvironmentalScore    null.Float  `json:"environmental_score"`
	SocialScore           null.Float  `json:"social_score"`
	GovernanceScore       null.Float  `json:"governance_score"`
	CarbonIntensity       null.Float  `json:"carbon_intensity"`
	FossilFuelInvolvement null.Float  `json:"fossil_fuel_involvement"`
	ESGQualityScore       null.Float  `json:"esg_quality_score"`
	SustainableInvestment null.String `json:"sustainable_investment"`
	LastUpdated           null.String `json:"last_updated"`
}

// MutualFundSustainabilityCompat represents legacy compatibility for existing sustainability data structure.
// Deprecated: Use MutualFundSustainabilityData instead.
type MutualFundSustainabilityCompat = MutualFundSustainabilityData
