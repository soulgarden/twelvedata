package response

import "github.com/guregu/null/v6"

// Usage represents API usage statistics including current usage and plan limits.
type Usage struct {
	TimeStamp    string   `json:"timestamp"`
	CurrentUsage null.Int `json:"current_usage"`
	PlanLimit    null.Int `json:"plan_limit"`
	PlanCategory string   `json:"plan_category"`
}
