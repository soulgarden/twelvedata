Twelve data go api client

[![Go Report Card](https://goreportcard.com/badge/github.com/soulgarden/twelvedata)](https://goreportcard.com/report/github.com/soulgarden/twelvedata)


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


## Core data

* Time series          ✅
* Exchange rate        ✅
* Currency conversion  ❌
* Quote                ✅
* Real-time price      ❌
* Edd of day price     ❌

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

    import (
        "github.com/rs/zerolog"
        "github.com/soulgarden/twelvedata"
        "github.com/valyala/fasthttp"
        "os"
    )

    logger := zerolog.New(os.Stdout)
    
    cli := twelvedata.NewCli(
        &twelvedata.Conf{APIKey: "demo"},
        &fasthttp.Client{},
        &logger,
    )
    
    resp, creditsLeft, creditsUsed := cli.GetEtfs("")