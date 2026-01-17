package response

import "github.com/guregu/null/v6"

// ETFRisk represents the response structure for ETF Risk data.
type ETFRisk struct {
	ETF    ETFRiskData `json:"etf"`
	Status string      `json:"status"`
}

// ETFRiskData contains the risk information for an ETF.
type ETFRiskData struct {
	Risk ETFWorldRisk `json:"risk"`
}

// ETFWorldRisk contains detailed risk analysis information for an ETF.
type ETFWorldRisk struct {
	VolatilityMeasures []ETFVolatilityMeasure `json:"volatility_measures"`
	ValuationMetrics   ETFValuationMetrics    `json:"valuation_metrics"`
}

// ETFVolatilityMeasure represents volatility and risk metrics for a specific period.
type ETFVolatilityMeasure struct {
	Period                   string     `json:"period"`
	Alpha                    null.Float `json:"alpha"`
	AlphaCategory            null.Float `json:"alpha_category"`
	Beta                     null.Float `json:"beta"`
	BetaCategory             null.Float `json:"beta_category"`
	MeanAnnualReturn         null.Float `json:"mean_annual_return"`
	MeanAnnualReturnCategory null.Float `json:"mean_annual_return_category"`
	RSquared                 null.Float `json:"r_squared"`
	RSquaredCategory         null.Float `json:"r_squared_category"`
	Std                      null.Float `json:"std"`
	StdCategory              null.Float `json:"std_category"`
	SharpeRatio              null.Float `json:"sharpe_ratio"`
	SharpeRatioCategory      null.Float `json:"sharpe_ratio_category"`
	TreynorRatio             null.Float `json:"treynor_ratio"`
	TreynorRatioCategory     null.Float `json:"treynor_ratio_category"`
}

// ETFValuationMetrics represents valuation metrics for an ETF.
type ETFValuationMetrics struct {
	PriceToEarnings null.Float `json:"price_to_earnings"`
	PriceToBook     null.Float `json:"price_to_book"`
	PriceToSales    null.Float `json:"price_to_sales"`
	PriceToCashflow null.Float `json:"price_to_cashflow"`
}
