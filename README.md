# Twelve data go api client

[![Go Report Card](https://goreportcard.com/badge/github.com/soulgarden/twelvedata)](https://goreportcard.com/report/github.com/soulgarden/twelvedata)
![Tests and linters](https://github.com/soulgarden/twelvedata/actions/workflows/main.yml/badge.svg)


# Covered:

## Reference data

* Stocks list                    ✅
* Forex Pairs List               ❌
* Cryptocurrencies List          ❌
* Etfs                           ✅
* Indices                        ✅
* Exchanges                      ✅
* Cryptocurrency exchanges       ❌
* Technical indicators interface ❌
* Symbol search                  ❌
* Earliest timestamp             ❌
* Market state                   ✅


## Core data

* Time series          ✅
* Exchange rate        ✅
* Currency conversion  ❌
* Quote                ✅
* Real-time price      ❌
* Edd of day price     ❌
* Market movers        ✅

## Fundamentals

* Logo                  ❌
* Profile               ✅
* Dividends             ✅
* Splits                ❌
* Earnings              ❌
* Earnings calendar     ✅
* IPO calendar          ❌
* Statistics            ✅
* Insider transactions  ✅
* Income statement      ✅
* Balance sheet         ✅
* Cash flow             ✅
* Options expiration    ❌
* Options chain         ❌
* Key executives        ❌
* Institutional holders ❌
* Fund holders          ❌

## WebSocket

* Price ✅

## Advanced

* Complex Data ❌
* Usage        ✅


# Usage

[http example](https://github.com/soulgarden/twelvedata/blob/main/examples/etfs.go)

[ws example](https://github.com/soulgarden/twelvedata/blob/main/examples/ws.go)
