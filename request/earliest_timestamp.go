package request

// GetEarliestTimestamp represents request parameters for earliest timestamp data.
type GetEarliestTimestamp struct {
	APIKey
	Symbol   string `schema:"symbol,omitempty"`
	Figi     string `schema:"figi,omitempty"`
	Isin     string `schema:"isin,omitempty"`
	Cusip    string `schema:"cusip,omitempty"`
	Interval string `schema:"interval"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Timezone string `schema:"timezone,omitempty"`
}
