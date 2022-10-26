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

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

type Cli struct {
	cfg     *Conf
	httpCli *HTTPCli
	logger  *zerolog.Logger
}

func NewCli(cfg *Conf, httpCli *HTTPCli, logger *zerolog.Logger) *Cli {
	return &Cli{cfg: cfg, httpCli: httpCli, logger: logger}
}

func (c *Cli) GetStocks(symbol, exchange, micCode, country, instrumentType string, showPlan, includeDelisted bool) (
	stocksResp *response.Stocks,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.StocksURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{type}", instrumentType, 1)
	uri = strings.Replace(uri, "{show_plan}", strconv.FormatBool(showPlan), 1)
	uri = strings.Replace(uri, "{include_delisted}", strconv.FormatBool(includeDelisted), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &stocksResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return stocksResp, creditsLeft, creditsUsed, nil
}

// nolint: cyclop
func (c *Cli) GetTimeSeries(
	symbol, interval, exchange, micCode, country, instrumentType string,
	outputSize int,
	prePost string,
	decimalPlaces int,
	order, timezone, date, startDate, endDate string,
	previousClose bool,
) (seriesResp *response.TimeSeries, creditsLeft int64, creditsUsed int64, err error) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.TimeSeriesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{interval}", interval, 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{type}", instrumentType, 1)
	uri = strings.Replace(uri, "{outputsize}", strconv.Itoa(outputSize), 1)
	uri = strings.Replace(uri, "{prepost}", prePost, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{order}", order, 1)
	uri = strings.Replace(uri, "{timezone}", timezone, 1)
	uri = strings.Replace(uri, "{date}", date, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)
	uri = strings.Replace(uri, "{previous_close}", strconv.FormatBool(previousClose), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	//nolint: nestif
	if errResp, err := c.CheckErrorInResponse(resp); err != nil {
		if errResp != nil && errResp.Code == http.StatusBadRequest {
			if strings.Contains(errResp.Message, dictionary.SymbolNotFoundMsg) ||
				strings.Contains(errResp.Message, dictionary.NewSymbolNotFoundMsg) {
				return nil, creditsLeft, creditsUsed, dictionary.ErrNotFound
			} else if strings.Contains(errResp.Message, dictionary.IsNotAvailableWithYourPlanMsg) {
				return nil, creditsLeft, creditsUsed, dictionary.ErrIsNotAvailableWithYourPlan
			}
		}

		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &seriesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return seriesResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetProfile(symbol, exchange, micCode, country string) (
	profileResp *response.Profile,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.ProfileURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &profileResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return profileResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetInsiderTransactions(symbol, exchange, micCode, country string) (
	insiderTransactionsResp *response.InsiderTransactions,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.InsiderTransactionsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &insiderTransactionsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return insiderTransactionsResp, creditsLeft, creditsUsed, nil
}

// nolint: varnamelen
func (c *Cli) GetDividends(symbol, exchange, micCode, country, r, startDate, endDate string) (
	dividendsResp *response.Dividends,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.DividendsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{range}", r, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &dividendsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshall json: %w", err)
	}

	return dividendsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetStatistics(symbol, exchange, micCode, country string) (
	statisticsResp *response.Statistics,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.StatisticsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &statisticsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return statisticsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetExchanges(instrumentType, name, code, country string, showPlan bool) (
	exchangesResp *response.Exchanges,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.ExchangesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{type}", url.QueryEscape(instrumentType), 1)
	uri = strings.Replace(uri, "{name}", url.QueryEscape(name), 1)
	uri = strings.Replace(uri, "{code}", url.QueryEscape(code), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{show_plan}", strconv.FormatBool(showPlan), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &exchangesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return exchangesResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetIndices(symbol, country string, showPlan, includeDelisted bool) (
	indicesResp *response.Indices,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.IndicesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{show_plan}", strconv.FormatBool(showPlan), 1)
	uri = strings.Replace(uri, "{include_delisted}", strconv.FormatBool(includeDelisted), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &indicesResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return indicesResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetEtfs(symbol, exchange, micCode, country string, showPlan, includeDelisted bool) (
	etfsResp *response.Etfs,
	creditsLeft, creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.EtfsURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{show_plan}", strconv.FormatBool(showPlan), 1)
	uri = strings.Replace(uri, "{include_delisted}", strconv.FormatBool(includeDelisted), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &etfsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return etfsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetQuote(
	symbol, interval, exchange, micCode, country, volumeTimePeriod, instrumentType string,
	eod bool,
	rollingPeriod string,
	decimalPlaces int,
	timezone string,
) (
	quotes *response.Quotes,
	creditsLeft, creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.QuotesURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{interval}", interval, 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{volume_time_period}", volumeTimePeriod, 1)
	uri = strings.Replace(uri, "{type}", instrumentType, 1)
	uri = strings.Replace(uri, "{eod}", strconv.FormatBool(eod), 1)
	uri = strings.Replace(uri, "{rolling_period}", rollingPeriod, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{timezone}", timezone, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	// nolint: nestif
	if errResp, err := c.CheckErrorInResponse(resp); err != nil {
		if errResp != nil && errResp.Code == http.StatusBadRequest {
			if strings.Contains(errResp.Message, dictionary.SymbolNotFoundMsg) ||
				strings.Contains(errResp.Message, dictionary.NewSymbolNotFoundMsg) {
				return nil, creditsLeft, creditsUsed, dictionary.ErrNotFound
			} else if strings.Contains(errResp.Message, dictionary.IsNotAvailableWithYourPlanMsg) {
				return nil, creditsLeft, creditsUsed, dictionary.ErrIsNotAvailableWithYourPlan
			}
		}

		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	quotes, err = c.processQuote(resp)

	return quotes, creditsLeft, creditsUsed, err
}

func (c *Cli) GetUsage() (
	usageResp *response.Usage,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Advanced.UsageURL, "{apikey}", c.cfg.APIKey, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &usageResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return usageResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetEarningsCalendar(exchange, micCode, country string, decimalPlaces int, startDate, endDate string) (
	earningsResp *response.Earnings,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.EarningsCalendarURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	if _, err = c.CheckErrorInResponse(resp); err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &earningsResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return earningsResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetExchangeRate(
	symbol, date, timeZone string, decimalPlaces int,
) (
	exchangeRate *response.ExchangeRate,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.ExchangeRateURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{date}", url.QueryEscape(date), 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)
	uri = strings.Replace(uri, "{timezone}", timeZone, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	errResp, err := c.CheckErrorInResponse(resp)
	if err != nil {
		if errResp != nil &&
			errResp.Code == http.StatusBadRequest &&
			(strings.Contains(errResp.Message, dictionary.SymbolNotFoundMsg) ||
				strings.Contains(errResp.Message, dictionary.NewSymbolNotFoundMsg)) {
			return nil, creditsLeft, creditsUsed, dictionary.ErrNotFound
		}

		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &exchangeRate); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return exchangeRate, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetIncomeStatement(symbol, exchange, micCode, country, period, startDate, endDate string) (
	incomeStatementResp *response.IncomeStatements,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.IncomeStatementURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &incomeStatementResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return incomeStatementResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetBalanceSheet(symbol, exchange, micCode, country, period, startDate, endDate string) (
	balanceSheetResp *response.BalanceSheets,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.BalanceSheetURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &balanceSheetResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return balanceSheetResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetCashFlow(symbol, exchange, micCode, country, startDate, endDate, period string) (
	cashFlowResp *response.CashFlows,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.Fundamentals.CashFlowURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{symbol}", url.QueryEscape(symbol), 1)
	uri = strings.Replace(uri, "{exchange}", url.QueryEscape(exchange), 1)
	uri = strings.Replace(uri, "{mic_code}", url.QueryEscape(micCode), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{period}", period, 1)
	uri = strings.Replace(uri, "{start_date}", startDate, 1)
	uri = strings.Replace(uri, "{end_date}", endDate, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &cashFlowResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return cashFlowResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetMarketMovers(instrument, direction string, outputSize int, country string, decimalPlaces int) (
	marketMoversResp *response.MarketMovers,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.CoreData.MarketMoversURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{instrument}", instrument, 1)
	uri = strings.Replace(uri, "{direction}", direction, 1)
	uri = strings.Replace(uri, "{outputsize}", strconv.Itoa(outputSize), 1)
	uri = strings.Replace(uri, "{country}", country, 1)
	uri = strings.Replace(uri, "{dp}", strconv.Itoa(decimalPlaces), 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil {
		if !errors.Is(err, dictionary.ErrTooManyRequests) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &marketMoversResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return marketMoversResp, creditsLeft, creditsUsed, nil
}

func (c *Cli) GetMarketState(exchange, code, country string) (
	marketStateResp []response.MarketState,
	creditsLeft int64,
	creditsUsed int64,
	err error,
) {
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)

	uri := strings.Replace(c.cfg.BaseURL+c.cfg.ReferenceData.MarketStateURL, "{apikey}", c.cfg.APIKey, 1)
	uri = strings.Replace(uri, "{exchange}", exchange, 1)
	uri = strings.Replace(uri, "{code}", code, 1)
	uri = strings.Replace(uri, "{country}", country, 1)

	if creditsLeft, creditsUsed, err = c.httpCli.makeRequest(uri, resp); err != nil {
		return nil, 0, 0, err
	}

	_, err = c.CheckErrorInResponse(resp)
	if err != nil && !errors.Is(err, dictionary.ErrUnmarshalResponse) {
		if !errors.Is(err, dictionary.ErrTooManyRequests) {
			c.logger.Err(err).Msg("check error in response")
		}

		return nil, creditsLeft, creditsUsed, err
	}

	if err := json.Unmarshal(resp.Body(), &marketStateResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, creditsLeft, creditsUsed, fmt.Errorf("unmarshal json: %w", err)
	}

	return marketStateResp, creditsLeft, creditsUsed, nil
}

// nolint: cyclop
func (c *Cli) processQuote(resp *fasthttp.Response) (quotes *response.Quotes, err error) {
	var errResp *response.Error

	// nolint: nestif
	if errResp, err = c.CheckErrorInResponse(resp); err != nil {
		if errResp != nil && errResp.Code == http.StatusBadRequest {
			if strings.Contains(errResp.Message, dictionary.SymbolNotFoundMsg) ||
				strings.Contains(errResp.Message, dictionary.NewSymbolNotFoundMsg) {
				return quotes, dictionary.ErrNotFound
			} else if strings.Contains(errResp.Message, dictionary.IsNotAvailableWithYourPlanMsg) {
				return quotes, dictionary.ErrIsNotAvailableWithYourPlan
			}
		}

		if !errors.Is(err, dictionary.ErrTooManyRequests) && !errors.Is(err, dictionary.ErrNotFound) {
			c.logger.Err(err).Msg("check error in response")
		}

		return quotes, err
	}

	data := map[string]json.RawMessage{}

	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return quotes, fmt.Errorf("unmarshal json: %w", err)
	}

	quotes = &response.Quotes{}

	quotesErr := c.parseQuotes(data, quotes)
	if quotesErr != nil {
		quoteErr := c.parseQuote(resp, quotes)
		if quoteErr != nil {
			return quotes, quotesErr
		}
	}

	return quotes, err
}

func (c *Cli) parseQuote(
	resp *fasthttp.Response,
	quotes *response.Quotes,
) error {
	var quote response.Quote

	if err := json.Unmarshal(resp.Body(), &quote); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return fmt.Errorf("unmarshal json: %w", err)
	}

	quotes.Data = append(quotes.Data, quote)

	return nil
}

func (c *Cli) parseQuotes(
	data map[string]json.RawMessage,
	quotes *response.Quotes,
) error {
	var quoteResp response.Quote

	for _, item := range data {
		var quoteErr response.QuoteError

		if bytes.Contains(item, []byte(`"code":`)) {
			if err := json.Unmarshal(item, &quoteErr); err != nil {
				c.logger.Err(err).Bytes("val", item).Msg("unmarshall")

				return fmt.Errorf("unmarshal json: %w", err)
			}

			quotes.Errors = append(quotes.Errors, quoteErr)

			continue
		}

		if err := json.Unmarshal(item, &quoteResp); err != nil {
			c.logger.Err(err).Bytes("val", item).Msg("unmarshall")

			return fmt.Errorf("unmarshal json: %w", err)
		}

		quotes.Data = append(quotes.Data, quoteResp)
	}

	return nil
}

func (c *Cli) CheckErrorInResponse(resp *fasthttp.Response) (*response.Error, error) {
	var errResp *response.Error

	if bytes.Equal(resp.Body(), []byte("[]")) {
		return nil, dictionary.ErrNotFound
	}

	if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
		c.logger.Err(err).Bytes("body", resp.Body()).Msg("unmarshall")

		return nil, dictionary.ErrUnmarshalResponse
	}

	switch {
	case errResp.Code == http.StatusBadRequest:
		return errResp, dictionary.ErrInvalidTwelveDataResponse
	case errResp.Code == http.StatusTooManyRequests:
		return nil, dictionary.ErrTooManyRequests
	case errResp.Code == http.StatusForbidden:
		return errResp, dictionary.ErrForbidden
	case errResp.Code == http.StatusNotFound:
		return errResp, dictionary.ErrNotFound
	}

	return errResp, nil
}
