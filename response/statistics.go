package response

import "gopkg.in/guregu/null.v4"

type Statistics struct {
	Meta       StatisticsMeta   `json:"meta"`
	Statistics StatisticsValues `json:"statistics"`
}

type StatisticsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

type StatisticsValues struct {
	ValuationsMetrics  StatisticsValuationsMetrics `json:"valuations_metrics"`
	Financials         StatisticsFinancials        `json:"financials"`
	StockStatistics    StockStatistics             `json:"stock_statistics"`
	StockPriceSummary  StockPriceSummary           `json:"stock_price_summary"`
	DividendsAndSplits DividendsAndSplits          `json:"dividends_and_splits"`
}

type StatisticsValuationsMetrics struct {
	MarketCapitalization null.Int   `json:"market_capitalization"`
	EnterpriseValue      null.Int   `json:"enterprise_value"`
	TrailingPe           null.Float `json:"trailing_pe"`
	ForwardPe            null.Float `json:"forward_pe"`
	PegRatio             null.Float `json:"peg_ratio"`
	PriceToSalesTtm      null.Float `json:"price_to_sales_ttm"`
	PriceToBookMrq       null.Float `json:"price_to_book_mrq"`
	EnterpriseToRevenue  null.Float `json:"enterprise_to_revenue"`
	EnterpriseToEbitda   null.Float `json:"enterprise_to_ebitda"`
}

type StatisticsFinancials struct {
	FiscalYearEnds    string                    `json:"fiscal_year_ends"`
	MostRecentQuarter string                    `json:"most_recent_quarter"`
	ProfitMargin      null.Float                `json:"profit_margin"`
	OperatingMargin   null.Float                `json:"operating_margin"`
	ReturnOnAssetsTtm null.Float                `json:"return_on_assets_ttm"`
	ReturnOnEquityTtm null.Float                `json:"return_on_equity_ttm"`
	IncomeStatement   StatisticsIncomeStatement `json:"income_statement"`
	BalanceSheet      StatisticsBalanceSheet    `json:"balance_sheet"`
	CashFlow          StatisticsCashFlow        `json:"cash_flow"`
}

type StatisticsIncomeStatement struct {
	RevenueTtm                 null.Int   `json:"revenue_ttm"`
	RevenuePerShareTtm         null.Float `json:"revenue_per_share_ttm"`
	QuarterlyRevenueGrowth     null.Float `json:"quarterly_revenue_growth"`
	GrossProfitTtm             null.Int   `json:"gross_profit_ttm"`
	Ebitda                     null.Int   `json:"ebitda"`
	NetIncomeToCommonTtm       null.Int   `json:"net_income_to_common_ttm"`
	DilutedEpsTtm              null.Float `json:"diluted_eps_ttm"`
	QuarterlyEarningsGrowthYoy null.Float `json:"quarterly_earnings_growth_yoy"`
}

type StatisticsBalanceSheet struct {
	RevenueTtm           null.Int   `json:"revenue_ttm"`
	TotalCashMrq         null.Int   `json:"total_cash_mrq"`
	TotalCashPerShareMrq null.Float `json:"total_cash_per_share_mrq"`
	TotalDebtMrq         null.Int   `json:"total_debt_mrq"`
	TotalDebtToEquityMrq null.Float `json:"total_debt_to_equity_mrq"`
	CurrentRatioMrq      null.Float `json:"current_ratio_mrq"`
	BookValuePerShareMrq null.Float `json:"book_value_per_share_mrq"`
}

type StatisticsCashFlow struct {
	OperatingCashFlowTtm   null.Int `json:"operating_cash_flow_ttm"`
	LeveredFreeCashFlowTtm null.Int `json:"levered_free_cash_flow_ttm"`
}

type StockStatistics struct {
	SharesOutstanding               null.Int   `json:"shares_outstanding"`
	FloatShares                     null.Int   `json:"float_shares"`
	Avg10Volume                     null.Int   `json:"avg_10_volume"`
	Avg30Volume                     null.Int   `json:"avg_30_volume"`
	SharesShort                     null.Int   `json:"shares_short"`
	ShortRatio                      null.Float `json:"short_ratio"`
	ShortPercentOfSharesOutstanding null.Float `json:"short_percent_of_shares_outstanding"`
	PercentHeldByInsiders           null.Float `json:"percent_held_by_insiders"`
	PercentHeldByInstitutions       null.Float `json:"percent_held_by_institutions"`
}

type StockPriceSummary struct {
	FiftyTwoWeekLow    null.Float `json:"fifty_two_week_low"`
	FiftyTwoWeekHigh   null.Float `json:"fifty_two_week_high"`
	FiftyTwoWeekChange null.Float `json:"fifty_two_week_change"`
	Beta               null.Float `json:"beta"`
	Day50Ma            null.Float `json:"day_50_ma"`
	Day200Ma           null.Float `json:"day_200_ma"`
}

type DividendsAndSplits struct {
	ForwardAnnualDividendRate   null.Float `json:"forward_annual_dividend_rate"`
	ForwardAnnualDividendYield  null.Float `json:"forward_annual_dividend_yield"`
	TrailingAnnualDividendRate  null.Float `json:"trailing_annual_dividend_rate"`
	TrailingAnnualDividendYield null.Float `json:"trailing_annual_dividend_yield"`
	YearAverageDividendYield    null.Float `json:"5_year_average_dividend_yield"`
	PayoutRatio                 null.Float `json:"payout_ratio"`
	DividendDate                string     `json:"dividend_date"`
	ExDividendDate              string     `json:"ex_dividend_date"`
	LastSplitFactor             string     `json:"last_split_factor"`
	LastSplitDate               string     `json:"last_split_date"`
}
