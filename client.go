package twelvedata

import (
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

type Client interface {
	// Reference Data - Asset Catalogs
	GetStocks(request.GetStock) (response.Stocks, response.Credits, error)
	GetForexPairs(request.GetForexPairs) (response.ForexPairs, response.Credits, error)
	GetCryptocurrencies(request.GetCryptocurrencies) (response.Cryptocurrencies, response.Credits, error)
	GetEtfs(request.GetEtfs) (response.Etfs, response.Credits, error)
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
	GetIndices(request.GetIndices) (response.Indices, response.Credits, error)
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
	GetExchangeRate(request.GetExchangeRate) (response.ExchangeRate, response.Credits, error)
	GetMarketMovers(request.GetMarketMovers) (response.MarketMovers, response.Credits, error)

	// Fundamentals
	GetProfile(request.GetProfile) (response.Profile, response.Credits, error)
	GetInsiderTransactions(request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error)
	GetDividends(request.GetDividends) (response.Dividends, response.Credits, error)
	GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error)
	GetEarningsCalendar(request.GetEarningsCalendar) (response.Earnings, response.Credits, error)
	GetIncomeStatement(request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error)
	GetBalanceSheet(request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error)
	GetCashFlow(request.GetCashFlow) (response.CashFlows, response.Credits, error)

	// Advanced
	GetUsage(request.GetUsage) (response.Usage, response.Credits, error)
}
