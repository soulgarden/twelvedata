package response

import "github.com/guregu/null/v6"

// MutualFundPurchaseInfoResponse represents the response structure for Mutual Fund Purchase Info data.
type MutualFundPurchaseInfoResponse struct {
	MutualFund MutualFundPurchaseInfoData `json:"mutual_fund"`
	Status     string                     `json:"status"`
}

// MutualFundPurchaseInfoData contains the purchase information for a mutual fund.
type MutualFundPurchaseInfoData struct {
	PurchaseInfo MutualFundPurchaseDetails `json:"purchase_info"`
}

// MutualFundPurchaseDetails contains detailed purchase information for a mutual fund.
type MutualFundPurchaseDetails struct {
	MinimumInvestment     null.Float                     `json:"minimum_investment"`
	MinimumAdditional     null.Float                     `json:"minimum_additional"`
	FrontEndLoad          null.Float                     `json:"front_end_load"`
	BackEndLoad           null.Float                     `json:"back_end_load"`
	TwelveB1Fee           null.Float                     `json:"12b1_fee"`
	ManagementFee         null.Float                     `json:"management_fee"`
	OtherFees             null.Float                     `json:"other_fees"`
	TotalAnnualOperating  null.Float                     `json:"total_annual_operating"`
	LoadStructures        MutualFundLoadStructures       `json:"load_structures"`
	PurchaseRestrictions  MutualFundPurchaseRestrictions `json:"purchase_restrictions"`
	AvailabilityInfo      MutualFundAvailabilityInfo     `json:"availability_info"`
	PurchaseConstraints   null.String                    `json:"purchase_constraints"`
	RedemptionConstraints null.String                    `json:"redemption_constraints"`
}

// MutualFundLoadStructures represents different load structures for the fund.
type MutualFundLoadStructures struct {
	ClassA          MutualFundLoadClass `json:"class_a"`
	ClassB          MutualFundLoadClass `json:"class_b"`
	ClassC          MutualFundLoadClass `json:"class_c"`
	Institutional   MutualFundLoadClass `json:"institutional"`
	LoadWaivers     []string            `json:"load_waivers"`
	BreakpointLevel null.Float          `json:"breakpoint_level"`
}

// MutualFundLoadClass represents load information for a specific fund class.
type MutualFundLoadClass struct {
	FrontEndLoad    null.Float  `json:"front_end_load"`
	BackEndLoad     null.Float  `json:"back_end_load"`
	MaxFrontEndLoad null.Float  `json:"max_front_end_load"`
	MaxBackEndLoad  null.Float  `json:"max_back_end_load"`
	LoadSchedule    null.String `json:"load_schedule"`
	TwelveB1Fee     null.Float  `json:"12b1_fee"`
	ManagementFee   null.Float  `json:"management_fee"`
	ExpenseRatio    null.Float  `json:"expense_ratio"`
}

// MutualFundPurchaseRestrictions represents purchase restrictions for the fund.
type MutualFundPurchaseRestrictions struct {
	MinimumInitialInvestment    null.Float  `json:"minimum_initial_investment"`
	MinimumSubsequentInvestment null.Float  `json:"minimum_subsequent_investment"`
	MaximumInvestment           null.Float  `json:"maximum_investment"`
	EligibleInvestors           []string    `json:"eligible_investors"`
	GeographicRestrictions      []string    `json:"geographic_restrictions"`
	AccreditedInvestorOnly      null.Bool   `json:"accredited_investor_only"`
	InstitutionalOnly           null.Bool   `json:"institutional_only"`
	TradingRestrictions         null.String `json:"trading_restrictions"`
}

// MutualFundAvailabilityInfo represents availability information for the fund.
type MutualFundAvailabilityInfo struct {
	AvailablePlatforms      []string  `json:"available_platforms"`
	BrokerageAvailability   []string  `json:"brokerage_availability"`
	DirectPurchaseAvailable null.Bool `json:"direct_purchase_available"`
	AutoInvestmentAvailable null.Bool `json:"auto_investment_available"`
	DividendReinvestment    null.Bool `json:"dividend_reinvestment"`
	SystematicWithdrawal    null.Bool `json:"systematic_withdrawal"`
	OnlineAccess            null.Bool `json:"online_access"`
	PhoneAccess             null.Bool `json:"phone_access"`
}

// MutualFundPurchaseInfo represents purchase information for a mutual fund (preserved for compatibility).
type MutualFundPurchaseInfo struct {
	MinimumInvestment     null.Float  `json:"minimum_investment"`
	MinimumAdditional     null.Float  `json:"minimum_additional"`
	FrontEndLoad          null.Float  `json:"front_end_load"`
	BackEndLoad           null.Float  `json:"back_end_load"`
	TwelveB1Fee           null.Float  `json:"12b1_fee"`
	ManagementFee         null.Float  `json:"management_fee"`
	OtherFees             null.Float  `json:"other_fees"`
	TotalAnnualOperating  null.Float  `json:"total_annual_operating"`
	PurchaseConstraints   null.String `json:"purchase_constraints"`
	RedemptionConstraints null.String `json:"redemption_constraints"`
}
