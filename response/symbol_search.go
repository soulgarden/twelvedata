package response

type SymbolSearch struct {
	Data   []*SymbolSearchResult `json:"data"`
	Status string                `json:"status"`
}

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

type SymbolSearchResultAccess struct {
	Global string `json:"global"`
	Plan   string `json:"plan"`
}
