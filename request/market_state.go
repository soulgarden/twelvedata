package request

type GetMarketState struct {
	ApiKey
	Exchange string `schema:"exchange,omitempty"`
	Code     string `schema:"code,omitempty"`
	Country  string `schema:"country,omitempty"`
}
