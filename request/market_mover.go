package request

type GetMarketMovers struct {
	instrument, direction string
	outputSize            int
	country               string
	decimalPlaces         int
}
