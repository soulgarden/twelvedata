package response

type Statistics struct {
	Meta       *StatisticsMeta   `json:"meta"`
	Statistics *StatisticsValues `json:"statistics"`
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
	ValuationsMetrics  *StatisticsValuationsMetrics `json:"valuations_metrics"`
	Financials         *StatisticsFinancials        `json:"financials"`
	StockStatistics    *StockStatistics             `json:"stock_statistics"`
	StockPriceSummary  *StockPriceSummary           `json:"stock_price_summary"`
	DividendsAndSplits *DividendsAndSplits          `json:"dividends_and_splits"`
}

type StatisticsValuationsMetrics struct {
	MarketCapitalization int64   `json:"market_capitalization"`
	EnterpriseValue      int64   `json:"enterprise_value"`
	TrailingPe           float64 `json:"trailing_pe"`
	ForwardPe            float64 `json:"forward_pe"`
	PegRatio             float64 `json:"peg_ratio"`
	PriceToSalesTtm      float64 `json:"price_to_sales_ttm"`
	PriceToBookMrq       float64 `json:"price_to_book_mrq"`
	EnterpriseToRevenue  float64 `json:"enterprise_to_revenue"`
	EnterpriseToEbitda   float64 `json:"enterprise_to_ebitda"`
}

type StatisticsFinancials struct {
	FiscalYearEnds    string                     `json:"fiscal_year_ends"`
	MostRecentQuarter string                     `json:"most_recent_quarter"`
	ProfitMargin      float64                    `json:"profit_margin"`
	OperatingMargin   float64                    `json:"operating_margin"`
	ReturnOnAssetsTtm float64                    `json:"return_on_assets_ttm"`
	ReturnOnEquityTtm float64                    `json:"return_on_equity_ttm"`
	IncomeStatement   *StatisticsIncomeStatement `json:"income_statement"`
	BalanceSheet      *StatisticsBalanceSheet    `json:"balance_sheet"`
	CashFlow          *StatisticsCashFlow        `json:"cash_flow"`
}

type StatisticsIncomeStatement struct {
	RevenueTtm                 int64   `json:"revenue_ttm"`
	RevenuePerShareTtm         float64 `json:"revenue_per_share_ttm"`
	QuarterlyRevenueGrowth     float64 `json:"quarterly_revenue_growth"`
	GrossProfitTtm             int64   `json:"gross_profit_ttm"`
	Ebitda                     int64   `json:"ebitda"`
	NetIncomeToCommonTtm       int64   `json:"net_income_to_common_ttm"`
	DilutedEpsTtm              float64 `json:"diluted_eps_ttm"`
	QuarterlyEarningsGrowthYoy float64 `json:"quarterly_earnings_growth_yoy"`
}

type StatisticsBalanceSheet struct {
	RevenueTtm           int64   `json:"revenue_ttm"`
	TotalCashMrq         int64   `json:"total_cash_mrq"`
	TotalCashPerShareMrq float64 `json:"total_cash_per_share_mrq"`
	TotalDebtMrq         int64   `json:"total_debt_mrq"`
	TotalDebtToEquityMrq float64 `json:"total_debt_to_equity_mrq"`
	CurrentRatioMrq      float64 `json:"current_ratio_mrq"`
	BookValuePerShareMrq float64 `json:"book_value_per_share_mrq"`
}

type StatisticsCashFlow struct {
	OperatingCashFlowTtm   int64 `json:"operating_cash_flow_ttm"`
	LeveredFreeCashFlowTtm int64 `json:"levered_free_cash_flow_ttm"`
}

type StockStatistics struct {
	SharesOutstanding               int64   `json:"shares_outstanding"`
	FloatShares                     int64   `json:"float_shares"`
	Avg10Volume                     int     `json:"avg_10_volume"`
	Avg30Volume                     int     `json:"avg_30_volume"`
	SharesShort                     int     `json:"shares_short"`
	ShortRatio                      float64 `json:"short_ratio"`
	ShortPercentOfSharesOutstanding float64 `json:"short_percent_of_shares_outstanding"`
	PercentHeldByInsiders           float64 `json:"percent_held_by_insiders"`
	PercentHeldByInstitutions       float64 `json:"percent_held_by_institutions"`
}

type StockPriceSummary struct {
	FiftyTwoWeekLow    float64 `json:"fifty_two_week_low"`
	FiftyTwoWeekHigh   float64 `json:"fifty_two_week_high"`
	FiftyTwoWeekChange float64 `json:"fifty_two_week_change"`
	Beta               float64 `json:"beta"`
	Day50Ma            float64 `json:"day_50_ma"`
	Day200Ma           float64 `json:"day_200_ma"`
}

type DividendsAndSplits struct {
	ForwardAnnualDividendRate   float64 `json:"forward_annual_dividend_rate"`
	ForwardAnnualDividendYield  float64 `json:"forward_annual_dividend_yield"`
	TrailingAnnualDividendRate  float64 `json:"trailing_annual_dividend_rate"`
	TrailingAnnualDividendYield float64 `json:"trailing_annual_dividend_yield"`
	YearAverageDividendYield    float64 `json:"5_year_average_dividend_yield"`
	PayoutRatio                 float64 `json:"payout_ratio"`
	DividendDate                string  `json:"dividend_date"`
	ExDividendDate              string  `json:"ex_dividend_date"`
	LastSplitFactor             string  `json:"last_split_factor"`
	LastSplitDate               string  `json:"last_split_date"`
}
