package response

import "github.com/guregu/null/v6"

// CashFlows represents the response structure for cash flow data.
type CashFlows struct {
	Meta     CashFlowsMeta `json:"meta"`
	CashFlow []CashFlow    `json:"cash_flow"`
}

// CashFlowsMeta contains metadata for cash flow data.
type CashFlowsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Period           string `json:"period"`
}

// CashFlow represents cash flow statement data for a specific fiscal period.
type CashFlow struct {
	FiscalDate          string                      `json:"fiscal_date"`
	Quarter             null.Int                    `json:"quarter"`
	Year                null.Int                    `json:"year"`
	OperatingActivities CashFlowOperatingActivities `json:"operating_activities"`
	InvestingActivities CashFlowInvestingActivities `json:"investing_activities"`
	FinancingActivities CashFlowFinancingActivities `json:"financing_activities"`
	EndCashPosition     null.Float                  `json:"end_cash_position"`
	IncomeTaxPaid       null.Float                  `json:"income_tax_paid"`
	InterestPaid        null.Float                  `json:"interest_paid"`
	FreeCashFlow        null.Float                  `json:"free_cash_flow"`
}

// CashFlowOperatingActivities represents operating activities section of cash flow statement.
type CashFlowOperatingActivities struct {
	NetIncome              null.Float `json:"net_income"`
	Depreciation           null.Float `json:"depreciation"`
	DeferredTaxes          null.Float `json:"deferred_taxes"`
	StockBasedCompensation null.Float `json:"stock_based_compensation"`
	OtherNonCashItems      null.Float `json:"other_non_cash_items"`
	AccountsReceivable     null.Float `json:"accounts_receivable"`
	AccountsPayable        null.Float `json:"accounts_payable"`
	OtherAssetsLiabilities null.Float `json:"other_assets_liabilities"`
	OperatingCashFlow      null.Float `json:"operating_cash_flow"`
}

// CashFlowInvestingActivities represents investing activities section of cash flow statement.
type CashFlowInvestingActivities struct {
	CapitalExpenditures    null.Float `json:"capital_expenditures"`
	NetIntangibles         null.Float `json:"net_intangibles"`
	NetAcquisitions        null.Float `json:"net_acquisitions"`
	PurchaseOfInvestments  null.Float `json:"purchase_of_investments"`
	SaleOfInvestments      null.Float `json:"sale_of_investments"`
	OtherInvestingActivity null.Float `json:"other_investing_activity"`
	InvestingCashFlow      null.Float `json:"investing_cash_flow"`
}

// CashFlowFinancingActivities represents financing activities section of cash flow statement.
type CashFlowFinancingActivities struct {
	LongTermDebtIssuance  null.Float `json:"long_term_debt_issuance"`
	LongTermDebtPayments  null.Float `json:"long_term_debt_payments"`
	ShortTermDebtIssuance null.Float `json:"short_term_debt_issuance"`
	CommonStockIssuance   null.Float `json:"common_stock_issuance"`
	CommonStockRepurchase null.Float `json:"common_stock_repurchase"`
	CommonDividends       null.Float `json:"common_dividends"`
	OtherFinancingCharges null.Float `json:"other_financing_charges"`
	FinancingCashFlow     null.Float `json:"financing_cash_flow"`
}
