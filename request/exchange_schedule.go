package request

// GetExchangeSchedule represents request parameters for exchange schedule data.
type GetExchangeSchedule struct {
	APIKey
	Date    string `schema:"date,omitempty"`
	MicName string `schema:"mic_name,omitempty"`
	MicCode string `schema:"mic_code,omitempty"`
	Country string `schema:"country,omitempty"`
}
