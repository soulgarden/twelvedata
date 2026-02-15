package request

// GetETFFamilies represents request parameters for ETF families endpoint.
type GetETFFamilies struct {
	APIKey
	Country    string `schema:"country,omitempty"`
	FundFamily string `schema:"fund_family,omitempty"`
}
