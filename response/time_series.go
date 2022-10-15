package response

type TimeSeries struct {
	Meta   *TimeSeriesMeta    `json:"meta"`
	Values []*TimeSeriesValue `json:"values"`
	Status string             `json:"status"`
}

type TimeSeriesMeta struct {
	Symbol           string `json:"symbol"`
	Interval         string `json:"interval"`
	Currency         string `json:"currency"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	Type             string `json:"type"`
}

type TimeSeriesValue struct {
	Datetime string `json:"datetime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}
