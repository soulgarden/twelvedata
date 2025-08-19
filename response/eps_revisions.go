package response

import "github.com/guregu/null/v6"

// EPSRevisions represents EPS revisions data.
type EPSRevisions struct {
	Symbol      string                `json:"symbol"`
	Exchange    string                `json:"exchange"`
	Summary     EPSRevisionsSummary   `json:"summary"`
	History     []EPSRevisionsHistory `json:"history"`
	Breakdown   EPSRevisionsBreakdown `json:"breakdown"`
	LastUpdated string                `json:"last_updated"`
}

// EPSRevisionsSummary represents overall EPS revisions summary.
type EPSRevisionsSummary struct {
	TotalRevisions    int        `json:"total_revisions"`
	UpRevisions       int        `json:"up_revisions"`
	DownRevisions     int        `json:"down_revisions"`
	NoChangeRevisions int        `json:"no_change_revisions"`
	NetRevision       null.Float `json:"net_revision"`
	AverageRevision   null.Float `json:"average_revision"`
	AnalystCount      int        `json:"analyst_count"`
}

// EPSRevisionsHistory represents historical EPS revision data.
type EPSRevisionsHistory struct {
	Date         string     `json:"date"`
	Period       string     `json:"period"`
	OldEstimate  null.Float `json:"old_estimate"`
	NewEstimate  null.Float `json:"new_estimate"`
	Revision     null.Float `json:"revision"`
	RevisionType string     `json:"revision_type"`
	Analyst      string     `json:"analyst"`
}

// EPSRevisionsBreakdown represents EPS revisions breakdown by time period.
type EPSRevisionsBreakdown struct {
	Last7Days  EPSRevisionsPeriod `json:"last_7_days"`
	Last30Days EPSRevisionsPeriod `json:"last_30_days"`
	Last60Days EPSRevisionsPeriod `json:"last_60_days"`
	Last90Days EPSRevisionsPeriod `json:"last_90_days"`
	LastYear   EPSRevisionsPeriod `json:"last_year"`
}

// EPSRevisionsPeriod represents EPS revision data for a specific time period.
type EPSRevisionsPeriod struct {
	UpRevisions     int        `json:"up_revisions"`
	DownRevisions   int        `json:"down_revisions"`
	NetRevision     null.Float `json:"net_revision"`
	AverageRevision null.Float `json:"average_revision"`
	TotalRevisions  int        `json:"total_revisions"`
}
