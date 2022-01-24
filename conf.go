package twelvedata

// nolint: tagliatelle
type Conf struct {
	BaseURL   string `json:"base_url" default:"https://api.twelvedata.com"`
	BaseWSURL string `json:"base_ws_url" default:"ws.twelvedata.com"`

	ReferenceData ReferenceData `json:"reference_data"`
	CoreData      CoreData      `json:"core_data"`
	Fundamentals  Fundamentals  `json:"fundamentals"`
	WebSocket     WebSocket     `json:"web_socket"`
	Advanced      Advanced      `json:"advanced"`

	APIKey  string `json:"api_key" default:"demo"`
	Timeout int    `json:"timeout" default:"15"`
}

// nolint: tagliatelle,lll
type ReferenceData struct {
	StocksURL    string `json:"stocks_url"   default:"/stocks?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}&type={type}"`
	ExchangesURL string `json:"exchange_url" default:"/exchanges?apikey={apikey}&type={type}&name={name}&code={code}&country={country}"`
	IndicesURL   string `json:"indices_url"  default:"/indices?apikey={apikey}&symbol={symbol}&country={country}"`
	EtfsURL      string `json:"etfs_url"     default:"/etf?apikey={apikey}&symbol={symbol}"`
}

// nolint: tagliatelle,lll
type CoreData struct {
	TimeSeriesURL   string `json:"time_series_url" default:"/time_series?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&country={country}&type={type}&outputsize={outputsize}&prepost={prepost}"`
	QuotesURL       string `json:"quotes_url"      default:"/quote?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&country={country}&volume_time_period={volume_time_period}&type={type}&prepost={prepost}"`
	ExchangeRateURL string `json:"exchange_rate_url" default:"/exchange_rate?apikey={apikey}&symbol={symbol}&precision={precision}&timezone={timezone}"`
}

// nolint: tagliatelle,lll
type Fundamentals struct {
	EarningsCalendarURL    string `json:"earnings_calendar_url"    default:"/earnings_calendar?apikey={apikey}&dp={dp}&start_date={start_date}&end_date={end_date}"`
	ProfileURL             string `json:"profile_url"              default:"/profile?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}"`
	InsiderTransactionsURL string `json:"insider_transactions_url" default:"/insider_transactions?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}"`
	IncomeStatementURL     string `json:"income_statement_url"     default:"/income_statement?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	BalanceSheetURL        string `json:"balance_sheet_url"        default:"/balance_sheet?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	CashFlowURL            string `json:"cash_flow_url"            default:"/cash_flow?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	DividendsURL           string `json:"dividends_url"            default:"/dividends?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}&range={range}&start_date={start_date}&end_date={end_date}"`
	StatisticsURL          string `json:"statistics_url"           default:"/statistics?apikey={apikey}&symbol={symbol}&exchange={exchange}&country={country}"`
}

// nolint: tagliatelle
type WebSocket struct {
	PriceURL string `json:"ws_price_url" default:"/v1/quotes/price"`
}

// nolint: tagliatelle
type Advanced struct {
	UsageURL string `json:"usage_url" default:"/api_usage?apikey={apikey}"`
}
