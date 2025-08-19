package response

// TimeSeriesCross represents cross time series data response with metadata and values.
type TimeSeriesCross struct {
	Meta   TimeSeriesMeta    `json:"meta"`
	Values []TimeSeriesValue `json:"values"`
	Status string            `json:"status"`
}
