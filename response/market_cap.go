package response

import "github.com/guregu/null/v6"

// MarketCap represents the response structure for market capitalization data.
type MarketCap struct {
	Meta      MarketCapMeta   `json:"meta"`
	MarketCap []MarketCapData `json:"market_cap"`
}

// MarketCapMeta contains metadata for market capitalization data.
type MarketCapMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// MarketCapData represents a single market capitalization data point with date and value.
type MarketCapData struct {
	Date  string   `json:"date"`
	Value null.Int `json:"value"`
}
