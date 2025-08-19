package response

// MutualFunds represents the response structure for the mutual funds directory endpoint.
type MutualFunds struct {
	Result MutualFundsResult `json:"result"`
	Status string            `json:"status"`
}

// MutualFundsResult contains the main result data for mutual funds.
type MutualFundsResult struct {
	Count int               `json:"count"`
	List  []MutualFundsData `json:"list"`
}

// MutualFundsData represents individual mutual fund information.
type MutualFundsData struct {
	Symbol   string            `json:"symbol"`
	Name     string            `json:"name"`
	Country  string            `json:"country"`
	Currency string            `json:"currency"`
	Exchange string            `json:"exchange"`
	MicCode  string            `json:"mic_code"`
	Type     string            `json:"type"`
	FigiCode string            `json:"figi_code"`
	CfiCode  string            `json:"cfi_code"`
	Isin     string            `json:"isin"`
	Cusip    string            `json:"cusip"`
	Access   MutualFundsAccess `json:"access"`
}

// MutualFundsAccess represents access information for mutual funds.
type MutualFundsAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
