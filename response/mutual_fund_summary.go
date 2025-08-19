package response

import "github.com/guregu/null/v6"

// MutualFundSummaryResponse represents the response structure for Mutual Fund Summary data.
type MutualFundSummaryResponse struct {
	MutualFund MutualFundSummaryData `json:"mutual_fund"`
	Status     string                `json:"status"`
}

// MutualFundSummaryData contains the summary information for a mutual fund.
type MutualFundSummaryData struct {
	Summary MutualFundSummaryInfo `json:"summary"`
}

// MutualFundSummaryInfo contains detailed summary information for a mutual fund.
type MutualFundSummaryInfo struct {
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
	YTDReturn     null.Float  `json:"ytd_return"`
	Yield         null.Float  `json:"yield"`
	NAV           null.Float  `json:"nav"`
	LastPrice     null.Float  `json:"last_price"`
	TurnoverRate  null.Float  `json:"turnover_rate"`
	Overview      string      `json:"overview"`
}
