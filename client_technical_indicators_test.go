//nolint:dupl // Technical indicator tests share similar structure by design
package twelvedata

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func buildExpectedURL(path string, params url.Values) string {
	if len(params) == 0 {
		return path
	}
	return path + "?" + params.Encode()
}

func technicalIndicatorParams(extra url.Values) url.Values {
	params := url.Values{
		"adjust":         []string{"splits"},
		"country":        []string{"United States"},
		"cusip":          []string{"037833100"},
		"date":           []string{"2024-01-02"},
		"delimiter":      []string{"comma"},
		"dp":             []string{"4"},
		"end_date":       []string{"2024-02-01"},
		"exchange":       []string{"NASDAQ"},
		"figi":           []string{"BBG000B9XRY4"},
		"format":         []string{"json"},
		"include_ohlc":   []string{"true"},
		"interval":       []string{"1day"},
		"isin":           []string{"US0378331005"},
		"mic_code":       []string{"XNGS"},
		"order":          []string{"asc"},
		"outputsize":     []string{"120"},
		"prepost":        []string{"true"},
		"previous_close": []string{"true"},
		"start_date":     []string{"2024-01-01"},
		"symbol":         []string{"AAPL"},
		"timezone":       []string{"America/New_York"},
		"type":           []string{"stock"},
	}

	for key, values := range extra {
		params[key] = values
	}

	return params
}

func Test_client_GetBBands(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetBBands, response.BBands]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetBBands]{
				req: request.GetBBands{
					APIKey:             request.APIKey{APIKey: ""},
					Symbol:             "AAPL",
					FIGI:               "BBG000B9XRY4",
					ISIN:               "US0378331005",
					CUSIP:              "037833100",
					Interval:           "1day",
					Exchange:           "NASDAQ",
					MICCode:            "XNGS",
					Country:            "United States",
					MAType:             "SMA",
					StandardDeviations: 2.5,
					SeriesType:         "close",
					TimePeriod:         20,
					Type:               "stock",
					OutputSize:         120,
					Format:             "json",
					Delimiter:          "comma",
					Prepost:            true,
					DP:                 4,
					Order:              "asc",
					IncludeOHLC:        true,
					Timezone:           "America/New_York",
					Date:               "2024-01-02",
					StartDate:          "2024-01-01",
					EndDate:            "2024-02-01",
					PreviousClose:      true,
					Adjust:             "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "BBANDS - Bollinger Bands®",
					      "series_type": "close",
					      "time_period": 20,
					      "sd": 2,
					      "ma_type": "SMA"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "upper_band": "241.35279",
					      "middle_band": "219.69950",
					      "lower_band": "198.046208"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/bbands",
						technicalIndicatorParams(url.Values{
							"ma_type":     []string{"SMA"},
							"sd":          []string{"2.500000"},
							"series_type": []string{"close"},
							"time_period": []string{"20"},
						}),
					),
				),
			},
			want: response.BBands{
				Meta: response.BBandsMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.BbandsIndicator{
						Name:       "BBANDS - Bollinger Bands®",
						SeriesType: "close",
						TimePeriod: 20,
						SD:         2,
						MAType:     "SMA",
					},
				},
				Values: []response.BbandsData{
					{
						Datetime:   "2025-08-21",
						UpperBand:  "241.35279",
						MiddleBand: "219.69950",
						LowerBand:  "198.046208",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/bbands",
		"GetBBands",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getBBands: NewEndpoint[request.GetBBands, response.BBands, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetBBands) (response.BBands, response.Credits, error) {
			return cli.(client).GetBBands(req)
		},
	)
}

func Test_client_GetSMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetSMA, response.SMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetSMA]{
				req: request.GetSMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    9,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "SMA - Simple Moving Average",
					      "series_type": "close",
					      "time_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "sma": "229.65444"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/sma",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"9"},
						}),
					),
				),
			},
			want: response.SMA{
				Meta: response.SMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.SMAIndicator{
						Name:       "SMA - Simple Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
					},
				},
				Values: []response.SMAData{
					{
						Datetime: "2025-08-21",
						SMA:      "229.65444",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/sma",
		"GetSMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getSMA: NewEndpoint[request.GetSMA, response.SMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetSMA) (response.SMA, response.Credits, error) {
			return cli.(client).GetSMA(req)
		},
	)
}

func Test_client_GetEMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetEMA, response.EMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetEMA]{
				req: request.GetEMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    9,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "EMA - Exponential Moving Average",
					      "series_type": "close",
					      "time_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "ema": "234.81765"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/ema",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"9"},
						}),
					),
				),
			},
			want: response.EMA{
				Meta: response.EMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.EMAIndicator{
						Name:       "EMA - Exponential Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
					},
				},
				Values: []response.EMAData{
					{
						Datetime: "2025-08-21",
						EMA:      "234.81765",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/ema",
		"GetEMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getEMA: NewEndpoint[request.GetEMA, response.EMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetEMA) (response.EMA, response.Credits, error) {
			return cli.(client).GetEMA(req)
		},
	)
}

func Test_client_GetMACD(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetMACD, response.MACD]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetMACD]{
				req: request.GetMACD{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					FastPeriod:    12,
					SlowPeriod:    26,
					SignalPeriod:  9,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "MACD - Moving Average Convergence/Divergence",
					      "series_type": "close",
					      "fast_period": 12,
					      "slow_period": 26,
					      "signal_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "macd": "9.15434",
					      "macd_signal": "7.89123",
					      "macd_hist": "1.26311"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/macd",
						technicalIndicatorParams(url.Values{
							"fast_period":   []string{"12"},
							"series_type":   []string{"close"},
							"signal_period": []string{"9"},
							"slow_period":   []string{"26"},
						}),
					),
				),
			},
			want: response.MACD{
				Meta: response.MACDMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.MACDIndicator{
						Name:         "MACD - Moving Average Convergence/Divergence",
						SeriesType:   "close",
						FastPeriod:   12,
						SlowPeriod:   26,
						SignalPeriod: 9,
					},
				},
				Values: []response.MACDData{
					{
						Datetime:   "2025-08-21",
						MACD:       "9.15434",
						MACDSignal: "7.89123",
						MACDHist:   "1.26311",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/macd",
		"GetMACD",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMACD: NewEndpoint[request.GetMACD, response.MACD, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMACD) (response.MACD, response.Credits, error) {
			return cli.(client).GetMACD(req)
		},
	)
}

func Test_client_GetRSI(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetRSI, response.RSI]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetRSI]{
				req: request.GetRSI{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "RSI - Relative Strength Index",
					      "series_type": "close",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "rsi": "57.99115"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/rsi",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.RSI{
				Meta: response.RSIMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.RSIIndicator{
						Name:       "RSI - Relative Strength Index",
						SeriesType: "close",
						TimePeriod: 14,
					},
				},
				Values: []response.RSIData{
					{
						Datetime: "2025-08-21",
						RSI:      "57.99115",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/rsi",
		"GetRSI",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getRSI: NewEndpoint[request.GetRSI, response.RSI, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetRSI) (response.RSI, response.Credits, error) {
			return cli.(client).GetRSI(req)
		},
	)
}

func Test_client_GetATR(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetATR, response.ATR]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetATR]{
				req: request.GetATR{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "ATR - Average True Range",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "atr": "5.42156"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/atr",
						technicalIndicatorParams(url.Values{
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.ATR{
				Meta: response.ATRMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.ATRIndicator{
						Name:       "ATR - Average True Range",
						TimePeriod: 14,
					},
				},
				Values: []response.ATRValue{
					{
						Datetime: "2025-08-21",
						ATR:      "5.42156",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/atr",
		"GetATR",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getATR: NewEndpoint[request.GetATR, response.ATR, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetATR) (response.ATR, response.Credits, error) {
			return cli.(client).GetATR(req)
		},
	)
}

func Test_client_GetCCI(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetCCI, response.CCI]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetCCI]{
				req: request.GetCCI{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					TimePeriod:    20,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "CCI - Commodity Channel Index",
					      "time_period": 20
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "cci": "123.45678"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/cci",
						technicalIndicatorParams(url.Values{
							"time_period": []string{"20"},
						}),
					),
				),
			},
			want: response.CCI{
				Meta: response.CCIMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.CCIIndicator{
						Name:       "CCI - Commodity Channel Index",
						TimePeriod: 20,
					},
				},
				Values: []response.CCIValue{
					{
						Datetime: "2025-08-21",
						CCI:      "123.45678",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/cci",
		"GetCCI",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getCCI: NewEndpoint[request.GetCCI, response.CCI, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetCCI) (response.CCI, response.Credits, error) {
			return cli.(client).GetCCI(req)
		},
	)
}

func Test_client_GetDEMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetDEMA, response.DEMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetDEMA]{
				req: request.GetDEMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    9,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "DEMA - Double Exponential Moving Average",
					      "series_type": "close",
					      "time_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "dema": "234.81765"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/dema",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"9"},
						}),
					),
				),
			},
			want: response.DEMA{
				Meta: response.DEMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.DEMAIndicator{
						Name:       "DEMA - Double Exponential Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
					},
				},
				Values: []response.DEMAValue{
					{
						Datetime: "2025-08-21",
						DEMA:     "234.81765",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/dema",
		"GetDEMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getDEMA: NewEndpoint[request.GetDEMA, response.DEMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetDEMA) (response.DEMA, response.Credits, error) {
			return cli.(client).GetDEMA(req)
		},
	)
}

func Test_client_GetKAMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetKAMA, response.KAMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetKAMA]{
				req: request.GetKAMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    10,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "KAMA - Kaufman Adaptive Moving Average",
					      "series_type": "close",
					      "time_period": 10
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "kama": "218.92345"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/kama",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"10"},
						}),
					),
				),
			},
			want: response.KAMA{
				Meta: response.KAMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.KAMAIndicator{
						Name:       "KAMA - Kaufman Adaptive Moving Average",
						SeriesType: "close",
						TimePeriod: 10,
					},
				},
				Values: []response.KAMAValue{
					{
						Datetime: "2025-08-21",
						KAMA:     "218.92345",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/kama",
		"GetKAMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getKAMA: NewEndpoint[request.GetKAMA, response.KAMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetKAMA) (response.KAMA, response.Credits, error) {
			return cli.(client).GetKAMA(req)
		},
	)
}

func Test_client_GetMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetMA, response.MA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetMA]{
				req: request.GetMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    9,
					MAType:        "SMA",
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "MA - Moving Average",
					      "series_type": "close",
					      "time_period": 9,
					      "ma_type": "SMA"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "ma": "219.69950"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/ma",
						technicalIndicatorParams(url.Values{
							"ma_type":     []string{"SMA"},
							"series_type": []string{"close"},
							"time_period": []string{"9"},
						}),
					),
				),
			},
			want: response.MA{
				Meta: response.MAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.MAIndicator{
						Name:       "MA - Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
						MAType:     "SMA",
					},
				},
				Values: []response.MAValue{
					{
						Datetime: "2025-08-21",
						MA:       "219.69950",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/ma",
		"GetMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMA: NewEndpoint[request.GetMA, response.MA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMA) (response.MA, response.Credits, error) {
			return cli.(client).GetMA(req)
		},
	)
}

func Test_client_GetSAR(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetSAR, response.SAR]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetSAR]{
				req: request.GetSAR{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					Acceleration:  0.02,
					Maximum:       0.2,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "SAR - Parabolic SAR",
					      "acceleration": 0.02,
					      "maximum": 0.2
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "sar": "210.45123"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/sar",
						technicalIndicatorParams(url.Values{
							"acceleration": []string{"0.020000"},
							"maximum":      []string{"0.200000"},
						}),
					),
				),
			},
			want: response.SAR{
				Meta: response.SARMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.SARIndicator{
						Name:         "SAR - Parabolic SAR",
						Acceleration: 0.02,
						Maximum:      0.2,
					},
				},
				Values: []response.SARValue{
					{
						Datetime: "2025-08-21",
						SAR:      "210.45123",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/sar",
		"GetSAR",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getSAR: NewEndpoint[request.GetSAR, response.SAR, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetSAR) (response.SAR, response.Credits, error) {
			return cli.(client).GetSAR(req)
		},
	)
}

func Test_client_GetTEMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetTEMA, response.TEMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetTEMA]{
				req: request.GetTEMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    10,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "TEMA - Triple Exponential Moving Average",
					      "series_type": "close",
					      "time_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "tema": "234.91876"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/tema",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"10"},
						}),
					),
				),
			},
			want: response.TEMA{
				Meta: response.TEMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.TEMAIndicator{
						Name:       "TEMA - Triple Exponential Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
					},
				},
				Values: []response.TEMAValue{
					{
						Datetime: "2025-08-21",
						TEMA:     "234.91876",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/tema",
		"GetTEMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getTEMA: NewEndpoint[request.GetTEMA, response.TEMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetTEMA) (response.TEMA, response.Credits, error) {
			return cli.(client).GetTEMA(req)
		},
	)
}

func Test_client_GetTRMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetTRMA, response.TRMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetTRMA]{
				req: request.GetTRMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "TRMA - Triangular Moving Average",
					      "series_type": "close",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "trma": "218.56789"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/trima",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.TRMA{
				Meta: response.TRMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.TRMAIndicator{
						Name:       "TRMA - Triangular Moving Average",
						SeriesType: "close",
						TimePeriod: 14,
					},
				},
				Values: []response.TRMAValue{
					{
						Datetime: "2025-08-21",
						TRMA:     "218.56789",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/trima",
		"GetTRMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getTRMA: NewEndpoint[request.GetTRMA, response.TRMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetTRMA) (response.TRMA, response.Credits, error) {
			return cli.(client).GetTRMA(req)
		},
	)
}

func Test_client_GetVWAP(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetVWAP, response.VWAP]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetVWAP]{
				req: request.GetVWAP{
					APIKey:             request.APIKey{APIKey: ""},
					Symbol:             "AAPL",
					FIGI:               "BBG000B9XRY4",
					ISIN:               "US0378331005",
					CUSIP:              "037833100",
					Interval:           "1day",
					Exchange:           "NASDAQ",
					MICCode:            "XNGS",
					Country:            "United States",
					StandardDeviations: 2.5,
					SDTimePeriod:       14,
					Type:               "stock",
					OutputSize:         120,
					Format:             "json",
					Delimiter:          "comma",
					Prepost:            true,
					DP:                 4,
					Order:              "asc",
					IncludeOHLC:        true,
					Timezone:           "America/New_York",
					Date:               "2024-01-02",
					StartDate:          "2024-01-01",
					EndDate:            "2024-02-01",
					PreviousClose:      true,
					Adjust:             "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "VWAP - Volume Weighted Average Price",
					      "sd_time_period": 14,
					      "sd": 2.5
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "vwap_lower": "219.12345",
					      "vwap": "220.12345",
					      "vwap_upper": "221.12345"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/vwap",
						technicalIndicatorParams(url.Values{
							"sd":             []string{"2.500000"},
							"sd_time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.VWAP{
				Meta: response.VWAPMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.VWAPIndicator{
						Name:         "VWAP - Volume Weighted Average Price",
						SDTimePeriod: null.IntFrom(14),
						SD:           response.FloatStringFrom(2.5),
					},
				},
				Values: []response.VWAPValue{
					{
						Datetime:  "2025-08-21",
						VWAPLower: response.FloatStringFrom(219.12345),
						VWAP:      response.FloatStringFrom(220.12345),
						VWAPUpper: response.FloatStringFrom(221.12345),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/vwap",
		"GetVWAP",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getVWAP: NewEndpoint[request.GetVWAP, response.VWAP, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetVWAP) (response.VWAP, response.Credits, error) {
			return cli.(client).GetVWAP(req)
		},
	)
}

func Test_client_GetWMA(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetWMA, response.WMA]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetWMA]{
				req: request.GetWMA{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    9,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "WMA - Weighted Moving Average",
					      "series_type": "close",
					      "time_period": 9
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "wma": "219.85432"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/wma",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"9"},
						}),
					),
				),
			},
			want: response.WMA{
				Meta: response.WMAMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.WMAIndicator{
						Name:       "WMA - Weighted Moving Average",
						SeriesType: "close",
						TimePeriod: 9,
					},
				},
				Values: []response.WMAValue{
					{
						Datetime: "2025-08-21",
						WMA:      "219.85432",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/wma",
		"GetWMA",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getWMA: NewEndpoint[request.GetWMA, response.WMA, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetWMA) (response.WMA, response.Credits, error) {
			return cli.(client).GetWMA(req)
		},
	)
}

func Test_client_GetADX(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetADX, response.ADX]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetADX]{
				req: request.GetADX{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "ADX - Average Directional Index",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "adx": "25.1234"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/adx",
						technicalIndicatorParams(url.Values{
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.ADX{
				Meta: response.ADXMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.ADXIndicator{
						Name:       "ADX - Average Directional Index",
						TimePeriod: 14,
					},
				},
				Values: []response.ADXData{
					{
						Datetime: "2025-08-21",
						ADX:      "25.1234",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/adx",
		"GetADX",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getADX: NewEndpoint[request.GetADX, response.ADX, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetADX) (response.ADX, response.Credits, error) {
			return cli.(client).GetADX(req)
		},
	)
}

func Test_client_GetStoch(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetStoch, response.Stoch]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetStoch]{
				req: request.GetStoch{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					FastKPeriod:   14,
					SlowKPeriod:   3,
					SlowDPeriod:   3,
					SlowKMAType:   "SMA",
					SlowDMAType:   "SMA",
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "STOCH - Stochastic Oscillator",
					      "fast_k_period": 14,
					      "slow_k_period": 3,
					      "slow_d_period": 3,
					      "slow_kma_type": "SMA",
					      "slow_dma_type": "SMA"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "slow_k": "78.1234",
					      "slow_d": "75.9876"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/stoch",
						technicalIndicatorParams(url.Values{
							"fast_k_period": []string{"14"},
							"slow_d_period": []string{"3"},
							"slow_dma_type": []string{"SMA"},
							"slow_k_period": []string{"3"},
							"slow_kma_type": []string{"SMA"},
						}),
					),
				),
			},
			want: response.Stoch{
				Meta: response.StochMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.StochIndicator{
						Name:        "STOCH - Stochastic Oscillator",
						FastKPeriod: 14,
						SlowKPeriod: 3,
						SlowDPeriod: 3,
						SlowKMAType: "SMA",
						SlowDMAType: "SMA",
					},
				},
				Values: []response.StochData{
					{
						Datetime: "2025-08-21",
						SlowK:    "78.1234",
						SlowD:    "75.9876",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/stoch",
		"GetStoch",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getStoch: NewEndpoint[request.GetStoch, response.Stoch, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetStoch) (response.Stoch, response.Credits, error) {
			return cli.(client).GetStoch(req)
		},
	)
}

func Test_client_GetPercentB(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetPercentB, response.PercentB]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetPercentB]{
				req: request.GetPercentB{
					APIKey:             request.APIKey{APIKey: ""},
					Symbol:             "AAPL",
					FIGI:               "BBG000B9XRY4",
					ISIN:               "US0378331005",
					CUSIP:              "037833100",
					Interval:           "1day",
					Exchange:           "NASDAQ",
					MICCode:            "XNGS",
					Country:            "United States",
					SeriesType:         "close",
					TimePeriod:         20,
					StandardDeviations: 2.5,
					MAType:             "SMA",
					Type:               "stock",
					OutputSize:         120,
					Format:             "json",
					Delimiter:          "comma",
					Prepost:            true,
					DP:                 4,
					Order:              "asc",
					IncludeOHLC:        true,
					Timezone:           "America/New_York",
					Date:               "2024-01-02",
					StartDate:          "2024-01-01",
					EndDate:            "2024-02-01",
					PreviousClose:      true,
					Adjust:             "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "%B - Bollinger Bands %B",
					      "series_type": "close",
					      "time_period": 20,
					      "sd": 2,
					      "ma_type": "SMA"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "percent_b": "0.42"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/percent_b",
						technicalIndicatorParams(url.Values{
							"ma_type":     []string{"SMA"},
							"sd":          []string{"2.500000"},
							"series_type": []string{"close"},
							"time_period": []string{"20"},
						}),
					),
				),
			},
			want: response.PercentB{
				Meta: response.PercentBMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.PercentBIndicator{
						Name:       "%B - Bollinger Bands %B",
						SeriesType: "close",
						TimePeriod: 20,
						SD:         2,
						MAType:     "SMA",
					},
				},
				Values: []response.PercentBData{
					{
						Datetime: "2025-08-21",
						PercentB: "0.42",
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/percent_b",
		"GetPercentB",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getPercentB: NewEndpoint[request.GetPercentB, response.PercentB, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetPercentB) (response.PercentB, response.Credits, error) {
			return cli.(client).GetPercentB(req)
		},
	)
}

func Test_client_GetWillR(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetWillR, response.WillR]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetWillR]{
				req: request.GetWillR{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "WILLR - Williams %R",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "williams_r": -35.5
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/willr",
						technicalIndicatorParams(url.Values{
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.WillR{
				Meta: response.WillRMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.WillRIndicator{
						Name:       "WILLR - Williams %R",
						TimePeriod: 14,
					},
				},
				Values: []response.WillRValue{
					{
						Datetime:  "2025-08-21",
						WilliamsR: null.FloatFrom(-35.5),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/willr",
		"GetWillR",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getWillR: NewEndpoint[request.GetWillR, response.WillR, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetWillR) (response.WillR, response.Credits, error) {
			return cli.(client).GetWillR(req)
		},
	)
}

func Test_client_GetROC(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetROC, response.ROC]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetROC]{
				req: request.GetROC{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    10,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "ROC - Rate of Change",
					      "series_type": "close",
					      "time_period": 10
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "roc": 1.23
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/roc",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"10"},
						}),
					),
				),
			},
			want: response.ROC{
				Meta: response.ROCMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.ROCIndicator{
						Name:       "ROC - Rate of Change",
						SeriesType: "close",
						TimePeriod: 10,
					},
				},
				Values: []response.ROCValue{
					{
						Datetime: "2025-08-21",
						ROC:      null.FloatFrom(1.23),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/roc",
		"GetROC",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getROC: NewEndpoint[request.GetROC, response.ROC, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetROC) (response.ROC, response.Credits, error) {
			return cli.(client).GetROC(req)
		},
	)
}

func Test_client_GetMOM(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetMOM, response.MOM]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetMOM]{
				req: request.GetMOM{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					TimePeriod:    10,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "MOM - Momentum",
					      "series_type": "close",
					      "time_period": 10
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "mom": 2.34
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/mom",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
							"time_period": []string{"10"},
						}),
					),
				),
			},
			want: response.MOM{
				Meta: response.MOMMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.MOMIndicator{
						Name:       "MOM - Momentum",
						SeriesType: "close",
						TimePeriod: 10,
					},
				},
				Values: []response.MOMValue{
					{
						Datetime: "2025-08-21",
						MOM:      null.FloatFrom(2.34),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/mom",
		"GetMOM",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getMOM: NewEndpoint[request.GetMOM, response.MOM, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetMOM) (response.MOM, response.Credits, error) {
			return cli.(client).GetMOM(req)
		},
	)
}

func Test_client_GetOBV(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetOBV, response.OBV]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetOBV]{
				req: request.GetOBV{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					SeriesType:    "close",
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "OBV - On Balance Volume",
					      "series_type": "close"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "obv": "123456.78"
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/obv",
						technicalIndicatorParams(url.Values{
							"series_type": []string{"close"},
						}),
					),
				),
			},
			want: response.OBV{
				Meta: response.OBVMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.OBVIndicator{
						Name:       "OBV - On Balance Volume",
						SeriesType: "close",
					},
				},
				Values: []response.OBVValue{
					{
						Datetime: "2025-08-21",
						OBV:      response.FloatStringFrom(123456.78),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/obv",
		"GetOBV",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getOBV: NewEndpoint[request.GetOBV, response.OBV, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetOBV) (response.OBV, response.Credits, error) {
			return cli.(client).GetOBV(req)
		},
	)
}

func Test_client_GetAD(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetAD, response.AD]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetAD]{
				req: request.GetAD{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "AD - Accumulation/Distribution"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "ad": 98765.43
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL("/ad", technicalIndicatorParams(url.Values{})),
				),
			},
			want: response.AD{
				Meta: response.ADMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.ADIndicator{
						Name: "AD - Accumulation/Distribution",
					},
				},
				Values: []response.ADValue{
					{
						Datetime: "2025-08-21",
						AD:       null.FloatFrom(98765.43),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/ad",
		"GetAD",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getAD: NewEndpoint[request.GetAD, response.AD, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetAD) (response.AD, response.Credits, error) {
			return cli.(client).GetAD(req)
		},
	)
}

func Test_client_GetNATR(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetNATR, response.NATR]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetNATR]{
				req: request.GetNATR{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					TimePeriod:    14,
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "NATR - Normalized Average True Range",
					      "time_period": 14
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "natr": 1.2345
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL(
						"/natr",
						technicalIndicatorParams(url.Values{
							"time_period": []string{"14"},
						}),
					),
				),
			},
			want: response.NATR{
				Meta: response.NATRMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.NATRIndicator{
						Name:       "NATR - Normalized Average True Range",
						TimePeriod: 14,
					},
				},
				Values: []response.NATRValue{
					{
						Datetime: "2025-08-21",
						NATR:     null.FloatFrom(1.2345),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/natr",
		"GetNATR",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getNATR: NewEndpoint[request.GetNATR, response.NATR, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetNATR) (response.NATR, response.Credits, error) {
			return cli.(client).GetNATR(req)
		},
	)
}

func Test_client_GetTR(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetTR, response.TR]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetTR]{
				req: request.GetTR{
					APIKey:        request.APIKey{APIKey: ""},
					Symbol:        "AAPL",
					FIGI:          "BBG000B9XRY4",
					ISIN:          "US0378331005",
					CUSIP:         "037833100",
					Interval:      "1day",
					Exchange:      "NASDAQ",
					MICCode:       "XNGS",
					Country:       "United States",
					Type:          "stock",
					OutputSize:    120,
					Format:        "json",
					Delimiter:     "comma",
					Prepost:       true,
					DP:            4,
					Order:         "asc",
					IncludeOHLC:   true,
					Timezone:      "America/New_York",
					Date:          "2024-01-02",
					StartDate:     "2024-01-01",
					EndDate:       "2024-02-01",
					PreviousClose: true,
					Adjust:        "splits",
				},
				url: mockServerWithURL(
					t,
					http.StatusOK,
					100,
					10,
					`{
					  "meta": {
					    "symbol": "AAPL",
					    "interval": "1day",
					    "currency": "USD",
					    "exchange_timezone": "America/New_York",
					    "exchange": "NASDAQ",
					    "mic_code": "XNGS",
					    "type": "Common Stock",
					    "indicator": {
					      "name": "TR - True Range"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "tr": 5.6789
					    }
					  ],
					  "status": "ok"
					}`,
					buildExpectedURL("/trange", technicalIndicatorParams(url.Values{})),
				),
			},
			want: response.TR{
				Meta: response.TRMeta{
					Symbol:           "AAPL",
					Interval:         "1day",
					Currency:         "USD",
					ExchangeTimezone: "America/New_York",
					Exchange:         "NASDAQ",
					MicCode:          "XNGS",
					Type:             "Common Stock",
					Indicator: response.TRIndicator{
						Name: "TR - True Range",
					},
				},
				Values: []response.TRValue{
					{
						Datetime: "2025-08-21",
						TR:       null.FloatFrom(5.6789),
					},
				},
				Status: "ok",
			},
			want1:   response.NewCreditsImpl(100, 10),
			wantErr: "",
		},
	}

	runTechnicalIndicatorTest(
		t,
		tests,
		"/trange",
		"GetTR",
		func(httpCli *HTTPCli, url string) interface{} {
			return client{
				getTR: NewEndpoint[request.GetTR, response.TR, response.Credits, error](httpCli, url),
			}
		},
		func(cli interface{}, req request.GetTR) (response.TR, response.Credits, error) {
			return cli.(client).GetTR(req)
		},
	)
}
