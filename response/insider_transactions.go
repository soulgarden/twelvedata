package response

type InsiderTransactions struct {
	Meta                InsiderTransactionsMeta `json:"meta"`
	InsiderTransactions []InsiderTransaction    `json:"insider_transactions"`
}

type InsiderTransactionsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

type InsiderTransaction struct {
	FullName     string `json:"full_name"`
	Position     string `json:"position"`
	DateReported string `json:"date_reported"`
	IsDirect     bool   `json:"is_direct"`
	Shares       int    `json:"shares"`
	Value        int    `json:"value"`
	Description  string `json:"description"`
}
