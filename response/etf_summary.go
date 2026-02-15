package response

import "github.com/guregu/null/v6"

// ETFWorldSummary represents the response structure for ETF World Summary data.
type ETFWorldSummary struct {
	ETF    ETFWorldSummaryData `json:"etf"`
	Status string              `json:"status"`
}

// ETFWorldSummaryData contains the summary information for an ETF.
type ETFWorldSummaryData struct {
	Summary ETFWorldSummaryInfo `json:"summary"`
}

// ETFWorldSummaryInfo contains detailed summary information for an ETF.
type ETFWorldSummaryInfo struct {
	Symbol                  string     `json:"symbol"`
	Name                    string     `json:"name"`
	FundFamily              string     `json:"fund_family"`
	FundType                string     `json:"fund_type"`
	Currency                string     `json:"currency"`
	ShareClassInceptionDate string     `json:"share_class_inception_date"`
	YTDReturn               null.Float `json:"ytd_return"`
	ExpenseRatioNet         null.Float `json:"expense_ratio_net"`
	Yield                   null.Float `json:"yield"`
	NAV                     null.Float `json:"nav"`
	LastPrice               null.Float `json:"last_price"`
	TurnoverRate            null.Float `json:"turnover_rate"`
	NetAssets               null.Int   `json:"net_assets"`
	Overview                string     `json:"overview"`
}
