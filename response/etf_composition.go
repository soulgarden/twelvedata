package response

import "github.com/guregu/null/v6"

// ETFComposition represents the response structure for ETF composition data.
type ETFComposition struct {
	ETF    ETFCompositionData `json:"etf"`
	Status string             `json:"status"`
}

// ETFCompositionData contains the composition information for an ETF.
type ETFCompositionData struct {
	Composition ETFWorldComposition `json:"composition"`
}

// ETFWorldComposition represents holdings and composition of an ETF.
type ETFWorldComposition struct {
	MajorMarketSectors []ETFSectorWeight      `json:"major_market_sectors"`
	CountryAllocation  []ETFCountryAllocation `json:"country_allocation"`
	AssetAllocation    ETFAssetAllocation     `json:"asset_allocation"`
	TopHoldings        []ETFTopHolding        `json:"top_holdings"`
	BondBreakdown      ETFBondBreakdown       `json:"bond_breakdown"`
}

// ETFSectorWeight represents allocation by major market sector.
type ETFSectorWeight struct {
	Sector string     `json:"sector"`
	Weight null.Float `json:"weight"`
}

// ETFCountryAllocation represents allocation by country.
type ETFCountryAllocation struct {
	Country    string     `json:"country"`
	Allocation null.Float `json:"allocation"`
}

// ETFAssetAllocation represents allocation by asset type.
type ETFAssetAllocation struct {
	Cash            null.Float `json:"cash"`
	Stocks          null.Float `json:"stocks"`
	PreferredStocks null.Float `json:"preferred_stocks"`
	Convertibles    null.Float `json:"convertibles"`
	Bonds           null.Float `json:"bonds"`
	Others          null.Float `json:"others"`
}

// ETFTopHolding represents a top holding in an ETF.
type ETFTopHolding struct {
	Symbol   string     `json:"symbol"`
	Name     string     `json:"name"`
	Exchange string     `json:"exchange"`
	MicCode  string     `json:"mic_code"`
	Weight   null.Float `json:"weight"`
}

// ETFBondBreakdown represents bond breakdown metrics for an ETF.
type ETFBondBreakdown struct {
	AverageMaturity ETFBondMetric      `json:"average_maturity"`
	AverageDuration ETFBondMetric      `json:"average_duration"`
	CreditQuality   []ETFCreditQuality `json:"credit_quality"`
}

// ETFBondMetric represents bond metric values for fund and category.
type ETFBondMetric struct {
	Fund     null.Float `json:"fund"`
	Category null.Float `json:"category"`
}

// ETFCreditQuality represents credit quality distribution.
type ETFCreditQuality struct {
	Grade  string     `json:"grade"`
	Weight null.Float `json:"weight"`
}

// Holding represents an individual holding in an ETF.
type Holding struct {
	Symbol      string     `json:"symbol"`
	Name        string     `json:"name"`
	Percentage  null.Float `json:"percentage"`
	Shares      null.Int   `json:"shares"`
	MarketValue null.Float `json:"market_value"`
}

// SectorAllocation represents allocation by sector.
type SectorAllocation struct {
	Sector     string     `json:"sector"`
	Percentage null.Float `json:"percentage"`
}

// CountryAllocation represents allocation by country.
type CountryAllocation struct {
	Country    string     `json:"country"`
	Percentage null.Float `json:"percentage"`
}

// AssetAllocation represents allocation by asset type.
type AssetAllocation struct {
	Stocks null.Float `json:"stocks"`
	Bonds  null.Float `json:"bonds"`
	Cash   null.Float `json:"cash"`
	Other  null.Float `json:"other"`
}
