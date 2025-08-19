package response

// MutualFundTypes represents the mutual fund types response structure.
type MutualFundTypes struct {
	Result map[string][]string `json:"result"`
	Status string              `json:"status"`
}
