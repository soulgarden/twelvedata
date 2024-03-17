package twelvedata

import (
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

type client struct {
	getStocks              *Endpoint[request.GetStock, response.Stocks, response.Credits, error]
	getTimeSeries          *Endpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error]
	getProfile             *Endpoint[request.GetProfile, response.Profile, response.Credits, error]
	getInsiderTransactions *Endpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error]
	getDividends           *Endpoint[request.GetDividends, response.Dividends, response.Credits, error]
	getStatistics          *Endpoint[request.GetStatistics, response.Statistics, response.Credits, error]
	getExchanges           *Endpoint[request.GetExchanges, response.Exchanges, response.Credits, error]
	getIndices             *Endpoint[request.GetIndices, response.Indices, response.Credits, error]
	getEtfs                *Endpoint[request.GetEtfs, response.Etfs, response.Credits, error]
	getQuote               *Endpoint[request.GetQuote, response.Quotes, response.Credits, error]
	getUsage               *Endpoint[request.GetUsage, response.Usage, response.Credits, error]
	getEarningsCalendar    *Endpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error]
	getExchangeRate        *Endpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error]
	getIncomeStatement     *Endpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error]
	getBalanceSheet        *Endpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error]
	getCashFlow            *Endpoint[request.GetCashFlow, response.CashFlows, response.Credits, error]
	getMarketMovers        *Endpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error]
	getMarketState         *Endpoint[request.GetMarketState, []response.MarketState, response.Credits, error]
}

func (client client) GetStocks(req request.GetStock) (response.Stocks, response.Credits, error) {
	return client.getStocks.Call(req)
}

func (client client) GetTimeSeries(req request.GetTimeSeries) (response.TimeSeries, response.Credits, error) {
	return client.getTimeSeries.Call(req)
}

func (client client) GetProfile(req request.GetProfile) (response.Profile, response.Credits, error) {
	return client.getProfile.Call(req)
}

func (client client) GetInsiderTransactions(req request.GetInsiderTransactions) (response.InsiderTransactions, response.Credits, error) {
	return client.getInsiderTransactions.Call(req)
}

func (client client) GetDividends(req request.GetDividends) (response.Dividends, response.Credits, error) {
	return client.getDividends.Call(req)
}

func (client client) GetStatistics(statistics request.GetStatistics) (response.Statistics, response.Credits, error) {
	return client.getStatistics.Call(statistics)
}

func (client client) GetExchanges(req request.GetExchanges) (response.Exchanges, response.Credits, error) {
	return client.getExchanges.Call(req)
}

func (client client) GetIndices(req request.GetIndices) (response.Indices, response.Credits, error) {
	return client.getIndices.Call(req)
}

func (client client) GetEtfs(req request.GetEtfs) (response.Etfs, response.Credits, error) {
	return client.getEtfs.Call(req)
}

func (client client) GetQuote(req request.GetQuote) (response.Quotes, response.Credits, error) {
	return client.getQuote.Call(req)
}

func (client client) GetUsage(req request.GetUsage) (response.Usage, response.Credits, error) {
	return client.getUsage.Call(req)
}

func (client client) GetEarningsCalendar(req request.GetEarningsCalendar) (response.Earnings, response.Credits, error) {
	return client.getEarningsCalendar.Call(req)
}

func (client client) GetExchangeRate(req request.GetExchangeRate) (response.ExchangeRate, response.Credits, error) {
	return client.getExchangeRate.Call(req)
}

func (client client) GetIncomeStatement(req request.GetIncomeStatement) (response.IncomeStatements, response.Credits, error) {
	return client.getIncomeStatement.Call(req)
}

func (client client) GetBalanceSheet(req request.GetBalanceSheet) (response.BalanceSheets, response.Credits, error) {
	return client.getBalanceSheet.Call(req)
}

func (client client) GetCashFlow(req request.GetCashFlow) (response.CashFlows, response.Credits, error) {
	return client.getCashFlow.Call(req)
}

func (client client) GetMarketMovers(req request.GetMarketMovers) (response.MarketMovers, response.Credits, error) {
	return client.getMarketMovers.Call(req)
}

func (client client) GetMarketState(req request.GetMarketState) ([]response.MarketState, response.Credits, error) {
	return client.getMarketState.Call(req)
}

func NewClient(httpCli *HTTPCli, cfg *Conf) Client {
	return client{
		getStocks: NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.StocksURL),
	}
}
