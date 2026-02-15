package response

import "github.com/guregu/null/v6"

// MutualFundRisk represents the response structure for mutual fund risk data.
type MutualFundRisk struct {
	MutualFund MutualFundRiskData `json:"mutual_fund"`
	Status     string             `json:"status"`
}

// MutualFundRiskData contains the risk analysis information for a mutual fund.
type MutualFundRiskData struct {
	Risk MutualFundRiskInfo `json:"risk"`
}

// MutualFundRiskInfo contains detailed risk metrics for a mutual fund.
type MutualFundRiskInfo struct {
	VolatilityMeasures []MutualFundVolatilityMeasure `json:"volatility_measures"`
	ValuationMetrics   MutualFundValuationMetrics    `json:"valuation_metrics"`
}

// MutualFundVolatilityMeasure represents volatility statistics of a fund.
type MutualFundVolatilityMeasure struct {
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

// MutualFundValuationMetrics represents valuation ratios and metrics of the fund and its category.
type MutualFundValuationMetrics struct {
	PriceToEarnings                    null.Float `json:"price_to_earnings"`
	PriceToEarningsCategory            null.Float `json:"price_to_earnings_category"`
	PriceToBook                        null.Float `json:"price_to_book"`
	PriceToBookCategory                null.Float `json:"price_to_book_category"`
	PriceToSales                       null.Float `json:"price_to_sales"`
	PriceToSalesCategory               null.Float `json:"price_to_sales_category"`
	PriceToCashflow                    null.Float `json:"price_to_cashflow"`
	PriceToCashflowCategory            null.Float `json:"price_to_cashflow_category"`
	MedianMarketCapitalization         null.Int   `json:"median_market_capitalization"`
	MedianMarketCapitalizationCategory null.Int   `json:"median_market_capitalization_category"`
	ThreeYearEarningsGrowth            null.Float `json:"3_year_earnings_growth"`
	ThreeYearEarningsGrowthCategory    null.Float `json:"3_year_earnings_growths_category"`
}
