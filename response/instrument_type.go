package response

// InstrumentType represents the response structure for instrument type data.
type InstrumentType struct {
	Result []string `json:"result"`
	Status string   `json:"status"`
}
