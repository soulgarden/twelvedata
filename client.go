// Package twelvedata provides a Go client for the Twelve Data API, offering comprehensive access to
// financial market data including real-time quotes, historical time series, fundamental data, and more.
package twelvedata

import (
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

// Client defines the interface for interacting with the Twelve Data API.
// It provides methods for accessing various financial data endpoints including
// reference data, core market data, fundamentals, currencies, and WebSocket streaming.
type Client interface {
	// Reference Data - Asset Catalogs
	GetStocks(request.GetStock) (response.Stocks, response.Credits, error)
	GetForexPairs(request.GetForexPairs) (response.ForexPairs, response.Credits, error)
	GetCryptocurrencies(request.GetCryptocurrencies) (response.Cryptocurrencies, response.Credits, error)
	GetETFs(request.GetETFs) (response.ETFs, response.Credits, error)
	GetFunds(request.GetFunds) (response.Funds, response.Credits, error)
	GetCommodities(request.GetCommodities) (response.Commodities, response.Credits, error)
	GetBonds(request.GetBonds) (response.Bonds, response.Credits, error)

	// Reference Data - Discovery
	GetSymbolSearch(request.GetSymbolSearch) (response.SymbolSearch, response.Credits, error)
	GetCrossListings(request.GetCrossListings) (response.CrossListings, response.Credits, error)
	GetEarliestTimestamp(request.GetEarliestTimestamp) (response.EarliestTimestamp, response.Credits, error)

	// Reference Data - Markets
	GetExchanges(request.GetExchanges) (response.Exchanges, response.Credits, error)
	GetExchangeSchedule(request.GetExchangeSchedule) (response.ExchangeSchedule, response.Credits, error)
	GetCryptocurrencyExchanges(request.GetCryptocurrencyExchanges) (response.CryptocurrencyExchanges, response.Credits, error)
	GetMarketState(request.GetMarketState) ([]response.MarketState, response.Credits, error)

	// Reference Data - Supporting Metadata
	GetCountries(request.GetCountries) (response.Countries, response.Credits, error)
	GetInstrumentType(request.GetInstrumentType) (response.InstrumentType, response.Credits, error)
	GetTechnicalIndicators(request.GetTechnicalIndicators) (response.TechnicalIndicators, response.Credits, error)

	// Core Data
	GetTimeSeries(request.GetTimeSeries) (response.TimeSeries, response.Credits, error)
	GetTimeSeriesCross(request.GetTimeSeriesCross) (response.TimeSeriesCross, response.Credits, error)
	GetQuote(request.GetQuote) (response.Quote, response.Credits, error)
	GetPrice(request.GetPrice) (response.Price, response.Credits, error)
	GetEOD(request.GetEOD) (response.EOD, response.Credits, error)
	GetMarketMovers(request.GetMarketMovers) (response.MarketMovers, response.Credits, error)

	// Fundamentals
	GetLogo(request.GetLogo) (response.Logo, response.Credits, error)
	GetProfile(request.GetProfile) (response.Profile, response.Credits, error)
	GetKeyExecutives(request.GetKeyExecutives) (response.KeyExecutives, response.Credits, error)
	GetDividends(request.GetDividends) (response.Dividends, response.Credits, error)
	GetDividendsCalendar(request.GetDividendsCalendar) (response.DividendsCalendar, response.Credits, error)
	GetEarnings(request.GetEarnings) (response.EarningsData, response.Credits, error)
	GetSplits(request.GetSplits) (response.Splits, response.Credits, error)
	GetSplitsCalendar(request.GetSplitsCalendar) (response.SplitsCalendar, response.Credits, error)
	GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error)
	GetEarningsCalendar(request.GetEarningsCalendar) (response.Earnings, response.Credits, error)
	GetIPOCalendar(request.GetIPOCalendar) (response.IPOCalendar, response.Credits, error)
	GetIncomeStatement(request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error)
	GetIncomeStatementConsolidated(request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error)
	GetBalanceSheet(request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error)
	GetBalanceSheetConsolidated(request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error)
	GetCashFlow(request.GetCashFlow) (response.CashFlows, response.Credits, error)
	GetCashFlowConsolidated(request.GetCashFlow) (response.CashFlows, response.Credits, error)
	GetMarketCap(request.GetMarketCap) (response.MarketCap, response.Credits, error)
	GetLastChange(request.GetLastChange) (response.LastChange, response.Credits, error)

	// Technical Indicators
	GetBBands(request.GetBBands) (response.BBands, response.Credits, error)
	GetSMA(request.GetSMA) (response.SMA, response.Credits, error)
	GetEMA(request.GetEMA) (response.EMA, response.Credits, error)
	GetADX(request.GetADX) (response.ADX, response.Credits, error)
	GetMACD(request.GetMACD) (response.MACD, response.Credits, error)
	GetRSI(request.GetRSI) (response.RSI, response.Credits, error)
	GetStoch(request.GetStoch) (response.Stoch, response.Credits, error)
	GetPercentB(request.GetPercentB) (response.PercentB, response.Credits, error)
	GetATR(request.GetATR) (response.ATR, response.Credits, error)
	GetVWAP(request.GetVWAP) (response.VWAP, response.Credits, error)
	GetMA(request.GetMA) (response.MA, response.Credits, error)
	GetWMA(request.GetWMA) (response.WMA, response.Credits, error)
	GetDEMA(request.GetDEMA) (response.DEMA, response.Credits, error)
	GetTEMA(request.GetTEMA) (response.TEMA, response.Credits, error)
	GetTRMA(request.GetTRMA) (response.TRMA, response.Credits, error)
	GetKAMA(request.GetKAMA) (response.KAMA, response.Credits, error)
	GetSAR(request.GetSAR) (response.SAR, response.Credits, error)
	GetCCI(request.GetCCI) (response.CCI, response.Credits, error)
	GetWillR(request.GetWillR) (response.WillR, response.Credits, error)
	GetROC(request.GetROC) (response.ROC, response.Credits, error)
	GetMOM(request.GetMOM) (response.MOM, response.Credits, error)
	GetOBV(request.GetOBV) (response.OBV, response.Credits, error)
	GetAD(request.GetAD) (response.AD, response.Credits, error)
	GetNATR(request.GetNATR) (response.NATR, response.Credits, error)
	GetTR(request.GetTR) (response.TR, response.Credits, error)

	// Currencies
	GetExchangeRate(request.GetExchangeRate) (response.ExchangeRate, response.Credits, error)
	GetCurrencyConversion(request.GetCurrencyConversion) (response.CurrencyConversion, response.Credits, error)

	// Advanced
	GetUsage(request.GetUsage) (response.Usage, response.Credits, error)
	GetBatches(request.GetBatches) (response.Batches, response.Credits, error)

	// ETFs
	GetETFsDirectory(request.GetETFsDirectory) (response.ETFsDirectory, response.Credits, error)
	GetETFFullData(request.GetETFFullData) (response.ETFFullData, response.Credits, error)
	GetETFSummary(request.GetETFSummary) (response.ETFWorldSummary, response.Credits, error)
	GetETFPerformance(request.GetETFPerformance) (response.ETFPerformance, response.Credits, error)
	GetETFRisk(request.GetETFRisk) (response.ETFRisk, response.Credits, error)
	GetETFComposition(request.GetETFComposition) (response.ETFComposition, response.Credits, error)
	GetETFFamilies(request.GetETFFamilies) (response.ETFFamilies, response.Credits, error)
	GetETFTypes(request.GetETFTypes) (response.ETFTypes, response.Credits, error)

	// Mutual Funds
	GetMutualFunds(request.GetMutualFunds) (response.MutualFunds, response.Credits, error)
	GetMutualFundFullData(request.GetMutualFundFullData) (response.MutualFundFullData, response.Credits, error)
	GetMutualFundSummary(request.GetMutualFundSummary) (response.MutualFundSummaryResponse, response.Credits, error)
	GetMutualFundPerformance(request.GetMutualFundPerformance) (response.MutualFundPerformance, response.Credits, error)
	GetMutualFundRisk(request.GetMutualFundRisk) (response.MutualFundRiskResponse, response.Credits, error)
	GetMutualFundRatings(request.GetMutualFundRatings) (response.MutualFundRatingsResponse, response.Credits, error)
	GetMutualFundComposition(request.GetMutualFundComposition) (response.MutualFundComposition, response.Credits, error)
	GetMutualFundPurchaseInfo(request.GetMutualFundPurchaseInfo) (response.MutualFundPurchaseInfoResponse, response.Credits, error)
	GetMutualFundSustainability(request.GetMutualFundSustainability) (response.MutualFundSustainability, response.Credits, error)
	GetMutualFundFamilies(request.GetMutualFundFamilies) (response.MutualFundFamilies, response.Credits, error)
	GetMutualFundTypes(request.GetMutualFundTypes) (response.MutualFundTypes, response.Credits, error)

	// Analysis
	GetRecommendations(request.GetRecommendations) (response.Recommendations, response.Credits, error)
	GetPriceTarget(request.GetPriceTarget) (response.PriceTarget, response.Credits, error)
	GetEarningsEstimate(request.GetEarningsEstimate) (response.EarningsEstimate, response.Credits, error)
	GetRevenueEstimate(request.GetRevenueEstimate) (response.RevenueEstimate, response.Credits, error)
	GetEPSTrend(request.GetEPSTrend) (response.EPSTrend, response.Credits, error)
	GetEPSRevisions(request.GetEPSRevisions) (response.EPSRevisions, response.Credits, error)
	GetGrowthEstimates(request.GetGrowthEstimates) (response.GrowthEstimates, response.Credits, error)
	GetAnalystRatingsSnapshot(request.GetAnalystRatingsSnapshot) (response.AnalystRatingsSnapshot, response.Credits, error)
	GetAnalystRatingsUSEquities(request.GetAnalystRatingsUSEquities) (response.AnalystRatingsUSEquities, response.Credits, error)

	// Regulatory
	GetInsiderTransactions(request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error)
	GetEDGARFillings(request.GetEDGARFillings) (response.EDGARFillings, response.Credits, error)
	GetInstitutionalHolders(request.GetInstitutionalHolders) (response.InstitutionalHolders, response.Credits, error)
	GetFundHolders(request.GetFundHolders) (response.FundHolders, response.Credits, error)
	GetDirectHolders(request.GetDirectHolders) (response.DirectHolders, response.Credits, error)
	GetTaxInformation(request.GetTaxInformation) (response.TaxInformation, response.Credits, error)
	GetSanctionedEntities(request.GetSanctionedEntities) (response.SanctionedEntities, response.Credits, error)
}
