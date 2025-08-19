//nolint:dupl // Technical indicator tests share similar structure by design
package twelvedata

import (
	"net/http"
	"testing"

	"github.com/soulgarden/twelvedata/request"
	"github.com/soulgarden/twelvedata/response"
)

func Test_client_GetBBands(t *testing.T) {
	tests := []TechnicalIndicatorTestCase[request.GetBBands, response.BBands]{
		{
			name: "success with demo API response",
			args: TechnicalIndicatorTestArgs[request.GetBBands]{
				req: request.GetBBands{
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/bbands?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/sma?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/ema?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/macd?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/rsi?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/atr?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/cci?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/dema?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/kama?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/ma?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/sar?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/tema?interval=1day&symbol=AAPL",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/trma?interval=1day&symbol=AAPL",
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
		"/trma",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					      "name": "VWAP - Volume Weighted Average Price"
					    }
					  },
					  "values": [
					    {
					      "datetime": "2025-08-21",
					      "vwap": "220.12345"
					    }
					  ],
					  "status": "ok"
					}`,
					"/vwap?interval=1day&symbol=AAPL",
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
						Name: "VWAP - Volume Weighted Average Price",
					},
				},
				Values: []response.VWAPValue{
					{
						Datetime: "2025-08-21",
						VWAP:     "220.12345",
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
					APIKey:   request.APIKey{APIKey: ""},
					Symbol:   "AAPL",
					Interval: "1day",
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
					"/wma?interval=1day&symbol=AAPL",
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
