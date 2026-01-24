package response

import "github.com/guregu/null/v6"

// ETFsDirectory represents the response structure for ETFs directory data.
type ETFsDirectory struct {
	Result ETFsDirectoryResult `json:"result"`
	Status string              `json:"status"`
}

// ETFsDirectoryResult contains the ETFs directory list and count.
type ETFsDirectoryResult struct {
	Count null.Int           `json:"count"`
	List  []ETFsDirectoryETF `json:"list"`
}

// ETFsDirectoryETF represents a single ETF in the directory list.
type ETFsDirectoryETF struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	MicCode    string `json:"mic_code"`
	FundFamily string `json:"fund_family"`
	FundType   string `json:"fund_type"`
}
