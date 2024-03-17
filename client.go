package twelvedata

import (
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

type Client interface {
	GetStocks(request.GetStock) (response.Stocks, response.Credits, error)
	GetTimeSeries(request.GetTimeSeries) (response.TimeSeries, response.Credits, error)
	GetProfile(request.GetProfile) (response.Profile, response.Credits, error)
	GetInsiderTransactions(request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error)
	GetDividends(request.GetDividends) (response.Dividends, response.Credits, error)
	GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error)
	GetExchanges(request.GetExchanges) (response.Exchanges, response.Credits, error)
	GetIndices(request.GetIndices) (response.Indices, response.Credits, error)
	GetEtfs(request.GetEtfs) (response.Etfs, response.Credits, error)
	GetQuote(request.GetQuote) (response.Quotes, response.Credits, error)
	GetUsage(request.GetUsage) (response.Usage, response.Credits, error)
	GetEarningsCalendar(request.GetEarningsCalendar) (response.Earnings, response.Credits, error)
	GetExchangeRate(request.GetExchangeRate) (response.ExchangeRate, response.Credits, error)
	GetIncomeStatement(request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error)
	GetBalanceSheet(request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error)
	GetCashFlow(request.GetCashFlow) (response.CashFlows, response.Credits, error)
	GetMarketMovers(request.GetMarketMovers) (response.MarketMovers, response.Credits, error)
	GetMarketState(request.GetMarketState) ([]response.MarketState, response.Credits, error)
}
