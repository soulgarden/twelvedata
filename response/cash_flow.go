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
	OperatingActivities CashFlowOperatingActivities `json:"operating_activities"`
	InvestingActivities CashFlowInvestingActivities `json:"investing_activities"`
	FinancingActivities CashFlowFinancingActivities `json:"financing_activities"`
	EndCashPosition     null.Int                    `json:"end_cash_position"`
	IncomeTaxPaid       null.Int                    `json:"income_tax_paid"`
	InterestPaid        null.Int                    `json:"interest_paid"`
	FreeCashFlow        null.Int                    `json:"free_cash_flow"`
}

// CashFlowOperatingActivities represents operating activities section of cash flow statement.
type CashFlowOperatingActivities struct {
	NetIncome              null.Int `json:"net_income"`
	Depreciation           null.Int `json:"depreciation"`
	DeferredTaxes          null.Int `json:"deferred_taxes"`
	StockBasedCompensation null.Int `json:"stock_based_compensation"`
	OtherNonCashItems      null.Int `json:"other_non_cash_items"`
	AccountsReceivable     null.Int `json:"accounts_receivable"`
	AccountsPayable        null.Int `json:"accounts_payable"`
	OtherAssetsLiabilities null.Int `json:"other_assets_liabilities"`
	OperatingCashFlow      null.Int `json:"operating_cash_flow"`
}

// CashFlowInvestingActivities represents investing activities section of cash flow statement.
type CashFlowInvestingActivities struct {
	CapitalExpenditures    null.Int `json:"capital_expenditures"`
	NetIntangibles         null.Int `json:"net_intangibles"`
	NetAcquisitions        null.Int `json:"net_acquisitions"`
	PurchaseOfInvestments  null.Int `json:"purchase_of_investments"`
	SaleOfInvestments      null.Int `json:"sale_of_investments"`
	OtherInvestingActivity null.Int `json:"other_investing_activity"`
	InvestingCashFlow      null.Int `json:"investing_cash_flow"`
}

// CashFlowFinancingActivities represents financing activities section of cash flow statement.
type CashFlowFinancingActivities struct {
	LongTermDebtIssuance  null.Int `json:"long_term_debt_issuance"`
	LongTermDebtPayments  null.Int `json:"long_term_debt_payments"`
	ShortTermDebtIssuance null.Int `json:"short_term_debt_issuance"`
	CommonStockIssuance   null.Int `json:"common_stock_issuance"`
	CommonStockRepurchase null.Int `json:"common_stock_repurchase"`
	CommonDividends       null.Int `json:"common_dividends"`
	OtherFinancingCharges null.Int `json:"other_financing_charges"`
	FinancingCashFlow     null.Int `json:"financing_cash_flow"`
}
