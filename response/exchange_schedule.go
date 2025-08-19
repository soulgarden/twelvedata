package response

// ExchangeSchedule represents the response structure for exchange schedule data.
type ExchangeSchedule struct {
	Data   []*ExchangeScheduleItem `json:"data"`
	Status string                  `json:"status"`
}

// ExchangeScheduleItem represents a single exchange with its schedule information.
type ExchangeScheduleItem struct {
	Title    string                     `json:"title"`
	Name     string                     `json:"name"`
	Code     string                     `json:"code"`
	Country  string                     `json:"country"`
	TimeZone string                     `json:"time_zone"`
	Sessions []*ExchangeScheduleSession `json:"sessions"`
}

// ExchangeScheduleSession represents a trading session with timing details.
type ExchangeScheduleSession struct {
	OpenTime    string `json:"open_time"`
	CloseTime   string `json:"close_time"`
	SessionName string `json:"session_name"`
	SessionType string `json:"session_type"`
}
