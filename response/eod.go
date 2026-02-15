package response

// EOD represents end-of-day price data for a financial instrument.
type EOD struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Currency string `json:"currency"`
	Datetime string `json:"datetime"`
	Close    string `json:"close"`
}
