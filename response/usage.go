package response

import "gopkg.in/guregu/null.v4"

type Usage struct {
	TimeStamp      string   `json:"timestamp"`
	CurrentUsage   null.Int `json:"current_usage"`
	PlanLimit      null.Int `json:"plan_limit"`
	DailyUsage     null.Int `json:"daily_usage"`
	PlanDailyLimit null.Int `json:"plan_daily_limit"`
}
