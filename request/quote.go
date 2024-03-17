package request

type GetQuote struct {
	symbol, interval, exchange, micCode, country, volumeTimePeriod, instrumentType, prepost string
	eod                                                                                     bool
	rollingPeriod, decimalPlaces                                                            int
	timezone                                                                                string
}
