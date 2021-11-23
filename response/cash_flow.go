package response

// nolint: tagliatelle
type CashFlow struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Period           string `json:"period"`
	} `json:"meta"`
	CashFlow []struct {
		FiscalDate          string `json:"fiscal_date"`
		OperatingActivities struct {
			NetIncome              int64 `json:"net_income"`
			Depreciation           int64 `json:"depreciation"`
			DeferredTaxes          int64 `json:"deferred_taxes"`
			StockBasedCompensation int64 `json:"stock_based_compensation"`
			OtherNonCashItems      int64 `json:"other_non_cash_items"`
			AccountsReceivable     int64 `json:"accounts_receivable"`
			AccountsPayable        int64 `json:"accounts_payable"`
			OtherAssetsLiabilities int64 `json:"other_assets_liabilities"`
			OperatingCashFlow      int64 `json:"operating_cash_flow"`
		} `json:"operating_activities"`
		InvestingActivities struct {
			CapitalExpenditures    int64 `json:"capital_expenditures"`
			NetIntangibles         int64 `json:"net_intangibles"`
			NetAcquisitions        int64 `json:"net_acquisitions"`
			PurchaseOfInvestments  int64 `json:"purchase_of_investments"`
			SaleOfInvestments      int64 `json:"sale_of_investments"`
			OtherInvestingActivity int64 `json:"other_investing_activity"`
			InvestingCashFlow      int64 `json:"investing_cash_flow"`
		} `json:"investing_activities"`
		FinancingActivities struct {
			LongTermDebtIssuance  int64 `json:"long_term_debt_issuance"`
			LongTermDebtPayments  int64 `json:"long_term_debt_payments"`
			ShortTermDebtIssuance int64 `json:"short_term_debt_issuance"`
			CommonStockIssuance   int64 `json:"common_stock_issuance"`
			CommonStockRepurchase int64 `json:"common_stock_repurchase"`
			CommonDividends       int64 `json:"common_dividends"`
			OtherFinancingCharges int64 `json:"other_financing_charges"`
			FinancingCashFlow     int64 `json:"financing_cash_flow"`
		} `json:"financing_activities"`
		EndCashPosition int64 `json:"end_cash_position"`
		IncomeTaxPaid   int64 `json:"income_tax_paid"`
		InterestPaid    int64 `json:"interest_paid"`
		FreeCashFlow    int64 `json:"free_cash_flow"`
	} `json:"cash_flow"`
}
