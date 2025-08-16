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
	getEtfs             *Endpoint[request.GetEtfs, response.Etfs, response.Credits, error]
	getFunds            *Endpoint[request.GetFunds, response.Funds, response.Credits, error]
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
	getIndices                 *Endpoint[request.GetIndices, response.Indices, response.Credits, error]
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
	getExchangeRate    *Endpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error]
	getMarketMovers    *Endpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error]

	// Fundamentals
	getProfile             *Endpoint[request.GetProfile, response.Profile, response.Credits, error]
	getInsiderTransactions *Endpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error]
	getDividends           *Endpoint[request.GetDividends, response.Dividends, response.Credits, error]
	getStatistics          *Endpoint[request.GetStatistics, response.Statistics, response.Credits, error]
	getEarningsCalendar    *Endpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error]
	getIncomeStatement     *Endpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error]
	getBalanceSheet        *Endpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error]
	getCashFlow            *Endpoint[request.GetCashFlow, response.CashFlows, response.Credits, error]

	// Advanced
	getUsage *Endpoint[request.GetUsage, response.Usage, response.Credits, error]
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

func (cli client) GetProfile(req request.GetProfile) (response.Profile, response.Credits, error) {
	return cli.getProfile.Call(req)
}

func (cli client) GetInsiderTransactions(req request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error) {
	return cli.getInsiderTransactions.Call(req)
}

func (cli client) GetDividends(req request.GetDividends) (response.Dividends, response.Credits, error) {
	return cli.getDividends.Call(req)
}

func (cli client) GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error) {
	return cli.getStatistics.Call(statistics)
}

func (cli client) GetExchanges(req request.GetExchanges) (response.Exchanges, response.Credits, error) {
	return cli.getExchanges.Call(req)
}

func (cli client) GetIndices(req request.GetIndices) (response.Indices, response.Credits, error) {
	return cli.getIndices.Call(req)
}

func (cli client) GetEtfs(req request.GetEtfs) (response.Etfs, response.Credits, error) {
	return cli.getEtfs.Call(req)
}

func (cli client) GetQuote(req request.GetQuote) (response.Quote, response.Credits, error) {
	return cli.getQuote.Call(req)
}

func (cli client) GetUsage(req request.GetUsage) (response.Usage, response.Credits, error) {
	return cli.getUsage.Call(req)
}

func (cli client) GetEarningsCalendar(req request.GetEarningsCalendar) (response.Earnings, response.Credits, error) {
	return cli.getEarningsCalendar.Call(req)
}

func (cli client) GetExchangeRate(req request.GetExchangeRate) (response.ExchangeRate, response.Credits, error) {
	return cli.getExchangeRate.Call(req)
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

func NewClient(httpCli *HTTPCli, cfg *Conf) Client {
	return client{
		// Reference Data - Asset Catalogs
		getStocks:           NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.StocksURL),
		getForexPairs:       NewEndpoint[request.GetForexPairs, response.ForexPairs, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.ForexPairsURL),
		getCryptocurrencies: NewEndpoint[request.GetCryptocurrencies, response.Cryptocurrencies, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.CryptocurrenciesURL),
		getEtfs:             NewEndpoint[request.GetEtfs, response.Etfs, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.EtfsURL),
		getFunds:            NewEndpoint[request.GetFunds, response.Funds, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.FundsURL),
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
		getIndices:                 NewEndpoint[request.GetIndices, response.Indices, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.IndicesURL),
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
		getExchangeRate:    NewEndpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.ExchangeRateURL),
		getMarketMovers:    NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.MarketMoversURL),

		// Fundamentals
		getProfile:             NewEndpoint[request.GetProfile, response.Profile, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.ProfileURL),
		getInsiderTransactions: NewEndpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.InsiderTransactionsURL),
		getDividends:           NewEndpoint[request.GetDividends, response.Dividends, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.DividendsURL),
		getStatistics:          NewEndpoint[request.GetStatistics, response.Statistics, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.StatisticsURL),
		getEarningsCalendar:    NewEndpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.EarningsCalendarURL),
		getIncomeStatement:     NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.IncomeStatementURL),
		getBalanceSheet:        NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.BalanceSheetURL),
		getCashFlow:            NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.CashFlowURL),

		// Advanced
		getUsage: NewEndpoint[request.GetUsage, response.Usage, response.Credits, error](httpCli, cfg.BaseURL+cfg.Advanced.UsageURL),
	}
}
