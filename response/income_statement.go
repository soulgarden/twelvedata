package response

import "gopkg.in/guregu/null.v4"

type IncomeStatements struct {
	Meta            IncomeStatementsMeta `json:"meta"`
	IncomeStatement []IncomeStatement    `json:"income_statement"`
}

type IncomeStatementsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
	Period           string `json:"period"`
}

type IncomeStatement struct {
	FiscalDate                    string                              `json:"fiscal_date"`
	Quarter                       null.Int                            `json:"quarter"`
	Sales                         null.Int                            `json:"sales"`
	CostOfGoods                   null.Int                            `json:"cost_of_goods"`
	GrossProfit                   null.Int                            `json:"gross_profit"`
	OperatingExpense              IncomeStatementOperatingExpense     `json:"operating_expense"`
	OperatingIncome               null.Int                            `json:"operating_income"`
	NonOperatingInterest          IncomeStatementNonOperatingInterest `json:"non_operating_interest"`
	OtherIncomeExpense            null.Int                            `json:"other_income_expense"`
	PretaxIncome                  null.Int                            `json:"pretax_income"`
	IncomeTax                     null.Int                            `json:"income_tax"`
	NetIncome                     null.Int                            `json:"net_income"`
	EpsBasic                      null.Float                          `json:"eps_basic"`
	EpsDiluted                    null.Float                          `json:"eps_diluted"`
	BasicSharesOutstanding        null.Int                            `json:"basic_shares_outstanding"`
	DilutedSharesOutstanding      null.Int                            `json:"diluted_shares_outstanding"`
	Ebitda                        null.Int                            `json:"ebitda"`
	NetIncomeContinuousOperations null.Int                            `json:"net_income_continuous_operations"`
	MinorityInterests             null.Int                            `json:"minority_interests"`
	PreferredStockDividends       null.Int                            `json:"preferred_stock_dividends"`
}

type IncomeStatementOperatingExpense struct {
	ResearchAndDevelopment          null.Int `json:"research_and_development"`
	SellingGeneralAndAdministrative null.Int `json:"selling_general_and_administrative"`
	OtherOperatingExpenses          null.Int `json:"other_operating_expenses"`
}

type IncomeStatementNonOperatingInterest struct {
	Income  null.Int `json:"income"`
	Expense null.Int `json:"expense"`
}
