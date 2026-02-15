package response

import "github.com/guregu/null/v6"

// InsiderTransactions represents the response structure for insider transactions data.
type InsiderTransactions struct {
	Meta                InsiderTransactionsMeta `json:"meta"`
	InsiderTransactions []InsiderTransaction    `json:"insider_transactions"`
}

// InsiderTransactionsMeta contains metadata for insider transactions data.
type InsiderTransactionsMeta struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Currency         string `json:"currency"`
	Exchange         string `json:"exchange"`
	MicCode          string `json:"mic_code"`
	ExchangeTimezone string `json:"exchange_timezone"`
}

// InsiderTransaction represents a single insider trading transaction.
type InsiderTransaction struct {
	FullName     string   `json:"full_name"`
	Position     string   `json:"position"`
	DateReported string   `json:"date_reported"`
	IsDirect     bool     `json:"is_direct"`
	Shares       null.Int `json:"shares"`
	Value        null.Int `json:"value"`
	Description  string   `json:"description"`
}
