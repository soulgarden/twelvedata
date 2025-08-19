package response

// MutualFundFamilies represents the mutual fund families response structure.
type MutualFundFamilies struct {
	Result map[string][]string `json:"result"`
	Status string              `json:"status"`
}
