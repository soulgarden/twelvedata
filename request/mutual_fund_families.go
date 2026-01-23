package request

// GetMutualFundFamilies represents request parameters for mutual fund families endpoint.
type GetMutualFundFamilies struct {
	APIKey
	Country    string `schema:"country,omitempty"`
	FundFamily string `schema:"fund_family,omitempty"`
}
