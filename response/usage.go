package response

import "gopkg.in/guregu/null.v4"

type Usage struct {
	TimeStamp    string   `json:"timestamp"`
	CurrentUsage null.Int `json:"current_usage"`
	PlanLimit    null.Int `json:"plan_limit"`
}
