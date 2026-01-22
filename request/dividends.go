package request

// GetDividends represents request parameters for dividends data.
type GetDividends struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	Figi      string `schema:"figi,omitempty"`
	Isin      string `schema:"isin,omitempty"`
	Cusip     string `schema:"cusip,omitempty"`
	Exchange  string `schema:"exchange,omitempty"`
	MicCode   string `schema:"mic_code,omitempty"`
	Country   string `schema:"country,omitempty"`
	Range     string `schema:"range,omitempty"`
	StartDate string `schema:"start_date,omitempty"`
	EndDate   string `schema:"end_date,omitempty"`
	Adjust    bool   `schema:"adjust,omitempty"`
}
