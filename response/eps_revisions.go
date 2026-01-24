package response

import "github.com/guregu/null/v6"

// EPSRevisions represents EPS revisions data.
type EPSRevisions struct {
	Meta        AnalysisMeta       `json:"meta"`
	EPSRevision []EPSRevisionEntry `json:"eps_revision"`
	Status      string             `json:"status"`
}

// EPSRevisionEntry represents a single EPS revision record.
type EPSRevisionEntry struct {
	Date          string   `json:"date"`
	Period        string   `json:"period"`
	UpLastWeek    null.Int `json:"up_last_week"`
	UpLastMonth   null.Int `json:"up_last_month"`
	DownLastWeek  null.Int `json:"down_last_week"`
	DownLastMonth null.Int `json:"down_last_month"`
}
