package response

import "github.com/guregu/null/v6"

// Quotes represents a collection of quote data and any associated errors.
type Quotes struct {
	Data   []Quote
	Errors []QuoteError
}

// Quote represents detailed market quote information for a financial instrument.
type Quote struct {
	Symbol                string             `json:"symbol"`
	Name                  string             `json:"name"`
	Exchange              string             `json:"exchange"`
	MicCode               string             `json:"mic_code"`
	Currency              string             `json:"currency"`
	Datetime              string             `json:"datetime"`
	Timestamp             int                `json:"timestamp"`
	LastQuoteAt           int                `json:"last_quote_at"`
	Open                  string             `json:"open"`
	High                  string             `json:"high"`
	Low                   string             `json:"low"`
	Close                 string             `json:"close"`
	Volume                string             `json:"volume"`
	PreviousClose         string             `json:"previous_close"`
	Change                string             `json:"change"`
	PercentChange         string             `json:"percent_change"`
	AverageVolume         string             `json:"average_volume"`
	Rolling1DChange       string             `json:"rolling_1d_change"`
	Rolling7DChange       string             `json:"rolling_7d_change"`
	RollingPeriodChange   string             `json:"rolling_period_change"`
	IsMarketOpen          bool               `json:"is_market_open"`
	FiftyTwoWeek          *QuoteFiftyTwoWeek `json:"fifty_two_week"`
	ExtendedChange        string             `json:"extended_change"`
	ExtendedPercentChange string             `json:"extended_percent_change"`
	ExtendedPrice         string             `json:"extended_price"`
	ExtendedTimestamp     null.Int           `json:"extended_timestamp"`
}

// QuoteFiftyTwoWeek represents 52-week high/low data for a quote.
type QuoteFiftyTwoWeek struct {
	Low               string `json:"low"`
	High              string `json:"high"`
	LowChange         string `json:"low_change"`
	HighChange        string `json:"high_change"`
	LowChangePercent  string `json:"low_change_percent"`
	HighChangePercent string `json:"high_change_percent"`
	Range             string `json:"range"`
}

// QuoteError represents an error that occurred while fetching quote data.
type QuoteError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Status  string          `json:"status"`
	Meta    *QuoteErrorMeta `json:"meta"`
}

// QuoteErrorMeta contains metadata about the symbol that caused a quote error.
type QuoteErrorMeta struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Exchange string `json:"exchange"`
}
