package twelvedata

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

// nolint: lll
type ReferenceData struct {
	StocksURL      string `json:"stocks_url"   default:"/stocks?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&type={type}&show_plan={show_plan}&include_delisted={include_delisted}"`
	ExchangesURL   string `json:"exchange_url" default:"/exchanges?apikey={apikey}&type={type}&name={name}&code={code}&country={country}&show_plan={show_plan}"`
	IndicesURL     string `json:"indices_url"  default:"/indices?apikey={apikey}&symbol={symbol}&country={country}&show_plan={show_plan}&include_delisted={include_delisted}"`
	EtfsURL        string `json:"etfs_url"     default:"/etf?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&show_plan={show_plan}&include_delisted={include_delisted}"`
	MarketStateURL string `json:"market_state_url" default:"/market_state?apikey={apikey}&exchange={exchange}&code={code}&country={country}"`
}

// nolint: lll
type CoreData struct {
	TimeSeriesURL string `json:"time_series_url" default:"/time_series?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&mic_code={mic_code}&country={country}&type={type}&outputsize={outputsize}&prepost={prepost}&dp={dp}&order={order}&timezone={timezone}&date={date}&start_date={start_date}&end_date={end_date}&previous_close={previous_close}"`
	QuotesURL     string `json:"quotes_url"      default:"/quote?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&mic_code={mic_code}&country={country}&volume_time_period={volume_time_period}&type={type}&prepost={prepost}&eod={eod}&rolling_period={rolling_period}&dp={dp}&timezone={timezone}"`

	ExchangeRateURL string `json:"exchange_rate_url" default:"/exchange_rate?apikey={apikey}&symbol={symbol}&date={date}&dp={dp}&timezone={timezone}"`
	MarketMoversURL string `json:"market_movers_url" default:"/market_movers/{instrument}?apikey={apikey}&direction={direction}&outputsize={outputsize}&country={country}&dp={dp}"`
}

// nolint: lll
type Fundamentals struct {
	EarningsCalendarURL    string `json:"earnings_calendar_url"    default:"/earnings_calendar?apikey={apikey}&exchange={exchange}&mic_code={mic_code}&country={country}&dp={dp}&start_date={start_date}&end_date={end_date}"`
	ProfileURL             string `json:"profile_url"              default:"/profile?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code;}&country={country}"`
	InsiderTransactionsURL string `json:"insider_transactions_url" default:"/insider_transactions?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}"`
	IncomeStatementURL     string `json:"income_statement_url"     default:"/income_statement?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	BalanceSheetURL        string `json:"balance_sheet_url"        default:"/balance_sheet?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	CashFlowURL            string `json:"cash_flow_url"            default:"/cash_flow?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"`
	DividendsURL           string `json:"dividends_url"            default:"/dividends?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&range={range}&start_date={start_date}&end_date={end_date}"`
	StatisticsURL          string `json:"statistics_url"           default:"/statistics?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}"`
}

type WebSocket struct {
	PriceURL string `json:"ws_price_url" default:"/v1/quotes/price"`
}

type Advanced struct {
	UsageURL string `json:"usage_url" default:"/api_usage?apikey={apikey}"`
}
