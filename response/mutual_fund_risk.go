package response

import "github.com/guregu/null/v6"

// MutualFundRisk represents risk metrics for a mutual fund.
type MutualFundRisk struct {
	Beta              null.Float `json:"beta"`
	StandardDeviation null.Float `json:"standard_deviation"`
	SharpeRatio       null.Float `json:"sharpe_ratio"`
	Alpha             null.Float `json:"alpha"`
	RSquared          null.Float `json:"r_squared"`
	MaxDrawdown       null.Float `json:"max_drawdown"`
	TrackingError     null.Float `json:"tracking_error"`
	InformationRatio  null.Float `json:"information_ratio"`
	Volatility        null.Float `json:"volatility"`
}

// MutualFundRiskResponse represents the response structure for Mutual Fund Risk data.
type MutualFundRiskResponse struct {
	MutualFund MutualFundRiskData `json:"mutual_fund"`
	Status     string             `json:"status"`
}

// MutualFundRiskData contains the risk analysis information for a mutual fund.
type MutualFundRiskData struct {
	Risk MutualFundRiskAnalysis `json:"risk"`
}

// MutualFundRiskAnalysis contains detailed risk metrics for a mutual fund.
type MutualFundRiskAnalysis struct {
	Beta              null.Float `json:"beta"`
	Alpha             null.Float `json:"alpha"`
	StandardDeviation null.Float `json:"standard_deviation"`
	SharpeRatio       null.Float `json:"sharpe_ratio"`
	RSquared          null.Float `json:"r_squared"`
	MaxDrawdown       null.Float `json:"max_drawdown"`
	TrackingError     null.Float `json:"tracking_error"`
	InformationRatio  null.Float `json:"information_ratio"`
	Volatility        null.Float `json:"volatility"`
	TreynorRatio      null.Float `json:"treynor_ratio"`
	SortinoRatio      null.Float `json:"sortino_ratio"`
	UpsideCapture     null.Float `json:"upside_capture"`
	DownsideCapture   null.Float `json:"downside_capture"`
	RiskCategory      string     `json:"risk_category"`
	RiskRating        null.Int   `json:"risk_rating"`
	BenchmarkName     string     `json:"benchmark_name"`
	BenchmarkSymbol   string     `json:"benchmark_symbol"`
}
