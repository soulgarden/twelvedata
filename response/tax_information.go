package response

import "github.com/guregu/null/v6"

// TaxInformation represents the response structure for tax information data.
type TaxInformation struct {
	Meta   TaxInformationMeta `json:"meta"`
	Data   TaxInformationData `json:"data"`
	Status string             `json:"status"`
}

// TaxInformationMeta contains metadata for tax information data.
type TaxInformationMeta struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Country  string `json:"country"`
}

// TaxInformationData represents tax-related data for a financial instrument.
type TaxInformationData struct {
	TaxIndicator null.String `json:"tax_indicator"`
}
