package request

// GetLastChange represents request for Last Changes endpoint /last_change/{endpoint}
// The endpoint parameter is substituted in the URL path.
type GetLastChange struct {
	APIKey
	Endpoint   string `schema:"-"`                    // URL path parameter, not query param
	Symbol     string `schema:"symbol,omitempty"`     // Symbol to track changes for
	Exchange   string `schema:"exchange,omitempty"`   // Exchange for the symbol
	MicCode    string `schema:"mic_code,omitempty"`   // Market Identifier Code (MIC)
	Country    string `schema:"country,omitempty"`    // Country filter
	StartDate  string `schema:"start_date,omitempty"` // Start date range
	Page       int    `schema:"page,omitempty"`       // Pagination page number
	OutputSize int    `schema:"outputsize,omitempty"` // Items per page
}
