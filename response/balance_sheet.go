package response

import "gopkg.in/guregu/null.v4"

type BalanceSheets struct {
	Meta         BalanceSheetsMeta `json:"meta"`
	BalanceSheet []BalanceSheet    `json:"balance_sheet"`
}

type BalanceSheetsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Period           string `json:"period"`
}

type BalanceSheet struct {
	FiscalDate         string                         `json:"fiscal_date"`
	Assets             BalanceSheetAssets             `json:"assets"`
	Liabilities        BalanceSheetLiabilities        `json:"liabilities"`
	ShareholdersEquity BalanceSheetShareholdersEquity `json:"shareholders_equity"`
}

type BalanceSheetAssets struct {
	CurrentAssets    BalanceSheetCurrentAssets    `json:"current_assets"`
	NonCurrentAssets BalanceSheetNonCurrentAssets `json:"non_current_assets"`
	TotalAssets      null.Int                     `json:"total_assets"`
}

type BalanceSheetCurrentAssets struct {
	Cash                      null.Int `json:"cash"`
	CashEquivalents           null.Int `json:"cash_equivalents"`
	CashAndCashEquivalents    null.Int `json:"cash_and_cash_equivalents"`
	OtherShortTermInvestments null.Int `json:"other_short_term_investments"`
	AccountsReceivable        null.Int `json:"accounts_receivable"`
	OtherReceivables          null.Int `json:"other_receivables"`
	Inventory                 null.Int `json:"inventory"`
	PrepaidAssets             null.Int `json:"prepaid_assets"`
	RestrictedCash            null.Int `json:"restricted_cash"`
	AssetsHeldForSale         null.Int `json:"assets_held_for_sale"`
	HedgingAssets             null.Int `json:"hedging_assets"`
	OtherCurrentAssets        null.Int `json:"other_current_assets"`
	TotalCurrentAssets        null.Int `json:"total_current_assets"`
}

type BalanceSheetNonCurrentAssets struct {
	Properties                  null.Int `json:"properties"`
	LandAndImprovements         null.Int `json:"land_and_improvements"`
	MachineryFurnitureEquipment null.Int `json:"machinery_furniture_equipment"`
	ConstructionInProgress      null.Int `json:"construction_in_progress"`
	Leases                      null.Int `json:"leases"`
	AccumulatedDepreciation     null.Int `json:"accumulated_depreciation"`
	Goodwill                    null.Int `json:"goodwill"`
	InvestmentProperties        null.Int `json:"investment_properties"`
	FinancialAssets             null.Int `json:"financial_assets"`
	IntangibleAssets            null.Int `json:"intangible_assets"`
	InvestmentsAndAdvances      null.Int `json:"investments_and_advances"`
	OtherNonCurrentAssets       null.Int `json:"other_non_current_assets"`
	TotalNonCurrentAssets       null.Int `json:"total_non_current_assets"`
}

type BalanceSheetLiabilities struct {
	CurrentLiabilities    BalanceSheetCurrentLiabilities    `json:"current_liabilities"`
	NonCurrentLiabilities BalanceSheetNonCurrentLiabilities `json:"non_current_liabilities"`
	TotalLiabilities      null.Int                          `json:"total_liabilities"`
}

type BalanceSheetCurrentLiabilities struct {
	AccountsPayable         null.Int `json:"accounts_payable"`
	AccruedExpenses         null.Int `json:"accrued_expenses"`
	ShortTermDebt           null.Int `json:"short_term_debt"`
	DeferredRevenue         null.Int `json:"deferred_revenue"`
	TaxPayable              null.Int `json:"tax_payable"`
	Pensions                null.Int `json:"pensions"`
	OtherCurrentLiabilities null.Int `json:"other_current_liabilities"`
	TotalCurrentLiabilities null.Int `json:"total_current_liabilities"`
}

type BalanceSheetNonCurrentLiabilities struct {
	LongTermProvisions           null.Int `json:"long_term_provisions"`
	LongTermDebt                 null.Int `json:"long_term_debt"`
	ProvisionForRisksAndCharges  null.Int `json:"provision_for_risks_and_charges"`
	DeferredLiabilities          null.Int `json:"deferred_liabilities"`
	DerivativeProductLiabilities null.Int `json:"derivative_product_liabilities"`
	OtherNonCurrentLiabilities   null.Int `json:"other_non_current_liabilities"`
	TotalNonCurrentLiabilities   null.Int `json:"total_non_current_liabilities"`
}

type BalanceSheetShareholdersEquity struct {
	CommonStock             null.Int `json:"common_stock"`
	RetainedEarnings        null.Int `json:"retained_earnings"`
	OtherShareholdersEquity null.Int `json:"other_shareholders_equity"`
	TotalShareholdersEquity null.Int `json:"total_shareholders_equity"`
	AdditionalPaidInCapital null.Int `json:"additional_paid_in_capital"`
	TreasuryStock           null.Int `json:"treasury_stock"`
	MinorityInterest        null.Int `json:"minority_interest"`
}
