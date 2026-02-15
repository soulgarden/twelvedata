package response

// ETFTypes represents the ETF types response structure.
type ETFTypes struct {
	Result map[string][]string `json:"result"`
	Status string              `json:"status"`
}
