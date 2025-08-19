package response

import "github.com/guregu/null/v6"

// MutualFundFullData represents the response structure for mutual fund full data.
type MutualFundFullData struct {
	Summary        MutualFundSummary        `json:"summary"`
	Performance    MutualFundPerformance    `json:"performance"`
	Risk           MutualFundRisk           `json:"risk"`
	Ratings        MutualFundRatings        `json:"ratings"`
	Composition    MutualFundComposition    `json:"composition"`
	PurchaseInfo   MutualFundPurchaseInfo   `json:"purchase_info"`
	Sustainability MutualFundSustainability `json:"sustainability"`
}

// MutualFundSummary represents summary information about a mutual fund.
type MutualFundSummary struct {
	Symbol        string      `json:"symbol"`
	Name          string      `json:"name"`
	Currency      string      `json:"currency"`
	Exchange      string      `json:"exchange"`
	Country       string      `json:"country"`
	AssetClass    string      `json:"asset_class"`
	Category      string      `json:"category"`
	FundFamily    string      `json:"fund_family"`
	NetAssets     null.Float  `json:"net_assets"`
	ExpenseRatio  null.Float  `json:"expense_ratio"`
	InceptionDate null.String `json:"inception_date"`
	LastUpdated   null.String `json:"last_updated"`
}
