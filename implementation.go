package twelvedata

import (
	"strings"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

type client struct {
	// Reference Data - Asset Catalogs
	getStocks           *Endpoint[request.GetStock, response.Stocks, response.Credits, error]
	getForexPairs       *Endpoint[request.GetForexPairs, response.ForexPairs, response.Credits, error]
	getCryptocurrencies *Endpoint[request.GetCryptocurrencies, response.Cryptocurrencies, response.Credits, error]
	getCommodities      *Endpoint[request.GetCommodities, response.Commodities, response.Credits, error]
	getBonds            *Endpoint[request.GetBonds, response.Bonds, response.Credits, error]

	// Reference Data - Discovery
	getSymbolSearch      *Endpoint[request.GetSymbolSearch, response.SymbolSearch, response.Credits, error]
	getCrossListings     *Endpoint[request.GetCrossListings, response.CrossListings, response.Credits, error]
	getEarliestTimestamp *Endpoint[request.GetEarliestTimestamp, response.EarliestTimestamp, response.Credits, error]

	// Reference Data - Markets
	getExchanges               *Endpoint[request.GetExchanges, response.Exchanges, response.Credits, error]
	getExchangeSchedule        *Endpoint[request.GetExchangeSchedule, response.ExchangeSchedule, response.Credits, error]
	getCryptocurrencyExchanges *Endpoint[request.GetCryptocurrencyExchanges, response.CryptocurrencyExchanges, response.Credits, error]
	getMarketState             *Endpoint[request.GetMarketState, []response.MarketState, response.Credits, error]

	// Reference Data - Supporting Metadata
	getCountries           *Endpoint[request.GetCountries, response.Countries, response.Credits, error]
	getInstrumentType      *Endpoint[request.GetInstrumentType, response.InstrumentType, response.Credits, error]
	getTechnicalIndicators *Endpoint[request.GetTechnicalIndicators, response.TechnicalIndicators, response.Credits, error]

	// Core Data
	getTimeSeries      *Endpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error]
	getTimeSeriesCross *Endpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error]
	getQuote           *Endpoint[request.GetQuote, response.Quote, response.Credits, error]
	getPrice           *Endpoint[request.GetPrice, response.Price, response.Credits, error]
	getEOD             *Endpoint[request.GetEOD, response.EOD, response.Credits, error]
	getMarketMovers    *Endpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error]

	// Fundamentals
	getLogo              *Endpoint[request.GetLogo, response.Logo, response.Credits, error]
	getProfile           *Endpoint[request.GetProfile, response.Profile, response.Credits, error]
	getKeyExecutives     *Endpoint[request.GetKeyExecutives, response.KeyExecutives, response.Credits, error]
	getDividends         *Endpoint[request.GetDividends, response.Dividends, response.Credits, error]
	getDividendsCalendar *Endpoint[request.GetDividendsCalendar, response.DividendsCalendar, response.Credits, error]
	getEarnings          *Endpoint[request.GetEarnings, response.EarningsData, response.Credits, error]
	getSplits            *Endpoint[request.GetSplits, response.Splits, response.Credits, error]
	getSplitsCalendar    *Endpoint[request.GetSplitsCalendar, response.SplitsCalendar, response.Credits, error]
	getStatistics        *Endpoint[request.GetStatistics, response.Statistics, response.Credits, error]
	getEarningsCalendar  *Endpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error]
	getIPOCalendar       *Endpoint[request.GetIPOCalendar, response.IPOCalendar, response.Credits, error]
	getIncomeStatement   *Endpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error]
	getBalanceSheet      *Endpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error]
	getCashFlow          *Endpoint[request.GetCashFlow, response.CashFlows, response.Credits, error]
	getMarketCap         *Endpoint[request.GetMarketCap, response.MarketCap, response.Credits, error]
	getLastChange        *Endpoint[request.GetLastChange, response.LastChange, response.Credits, error]

	// Technical Indicators
	getBBands   *Endpoint[request.GetBBands, response.BBands, response.Credits, error]
	getSMA      *Endpoint[request.GetSMA, response.SMA, response.Credits, error]
	getEMA      *Endpoint[request.GetEMA, response.EMA, response.Credits, error]
	getADX      *Endpoint[request.GetADX, response.ADX, response.Credits, error]
	getMACD     *Endpoint[request.GetMACD, response.MACD, response.Credits, error]
	getRSI      *Endpoint[request.GetRSI, response.RSI, response.Credits, error]
	getStoch    *Endpoint[request.GetStoch, response.Stoch, response.Credits, error]
	getPercentB *Endpoint[request.GetPercentB, response.PercentB, response.Credits, error]
	getATR      *Endpoint[request.GetATR, response.ATR, response.Credits, error]
	getVWAP     *Endpoint[request.GetVWAP, response.VWAP, response.Credits, error]
	getMA       *Endpoint[request.GetMA, response.MA, response.Credits, error]
	getWMA      *Endpoint[request.GetWMA, response.WMA, response.Credits, error]
	getDEMA     *Endpoint[request.GetDEMA, response.DEMA, response.Credits, error]
	getTEMA     *Endpoint[request.GetTEMA, response.TEMA, response.Credits, error]
	getTRMA     *Endpoint[request.GetTRMA, response.TRMA, response.Credits, error]
	getKAMA     *Endpoint[request.GetKAMA, response.KAMA, response.Credits, error]
	getSAR      *Endpoint[request.GetSAR, response.SAR, response.Credits, error]
	getCCI      *Endpoint[request.GetCCI, response.CCI, response.Credits, error]
	getWillR    *Endpoint[request.GetWillR, response.WillR, response.Credits, error]
	getROC      *Endpoint[request.GetROC, response.ROC, response.Credits, error]
	getMOM      *Endpoint[request.GetMOM, response.MOM, response.Credits, error]
	getOBV      *Endpoint[request.GetOBV, response.OBV, response.Credits, error]
	getAD       *Endpoint[request.GetAD, response.AD, response.Credits, error]
	getNATR     *Endpoint[request.GetNATR, response.NATR, response.Credits, error]
	getTR       *Endpoint[request.GetTR, response.TR, response.Credits, error]

	// Currencies
	getExchangeRate       *Endpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error]
	getCurrencyConversion *Endpoint[request.GetCurrencyConversion, response.CurrencyConversion, response.Credits, error]

	// Advanced
	getUsage   *Endpoint[request.GetUsage, response.Usage, response.Credits, error]
	getBatches *Endpoint[request.GetBatches, response.Batches, response.Credits, error]

	// ETFs
	getETFsDirectory  *Endpoint[request.GetETFsDirectory, response.ETFsDirectory, response.Credits, error]
	getETFFullData    *Endpoint[request.GetETFFullData, response.ETFFullData, response.Credits, error]
	getETFSummary     *Endpoint[request.GetETFSummary, response.ETFWorldSummary, response.Credits, error]
	getETFPerformance *Endpoint[request.GetETFPerformance, response.ETFPerformance, response.Credits, error]
	getETFRisk        *Endpoint[request.GetETFRisk, response.ETFRisk, response.Credits, error]
	getETFComposition *Endpoint[request.GetETFComposition, response.ETFComposition, response.Credits, error]
	getETFFamilies    *Endpoint[request.GetETFFamilies, response.ETFFamilies, response.Credits, error]
	getETFTypes       *Endpoint[request.GetETFTypes, response.ETFTypes, response.Credits, error]

	// Mutual Funds
	getFunds                    *Endpoint[request.GetFunds, response.Funds, response.Credits, error]
	getMutualFunds              *Endpoint[request.GetMutualFunds, response.MutualFunds, response.Credits, error]
	getMutualFundFullData       *Endpoint[request.GetMutualFundFullData, response.MutualFundFullData, response.Credits, error]
	getMutualFundSummary        *Endpoint[request.GetMutualFundSummary, response.MutualFundSummaryResponse, response.Credits, error]
	getMutualFundPerformance    *Endpoint[request.GetMutualFundPerformance, response.MutualFundPerformance, response.Credits, error]
	getMutualFundRisk           *Endpoint[request.GetMutualFundRisk, response.MutualFundRiskResponse, response.Credits, error]
	getMutualFundRatings        *Endpoint[request.GetMutualFundRatings, response.MutualFundRatingsResponse, response.Credits, error]
	getMutualFundComposition    *Endpoint[request.GetMutualFundComposition, response.MutualFundComposition, response.Credits, error]
	getMutualFundPurchaseInfo   *Endpoint[request.GetMutualFundPurchaseInfo, response.MutualFundPurchaseInfoResponse, response.Credits, error]
	getMutualFundSustainability *Endpoint[request.GetMutualFundSustainability, response.MutualFundSustainability, response.Credits, error]
	getMutualFundFamilies       *Endpoint[request.GetMutualFundFamilies, response.MutualFundFamilies, response.Credits, error]
	getMutualFundTypes          *Endpoint[request.GetMutualFundTypes, response.MutualFundTypes, response.Credits, error]

	// Analysis
	getRecommendations          *Endpoint[request.GetRecommendations, response.Recommendations, response.Credits, error]
	getPriceTarget              *Endpoint[request.GetPriceTarget, response.PriceTarget, response.Credits, error]
	getEarningsEstimate         *Endpoint[request.GetEarningsEstimate, response.EarningsEstimate, response.Credits, error]
	getRevenueEstimate          *Endpoint[request.GetRevenueEstimate, response.RevenueEstimate, response.Credits, error]
	getEPSTrend                 *Endpoint[request.GetEPSTrend, response.EPSTrend, response.Credits, error]
	getEPSRevisions             *Endpoint[request.GetEPSRevisions, response.EPSRevisions, response.Credits, error]
	getGrowthEstimates          *Endpoint[request.GetGrowthEstimates, response.GrowthEstimates, response.Credits, error]
	getAnalystRatingsSnapshot   *Endpoint[request.GetAnalystRatingsSnapshot, response.AnalystRatingsSnapshot, response.Credits, error]
	getAnalystRatingsUSEquities *Endpoint[request.GetAnalystRatingsUSEquities, response.AnalystRatingsUSEquities, response.Credits, error]

	// Regulatory
	getInsiderTransactions  *Endpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error]
	getEDGARFillings        *Endpoint[request.GetEDGARFillings, response.EDGARFillings, response.Credits, error]
	getInstitutionalHolders *Endpoint[request.GetInstitutionalHolders, response.InstitutionalHolders, response.Credits, error]
	getFundHolders          *Endpoint[request.GetFundHolders, response.FundHolders, response.Credits, error]
	getDirectHolders        *Endpoint[request.GetDirectHolders, response.DirectHolders, response.Credits, error]
	getTaxInformation       *Endpoint[request.GetTaxInformation, response.TaxInformation, response.Credits, error]
	getSanctionedEntities   *Endpoint[request.GetSanctionedEntities, response.SanctionedEntities, response.Credits, error]
}

func (cli client) GetStocks(req request.GetStock) (response.Stocks, response.Credits, error) {
	return cli.getStocks.Call(req)
}

func (cli client) GetTimeSeries(req request.GetTimeSeries) (response.TimeSeries, response.Credits, error) {
	return cli.getTimeSeries.Call(req)
}

func (cli client) GetTimeSeriesCross(req request.GetTimeSeriesCross) (response.TimeSeriesCross, response.Credits, error) {
	return cli.getTimeSeriesCross.Call(req)
}

func (cli client) GetLogo(req request.GetLogo) (response.Logo, response.Credits, error) {
	return cli.getLogo.Call(req)
}

func (cli client) GetProfile(req request.GetProfile) (response.Profile, response.Credits, error) {
	return cli.getProfile.Call(req)
}

func (cli client) GetKeyExecutives(req request.GetKeyExecutives) (response.KeyExecutives, response.Credits, error) {
	return cli.getKeyExecutives.Call(req)
}

func (cli client) GetInsiderTransactions(req request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error) {
	return cli.getInsiderTransactions.Call(req)
}
func (cli client) GetEDGARFillings(req request.GetEDGARFillings) (response.EDGARFillings, response.Credits, error) {
	return cli.getEDGARFillings.Call(req)
}
func (cli client) GetInstitutionalHolders(req request.GetInstitutionalHolders) (response.InstitutionalHolders, response.Credits, error) {
	return cli.getInstitutionalHolders.Call(req)
}
func (cli client) GetFundHolders(req request.GetFundHolders) (response.FundHolders, response.Credits, error) {
	return cli.getFundHolders.Call(req)
}
func (cli client) GetDirectHolders(req request.GetDirectHolders) (response.DirectHolders, response.Credits, error) {
	return cli.getDirectHolders.Call(req)
}
func (cli client) GetTaxInformation(req request.GetTaxInformation) (response.TaxInformation, response.Credits, error) {
	return cli.getTaxInformation.Call(req)
}
func (cli client) GetSanctionedEntities(req request.GetSanctionedEntities) (response.SanctionedEntities, response.Credits, error) {
	return cli.getSanctionedEntities.Call(req)
}

func (cli client) GetDividends(req request.GetDividends) (response.Dividends, response.Credits, error) {
	return cli.getDividends.Call(req)
}

func (cli client) GetDividendsCalendar(req request.GetDividendsCalendar) (response.DividendsCalendar, response.Credits, error) {
	return cli.getDividendsCalendar.Call(req)
}

func (cli client) GetEarnings(req request.GetEarnings) (response.EarningsData, response.Credits, error) {
	return cli.getEarnings.Call(req)
}

func (cli client) GetSplits(req request.GetSplits) (response.Splits, response.Credits, error) {
	return cli.getSplits.Call(req)
}

func (cli client) GetSplitsCalendar(req request.GetSplitsCalendar) (response.SplitsCalendar, response.Credits, error) {
	return cli.getSplitsCalendar.Call(req)
}

func (cli client) GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error) {
	return cli.getStatistics.Call(statistics)
}

func (cli client) GetExchanges(req request.GetExchanges) (response.Exchanges, response.Credits, error) {
	return cli.getExchanges.Call(req)
}

func (cli client) GetQuote(req request.GetQuote) (response.Quote, response.Credits, error) {
	return cli.getQuote.Call(req)
}

func (cli client) GetUsage(req request.GetUsage) (response.Usage, response.Credits, error) {
	return cli.getUsage.Call(req)
}
func (cli client) GetBatches(req request.GetBatches) (response.Batches, response.Credits, error) {
	return cli.getBatches.Call(req)
}

func (cli client) GetLastChange(req request.GetLastChange) (response.LastChange, response.Credits, error) {
	// Replace {endpoint} placeholder with actual endpoint value
	url := strings.Replace(cli.getLastChange.URL, "{endpoint}", req.Endpoint, 1)
	lastChangeEndpoint := NewEndpoint[request.GetLastChange, response.LastChange, response.Credits, error](
		cli.getLastChange.httpCli,
		url,
	)

	return lastChangeEndpoint.Call(req)
}

func (cli client) GetEarningsCalendar(req request.GetEarningsCalendar) (response.Earnings, response.Credits, error) {
	return cli.getEarningsCalendar.Call(req)
}

func (cli client) GetIPOCalendar(req request.GetIPOCalendar) (response.IPOCalendar, response.Credits, error) {
	return cli.getIPOCalendar.Call(req)
}

func (cli client) GetExchangeRate(req request.GetExchangeRate) (response.ExchangeRate, response.Credits, error) {
	return cli.getExchangeRate.Call(req)
}

func (cli client) GetCurrencyConversion(req request.GetCurrencyConversion) (response.CurrencyConversion, response.Credits, error) {
	return cli.getCurrencyConversion.Call(req)
}

func (cli client) GetIncomeStatement(req request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error) {
	return cli.getIncomeStatement.Call(req)
}

func (cli client) GetBalanceSheet(req request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error) {
	return cli.getBalanceSheet.Call(req)
}

func (cli client) GetCashFlow(req request.GetCashFlow) (response.CashFlows, response.Credits, error) {
	return cli.getCashFlow.Call(req)
}

func (cli client) GetMarketCap(req request.GetMarketCap) (response.MarketCap, response.Credits, error) {
	return cli.getMarketCap.Call(req)
}

func (cli client) GetMarketMovers(req request.GetMarketMovers) (response.MarketMovers, response.Credits, error) {
	// Replace {market} placeholder with actual market value
	url := strings.Replace(cli.getMarketMovers.URL, "{market}", req.Market, 1)
	marketEndpoint := NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](
		cli.getMarketMovers.httpCli,
		url,
	)

	return marketEndpoint.Call(req)
}

func (cli client) GetMarketState(req request.GetMarketState) ([]response.MarketState, response.Credits, error) {
	return cli.getMarketState.Call(req)
}

func (cli client) GetPrice(req request.GetPrice) (response.Price, response.Credits, error) {
	return cli.getPrice.Call(req)
}

func (cli client) GetEOD(req request.GetEOD) (response.EOD, response.Credits, error) {
	return cli.getEOD.Call(req)
}

// Reference Data - Asset Catalogs.
func (cli client) GetForexPairs(req request.GetForexPairs) (response.ForexPairs, response.Credits, error) {
	return cli.getForexPairs.Call(req)
}

func (cli client) GetCryptocurrencies(req request.GetCryptocurrencies) (response.Cryptocurrencies, response.Credits, error) {
	return cli.getCryptocurrencies.Call(req)
}

func (cli client) GetFunds(req request.GetFunds) (response.Funds, response.Credits, error) {
	return cli.getFunds.Call(req)
}

func (cli client) GetCommodities(req request.GetCommodities) (response.Commodities, response.Credits, error) {
	return cli.getCommodities.Call(req)
}

func (cli client) GetBonds(req request.GetBonds) (response.Bonds, response.Credits, error) {
	return cli.getBonds.Call(req)
}

// Reference Data - Discovery.
func (cli client) GetSymbolSearch(req request.GetSymbolSearch) (response.SymbolSearch, response.Credits, error) {
	return cli.getSymbolSearch.Call(req)
}

func (cli client) GetCrossListings(req request.GetCrossListings) (response.CrossListings, response.Credits, error) {
	return cli.getCrossListings.Call(req)
}

func (cli client) GetEarliestTimestamp(req request.GetEarliestTimestamp) (response.EarliestTimestamp, response.Credits, error) {
	return cli.getEarliestTimestamp.Call(req)
}

// Reference Data - Markets.
func (cli client) GetExchangeSchedule(req request.GetExchangeSchedule) (response.ExchangeSchedule, response.Credits, error) {
	return cli.getExchangeSchedule.Call(req)
}

func (cli client) GetCryptocurrencyExchanges(req request.GetCryptocurrencyExchanges) (response.CryptocurrencyExchanges, response.Credits, error) {
	return cli.getCryptocurrencyExchanges.Call(req)
}

// Reference Data - Supporting Metadata.
func (cli client) GetCountries(req request.GetCountries) (response.Countries, response.Credits, error) {
	return cli.getCountries.Call(req)
}

func (cli client) GetInstrumentType(req request.GetInstrumentType) (response.InstrumentType, response.Credits, error) {
	return cli.getInstrumentType.Call(req)
}

func (cli client) GetTechnicalIndicators(req request.GetTechnicalIndicators) (response.TechnicalIndicators, response.Credits, error) {
	return cli.getTechnicalIndicators.Call(req)
}

// Technical Indicators.
func (cli client) GetBBands(req request.GetBBands) (response.BBands, response.Credits, error) {
	return cli.getBBands.Call(req)
}

func (cli client) GetSMA(req request.GetSMA) (response.SMA, response.Credits, error) {
	return cli.getSMA.Call(req)
}

func (cli client) GetEMA(req request.GetEMA) (response.EMA, response.Credits, error) {
	return cli.getEMA.Call(req)
}

func (cli client) GetADX(req request.GetADX) (response.ADX, response.Credits, error) {
	return cli.getADX.Call(req)
}

func (cli client) GetMACD(req request.GetMACD) (response.MACD, response.Credits, error) {
	return cli.getMACD.Call(req)
}

func (cli client) GetRSI(req request.GetRSI) (response.RSI, response.Credits, error) {
	return cli.getRSI.Call(req)
}

func (cli client) GetStoch(req request.GetStoch) (response.Stoch, response.Credits, error) {
	return cli.getStoch.Call(req)
}

func (cli client) GetPercentB(req request.GetPercentB) (response.PercentB, response.Credits, error) {
	return cli.getPercentB.Call(req)
}
func (cli client) GetATR(req request.GetATR) (response.ATR, response.Credits, error) {
	return cli.getATR.Call(req)
}
func (cli client) GetVWAP(req request.GetVWAP) (response.VWAP, response.Credits, error) {
	return cli.getVWAP.Call(req)
}
func (cli client) GetMA(req request.GetMA) (response.MA, response.Credits, error) {
	return cli.getMA.Call(req)
}
func (cli client) GetWMA(req request.GetWMA) (response.WMA, response.Credits, error) {
	return cli.getWMA.Call(req)
}
func (cli client) GetDEMA(req request.GetDEMA) (response.DEMA, response.Credits, error) {
	return cli.getDEMA.Call(req)
}
func (cli client) GetTEMA(req request.GetTEMA) (response.TEMA, response.Credits, error) {
	return cli.getTEMA.Call(req)
}
func (cli client) GetTRMA(req request.GetTRMA) (response.TRMA, response.Credits, error) {
	return cli.getTRMA.Call(req)
}
func (cli client) GetKAMA(req request.GetKAMA) (response.KAMA, response.Credits, error) {
	return cli.getKAMA.Call(req)
}
func (cli client) GetSAR(req request.GetSAR) (response.SAR, response.Credits, error) {
	return cli.getSAR.Call(req)
}
func (cli client) GetCCI(req request.GetCCI) (response.CCI, response.Credits, error) {
	return cli.getCCI.Call(req)
}
func (cli client) GetWillR(req request.GetWillR) (response.WillR, response.Credits, error) {
	return cli.getWillR.Call(req)
}
func (cli client) GetROC(req request.GetROC) (response.ROC, response.Credits, error) {
	return cli.getROC.Call(req)
}
func (cli client) GetMOM(req request.GetMOM) (response.MOM, response.Credits, error) {
	return cli.getMOM.Call(req)
}
func (cli client) GetOBV(req request.GetOBV) (response.OBV, response.Credits, error) {
	return cli.getOBV.Call(req)
}
func (cli client) GetAD(req request.GetAD) (response.AD, response.Credits, error) {
	return cli.getAD.Call(req)
}
func (cli client) GetNATR(req request.GetNATR) (response.NATR, response.Credits, error) {
	return cli.getNATR.Call(req)
}
func (cli client) GetTR(req request.GetTR) (response.TR, response.Credits, error) {
	return cli.getTR.Call(req)
}

func (cli client) GetETFsDirectory(req request.GetETFsDirectory) (response.ETFsDirectory, response.Credits, error) {
	return cli.getETFsDirectory.Call(req)
}

func (cli client) GetETFFullData(req request.GetETFFullData) (response.ETFFullData, response.Credits, error) {
	return cli.getETFFullData.Call(req)
}

func (cli client) GetETFPerformance(req request.GetETFPerformance) (response.ETFPerformance, response.Credits, error) {
	return cli.getETFPerformance.Call(req)
}

func (cli client) GetETFComposition(req request.GetETFComposition) (response.ETFComposition, response.Credits, error) {
	return cli.getETFComposition.Call(req)
}
func (cli client) GetETFSummary(req request.GetETFSummary) (response.ETFWorldSummary, response.Credits, error) {
	return cli.getETFSummary.Call(req)
}
func (cli client) GetETFRisk(req request.GetETFRisk) (response.ETFRisk, response.Credits, error) {
	return cli.getETFRisk.Call(req)
}
func (cli client) GetETFFamilies(req request.GetETFFamilies) (response.ETFFamilies, response.Credits, error) {
	return cli.getETFFamilies.Call(req)
}
func (cli client) GetETFTypes(req request.GetETFTypes) (response.ETFTypes, response.Credits, error) {
	return cli.getETFTypes.Call(req)
}

func (cli client) GetMutualFunds(req request.GetMutualFunds) (response.MutualFunds, response.Credits, error) {
	return cli.getMutualFunds.Call(req)
}

func (cli client) GetMutualFundFullData(req request.GetMutualFundFullData) (response.MutualFundFullData, response.Credits, error) {
	return cli.getMutualFundFullData.Call(req)
}

func (cli client) GetMutualFundPerformance(req request.GetMutualFundPerformance) (response.MutualFundPerformance, response.Credits, error) {
	return cli.getMutualFundPerformance.Call(req)
}

func (cli client) GetMutualFundComposition(req request.GetMutualFundComposition) (response.MutualFundComposition, response.Credits, error) {
	return cli.getMutualFundComposition.Call(req)
}
func (cli client) GetMutualFundSummary(req request.GetMutualFundSummary) (response.MutualFundSummaryResponse, response.Credits, error) {
	return cli.getMutualFundSummary.Call(req)
}
func (cli client) GetMutualFundRisk(req request.GetMutualFundRisk) (response.MutualFundRiskResponse, response.Credits, error) {
	return cli.getMutualFundRisk.Call(req)
}
func (cli client) GetMutualFundRatings(req request.GetMutualFundRatings) (response.MutualFundRatingsResponse, response.Credits, error) {
	return cli.getMutualFundRatings.Call(req)
}
func (cli client) GetMutualFundPurchaseInfo(req request.GetMutualFundPurchaseInfo) (response.MutualFundPurchaseInfoResponse, response.Credits, error) {
	return cli.getMutualFundPurchaseInfo.Call(req)
}
func (cli client) GetMutualFundSustainability(req request.GetMutualFundSustainability) (response.MutualFundSustainability, response.Credits, error) {
	return cli.getMutualFundSustainability.Call(req)
}
func (cli client) GetMutualFundFamilies(req request.GetMutualFundFamilies) (response.MutualFundFamilies, response.Credits, error) {
	return cli.getMutualFundFamilies.Call(req)
}
func (cli client) GetMutualFundTypes(req request.GetMutualFundTypes) (response.MutualFundTypes, response.Credits, error) {
	return cli.getMutualFundTypes.Call(req)
}

func (cli client) GetRecommendations(req request.GetRecommendations) (response.Recommendations, response.Credits, error) {
	return cli.getRecommendations.Call(req)
}

func (cli client) GetPriceTarget(req request.GetPriceTarget) (response.PriceTarget, response.Credits, error) {
	return cli.getPriceTarget.Call(req)
}

func (cli client) GetEarningsEstimate(req request.GetEarningsEstimate) (response.EarningsEstimate, response.Credits, error) {
	return cli.getEarningsEstimate.Call(req)
}
func (cli client) GetRevenueEstimate(req request.GetRevenueEstimate) (response.RevenueEstimate, response.Credits, error) {
	return cli.getRevenueEstimate.Call(req)
}
func (cli client) GetEPSTrend(req request.GetEPSTrend) (response.EPSTrend, response.Credits, error) {
	return cli.getEPSTrend.Call(req)
}
func (cli client) GetEPSRevisions(req request.GetEPSRevisions) (response.EPSRevisions, response.Credits, error) {
	return cli.getEPSRevisions.Call(req)
}
func (cli client) GetGrowthEstimates(req request.GetGrowthEstimates) (response.GrowthEstimates, response.Credits, error) {
	return cli.getGrowthEstimates.Call(req)
}
func (cli client) GetAnalystRatingsSnapshot(req request.GetAnalystRatingsSnapshot) (response.AnalystRatingsSnapshot, response.Credits, error) {
	return cli.getAnalystRatingsSnapshot.Call(req)
}
func (cli client) GetAnalystRatingsUSEquities(req request.GetAnalystRatingsUSEquities) (response.AnalystRatingsUSEquities, response.Credits, error) {
	return cli.getAnalystRatingsUSEquities.Call(req)
}

// NewClient creates a new Twelve Data API client instance with the provided HTTP client and configuration.
// The httpCli parameter should be configured with appropriate timeout and other HTTP settings,
// while cfg contains the API endpoints, authentication, and other client configuration.
func NewClient(httpCli *HTTPCli, cfg *Conf) Client {
	return client{
		// Reference Data - Asset Catalogs
		getStocks:           NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.StocksURL),
		getForexPairs:       NewEndpoint[request.GetForexPairs, response.ForexPairs, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.ForexPairsURL),
		getCryptocurrencies: NewEndpoint[request.GetCryptocurrencies, response.Cryptocurrencies, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CryptocurrenciesURL),
		getCommodities:      NewEndpoint[request.GetCommodities, response.Commodities, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CommoditiesURL),
		getBonds:            NewEndpoint[request.GetBonds, response.Bonds, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.BondsURL),

		// Reference Data - Discovery
		getSymbolSearch:      NewEndpoint[request.GetSymbolSearch, response.SymbolSearch, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.SymbolSearchURL),
		getCrossListings:     NewEndpoint[request.GetCrossListings, response.CrossListings, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CrossListingsURL),
		getEarliestTimestamp: NewEndpoint[request.GetEarliestTimestamp, response.EarliestTimestamp, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.EarliestTimestampURL),

		// Reference Data - Markets
		getExchanges:               NewEndpoint[request.GetExchanges, response.Exchanges, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.ExchangesURL),
		getExchangeSchedule:        NewEndpoint[request.GetExchangeSchedule, response.ExchangeSchedule, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.ExchangeScheduleURL),
		getCryptocurrencyExchanges: NewEndpoint[request.GetCryptocurrencyExchanges, response.CryptocurrencyExchanges, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CryptocurrencyExchangesURL),
		getMarketState:             NewEndpoint[request.GetMarketState, []response.MarketState, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.MarketStateURL),

		// Reference Data - Supporting Metadata
		getCountries:           NewEndpoint[request.GetCountries, response.Countries, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CountriesURL),
		getInstrumentType:      NewEndpoint[request.GetInstrumentType, response.InstrumentType, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.InstrumentTypeURL),
		getTechnicalIndicators: NewEndpoint[request.GetTechnicalIndicators, response.TechnicalIndicators, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.TechnicalIndicatorsURL),

		// Core Data
		getTimeSeries:      NewEndpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.TimeSeriesURL),
		getTimeSeriesCross: NewEndpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.TimeSeriesCrossURL),
		getQuote:           NewEndpoint[request.GetQuote, response.Quote, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.QuotesURL),
		getPrice:           NewEndpoint[request.GetPrice, response.Price, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.PriceURL),
		getEOD:             NewEndpoint[request.GetEOD, response.EOD, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.EODURL),
		getMarketMovers:    NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.MarketMoversURL),

		// Fundamentals
		getLogo:              NewEndpoint[request.GetLogo, response.Logo, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.LogoURL),
		getProfile:           NewEndpoint[request.GetProfile, response.Profile, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.ProfileURL),
		getKeyExecutives:     NewEndpoint[request.GetKeyExecutives, response.KeyExecutives, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.KeyExecutivesURL),
		getDividends:         NewEndpoint[request.GetDividends, response.Dividends, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.DividendsURL),
		getDividendsCalendar: NewEndpoint[request.GetDividendsCalendar, response.DividendsCalendar, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.DividendsCalendarURL),
		getEarnings:          NewEndpoint[request.GetEarnings, response.EarningsData, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.EarningsURL),
		getSplits:            NewEndpoint[request.GetSplits, response.Splits, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.SplitsURL),
		getSplitsCalendar:    NewEndpoint[request.GetSplitsCalendar, response.SplitsCalendar, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.SplitsCalendarURL),
		getStatistics:        NewEndpoint[request.GetStatistics, response.Statistics, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.StatisticsURL),
		getEarningsCalendar:  NewEndpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.EarningsCalendarURL),
		getIPOCalendar:       NewEndpoint[request.GetIPOCalendar, response.IPOCalendar, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.IPOCalendarURL),
		getIncomeStatement:   NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.IncomeStatementURL),
		getBalanceSheet:      NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.BalanceSheetURL),
		getCashFlow:          NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.CashFlowURL),
		getMarketCap:         NewEndpoint[request.GetMarketCap, response.MarketCap, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.MarketCapURL),
		getLastChange:        NewEndpoint[request.GetLastChange, response.LastChange, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.LastChangeURL),

		// Technical Indicators
		getBBands:   NewEndpoint[request.GetBBands, response.BBands, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.BbandsURL),
		getSMA:      NewEndpoint[request.GetSMA, response.SMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.SMAURL),
		getEMA:      NewEndpoint[request.GetEMA, response.EMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.EMAURL),
		getADX:      NewEndpoint[request.GetADX, response.ADX, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.ADXURL),
		getMACD:     NewEndpoint[request.GetMACD, response.MACD, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.MACDURL),
		getRSI:      NewEndpoint[request.GetRSI, response.RSI, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.RSIURL),
		getStoch:    NewEndpoint[request.GetStoch, response.Stoch, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.StochURL),
		getPercentB: NewEndpoint[request.GetPercentB, response.PercentB, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.PercentBURL),
		getATR:      NewEndpoint[request.GetATR, response.ATR, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.ATRURL),
		getVWAP:     NewEndpoint[request.GetVWAP, response.VWAP, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.VWAPURL),
		getMA:       NewEndpoint[request.GetMA, response.MA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.MAURL),
		getWMA:      NewEndpoint[request.GetWMA, response.WMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.WMAURL),
		getDEMA:     NewEndpoint[request.GetDEMA, response.DEMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.DEMAURL),
		getTEMA:     NewEndpoint[request.GetTEMA, response.TEMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.TEMAURL),
		getTRMA:     NewEndpoint[request.GetTRMA, response.TRMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.TRMAURL),
		getKAMA:     NewEndpoint[request.GetKAMA, response.KAMA, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.KAMAURL),
		getSAR:      NewEndpoint[request.GetSAR, response.SAR, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.SARURL),
		getCCI:      NewEndpoint[request.GetCCI, response.CCI, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.CCIURL),
		getWillR:    NewEndpoint[request.GetWillR, response.WillR, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.WilliamsRURL),
		getROC:      NewEndpoint[request.GetROC, response.ROC, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.ROCURL),
		getMOM:      NewEndpoint[request.GetMOM, response.MOM, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.MomURL),
		getOBV:      NewEndpoint[request.GetOBV, response.OBV, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.OBVURL),
		getAD:       NewEndpoint[request.GetAD, response.AD, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.ADURL),
		getNATR:     NewEndpoint[request.GetNATR, response.NATR, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.NATRURL),
		getTR:       NewEndpoint[request.GetTR, response.TR, response.Credits, error](httpCli, cfg.BaseURL+cfg.TechnicalIndicators.TRURL),

		// Currencies
		getExchangeRate:       NewEndpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error](httpCli, cfg.BaseURL+cfg.Currencies.ExchangeRateURL),
		getCurrencyConversion: NewEndpoint[request.GetCurrencyConversion, response.CurrencyConversion, response.Credits, error](httpCli, cfg.BaseURL+cfg.Currencies.CurrencyConversionURL),

		// Advanced
		getUsage:   NewEndpoint[request.GetUsage, response.Usage, response.Credits, error](httpCli, cfg.BaseURL+cfg.Advanced.UsageURL),
		getBatches: NewEndpoint[request.GetBatches, response.Batches, response.Credits, error](httpCli, cfg.BaseURL+cfg.Advanced.BatchesURL),

		// ETFs
		getETFsDirectory:  NewEndpoint[request.GetETFsDirectory, response.ETFsDirectory, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsDirectoryURL),
		getETFFullData:    NewEndpoint[request.GetETFFullData, response.ETFFullData, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsFullDataURL),
		getETFSummary:     NewEndpoint[request.GetETFSummary, response.ETFWorldSummary, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsSummaryURL),
		getETFPerformance: NewEndpoint[request.GetETFPerformance, response.ETFPerformance, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsPerformanceURL),
		getETFRisk:        NewEndpoint[request.GetETFRisk, response.ETFRisk, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsRiskURL),
		getETFComposition: NewEndpoint[request.GetETFComposition, response.ETFComposition, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsCompositionURL),
		getETFFamilies:    NewEndpoint[request.GetETFFamilies, response.ETFFamilies, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsFamiliesURL),
		getETFTypes:       NewEndpoint[request.GetETFTypes, response.ETFTypes, response.Credits, error](httpCli, cfg.BaseURL+cfg.ETFs.ETFsTypesURL),

		// Mutual Funds
		getFunds:                    NewEndpoint[request.GetFunds, response.Funds, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsDirectoryURL),
		getMutualFunds:              NewEndpoint[request.GetMutualFunds, response.MutualFunds, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsDirectoryURL),
		getMutualFundFullData:       NewEndpoint[request.GetMutualFundFullData, response.MutualFundFullData, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsFullDataURL),
		getMutualFundSummary:        NewEndpoint[request.GetMutualFundSummary, response.MutualFundSummaryResponse, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsSummaryURL),
		getMutualFundPerformance:    NewEndpoint[request.GetMutualFundPerformance, response.MutualFundPerformance, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsPerformanceURL),
		getMutualFundRisk:           NewEndpoint[request.GetMutualFundRisk, response.MutualFundRiskResponse, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsRiskURL),
		getMutualFundRatings:        NewEndpoint[request.GetMutualFundRatings, response.MutualFundRatingsResponse, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsRatingsURL),
		getMutualFundComposition:    NewEndpoint[request.GetMutualFundComposition, response.MutualFundComposition, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsCompositionURL),
		getMutualFundPurchaseInfo:   NewEndpoint[request.GetMutualFundPurchaseInfo, response.MutualFundPurchaseInfoResponse, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsPurchaseInfoURL),
		getMutualFundSustainability: NewEndpoint[request.GetMutualFundSustainability, response.MutualFundSustainability, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsSustainabilityURL),
		getMutualFundFamilies:       NewEndpoint[request.GetMutualFundFamilies, response.MutualFundFamilies, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsFamiliesURL),
		getMutualFundTypes:          NewEndpoint[request.GetMutualFundTypes, response.MutualFundTypes, response.Credits, error](httpCli, cfg.BaseURL+cfg.MutualFunds.MutualFundsTypesURL),

		// Analysis
		getRecommendations:          NewEndpoint[request.GetRecommendations, response.Recommendations, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.RecommendationsURL),
		getPriceTarget:              NewEndpoint[request.GetPriceTarget, response.PriceTarget, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.PriceTargetURL),
		getEarningsEstimate:         NewEndpoint[request.GetEarningsEstimate, response.EarningsEstimate, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.EarningsEstimateURL),
		getRevenueEstimate:          NewEndpoint[request.GetRevenueEstimate, response.RevenueEstimate, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.RevenueEstimateURL),
		getEPSTrend:                 NewEndpoint[request.GetEPSTrend, response.EPSTrend, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.EPSTrendURL),
		getEPSRevisions:             NewEndpoint[request.GetEPSRevisions, response.EPSRevisions, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.EPSRevisionsURL),
		getGrowthEstimates:          NewEndpoint[request.GetGrowthEstimates, response.GrowthEstimates, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.GrowthEstimatesURL),
		getAnalystRatingsSnapshot:   NewEndpoint[request.GetAnalystRatingsSnapshot, response.AnalystRatingsSnapshot, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.AnalystRatingsSnapshotURL),
		getAnalystRatingsUSEquities: NewEndpoint[request.GetAnalystRatingsUSEquities, response.AnalystRatingsUSEquities, response.Credits, error](httpCli, cfg.BaseURL+cfg.Analysis.AnalystRatingsUSEquitiesURL),

		// Regulatory
		getInsiderTransactions:  NewEndpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.InsiderTransactionsURL),
		getEDGARFillings:        NewEndpoint[request.GetEDGARFillings, response.EDGARFillings, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.EDGARFillingsURL),
		getInstitutionalHolders: NewEndpoint[request.GetInstitutionalHolders, response.InstitutionalHolders, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.InstitutionalHoldersURL),
		getFundHolders:          NewEndpoint[request.GetFundHolders, response.FundHolders, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.FundHoldersURL),
		getDirectHolders:        NewEndpoint[request.GetDirectHolders, response.DirectHolders, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.DirectHoldersURL),
		getTaxInformation:       NewEndpoint[request.GetTaxInformation, response.TaxInformation, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.TaxInformationURL),
		getSanctionedEntities:   NewEndpoint[request.GetSanctionedEntities, response.SanctionedEntities, response.Credits, error](httpCli, cfg.BaseURL+cfg.Regulatory.SanctionedEntitiesURL),
	}
}
