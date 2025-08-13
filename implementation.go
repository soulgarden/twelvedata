package twelvedata

import (
	"strings"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

type client struct {
	getStocks              *Endpoint[request.GetStock, response.Stocks, response.Credits, error]
	getTimeSeries          *Endpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error]
	getTimeSeriesCross     *Endpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error]
	getProfile             *Endpoint[request.GetProfile, response.Profile, response.Credits, error]
	getInsiderTransactions *Endpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error]
	getDividends           *Endpoint[request.GetDividends, response.Dividends, response.Credits, error]
	getStatistics          *Endpoint[request.GetStatistics, response.Statistics, response.Credits, error]
	getExchanges           *Endpoint[request.GetExchanges, response.Exchanges, response.Credits, error]
	getIndices             *Endpoint[request.GetIndices, response.Indices, response.Credits, error]
	getEtfs                *Endpoint[request.GetEtfs, response.Etfs, response.Credits, error]
	getQuote               *Endpoint[request.GetQuote, response.Quote, response.Credits, error]
	getUsage               *Endpoint[request.GetUsage, response.Usage, response.Credits, error]
	getEarningsCalendar    *Endpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error]
	getExchangeRate        *Endpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error]
	getMarketMovers        *Endpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error]
	getIncomeStatement     *Endpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error]
	getBalanceSheet        *Endpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error]
	getCashFlow            *Endpoint[request.GetCashFlow, response.CashFlows, response.Credits, error]
	getMarketState         *Endpoint[request.GetMarketState, []response.MarketState, response.Credits, error]
	getPrice               *Endpoint[request.GetPrice, response.Price, response.Credits, error]
	getEOD                 *Endpoint[request.GetEOD, response.EOD, response.Credits, error]
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

func NewClient(httpCli *HTTPCli, cfg *Conf) Client {
	return client{
		getStocks:              NewEndpoint[request.GetStock, response.Stocks, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.StocksURL),
		getTimeSeries:          NewEndpoint[request.GetTimeSeries, response.TimeSeries, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.TimeSeriesURL),
		getTimeSeriesCross:     NewEndpoint[request.GetTimeSeriesCross, response.TimeSeriesCross, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.TimeSeriesCrossURL),
		getProfile:             NewEndpoint[request.GetProfile, response.Profile, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.ProfileURL),
		getInsiderTransactions: NewEndpoint[request.GetInsiderTransactions, response.InsiderTransactions, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.InsiderTransactionsURL),
		getDividends:           NewEndpoint[request.GetDividends, response.Dividends, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.DividendsURL),
		getStatistics:          NewEndpoint[request.GetStatistics, response.Statistics, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.StatisticsURL),
		getExchanges:           NewEndpoint[request.GetExchanges, response.Exchanges, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.ExchangesURL),
		getIndices:             NewEndpoint[request.GetIndices, response.Indices, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.IndicesURL),
		getEtfs:                NewEndpoint[request.GetEtfs, response.Etfs, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.EtfsURL),
		getQuote:               NewEndpoint[request.GetQuote, response.Quote, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.QuotesURL),
		getUsage:               NewEndpoint[request.GetUsage, response.Usage, response.Credits, error](httpCli, cfg.BaseURL+cfg.Advanced.UsageURL),
		getEarningsCalendar:    NewEndpoint[request.GetEarningsCalendar, response.Earnings, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.EarningsCalendarURL),
		getExchangeRate:        NewEndpoint[request.GetExchangeRate, response.ExchangeRate, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.ExchangeRateURL),
		getMarketMovers:        NewEndpoint[request.GetMarketMovers, response.MarketMovers, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.MarketMoversURL),
		getIncomeStatement:     NewEndpoint[request.GetIncomeStatement, response.IncomeStatements, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.IncomeStatementURL),
		getBalanceSheet:        NewEndpoint[request.GetBalanceSheet, response.BalanceSheets, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.BalanceSheetURL),
		getCashFlow:            NewEndpoint[request.GetCashFlow, response.CashFlows, response.Credits, error](httpCli, cfg.BaseURL+cfg.Fundamentals.CashFlowURL),
		getMarketState:         NewEndpoint[request.GetMarketState, []response.MarketState, response.Credits, error](httpCli, cfg.BaseURL+cfg.ReferenceData.MarketStateURL),
		getPrice:               NewEndpoint[request.GetPrice, response.Price, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.PriceURL),
		getEOD:                 NewEndpoint[request.GetEOD, response.EOD, response.Credits, error](httpCli, cfg.BaseURL+cfg.CoreData.EODURL),
	}
}
