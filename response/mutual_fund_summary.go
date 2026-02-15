package response

import "github.com/guregu/null/v6"

// MutualFundSummary represents the response structure for mutual fund summary data.
type MutualFundSummary struct {
	MutualFund MutualFundSummaryData `json:"mutual_fund"`
	Status     string                `json:"status"`
}

// MutualFundSummaryData contains the summary information for a mutual fund.
type MutualFundSummaryData struct {
	Summary MutualFundSummaryInfo `json:"summary"`
}

// MutualFundSummaryInfo contains detailed summary information for a mutual fund.
type MutualFundSummaryInfo struct {
	Symbol                  string              `json:"symbol"`
	Name                    string              `json:"name"`
	FundFamily              string              `json:"fund_family"`
	FundType                string              `json:"fund_type"`
	Currency                string              `json:"currency"`
	ShareClassInceptionDate string              `json:"share_class_inception_date"`
	YTDReturn               null.Float          `json:"ytd_return"`
	ExpenseRatioNet         null.Float          `json:"expense_ratio_net"`
	Yield                   null.Float          `json:"yield"`
	NAV                     null.Float          `json:"nav"`
	MinInvestment           null.Int            `json:"min_investment"`
	TurnoverRate            null.Float          `json:"turnover_rate"`
	NetAssets               null.Int            `json:"net_assets"`
	Overview                string              `json:"overview"`
	People                  []MutualFundManager `json:"people"`
}

// MutualFundManager represents a mutual fund manager.
type MutualFundManager struct {
	Name        string `json:"name"`
	TenureSince string `json:"tenure_since"`
}
