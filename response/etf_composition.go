package response

import "github.com/guregu/null/v6"

// ETFComposition represents the holdings and composition of an ETF.
type ETFComposition struct {
	TopHoldings       []Holding           `json:"top_holdings"`
	SectorAllocation  []SectorAllocation  `json:"sector_allocation"`
	CountryAllocation []CountryAllocation `json:"country_allocation"`
	AssetAllocation   AssetAllocation     `json:"asset_allocation"`
	LastUpdated       string              `json:"last_updated"`
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
