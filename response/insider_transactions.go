package response

// nolint: tagliatelle
type InsiderTransactions struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Currency         string `json:"currency"`
		Exchange         string `json:"exchange"`
		ExchangeTimezone string `json:"exchange_timezone"`
	} `json:"meta"`
	InsiderTransactions []struct {
		FullName     string `json:"full_name"`
		Position     string `json:"position"`
		DateReported string `json:"date_reported"`
		IsDirect     bool   `json:"is_direct"`
		Shares       int    `json:"shares"`
		Value        int    `json:"value"`
		Description  string `json:"description"`
	} `json:"insider_transactions"`
}
