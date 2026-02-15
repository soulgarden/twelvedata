package request

// GetETFsDirectory represents request parameters for ETFs directory data.
type GetETFsDirectory struct {
	APIKey
	Symbol     string `schema:"symbol,omitempty"`
	FIGI       string `schema:"figi,omitempty"`
	ISIN       string `schema:"isin,omitempty"`
	CUSIP      string `schema:"cusip,omitempty"`
	CIK        string `schema:"cik,omitempty"`
	Country    string `schema:"country,omitempty"`
	FundFamily string `schema:"fund_family,omitempty"`
	FundType   string `schema:"fund_type,omitempty"`
	Page       int    `schema:"page,omitempty"`
	OutputSize int    `schema:"outputsize,omitempty"`
}
