package response

// AnalystRatingsSnapshot represents analyst ratings snapshot data.
type AnalystRatingsSnapshot struct {
	Meta    AnalysisMeta                  `json:"meta"`
	Ratings []AnalystRatingsSnapshotEntry `json:"ratings"`
	Status  string                        `json:"status"`
}

// AnalystRatingsSnapshotEntry represents a single analyst rating snapshot.
type AnalystRatingsSnapshotEntry struct {
	Date          string `json:"date"`
	Firm          string `json:"firm"`
	RatingChange  string `json:"rating_change"`
	RatingCurrent string `json:"rating_current"`
	RatingPrior   string `json:"rating_prior"`
}
