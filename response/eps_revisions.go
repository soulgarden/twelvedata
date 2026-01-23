package response

// EPSRevisions represents EPS revisions data.
type EPSRevisions struct {
	Meta        AnalysisMeta       `json:"meta"`
	EPSRevision []EPSRevisionEntry `json:"eps_revision"`
	Status      string             `json:"status"`
}

// EPSRevisionEntry represents a single EPS revision record.
type EPSRevisionEntry struct {
	Date          string `json:"date"`
	Period        string `json:"period"`
	UpLastWeek    int    `json:"up_last_week"`
	UpLastMonth   int    `json:"up_last_month"`
	DownLastWeek  int    `json:"down_last_week"`
	DownLastMonth int    `json:"down_last_month"`
}
