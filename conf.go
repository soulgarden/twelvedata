package twelvedata

type Conf struct {
	BaseURL   string `default:"https://api.twelvedata.com" json:"base_url"`
	BaseWSURL string `default:"ws.twelvedata.com"          json:"base_ws_url"`

	ReferenceData ReferenceData `json:"reference_data"`
	CoreData      CoreData      `json:"core_data"`
	Fundamentals  Fundamentals  `json:"fundamentals"`
	WebSocket     WebSocket     `json:"web_socket"`
	Advanced      Advanced      `json:"advanced"`

	APIKey  string `default:"demo" json:"api_key"`
	Timeout int    `default:"15"   json:"timeout"`
}

// nolint: lll
type ReferenceData struct {
	StocksURL      string `default:"/stocks"       json:"stocks_url"`
	ExchangesURL   string `default:"/exchanges"    json:"exchange_url"`
	IndicesURL     string `default:"/indices"      json:"indices_url"`
	EtfsURL        string `default:"/etf"          json:"etfs_url"`
	MarketStateURL string `default:"/market_state" json:"market_state_url"`
}

// nolint: lll
type CoreData struct {
	TimeSeriesURL string `default:"/time_series" json:"time_series_url"`
	QuotesURL     string `default:"/quote"       json:"quotes_url"`

	ExchangeRateURL string `default:"/exchange_rate" json:"exchange_rate_url"`
	MarketMoversURL string `default:"/market_movers" json:"market_movers_url"`
}

// nolint: lll
type Fundamentals struct {
	EarningsCalendarURL    string `default:"/earnings_calendar"    json:"earnings_calendar_url"`
	ProfileURL             string `default:"/profile"              json:"profile_url"`
	InsiderTransactionsURL string `default:"/insider_transactions" json:"insider_transactions_url"`
	IncomeStatementURL     string `default:"/income_statement"     json:"income_statement_url"`
	BalanceSheetURL        string `default:"/balance_sheet"        json:"balance_sheet_url"`
	CashFlowURL            string `default:"/cash_flow"            json:"cash_flow_url"`
	DividendsURL           string `default:"/dividends"            json:"dividends_url"`
	StatisticsURL          string `default:"/statistics"           json:"statistics_url"`
}

type WebSocket struct {
	PriceURL string `default:"/v1/quotes/price" json:"ws_price_url"`
}

type Advanced struct {
	UsageURL string `default:"/api_usage" json:"usage_url"`
}
