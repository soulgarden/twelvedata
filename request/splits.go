package request

// GetSplits represents request parameters for splits data.
type GetSplits struct {
	APIKey
	Symbol    string `schema:"symbol,omitempty"`
	FIGI      string `schema:"figi,omitempty"`
	ISIN      string `schema:"isin,omitempty"`
	CUSIP     string `schema:"cusip,omitempty"`
	Exchange  string `schema:"exchange,omitempty"`
	MicCode   string `schema:"mic_code,omitempty"`
	Country   string `schema:"country,omitempty"`
	Range     string `schema:"range,omitempty"`
	StartDate string `schema:"start_date,omitempty"`
	EndDate   string `schema:"end_date,omitempty"`
}
