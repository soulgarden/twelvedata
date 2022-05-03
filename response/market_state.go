package response

type MarketState struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Country      string `json:"country"`
	IsMarketOpen bool   `json:"is_market_open"`
	TimeToOpen   string `json:"time_to_open"`
	TimeToClose  string `json:"time_to_close"`
}
