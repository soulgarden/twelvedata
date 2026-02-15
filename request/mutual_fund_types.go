package request

// GetMutualFundTypes represents request parameters for mutual fund types endpoint.
type GetMutualFundTypes struct {
	APIKey
	Country  string `schema:"country,omitempty"`
	FundType string `schema:"fund_type,omitempty"`
}
