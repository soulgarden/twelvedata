package response

type TimeSeries struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Interval         string `json:"interval"`
		Currency         string `json:"currency"`
		ExchangeTimezone string `json:"exchange_timezone"` // nolint: tagliatelle
		Exchange         string `json:"exchange"`
		Type             string `json:"type"`
	} `json:"meta"`
	Values []*TimeSeriesValue `json:"values"`
	Status string             `json:"status"`
}

type TimeSeriesValue struct {
	Datetime string `json:"datetime"`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
}
