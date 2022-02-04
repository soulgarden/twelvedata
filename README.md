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

### HTTP

    import (
        "fmt"
        "github.com/jinzhu/configor"
        "github.com/rs/zerolog"
        "github.com/soulgarden/twelvedata"
        "github.com/valyala/fasthttp"
        "os"
    )

    logger := zerolog.New(os.Stdout)
    
    cfg := &twelvedata.Conf{
        APIKey: "4e0133f255164c499a387977ce017ebc",
    }
    
    if err := configor.New(&configor.Config{}).Load(cfg); err != nil {
        logger.Err(err).Msg("init config")
        
        return
    }
    
    cli := twelvedata.NewCli(
        cfg,
        &fasthttp.Client{},
        &logger,
    )
    
    resp, creditsLeft, creditsUsed, err := cli.GetEtfs("")
    
    fmt.Println(resp, creditsLeft, creditsUsed, err)

### WebSocket

    import (
        "context"
        "fmt"
        "github.com/jinzhu/configor"
        "github.com/rs/zerolog"
        "github.com/soulgarden/twelvedata"
        "os"
    )

    logger := zerolog.New(os.Stdout)

    ctx := context.Background()
    
    cfg := &twelvedata.Conf{
        APIKey: "4e0133f255164c499a387977ce017ebc",
    }
    
    if err := configor.New(&configor.Config{}).Load(cfg); err != nil {
        logger.Err(err).Msg("init config")
    
        return
    }
    
    wsCli := twelvedata.NewWS(
        cfg,
        &logger,
        nil,
    )
    
    go wsCli.Subscribe(ctx, []string{"AAPL", "FB"})
    
    for e := range wsCli.Consume() {
        fmt.Println(e)
    }