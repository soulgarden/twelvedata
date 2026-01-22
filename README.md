# Twelve data go api client

[![Go Report Card](https://goreportcard.com/badge/github.com/soulgarden/twelvedata)](https://goreportcard.com/report/github.com/soulgarden/twelvedata)
![Tests and linters](https://github.com/soulgarden/twelvedata/actions/workflows/main.yml/badge.svg)


# Covered:

## Core Data

* Time series          ✅ **High demand**
* Time series cross    ✅
* Quote                ✅ **High demand**
* Latest price         ✅ **High demand**
* End of day price     ✅
* Market movers        ✅ **High demand**

## Reference Data

### Asset Catalogs
* Stocks list           ✅
* Forex Pairs List      ✅
* Cryptocurrencies List ✅
* ETFs                  ✅
* Funds                 ✅
* Commodities           ✅
* Fixed Income          ✅

### Discovery
* Symbol search         ✅ **High demand**
* Cross Listings        ✅
* Earliest timestamp    ✅

### Markets
* Exchanges                ✅ **High demand**
* Exchanges Schedule       ✅
* Cryptocurrency exchanges ✅
* Market state             ✅

### Supporting Metadata
* Countries             ✅
* Instrument Type       ✅
* Technical indicators  ✅

## Fundamentals

* Logo                  ✅
* Profile               ✅ **Useful**
* Dividends             ✅
* Dividends Calendar    ✅
* Splits                ✅
* Splits Calendar       ✅
* Earnings              ✅
* Earnings calendar     ✅
* IPO calendar          ✅
* Statistics            ✅
* Income statement      ✅
* Income statement consolidated ✅ **New**
* Balance sheet         ✅
* Balance sheet consolidated ✅ **New**
* Cash flow             ✅
* Cash flow consolidated ✅ **New**
* Insider transactions  ✅
* Key executives        ✅ **Useful**
* Market Capitalization ✅ **New**
* Last Changes          ✅ **New**

## Currencies

* Exchange rate         ✅
* Currency conversion   ✅ **Useful**

## WebSocket

* Real-time price ✅ **Useful**

## ETFs

* ETFs Directory        ✅ **Useful**
* ETF Full Data         ✅ **High demand**
* ETF Summary           ✅
* ETF Performance       ✅ **High demand**
* ETF Risk              ✅
* ETF Composition       ✅ **High demand**
* ETF Families          ✅
* ETF Types             ✅

## Mutual Funds

* Mutual Funds Directory         ✅ **Useful**
* Mutual Fund Full Data          ✅ **High demand**
* Mutual Fund Summary            ✅
* Mutual Fund Performance        ✅ **High demand**
* Mutual Fund Risk               ✅
* Mutual Fund Ratings            ✅
* Mutual Fund Composition        ✅ **High demand**
* Mutual Fund Purchase Info      ✅
* Mutual Fund Sustainability     ✅
* Mutual Fund Families           ✅
* Mutual Fund Types              ✅

## Technical Indicators

### Overlap Studies
* Bollinger Bands                               ✅ **High demand**
* Double Exponential Moving Average (DEMA)      ✅
* Exponential Moving Average (EMA)              ✅ **High demand**
* Hilbert Transform Instantaneous Trendline     ❌
* Ichimoku Cloud                                ❌
* Kaufman Adaptive Moving Average (KAMA)        ✅
* Keltner Channel                               ❌
* Moving Average                                ✅
* MESA Adaptive Moving Average (MAMA)           ❌
* McGinley Dynamic Indicator                    ❌
* Midpoint                                      ❌
* Midprice                                      ❌
* Pivot Points High Low                         ❌
* Parabolic Stop and Reverse (SAR)              ✅
* Parabolic Stop and Reverse Extended (SAREXT)  ❌
* Simple Moving Average (SMA)                   ✅ **High demand**
* Triple Exponential Moving Average (T3MA)      ❌
* Triple Exponential Moving Average (TEMA)      ✅
* Triangular Moving Average                     ✅
* Volume Weighted Average Price                 ✅
* Weighted Moving Average                       ✅

### Momentum Indicators
* Average Directional Index (ADX)               ✅ **High demand**
* Average Directional Movement Index Rating     ❌
* Absolute Price Oscillator                     ❌
* Aroon Indicator                               ❌
* Aroon Oscillator                              ❌
* Balance of Power                              ❌
* Commodity Channel Index                       ✅
* Chande Momentum Oscillator                    ❌
* Coppock Curve                                 ❌
* Connors Relative Strength Index               ❌
* Detrended Price Oscillator                    ❌
* Directional Movement Index                    ❌
* Know Sure Thing                               ❌
* Moving Average Convergence Divergence (MACD)  ✅ **High demand**
* MACD Slope                                    ❌
* MACD Extension                                ❌
* Money Flow Index                              ❌
* Minus Directional Indicator                   ❌
* Minus Directional Movement                    ❌
* Momentum                                      ✅
* Percent B                                     ✅ **High demand**
* Plus Directional Indicator                    ❌
* Plus Directional Movement                     ❌
* Percentage Price Oscillator                   ❌
* Rate of Change                                ✅
* Rate of Change Percentage                     ❌
* Rate of Change Ratio                          ❌
* Rate of Change Ratio 100                      ❌
* Relative Strength Index (RSI)                 ✅ **High demand**
* Stochastic Oscillator                         ✅ **High demand**
* Stochastic Fast                               ❌
* Stochastic Relative Strength Index            ❌
* Ultimate Oscillator                           ❌
* Williams %R                                   ✅

### Volume Indicators
* Accumulation/Distribution                     ✅
* Accumulation/Distribution Oscillator          ❌
* On Balance Volume                             ✅
* Relative Volume                               ❌

### Volatility Indicators
* Average True Range (ATR)                      ✅
* Normalized Average True Range (NATR)          ✅
* Supertrend                                    ❌
* Supertrend Heikin Ashi candles                ❌
* True Range                                    ✅

### Price Transform
* Addition                                      ❌
* Average                                       ❌
* Average Price                                 ❌
* Ceiling                                       ❌
* Division                                      ❌
* Exponential                                   ❌
* Floor                                         ❌
* Heikinashi Candles                            ❌
* High, Low, Close Average                      ❌
* Natural Logarithm                             ❌
* Base-10 Logarithm                             ❌
* Median Price                                  ❌
* Multiplication                                ❌
* Square Root                                   ❌
* Subtraction                                   ❌
* Summation                                     ❌
* Typical Price                                 ❌
* Weighted Close Price                          ❌

### Cycle Indicators
* Hilbert Transform Dominant Cycle Period       ❌
* Hilbert Transform Dominant Cycle Phase        ❌
* Hilbert Transform Phasor Components           ❌
* Hilbert Transform Sine Wave                   ❌
* Hilbert Transform Trend vs Cycle Mode         ❌

### Statistic Functions
* Beta Indicator                                ❌
* Correlation                                   ❌
* Linear Regression                             ❌
* Linear Regression Angle                       ❌
* Linear Regression Intercept                   ❌
* Linear Regression Slope                       ❌
* Maximum                                       ❌
* Maximum Index                                 ❌
* Minimum                                       ❌
* Minimum Index                                 ❌
* Minimum and Maximum                           ❌
* Minimum and Maximum Index                     ❌
* Standard Deviation                            ❌
* Time Series Forecast                          ❌
* Variance                                      ❌

## Analysis

* Earnings estimate           ✅ **Useful**
* Revenue estimate            ✅
* EPS trend                   ✅
* EPS revisions               ✅
* Growth estimates            ✅
* Recommendations             ✅ **High demand**
* Price target                ✅ **High demand**
* Analyst ratings snapshot    ✅
* Analyst ratings US equities ✅

## Regulatory

* EDGAR fillings        ✅
* Insider transaction   ✅
* Institutional holders ✅
* Fund holders          ✅
* Direct holders        ✅
* Tax information       ✅
* Sanctioned entities   ✅

## Advanced

* Batches      ✅ **Useful**
* Usage        ✅


# Usage

[http example](https://github.com/soulgarden/twelvedata/blob/main/examples/etfs.go)

[ws example](https://github.com/soulgarden/twelvedata/blob/main/examples/ws.go)
