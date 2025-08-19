package response

import "github.com/guregu/null/v6"

// MutualFundComposition represents the holdings and composition of a mutual fund.
type MutualFundComposition struct {
	TopHoldings       []Holding           `json:"top_holdings"`
	SectorAllocation  []SectorAllocation  `json:"sector_allocation"`
	CountryAllocation []CountryAllocation `json:"country_allocation"`
	AssetAllocation   AssetAllocation     `json:"asset_allocation"`
	RegionAllocation  []RegionAllocation  `json:"region_allocation"`
	LastUpdated       string              `json:"last_updated"`
}

// RegionAllocation represents allocation by geographic region.
type RegionAllocation struct {
	Region     string     `json:"region"`
	Percentage null.Float `json:"percentage"`
}
