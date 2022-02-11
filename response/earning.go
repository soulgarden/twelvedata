package response

import "gopkg.in/guregu/null.v4"

type Earnings struct {
	Earnings map[string][]*Earning `json:"earnings"`
}

type Earning struct {
	Symbol      string     `json:"symbol"`
	Name        string     `json:"name"`
	Currency    string     `json:"currency"`
	Exchange    string     `json:"exchange"`
	Time        string     `json:"time"`
	EpsEstimate null.Float `json:"eps_estimate"`
	EpsActual   null.Float `json:"eps_actual"`
	Difference  null.Float `json:"difference"`
	SurprisePrc null.Float `json:"surprise_prc"`
}
