package response

// nolint: tagliatelle
type Usage struct {
	TimeStamp      string `json:"timestamp"`
	CurrentUsage   int    `json:"current_usage"`
	PlanLimit      int    `json:"plan_limit"`
	DailyUsage     int    `json:"daily_usage"`
	PlanDailyLimit int    `json:"plan_daily_limit"`
}
