package twelvedata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

type Cli struct {
	cfg     *Conf
	httpCli *fasthttp.Client
	logger  *zerolog.Logger
}

func NewCli(cfg *Conf, httpCli *fasthttp.Client, logger *zerolog.Logger) *Cli {
	return &Cli{cfg: cfg, httpCli: httpCli, logger: logger}
}

func (c *Cli) GetStocks(symbol, exchange, country, instrumentType string) (
	stocksResp *response.Stocks,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.StocksURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{type}", instrumentType, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &stocksResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return stocksResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetTimeSeries(
	symbol, interval, exchange, country, instrumentType string,
	outputSize int,
	prePost string,
) (
	seriesResp *response.TimeSeries,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.TimeSeriesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{interval}", interval, 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{type}", instrumentType, 1)
	uri = strings.Replace(uri, "{outputsize}", strconv.Itoa(outputSize), 1)
	uri = strings.Replace(uri, "{prepost}", prePost, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, dictionary.ErrInvalidTwelveDataResponse
	}

	if err := json.Unmarshal(resp.Body(), &seriesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return seriesResp, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetProfile(symbol, exchange, country string) (
	profileResp *response.Profile,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.ProfileURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &profileResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return profileResp, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetInsiderTransactions(symbol, exchange, country string) (
	insiderTransactionsResp *response.InsiderTransactions,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.InsiderTransactionsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &insiderTransactionsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return insiderTransactionsResp, creditsLeft, creditsUsed, nil
}

// nolint: varnamelen,dupl
func (c *Cli) GetDividends(symbol, exchange, country, r, startTime, endTime string) (
	dividendsResp *response.Dividends,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.DividendsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{range}", r, 1)
	uri = strings.Replace(uri, "{start_time}", startTime, 1)
	uri = strings.Replace(uri, "{end_time}", endTime, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &dividendsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return dividendsResp, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetStatistics(symbol, exchange, country string) (
	statisticsResp *response.Statistics,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.StatisticsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &statisticsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return statisticsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetExchanges(instrumentType, name, code, country string) (
	exchangesResp *response.Exchanges,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.ExchangesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{type}", url.QueryEscape(instrumentType), 1)
	uri = strings.Replace(uri, "{name}", url.QueryEscape(name), 1)
	uri = strings.Replace(uri, "{code}", url.QueryEscape(code), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &exchangesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return exchangesResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetIndices(symbol, country string) (
	indicesResp *response.Indices,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.IndicesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &indicesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return indicesResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetEtfs(symbol string) (
	etfsResp *response.Etfs,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.EtfsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &etfsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return etfsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetQuotes(
	interval, exchange, country, volumeTimePeriod, instrumentType, prePost, timezone string,
	decimalPlaces int,
	symbols []string,
) (
	*response.Quotes,
	int,
	int,
	error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	var creditsLeft, creditsUsed int

	var err error

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.QuotesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(strings.Join(symbols, ",")), 1)
	uri = strings.Replace(uri, "{interval}", interval, 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{volume_time_period}", volumeTimePeriod, 1)
	uri = strings.Replace(uri, "{type}", url.QueryEscape(instrumentType), 1)
	uri = strings.Replace(uri, "{prepost}", prePost, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{timezone}", timezone, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	quotes, err := c.processQuotes(resp, symbols)

	return quotes, creditsLeft, creditsUsed, err
}

func (c *Cli) GetUsage() (
	usageResp *response.Usage,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Advanced.UsageURL, "{apikey}", c.cfg.APIKey, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, dictionary.ErrInvalidTwelveDataResponse
	}

	if err := json.Unmarshal(resp.Body(), &usageResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return usageResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetEarningsCalendar(decimalPlaces int, startDate, endDate string) (
	earningsResp *response.Earnings,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.EarningsCalendarURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &earningsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return earningsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetExchangeRate(
	symbol, timeZone string, precision int,
) (
	exchangeRate *response.ExchangeRate,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.ExchangeRateURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{precision}", strconv.Itoa(precision), 1)
	uri = strings.Replace(uri, "{timeZone}", timeZone, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, dictionary.ErrInvalidTwelveDataResponse
	}

	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return exchangeRate, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetIncomeStatement(symbol, exchange, country, period, startDate, endDate string) (
	incomeStatementResp *response.IncomeStatement,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.IncomeStatementURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &incomeStatementResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return incomeStatementResp, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetBalanceSheet(symbol, exchange, country, startDate, endDate string, period string) (
	balanceSheetResp *response.BalanceSheet,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.BalanceSheetURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &balanceSheetResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return balanceSheetResp, creditsLeft, creditsUsed, nil
}

// nolint:dupl
func (c *Cli) GetCashFlow(symbol, exchange, country, startDate, endDate string, period string) (
	cashFlowResp *response.CashFlow,
	creditsLeft int,
	creditsUsed int,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.CashFlowURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, creditsLeft, creditsUsed, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, creditsLeft, creditsUsed, nil
	}

	if err := json.Unmarshal(resp.Body(), &cashFlowResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return cashFlowResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) makeRequest(uri string, resp *fasthttp.Response) (int, int, error) {
	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri)

	start := time.Now()

	if err := c.httpCli.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
		c.logRequest(req, resp, time.Since(start), err)

		if !errors.Is(err, fasthttp.ErrDialTimeout) {
			return 0, 0, fmt.Errorf("http request: %w", err)
		}

		if err := c.httpCli.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
			return 0, 0, fmt.Errorf("http cli request: %w", err)
		}
	}

	if resp.StatusCode() != http.StatusOK {
		c.logRequest(req, resp, time.Since(start), dictionary.ErrBadStatusCode)

		return 0, 0, dictionary.ErrBadStatusCode
	}

	c.logRequest(req, resp, time.Since(start), nil)

	creditsLeft, creditsUsed, err := c.getCredits(resp)
	c.logger.Err(err).Msg("get credits")

	return creditsLeft, creditsUsed, nil
}

func (c *Cli) logRequest(
	req *fasthttp.Request,
	resp *fasthttp.Response,
	duration time.Duration,
	err error,
) {
	c.logger.Err(err).
		Str("request headers", req.Header.String()).
		Int("response code", resp.StatusCode()).
		Dur("duration", duration).
		Msg("request")
}

func (c *Cli) processQuotes(resp *fasthttp.Response, symbols []string) (*response.Quotes, error) {
	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		c.logger.Err(err).Msg("check error in response")

		return nil, err
	}

	if errResp.Code == http.StatusNotFound {
		return nil, dictionary.ErrInvalidTwelveDataResponse
	}

	data := map[string]json.RawMessage{}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	quotes := &response.Quotes{Data: []*response.Quote{}, Errors: []*response.QuoteError{}}

	if len(data) > len(symbols) { // one item
		var quote *response.Quote

		if err := json.Unmarshal(resp.Body(), quote); err != nil {
			c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")
		}

		quotes.Data = append(quotes.Data, quote)

		return quotes, nil
	}

	var quoteErr *response.QuoteError

	var quoteResp *response.Quote

	for _, item := range data {
		if bytes.Contains(item, []byte(`{"code":`)) {
			if err := json.Unmarshal(item, &quoteErr); err != nil {
				c.logger.Err(err).Bytes("val", item).Msg("unmarshall")

				return nil, fmt.Errorf("unmarshal json: %w", err)
			}

			quotes.Errors = append(quotes.Errors, quoteErr)

			continue
		}

		if err := json.Unmarshal(item, &quoteResp); err != nil {
			c.logger.Err(err).Bytes("val", item).Msg("unmarshall")

			return quotes, fmt.Errorf("unmarshal json: %w", err)
		}

		quotes.Data = append(quotes.Data, quoteResp)
	}

	return quotes, nil
}

func (c *Cli) CheckErrorInResponse(resp *fasthttp.Response) (*response.Error, error) {
	var errResp *response.Error

	if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	if errResp.Code == http.StatusBadRequest {
		return nil, dictionary.ErrInvalidTwelveDataResponse
	}

	return errResp, nil
}

func (c *Cli) getCredits(resp *fasthttp.Response) (creditsLeft int, creditsUsed int, err error) {
	creditsLeftStr := string(resp.Header.Peek("api-credits-left"))

	if creditsLeftStr != "" {
		creditsLeft, err = strconv.Atoi(creditsLeftStr)
		c.logger.Err(err).Str("val", creditsLeftStr).Msg("str to int")

		if err != nil {
			return 0, 0, fmt.Errorf("str to int: %w", err)
		}
	}

	creditsUsedStr := string(resp.Header.Peek("api-credits-used"))

	if creditsUsedStr != "" {
		creditsUsed, err = strconv.Atoi(creditsUsedStr)
		c.logger.Err(err).Str("val", creditsUsedStr).Msg("str to int")

		if err != nil {
			return creditsLeft, 0, fmt.Errorf("str to int: %w", err)
		}
	}

	return creditsLeft, creditsUsed, nil
}
