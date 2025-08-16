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
	// Asset Catalogs
	StocksURL           string `default:"/stocks"           json:"stocks_url"`
	ForexPairsURL       string `default:"/forex_pairs"      json:"forex_pairs_url"`
	CryptocurrenciesURL string `default:"/cryptocurrencies" json:"cryptocurrencies_url"`
	EtfsURL             string `default:"/etf"              json:"etfs_url"`
	FundsURL            string `default:"/funds"            json:"funds_url"`
	CommoditiesURL      string `default:"/commodities"      json:"commodities_url"`
	BondsURL            string `default:"/bonds"            json:"bonds_url"`

	// Discovery
	SymbolSearchURL      string `default:"/symbol_search"      json:"symbol_search_url"`
	CrossListingsURL     string `default:"/cross_listings"     json:"cross_listings_url"`
	EarliestTimestampURL string `default:"/earliest_timestamp" json:"earliest_timestamp_url"`

	// Markets
	ExchangesURL               string `default:"/exchanges"                json:"exchange_url"`
	ExchangeScheduleURL        string `default:"/exchange_schedule"        json:"exchange_schedule_url"`
	CryptocurrencyExchangesURL string `default:"/cryptocurrency_exchanges" json:"cryptocurrency_exchanges_url"`
	IndicesURL                 string `default:"/indices"                  json:"indices_url"`
	MarketStateURL             string `default:"/market_state"             json:"market_state_url"`

	// Supporting Metadata
	CountriesURL           string `default:"/countries"            json:"countries_url"`
	InstrumentTypeURL      string `default:"/instrument_type"      json:"instrument_type_url"`
	TechnicalIndicatorsURL string `default:"/technical_indicators" json:"technical_indicators_url"`
}

// nolint: lll
type CoreData struct {
	TimeSeriesURL      string `default:"/time_series"       json:"time_series_url"`
	TimeSeriesCrossURL string `default:"/time_series/cross" json:"time_series_cross_url"`
	QuotesURL          string `default:"/quote"             json:"quotes_url"`
	PriceURL           string `default:"/price"             json:"price_url"`
	EODURL             string `default:"/eod"               json:"eod_url"`

	ExchangeRateURL string `default:"/exchange_rate"          json:"exchange_rate_url"`
	MarketMoversURL string `default:"/market_movers/{market}" json:"market_movers_url"`
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
