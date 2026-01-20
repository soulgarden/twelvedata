package response

// TimeSeriesCross represents cross time series data response with metadata and values.
type TimeSeriesCross struct {
	Meta   TimeSeriesCrossMeta    `json:"meta"`
	Values []TimeSeriesCrossValue `json:"values"`
}

// TimeSeriesCrossMeta contains metadata for cross time series data.
type TimeSeriesCrossMeta struct {
	BaseInstrument  string `json:"base_instrument"`
	BaseCurrency    string `json:"base_currency"`
	BaseExchange    string `json:"base_exchange"`
	Interval        string `json:"interval"`
	QuoteInstrument string `json:"quote_instrument"`
	QuoteCurrency   string `json:"quote_currency"`
	QuoteExchange   string `json:"quote_exchange"`
}

// TimeSeriesCrossValue represents a single cross time series data point.
type TimeSeriesCrossValue struct {
	Datetime string `json:"datetime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
}
