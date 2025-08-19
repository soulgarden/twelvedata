package response

import "github.com/guregu/null/v6"

// EPSTrend represents EPS trend analysis data.
type EPSTrend struct {
	Symbol      string            `json:"symbol"`
	Exchange    string            `json:"exchange"`
	Current     EPSTrendCurrent   `json:"current"`
	Trends      EPSTrendAnalysis  `json:"trends"`
	Revisions   EPSTrendRevisions `json:"revisions"`
	LastUpdated string            `json:"last_updated"`
}

// EPSTrendCurrent represents current EPS trend information.
type EPSTrendCurrent struct {
	CurrentQuarterEPS null.Float `json:"current_quarter_eps"`
	NextQuarterEPS    null.Float `json:"next_quarter_eps"`
	CurrentYearEPS    null.Float `json:"current_year_eps"`
	NextYearEPS       null.Float `json:"next_year_eps"`
	LongTermGrowth    null.Float `json:"long_term_growth"`
	AnalystCount      int        `json:"analyst_count"`
}

// EPSTrendAnalysis represents EPS trend analysis over different periods.
type EPSTrendAnalysis struct {
	CurrentQuarter EPSTrendPeriod `json:"current_quarter"`
	NextQuarter    EPSTrendPeriod `json:"next_quarter"`
	CurrentYear    EPSTrendPeriod `json:"current_year"`
	NextYear       EPSTrendPeriod `json:"next_year"`
}

// EPSTrendPeriod represents EPS trend data for a specific period.
type EPSTrendPeriod struct {
	Period        string     `json:"period"`
	EPSEstimate   null.Float `json:"eps_estimate"`
	EPSGrowth     null.Float `json:"eps_growth"`
	RevenueGrowth null.Float `json:"revenue_growth"`
	AnalystCount  int        `json:"analyst_count"`
}

// EPSTrendRevisions represents EPS revision trends over different time periods.
type EPSTrendRevisions struct {
	UpRevisions7d    int        `json:"up_revisions_7d"`
	DownRevisions7d  int        `json:"down_revisions_7d"`
	UpRevisions30d   int        `json:"up_revisions_30d"`
	DownRevisions30d int        `json:"down_revisions_30d"`
	UpRevisions60d   int        `json:"up_revisions_60d"`
	DownRevisions60d int        `json:"down_revisions_60d"`
	UpRevisions90d   int        `json:"up_revisions_90d"`
	DownRevisions90d int        `json:"down_revisions_90d"`
	NetRevision      null.Float `json:"net_revision"`
}
