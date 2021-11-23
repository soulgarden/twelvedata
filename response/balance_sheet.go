package response

// nolint: tagliatelle
type BalanceSheet struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Period           string `json:"period"`
	} `json:"meta"`
	BalanceSheet []struct {
		FiscalDate string `json:"fiscal_date"`
		Assets     struct {
			CurrentAssets struct {
				Cash                      int64 `json:"cash"`
				CashEquivalents           int64 `json:"cash_equivalents"`
				OtherShortTermInvestments int64 `json:"other_short_term_investments"`
				AccountsReceivable        int64 `json:"accounts_receivable"`
				OtherReceivables          int64 `json:"other_receivables"`
				Inventory                 int64 `json:"inventory"`
				OtherCurrentAssets        int64 `json:"other_current_assets"`
				TotalCurrentAssets        int64 `json:"total_current_assets"`
			} `json:"current_assets"`
			NonCurrentAssets struct {
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
			} `json:"non_current_assets"`
			TotalAssets int64 `json:"total_assets"`
		} `json:"assets"`
		Liabilities struct {
			CurrentLiabilities struct {
				AccountsPayable         int64 `json:"accounts_payable"`
				AccruedExpenses         int64 `json:"accrued_expenses"`
				ShortTermDebt           int64 `json:"short_term_debt"`
				DeferredRevenue         int64 `json:"deferred_revenue"`
				OtherCurrentLiabilities int64 `json:"other_current_liabilities"`
				TotalCurrentLiabilities int64 `json:"total_current_liabilities"`
			} `json:"current_liabilities"`
			NonCurrentLiabilities struct {
				LongTermDebt                int64 `json:"long_term_debt"`
				ProvisionForRisksAndCharges int64 `json:"provision_for_risks_and_charges"`
				DeferredLiabilities         int64 `json:"deferred_liabilities"`
				OtherNonCurrentLiabilities  int64 `json:"other_non_current_liabilities"`
				TotalNonCurrentLiabilities  int64 `json:"total_non_current_liabilities"`
			} `json:"non_current_liabilities"`
			TotalLiabilities int64 `json:"total_liabilities"`
		} `json:"liabilities"`
		ShareholdersEquity struct {
			CommonStock             int64 `json:"common_stock"`
			RetainedEarnings        int64 `json:"retained_earnings"`
			OtherShareholdersEquity int64 `json:"other_shareholders_equity"`
			TotalShareholdersEquity int64 `json:"total_shareholders_equity"`
		} `json:"shareholders_equity"`
	} `json:"balance_sheet"`
}
