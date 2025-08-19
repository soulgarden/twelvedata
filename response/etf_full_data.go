package response

import "github.com/guregu/null/v6"

// ETFFullData represents the response structure for ETF full data.
type ETFFullData struct {
	Summary     ETFSummary      `json:"summary"`
	Performance ETFPerformance  `json:"performance"`
	Risk        ETFFullDataRisk `json:"risk"`
	Composition ETFComposition  `json:"composition"`
}

// ETFFullDataRisk represents the risk data in ETF full data response (simpler structure than standalone risk endpoint).
type ETFFullDataRisk struct {
	Beta              null.Float `json:"beta"`
	Alpha             null.Float `json:"alpha"`
	StandardDeviation null.Float `json:"standard_deviation"`
	SharpeRatio       null.Float `json:"sharpe_ratio"`
	Volatility        null.Float `json:"volatility"`
	RSquared          null.Float `json:"r_squared"`
}

// ETFSummary represents summary information about an ETF.
type ETFSummary struct {
	Symbol        string      `json:"symbol"`
	Name          string      `json:"name"`
	Currency      string      `json:"currency"`
	Exchange      string      `json:"exchange"`
	Country       string      `json:"country"`
	AssetClass    string      `json:"asset_class"`
	NetAssets     null.Float  `json:"net_assets"`
	ExpenseRatio  null.Float  `json:"expense_ratio"`
	InceptionDate null.String `json:"inception_date"`
	LastUpdated   null.String `json:"last_updated"`
}
