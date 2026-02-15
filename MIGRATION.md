# Migration guide (v0.1.x → v1.0.0)

This release promotes the new client architecture into the mainline.
It includes breaking API changes compared to `v0.1.x` tags.

## Key changes

- `twelvedata.Cli` is removed. Use `twelvedata.NewClient(...)` (typed client).
- Endpoints no longer take long positional argument lists; they take `request.*` structs.
- Credits are returned as `response.Credits` (instead of `(creditsLeft, creditsUsed int64)`).
- `Conf` endpoint URLs are now paths (for example `"/time_series"`), not URL templates with `{placeholder}` query strings.
- Error handling now provides classification helpers (for example `twelvedata.IsUnauthorizedError(err)`).

## Before/after sketch

### Old (v0.1.x)

```go
cfg := &twelvedata.Conf{APIKey: "demo"}
httpCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, cfg, logger)
cli := twelvedata.NewCli(cfg, httpCli, logger)

resp, creditsLeft, creditsUsed, err := cli.GetTimeSeries("AAPL", "1day", "", "", "", "", 30, "", 2, "", "", "", "", "", false)
_ = resp
_ = creditsLeft
_ = creditsUsed
_ = err
```

### New (v1.0.0)

```go
cfg := &twelvedata.Conf{APIKey: "demo"}
httpCli := twelvedata.NewHTTPCli(&fasthttp.Client{}, cfg, logger)
cli := twelvedata.NewClient(httpCli, cfg)

resp, credits, err := cli.GetTimeSeries(request.GetTimeSeries{
	APIKey: request.APIKey{APIKey: cfg.APIKey},
	Symbol: "AAPL",
	Interval: "1day",
})
_ = resp
_ = credits
_ = err
```

## Examples

- HTTP: `examples/etf_example/etfs.go`
- WebSocket: `examples/ws_example/ws.go`
- Error handling: `examples/error_handling/error_handling.go`
