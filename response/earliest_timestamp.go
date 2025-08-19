package response

// EarliestTimestamp represents the earliest available data timestamp for a symbol.
type EarliestTimestamp struct {
	Datetime string `json:"datetime"`
	UnixTime int64  `json:"unix_time"`
}
