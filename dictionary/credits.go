// Package dictionary contains constant definitions used throughout the Twelve Data API client,
// including API credit costs for various endpoints and other shared constants.
package dictionary

const (
	// TimeSeries represents the API credit cost for time series data requests.
	// Market Data endpoints.
	TimeSeries = 1
	// TimeSeriesCross represents the API credit cost for cross-currency time series requests.
	TimeSeriesCross = 5
	// Quote represents the API credit cost for quote data requests.
	Quote = 1
	// Price represents the API credit cost for price data requests.
	Price = 1
	// EOD represents the API credit cost for end-of-day data requests.
	EOD = 1
	// MarketMovers represents the API credit cost for market movers requests.
	MarketMovers = 100

	// Stocks represents the API credit cost for stocks catalog requests.
	// Reference Data - Asset Catalogs.
	Stocks = 1
	// ForexPairs represents the API credit cost for forex pairs catalog requests.
	ForexPairs = 1
	// Cryptocurrencies represents the API credit cost for cryptocurrencies catalog requests.
	Cryptocurrencies = 1
	// ETFs represents the API credit cost for ETFs catalog requests.
	ETFs = 1
	// ETFsDirectory represents the API credit cost for ETFs directory requests.
	ETFsDirectory = 1
	// Funds represents the API credit cost for funds catalog requests.
	Funds = 1
	// MutualFunds represents the API credit cost for mutual funds directory requests.
	MutualFunds = 1
	// Commodities represents the API credit cost for commodities catalog requests.
	Commodities = 1
	// Bonds represents the API credit cost for bonds catalog requests.
	Bonds = 1

	// SymbolSearch represents the API credit cost for symbol search requests.
	// Reference Data - Discovery.
	SymbolSearch = 1
	// CrossListings represents the API credit cost for cross-listings requests.
	CrossListings = 40
	// EarliestTimestamp represents the API credit cost for earliest timestamp requests.
	EarliestTimestamp = 1

	// Exchanges represents the API credit cost for exchanges list requests.
	// Reference Data - Markets.
	Exchanges = 1
	// ExchangesSchedule represents the API credit cost for exchanges schedule requests.
	ExchangesSchedule = 100
	// CryptocurrencyExchanges represents the API credit cost for cryptocurrency exchanges requests.
	CryptocurrencyExchanges = 1
	// MarketState represents the API credit cost for market state requests.
	MarketState = 1

	// Countries represents the API credit cost for countries list requests.
	// Reference Data - Supporting Metadata.
	Countries = 1
	// InstrumentType represents the API credit cost for instrument type requests.
	InstrumentType = 1
	// TechnicalIndicatorsList represents the API credit cost for technical indicators list requests.
	TechnicalIndicatorsList = 1

	// Logo represents the API credit cost for company logo requests.
	// Fundamentals.
	Logo = 1
	// Profile represents the API credit cost for company profile requests.
	Profile = 10
	// Dividends represents the API credit cost for dividends data requests.
	Dividends = 20
	// DividendsCalendar represents the API credit cost for dividends calendar requests.
	DividendsCalendar = 40
	// Splits represents the API credit cost for stock splits data requests.
	Splits = 20
	// SplitsCalendar represents the API credit cost for stock splits calendar requests.
	SplitsCalendar = 40
	// Earnings represents the API credit cost for earnings data requests.
	Earnings = 20
	// EarningsCalendar represents the API credit cost for earnings calendar requests.
	EarningsCalendar = 40
	// IPOCalendar represents the API credit cost for IPO calendar requests.
	IPOCalendar = 40
	// Statistics represents the API credit cost for company statistics requests.
	Statistics = 50
	// PressReleases represents the API credit cost for press releases requests.
	PressReleases = 50
	// IncomeStatement represents the API credit cost for income statement requests.
	IncomeStatement = 100
	// IncomeStatementConsolidated represents the API credit cost for consolidated income statement requests.
	IncomeStatementConsolidated = 100
	// BalanceSheet represents the API credit cost for balance sheet requests.
	BalanceSheet = 100
	// BalanceSheetConsolidated represents the API credit cost for consolidated balance sheet requests.
	BalanceSheetConsolidated = 100
	// CashFlow represents the API credit cost for cash flow statement requests.
	CashFlow = 100
	// CashFlowConsolidated represents the API credit cost for consolidated cash flow statement requests.
	CashFlowConsolidated = 100
	// OptionsExpiration represents the API credit cost for options expiration requests.
	OptionsExpiration = 50
	// OptionsChain represents the API credit cost for options chain requests.
	OptionsChain = 100
	// KeyExecutives represents the API credit cost for key executives requests.
	KeyExecutives = 1000
	// MarketCapitalization represents the API credit cost for market capitalization requests.
	MarketCapitalization = 5
	// LastChanges represents the API credit cost for last changes requests.
	LastChanges = 50

	// ExchangeRate represents the API credit cost for exchange rate requests.
	// Currencies.
	ExchangeRate = 1
	// CurrencyConversion represents the API credit cost for currency conversion requests.
	CurrencyConversion = 1

	// ETFFullData represents the API credit cost for ETF full data requests.
	// ETFs.
	ETFFullData = 800
	// ETFSummary represents the API credit cost for ETF summary requests.
	ETFSummary = 50
	// ETFPerformance represents the API credit cost for ETF performance requests.
	ETFPerformance = 200
	// ETFRisk represents the API credit cost for ETF risk data requests.
	ETFRisk = 50
	// ETFComposition represents the API credit cost for ETF composition requests.
	ETFComposition = 100
	// ETFsFamilies represents the API credit cost for ETF families requests.
	ETFsFamilies = 10
	// ETFsTypes represents the API credit cost for ETF types requests.
	ETFsTypes = 10

	// MFsDirectory represents the API credit cost for mutual funds directory requests.
	// Mutual Funds.
	MFsDirectory = 50
	// MFFullData represents the API credit cost for mutual fund full data requests.
	MFFullData = 1000
	// MFSummary represents the API credit cost for mutual fund summary requests.
	MFSummary = 50
	// MFPerformance represents the API credit cost for mutual fund performance requests.
	MFPerformance = 200
	// MFRisk represents the API credit cost for mutual fund risk data requests.
	MFRisk = 50
	// MFRatings represents the API credit cost for mutual fund ratings requests.
	MFRatings = 50
	// MFComposition represents the API credit cost for mutual fund composition requests.
	MFComposition = 200
	// MFPurchaseInfo represents the API credit cost for mutual fund purchase info requests.
	MFPurchaseInfo = 50
	// MFSustainability represents the API credit cost for mutual fund sustainability requests.
	MFSustainability = 50
	// MFsFamilies represents the API credit cost for mutual fund families requests.
	MFsFamilies = 10
	// MFsTypes represents the API credit cost for mutual fund types requests.
	MFsTypes = 10

	// IndividualIndicators represents the API credit cost for individual technical indicator requests.
	// Technical Indicators.
	IndividualIndicators = 10
	// BBands represents the API credit cost for Bollinger Bands technical indicator requests.
	BBands = 10
	// SMA represents the API credit cost for Simple Moving Average technical indicator requests.
	SMA = 10
	// EMA represents the API credit cost for Exponential Moving Average technical indicator requests.
	EMA = 10
	// ADX represents the API credit cost for Average Directional Index technical indicator requests.
	ADX = 10
	// MACD represents the API credit cost for Moving Average Convergence Divergence technical indicator requests.
	MACD = 10
	// RSI represents the API credit cost for Relative Strength Index technical indicator requests.
	RSI = 10
	// Stoch represents the API credit cost for Stochastic Oscillator technical indicator requests.
	Stoch = 10
	// PercentB represents the API credit cost for %B technical indicator requests.
	PercentB = 10
	// ATR represents the API credit cost for Average True Range technical indicator requests.
	ATR = 10
	// VWAP represents the API credit cost for Volume Weighted Average Price technical indicator requests.
	VWAP = 10
	// MA represents the API credit cost for Moving Average technical indicator requests.
	MA = 10
	// WMA represents the API credit cost for Weighted Moving Average technical indicator requests.
	WMA = 10
	// DEMA represents the API credit cost for Double Exponential Moving Average technical indicator requests.
	DEMA = 10
	// TEMA represents the API credit cost for Triple Exponential Moving Average technical indicator requests.
	TEMA = 10
	// TRMA represents the API credit cost for Triangular Moving Average technical indicator requests.
	TRMA = 10
	// KAMA represents the API credit cost for Kaufman Adaptive Moving Average technical indicator requests.
	KAMA = 10
	// SAR represents the API credit cost for Parabolic SAR technical indicator requests.
	SAR = 10
	// CCI represents the API credit cost for Commodity Channel Index technical indicator requests.
	CCI = 10
	// WillR represents the API credit cost for Williams %R technical indicator requests.
	WillR = 10
	// ROC represents the API credit cost for Rate of Change technical indicator requests.
	ROC = 10
	// MOM represents the API credit cost for Momentum technical indicator requests.
	MOM = 10
	// OBV represents the API credit cost for On Balance Volume technical indicator requests.
	OBV = 10
	// AD represents the API credit cost for Accumulation/Distribution technical indicator requests.
	AD = 10
	// NATR represents the API credit cost for Normalized Average True Range technical indicator requests.
	NATR = 10
	// TR represents the API credit cost for True Range technical indicator requests.
	TR = 10
	// CustomIndicators represents the API credit cost for custom technical indicators requests.
	CustomIndicators = 20

	// EarningsEstimate represents the API credit cost for earnings estimate requests.
	// Analysis.
	EarningsEstimate = 100
	// RevenueEstimate represents the API credit cost for revenue estimate requests.
	RevenueEstimate = 100
	// EPSTrend represents the API credit cost for EPS trend requests.
	EPSTrend = 100
	// EPSRevisions represents the API credit cost for EPS revisions requests.
	EPSRevisions = 100
	// GrowthEstimates represents the API credit cost for growth estimates requests.
	GrowthEstimates = 100
	// Recommendations represents the API credit cost for analyst recommendations requests.
	Recommendations = 100
	// PriceTarget represents the API credit cost for price target requests.
	PriceTarget = 100
	// AnalystRatingsSnapshot represents the API credit cost for analyst ratings snapshot requests.
	AnalystRatingsSnapshot = 200
	// AnalystRatingsUSEquities represents the API credit cost for US equities analyst ratings requests.
	AnalystRatingsUSEquities = 200

	// InsiderTransactions represents the API credit cost for insider transactions requests.
	// Regulatory.
	InsiderTransactions = 200
	// EDGARFilings represents the API credit cost for EDGAR filings requests.
	EDGARFilings = 50
	// InstitutionalHolders represents the API credit cost for institutional holders requests.
	InstitutionalHolders = 1500
	// FundHolders represents the API credit cost for fund holders requests.
	FundHolders = 1500
	// DirectHolders represents the API credit cost for direct holders requests.
	DirectHolders = 1500
	// TaxInformation represents the API credit cost for tax information requests.
	TaxInformation = 50
	// SanctionedEntities represents the API credit cost for sanctioned entities requests.
	SanctionedEntities = 50

	// Usage represents the API credit cost for usage tracking requests.
	// Advanced.
	Usage = 1

	// RealTimePrice represents the API credit cost for real-time price WebSocket connections.
	// WebSocket.
	RealTimePrice = 1
)
