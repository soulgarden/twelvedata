package response

import "github.com/guregu/null/v6"

// FundHolders represents the response structure for fund holders data.
type FundHolders struct {
	Meta        FundHoldersMeta `json:"meta"`
	FundHolders []FundHolder    `json:"fund_holders"`
}

// FundHoldersMeta contains metadata for fund holders data.
type FundHoldersMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// FundHolder represents a single mutual fund holder record.
type FundHolder struct {
	FundName     string     `json:"fund_name"`
	DateReported string     `json:"date_reported"`
	Shares       int64      `json:"shares"`
	Value        int64      `json:"value"`
	PercentHeld  null.Float `json:"percent_held"`
}
