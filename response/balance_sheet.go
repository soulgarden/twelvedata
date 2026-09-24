package response

import "github.com/guregu/null/v6"

// BalanceSheets represents the response structure for balance sheet data.
type BalanceSheets struct {
	Meta         BalanceSheetsMeta `json:"meta"`
	BalanceSheet []BalanceSheet    `json:"balance_sheet"`
}

// BalanceSheetsMeta contains metadata for balance sheet data.
type BalanceSheetsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Period           string `json:"period"`
}

// BalanceSheet represents financial balance sheet data for a specific fiscal period.
type BalanceSheet struct {
	FiscalDate         string                         `json:"fiscal_date"`
	Year               null.Int                       `json:"year"`
	Assets             BalanceSheetAssets             `json:"assets"`
	Liabilities        BalanceSheetLiabilities        `json:"liabilities"`
	ShareholdersEquity BalanceSheetShareholdersEquity `json:"shareholders_equity"`
}

// BalanceSheetAssets represents the assets section of a balance sheet.
type BalanceSheetAssets struct {
	CurrentAssets    BalanceSheetCurrentAssets    `json:"current_assets"`
	NonCurrentAssets BalanceSheetNonCurrentAssets `json:"non_current_assets"`
	TotalAssets      null.Float                   `json:"total_assets"`
}

// BalanceSheetCurrentAssets represents current assets in a balance sheet.
type BalanceSheetCurrentAssets struct {
	Cash                      null.Float `json:"cash"`
	CashEquivalents           null.Float `json:"cash_equivalents"`
	CashAndCashEquivalents    null.Float `json:"cash_and_cash_equivalents"`
	OtherShortTermInvestments null.Float `json:"other_short_term_investments"`
	AccountsReceivable        null.Float `json:"accounts_receivable"`
	OtherReceivables          null.Float `json:"other_receivables"`
	Inventory                 null.Float `json:"inventory"`
	PrepaidAssets             null.Float `json:"prepaid_assets"`
	RestrictedCash            null.Float `json:"restricted_cash"`
	AssetsHeldForSale         null.Float `json:"assets_held_for_sale"`
	HedgingAssets             null.Float `json:"hedging_assets"`
	OtherCurrentAssets        null.Float `json:"other_current_assets"`
	TotalCurrentAssets        null.Float `json:"total_current_assets"`
}

// BalanceSheetNonCurrentAssets represents non-current assets in a balance sheet.
type BalanceSheetNonCurrentAssets struct {
	Properties                  null.Float `json:"properties"`
	LandAndImprovements         null.Float `json:"land_and_improvements"`
	MachineryFurnitureEquipment null.Float `json:"machinery_furniture_equipment"`
	ConstructionInProgress      null.Float `json:"construction_in_progress"`
	Leases                      null.Float `json:"leases"`
	AccumulatedDepreciation     null.Float `json:"accumulated_depreciation"`
	Goodwill                    null.Float `json:"goodwill"`
	InvestmentProperties        null.Float `json:"investment_properties"`
	FinancialAssets             null.Float `json:"financial_assets"`
	IntangibleAssets            null.Float `json:"intangible_assets"`
	InvestmentsAndAdvances      null.Float `json:"investments_and_advances"`
	OtherNonCurrentAssets       null.Float `json:"other_non_current_assets"`
	TotalNonCurrentAssets       null.Float `json:"total_non_current_assets"`
}

// BalanceSheetLiabilities represents the liabilities section of a balance sheet.
type BalanceSheetLiabilities struct {
	CurrentLiabilities    BalanceSheetCurrentLiabilities    `json:"current_liabilities"`
	NonCurrentLiabilities BalanceSheetNonCurrentLiabilities `json:"non_current_liabilities"`
	TotalLiabilities      null.Float                        `json:"total_liabilities"`
}

// BalanceSheetCurrentLiabilities represents current liabilities in a balance sheet.
type BalanceSheetCurrentLiabilities struct {
	AccountsPayable         null.Float `json:"accounts_payable"`
	AccruedExpenses         null.Float `json:"accrued_expenses"`
	ShortTermDebt           null.Float `json:"short_term_debt"`
	DeferredRevenue         null.Float `json:"deferred_revenue"`
	TaxPayable              null.Float `json:"tax_payable"`
	Pensions                null.Float `json:"pensions"`
	OtherCurrentLiabilities null.Float `json:"other_current_liabilities"`
	TotalCurrentLiabilities null.Float `json:"total_current_liabilities"`
}

// BalanceSheetNonCurrentLiabilities represents non-current liabilities in a balance sheet.
type BalanceSheetNonCurrentLiabilities struct {
	LongTermProvisions           null.Float `json:"long_term_provisions"`
	LongTermDebt                 null.Float `json:"long_term_debt"`
	ProvisionForRisksAndCharges  null.Float `json:"provision_for_risks_and_charges"`
	DeferredLiabilities          null.Float `json:"deferred_liabilities"`
	DerivativeProductLiabilities null.Float `json:"derivative_product_liabilities"`
	OtherNonCurrentLiabilities   null.Float `json:"other_non_current_liabilities"`
	TotalNonCurrentLiabilities   null.Float `json:"total_non_current_liabilities"`
}

// BalanceSheetShareholdersEquity represents shareholders' equity section of a balance sheet.
type BalanceSheetShareholdersEquity struct {
	CommonStock             null.Float `json:"common_stock"`
	RetainedEarnings        null.Float `json:"retained_earnings"`
	OtherShareholdersEquity null.Float `json:"other_shareholders_equity"`
	TotalShareholdersEquity null.Float `json:"total_shareholders_equity"`
	AdditionalPaidInCapital null.Float `json:"additional_paid_in_capital"`
	TreasuryStock           null.Float `json:"treasury_stock"`
	MinorityInterest        null.Float `json:"minority_interest"`
}
