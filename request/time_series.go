package request

type GetTimeSeries struct {
	symbol, interval, exchange, micCode, country, instrumentType string
	outputSize                                                   int
	prePost                                                      string
	decimalPlaces                                                int
	order, timezone, date, startDate, endDate                    string
	previousClose                                                bool
}
