package request

type GetIndices struct {
	symbol, country           string
	showPlan, includeDelisted bool
}
