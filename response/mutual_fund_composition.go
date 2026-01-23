package response

import "github.com/guregu/null/v6"

// MutualFundComposition represents the response structure for mutual fund composition data.
type MutualFundComposition struct {
	MutualFund MutualFundCompositionData `json:"mutual_fund"`
	Status     string                    `json:"status"`
}

// MutualFundCompositionData contains the composition information for a mutual fund.
type MutualFundCompositionData struct {
	Composition MutualFundCompositionInfo `json:"composition"`
}

// MutualFundCompositionInfo represents holdings and composition of a mutual fund.
type MutualFundCompositionInfo struct {
	MajorMarketSectors []MutualFundSectorWeight  `json:"major_market_sectors"`
	AssetAllocation    MutualFundAssetAllocation `json:"asset_allocation"`
	TopHoldings        []MutualFundTopHolding    `json:"top_holdings"`
	BondBreakdown      MutualFundBondBreakdown   `json:"bond_breakdown"`
}

// MutualFundSectorWeight represents allocation by major market sector.
type MutualFundSectorWeight struct {
	Sector string     `json:"sector"`
	Weight null.Float `json:"weight"`
}

// MutualFundAssetAllocation represents allocation by asset type.
type MutualFundAssetAllocation struct {
	Cash            null.Float `json:"cash"`
	Stocks          null.Float `json:"stocks"`
	PreferredStocks null.Float `json:"preferred_stocks"`
	Convertibles    null.Float `json:"convertibles"`
	Bonds           null.Float `json:"bonds"`
	Others          null.Float `json:"others"`
}

// MutualFundTopHolding represents a top holding in a mutual fund.
type MutualFundTopHolding struct {
	Symbol   string     `json:"symbol"`
	Name     string     `json:"name"`
	Exchange string     `json:"exchange"`
	MicCode  string     `json:"mic_code"`
	Weight   null.Float `json:"weight"`
}

// MutualFundBondBreakdown represents bond breakdown metrics for a mutual fund.
type MutualFundBondBreakdown struct {
	AverageMaturity MutualFundBondMetric      `json:"average_maturity"`
	AverageDuration MutualFundBondMetric      `json:"average_duration"`
	CreditQuality   []MutualFundCreditQuality `json:"credit_quality"`
}

// MutualFundBondMetric represents bond metric values for fund and category.
type MutualFundBondMetric struct {
	Fund     null.Float `json:"fund"`
	Category null.Float `json:"category"`
}

// MutualFundCreditQuality represents credit quality distribution.
type MutualFundCreditQuality struct {
	Grade  string     `json:"grade"`
	Weight null.Float `json:"weight"`
}
