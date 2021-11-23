package response

type Quotes struct {
	Data   []*Quote
	Errors []*QuoteError
}

// nolint: tagliatelle
type Quote struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Exchange      string `json:"exchange"`
	Currency      string `json:"currency"`
	Datetime      string `json:"datetime"`
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Close         string `json:"close"`
	Volume        string `json:"volume"`
	PreviousClose string `json:"previous_close"`
	Change        string `json:"change"`
	PercentChange string `json:"percent_change"`
	AverageVolume string `json:"average_volume"`

	FiftyTwoWeek struct {
		Low               string `json:"low"`
		High              string `json:"high"`
		LowChange         string `json:"low_change"`
		HighChange        string `json:"high_change"`
		LowChangePercent  string `json:"low_change_percent"`
		HighChangePercent string `json:"high_change_percent"`
		Range             string `json:"range"`
	} `json:"fifty_two_week"`
}

type QuoteError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Meta    struct {
		Symbol   string `json:"symbol"`
		Interval string `json:"interval"`
		Exchange string `json:"exchange"`
	} `json:"meta"`
}
