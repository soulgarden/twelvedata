package response

type ExchangeSchedule struct {
	Data   []*ExchangeScheduleItem `json:"data"`
	Status string                  `json:"status"`
}

type ExchangeScheduleItem struct {
	Title    string                     `json:"title"`
	Name     string                     `json:"name"`
	Code     string                     `json:"code"`
	Country  string                     `json:"country"`
	TimeZone string                     `json:"time_zone"`
	Sessions []*ExchangeScheduleSession `json:"sessions"`
}

type ExchangeScheduleSession struct {
	OpenTime    string `json:"open_time"`
	CloseTime   string `json:"close_time"`
	SessionName string `json:"session_name"`
	SessionType string `json:"session_type"`
}
