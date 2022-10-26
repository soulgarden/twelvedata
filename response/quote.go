package response

import "gopkg.in/guregu/null.v4"

type Quotes struct {
	Data   []Quote
	Errors []QuoteError
}

type Quote struct {
	Symbol                string        `json:"symbol"`
	Name                  string        `json:"name"`
	Exchange              string        `json:"exchange"`
	MicCode               string        `json:"mic_code"`
	Currency              string        `json:"currency"`
	Datetime              string        `json:"datetime"`
	Timestamp             int           `json:"timestamp"`
	Open                  string        `json:"open"`
	High                  string        `json:"high"`
	Low                   string        `json:"low"`
	Close                 string        `json:"close"`
	Volume                string        `json:"volume"`
	PreviousClose         string        `json:"previous_close"`
	Change                string        `json:"change"`
	PercentChange         string        `json:"percent_change"`
	AverageVolume         string        `json:"average_volume"`
	Rolling1DChange       string        `json:"rolling_1d_change"`
	Rolling7DChange       string        `json:"rolling_7d_change"`
	RollingPeriodChange   string        `json:"rolling_period_change"`
	IsMarketOpen          bool          `json:"is_market_open"`
	FiftyTwoWeek          *FiftyTwoWeek `json:"fifty_two_week"`
	ExtendedChange        string        `json:"extended_change"`
	ExtendedPercentChange string        `json:"extended_percent_change"`
	ExtendedPrice         string        `json:"extended_price"`
	ExtendedTimestamp     null.Int      `json:"extended_timestamp"`
}

type FiftyTwoWeek struct {
	Low               string `json:"low"`
	High              string `json:"high"`
	LowChange         string `json:"low_change"`
	HighChange        string `json:"high_change"`
	LowChangePercent  string `json:"low_change_percent"`
	HighChangePercent string `json:"high_change_percent"`
	Range             string `json:"range"`
}

type QuoteError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Status  string          `json:"status"`
	Meta    *QuoteErrorMeta `json:"meta"`
}

type QuoteErrorMeta struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Exchange string `json:"exchange"`
}
