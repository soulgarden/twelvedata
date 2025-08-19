package response

import "github.com/guregu/null/v6"

// Usage represents API usage statistics including current usage and plan limits.
type Usage struct {
	TimeStamp      string   `json:"timestamp"`
	CurrentUsage   null.Int `json:"current_usage"`
	PlanLimit      null.Int `json:"plan_limit"`
	DailyUsage     null.Int `json:"daily_usage"`
	PlanDailyLimit null.Int `json:"plan_daily_limit"`
	PlanCategory   string   `json:"plan_category"`
}
