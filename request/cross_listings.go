package request

type GetCrossListings struct {
	ApiKey
	Symbol   string `schema:"symbol"`
	Exchange string `schema:"exchange,omitempty"`
	MicCode  string `schema:"mic_code,omitempty"`
	Country  string `schema:"country,omitempty"`
}
