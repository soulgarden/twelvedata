package request

// GetETFTypes represents request parameters for ETF types endpoint.
type GetETFTypes struct {
	APIKey
	Country  string `schema:"country,omitempty"`
	FundType string `schema:"fund_type,omitempty"`
}
