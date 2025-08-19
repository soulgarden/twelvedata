package response

// ETFFamilies represents the ETF families response structure.
type ETFFamilies struct {
	Result map[string][]string `json:"result"`
	Status string              `json:"status"`
}
