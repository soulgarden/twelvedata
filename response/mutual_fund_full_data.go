package response

// MutualFundFullData represents the response structure for mutual fund full data.
type MutualFundFullData struct {
	MutualFund MutualFundFullDataData `json:"mutual_fund"`
	Status     string                 `json:"status"`
}

// MutualFundFullDataData contains the full mutual fund data payload.
type MutualFundFullDataData struct {
	Summary        MutualFundSummaryInfo           `json:"summary"`
	Performance    MutualFundPerformanceInfo       `json:"performance"`
	Risk           MutualFundRiskInfo              `json:"risk"`
	Ratings        MutualFundRatingsInfo           `json:"ratings"`
	Composition    MutualFundCompositionInfo       `json:"composition"`
	PurchaseInfo   MutualFundPurchaseInfoDetails   `json:"purchase_info"`
	Sustainability MutualFundSustainabilityDetails `json:"sustainability"`
}
