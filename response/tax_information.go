package response

import "github.com/guregu/null/v6"

// TaxInformation represents the response structure for tax information data.
type TaxInformation struct {
	Meta TaxInformationMeta `json:"meta"`
	Data TaxInformationData `json:"data"`
}

// TaxInformationMeta contains metadata for tax information data.
type TaxInformationMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// TaxInformationData represents tax-related data for a financial instrument.
type TaxInformationData struct {
	TaxIndicator        null.String `json:"tax_indicator"`
	TaxRate             null.Float  `json:"tax_rate"`
	TaxCode             null.String `json:"tax_code"`
	WithholdingTaxRate  null.Float  `json:"withholding_tax_rate"`
	CapitalGainsTaxRate null.Float  `json:"capital_gains_tax_rate"`
	DividendTaxRate     null.Float  `json:"dividend_tax_rate"`
	TaxCountry          null.String `json:"tax_country"`
	TaxRegion           null.String `json:"tax_region"`
}
