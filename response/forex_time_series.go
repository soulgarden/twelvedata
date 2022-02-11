package response

type ForexTimeSeries struct {
	Meta struct {
		Symbol        string `json:"symbol"`
		Interval      string `json:"interval"`
		CurrencyBase  string `json:"currency_base"`
		CurrencyQuote string `json:"currency_quote"`
		Type          string `json:"type"`
	} `json:"meta"`
	Values []struct {
		Datetime string `json:"datetime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
	} `json:"values"`
	Status string `json:"status"`
}
