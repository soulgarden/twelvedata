package response

// nolint: tagliatelle
type IncomeStatement struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Period           string `json:"period"`
	} `json:"meta"`
	IncomeStatement []struct {
		FiscalDate       string `json:"fiscal_date"`
		Sales            int64  `json:"sales"`
		CostOfGoods      int64  `json:"cost_of_goods"`
		GrossProfit      int64  `json:"gross_profit"`
		OperatingExpense struct {
			ResearchAndDevelopment          int64 `json:"research_and_development"`
			SellingGeneralAndAdministrative int64 `json:"selling_general_and_administrative"`
		} `json:"operating_expense"`
		OperatingIncome      int64 `json:"operating_income"`
		NonOperatingInterest struct {
			Income  int64 `json:"income"`
			Expense int64 `json:"expense"`
		} `json:"non_operating_interest"`
		OtherIncomeExpense       int64   `json:"other_income_expense"`
		PretaxIncome             int64   `json:"pretax_income"`
		IncomeTax                int64   `json:"income_tax"`
		NetIncome                int64   `json:"net_income"`
		EpsBasic                 float64 `json:"eps_basic"`
		EpsDiluted               float64 `json:"eps_diluted"`
		BasicSharesOutstanding   int64   `json:"basic_shares_outstanding"`
		DilutedSharesOutstanding int64   `json:"diluted_shares_outstanding"`
		Ebitda                   int64   `json:"ebitda"`
	} `json:"income_statement"`
}
