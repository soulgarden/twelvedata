package response

type EarliestTimestamp struct {
	Datetime string `json:"datetime"`
	UnixTime int64  `json:"unix_time"`
}
