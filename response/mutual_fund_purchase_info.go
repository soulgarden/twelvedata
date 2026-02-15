package response

import "github.com/guregu/null/v6"

// MutualFundPurchaseInfo represents the response structure for mutual fund purchase info data.
type MutualFundPurchaseInfo struct {
	MutualFund MutualFundPurchaseInfoData `json:"mutual_fund"`
	Status     string                     `json:"status"`
}

// MutualFundPurchaseInfoData contains the purchase information for a mutual fund.
type MutualFundPurchaseInfoData struct {
	PurchaseInfo MutualFundPurchaseInfoDetails `json:"purchase_info"`
}

// MutualFundPurchaseInfoDetails contains detailed purchase information for a mutual fund.
type MutualFundPurchaseInfoDetails struct {
	Expenses   MutualFundPurchaseExpenses `json:"expenses"`
	Minimums   MutualFundPurchaseMinimums `json:"minimums"`
	Pricing    MutualFundPurchasePricing  `json:"pricing"`
	Brokerages []string                   `json:"brokerages"`
}

// MutualFundPurchaseExpenses represents expense ratios for a mutual fund.
type MutualFundPurchaseExpenses struct {
	ExpenseRatioGross null.Float `json:"expense_ratio_gross"`
	ExpenseRatioNet   null.Float `json:"expense_ratio_net"`
}

// MutualFundPurchaseMinimums represents minimum investment amounts.
type MutualFundPurchaseMinimums struct {
	InitialInvestment       null.Int    `json:"initial_investment"`
	AdditionalInvestment    null.Int    `json:"additional_investment"`
	InitialIRAInvestment    null.String `json:"initial_ira_investment"`
	AdditionalIRAInvestment null.String `json:"additional_ira_investment"`
}

// MutualFundPurchasePricing represents pricing information for the mutual fund.
type MutualFundPurchasePricing struct {
	NAV             null.Float `json:"nav"`
	TwelveMonthLow  null.Float `json:"12_month_low"`
	TwelveMonthHigh null.Float `json:"12_month_high"`
	LastMonth       null.Float `json:"last_month"`
}
