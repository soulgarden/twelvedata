package response

type TimeSeriesCross struct {
	Meta   TimeSeriesMeta    `json:"meta"`
	Values []TimeSeriesValue `json:"values"`
	Status string            `json:"status"`
}
