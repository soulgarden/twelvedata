package response

// SymbolSearch represents the response structure for symbol search data.
type SymbolSearch struct {
	Data   []*SymbolSearchResult `json:"data"`
	Status string                `json:"status"`
}

// SymbolSearchResult represents a single search result for a financial symbol.
type SymbolSearchResult struct {
	Symbol           string                    `json:"symbol"`
	InstrumentName   string                    `json:"instrument_name"`
	Exchange         string                    `json:"exchange"`
	MicCode          string                    `json:"mic_code"`
	ExchangeTimezone string                    `json:"exchange_timezone"`
	InstrumentType   string                    `json:"instrument_type"`
	Country          string                    `json:"country"`
	Currency         string                    `json:"currency"`
	Access           *SymbolSearchResultAccess `json:"access"`
}

// SymbolSearchResultAccess represents access level information for symbol search results.
type SymbolSearchResultAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
