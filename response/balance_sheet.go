package response

type BalanceSheets struct {
	Meta         *BalanceSheetsMeta `json:"meta"`
	BalanceSheet []*BalanceSheet    `json:"balance_sheet"`
}

type BalanceSheetsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Period           string `json:"period"`
}

type BalanceSheet struct {
	FiscalDate         string                          `json:"fiscal_date"`
	Assets             *BalanceSheetAssets             `json:"assets"`
	Liabilities        *BalanceSheetLiabilities        `json:"liabilities"`
	ShareholdersEquity *BalanceSheetShareholdersEquity `json:"shareholders_equity"`
}

type BalanceSheetAssets struct {
	CurrentAssets    *BalanceSheetCurrentAssets    `json:"current_assets"`
	NonCurrentAssets *BalanceSheetNonCurrentAssets `json:"non_current_assets"`
	TotalAssets      int64                         `json:"total_assets"`
}

type BalanceSheetCurrentAssets struct {
	Cash                      int64 `json:"cash"`
	CashEquivalents           int64 `json:"cash_equivalents"`
	CashAndCashEquivalents    int64 `json:"cash_and_cash_equivalents"`
	OtherShortTermInvestments int64 `json:"other_short_term_investments"`
	AccountsReceivable        int64 `json:"accounts_receivable"`
	OtherReceivables          int64 `json:"other_receivables"`
	Inventory                 int64 `json:"inventory"`
	PrepaidAssets             int64 `json:"prepaid_assets"`
	OtherCurrentAssets        int64 `json:"other_current_assets"`
	TotalCurrentAssets        int64 `json:"total_current_assets"`
}

type BalanceSheetNonCurrentAssets struct {
	Properties                  int64 `json:"properties"`
	LandAndImprovements         int64 `json:"land_and_improvements"`
	MachineryFurnitureEquipment int64 `json:"machinery_furniture_equipment"`
	Leases                      int64 `json:"leases"`
	AccumulatedDepreciation     int64 `json:"accumulated_depreciation"`
	Goodwill                    int64 `json:"goodwill"`
	IntangibleAssets            int64 `json:"intangible_assets"`
	InvestmentsAndAdvances      int64 `json:"investments_and_advances"`
	OtherNonCurrentAssets       int64 `json:"other_non_current_assets"`
	TotalNonCurrentAssets       int64 `json:"total_non_current_assets"`
}

type BalanceSheetLiabilities struct {
	CurrentLiabilities    *BalanceSheetCurrentLiabilities    `json:"current_liabilities"`
	NonCurrentLiabilities *BalanceSheetNonCurrentLiabilities `json:"non_current_liabilities"`
	TotalLiabilities      int64                              `json:"total_liabilities"`
}

type BalanceSheetCurrentLiabilities struct {
	AccountsPayable         int64 `json:"accounts_payable"`
	AccruedExpenses         int64 `json:"accrued_expenses"`
	ShortTermDebt           int64 `json:"short_term_debt"`
	DeferredRevenue         int64 `json:"deferred_revenue"`
	OtherCurrentLiabilities int64 `json:"other_current_liabilities"`
	TotalCurrentLiabilities int64 `json:"total_current_liabilities"`
	TaxPayable              int64 `json:"tax_payable"`
}

type BalanceSheetNonCurrentLiabilities struct {
	LongTermDebt                int64 `json:"long_term_debt"`
	ProvisionForRisksAndCharges int64 `json:"provision_for_risks_and_charges"`
	DeferredLiabilities         int64 `json:"deferred_liabilities"`
	OtherNonCurrentLiabilities  int64 `json:"other_non_current_liabilities"`
	TotalNonCurrentLiabilities  int64 `json:"total_non_current_liabilities"`
	LongTermProvisions          int64 `json:"long_term_provisions"`
}

type BalanceSheetShareholdersEquity struct {
	CommonStock             int64 `json:"common_stock"`
	RetainedEarnings        int64 `json:"retained_earnings"`
	OtherShareholdersEquity int64 `json:"other_shareholders_equity"`
	TotalShareholdersEquity int64 `json:"total_shareholders_equity"`
	AdditionalPaidInCapital int64 `json:"additional_paid_in_capital"`
	TreasuryStock           int64 `json:"treasury_stock"`
	MinorityInterest        int64 `json:"minority_interest"`
}
