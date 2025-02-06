package request

type GetMarketMovers struct {
	ApiKey
	Instrument    string `schema:"instrument,omitempty"`
	Direction     string `schema:"direction,omitempty"`
	OutputSize    int    `schema:"output_size,omitempty"`
	Country       string `schema:"country,omitempty"`
	DecimalPlaces int    `schema:"dp,omitempty"`
}
