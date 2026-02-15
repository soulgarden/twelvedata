package response

// ETFFullData represents the response structure for ETF full data.
type ETFFullData struct {
	ETF    ETFFullDataData `json:"etf"`
	Status string          `json:"status"`
}

// ETFFullDataData contains the full ETF data payload.
type ETFFullDataData struct {
	Summary     ETFWorldSummaryInfo `json:"summary"`
	Performance ETFWorldPerformance `json:"performance"`
	Risk        ETFWorldRisk        `json:"risk"`
	Composition ETFWorldComposition `json:"composition"`
}
