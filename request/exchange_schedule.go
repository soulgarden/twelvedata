package request

type GetExchangeSchedule struct {
	ApiKey
	Date    string `schema:"date,omitempty"`
	MicName string `schema:"mic_name,omitempty"`
	MicCode string `schema:"mic_code,omitempty"`
	Country string `schema:"country,omitempty"`
}
