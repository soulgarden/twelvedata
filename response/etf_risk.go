package response

import "github.com/guregu/null/v6"

// ETFRisk represents the response structure for ETF Risk data.
type ETFRisk struct {
	ETF    ETFRiskData `json:"etf"`
	Status string      `json:"status"`
}

// ETFRiskData contains the risk information for an ETF.
type ETFRiskData struct {
	Risk ETFRiskInfo `json:"risk"`
}

// ETFRiskInfo contains detailed risk analysis information for an ETF.
type ETFRiskInfo struct {
	Symbol            string     `json:"symbol"`
	Name              string     `json:"name"`
	Currency          string     `json:"currency"`
	Exchange          string     `json:"exchange"`
	Country           string     `json:"country"`
	Beta              null.Float `json:"beta"`
	Alpha             null.Float `json:"alpha"`
	StandardDeviation null.Float `json:"standard_deviation"`
	SharpeRatio       null.Float `json:"sharpe_ratio"`
	Volatility        null.Float `json:"volatility"`
	RSquared          null.Float `json:"r_squared"`
	TrackingError     null.Float `json:"tracking_error"`
	InformationRatio  null.Float `json:"information_ratio"`
	UpsideCapture     null.Float `json:"upside_capture"`
	DownsideCapture   null.Float `json:"downside_capture"`
	MaxDrawdown       null.Float `json:"max_drawdown"`
	ValueAtRisk       null.Float `json:"value_at_risk"`
	RiskRating        string     `json:"risk_rating"`
	RiskCategory      string     `json:"risk_category"`
	VolatilityRating  string     `json:"volatility_rating"`
	LastUpdated       string     `json:"last_updated"`
}
