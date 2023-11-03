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
	StocksURL      string `default:"/stocks?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&type={type}&show_plan={show_plan}&include_delisted={include_delisted}" json:"stocks_url"`
	ExchangesURL   string `default:"/exchanges?apikey={apikey}&type={type}&name={name}&code={code}&country={country}&show_plan={show_plan}"                                                                  json:"exchange_url"`
	IndicesURL     string `default:"/indices?apikey={apikey}&symbol={symbol}&country={country}&show_plan={show_plan}&include_delisted={include_delisted}"                                                    json:"indices_url"`
	EtfsURL        string `default:"/etf?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&show_plan={show_plan}&include_delisted={include_delisted}"                json:"etfs_url"`
	MarketStateURL string `default:"/market_state?apikey={apikey}&exchange={exchange}&code={code}&country={country}"                                                                                         json:"market_state_url"`
}

// nolint: lll
type CoreData struct {
	TimeSeriesURL string `default:"/time_series?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&mic_code={mic_code}&country={country}&type={type}&outputsize={outputsize}&prepost={prepost}&dp={dp}&order={order}&timezone={timezone}&date={date}&start_date={start_date}&end_date={end_date}&previous_close={previous_close}" json:"time_series_url"`
	QuotesURL     string `default:"/quote?apikey={apikey}&symbol={symbol}&interval={interval}&exchange={exchange}&mic_code={mic_code}&country={country}&volume_time_period={volume_time_period}&type={type}&prepost={prepost}&eod={eod}&rolling_period={rolling_period}&dp={dp}&timezone={timezone}"                                                   json:"quotes_url"`

	ExchangeRateURL string `default:"/exchange_rate?apikey={apikey}&symbol={symbol}&date={date}&dp={dp}&timezone={timezone}"                              json:"exchange_rate_url"`
	MarketMoversURL string `default:"/market_movers/{instrument}?apikey={apikey}&direction={direction}&outputsize={outputsize}&country={country}&dp={dp}" json:"market_movers_url"`
}

// nolint: lll
type Fundamentals struct {
	EarningsCalendarURL    string `default:"/earnings_calendar?apikey={apikey}&exchange={exchange}&mic_code={mic_code}&country={country}&dp={dp}&start_date={start_date}&end_date={end_date}"                        json:"earnings_calendar_url"`
	ProfileURL             string `default:"/profile?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}"                                                                      json:"profile_url"`
	InsiderTransactionsURL string `default:"/insider_transactions?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}"                                                         json:"insider_transactions_url"`
	IncomeStatementURL     string `default:"/income_statement?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}" json:"income_statement_url"`
	BalanceSheetURL        string `default:"/balance_sheet?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"    json:"balance_sheet_url"`
	CashFlowURL            string `default:"/cash_flow?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&period={period}&start_date={start_date}&end_date={end_date}"        json:"cash_flow_url"`
	DividendsURL           string `default:"/dividends?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}&range={range}&start_date={start_date}&end_date={end_date}"          json:"dividends_url"`
	StatisticsURL          string `default:"/statistics?apikey={apikey}&symbol={symbol}&exchange={exchange}&mic_code={mic_code}&country={country}"                                                                   json:"statistics_url"`
}

type WebSocket struct {
	PriceURL string `default:"/v1/quotes/price" json:"ws_price_url"`
}

type Advanced struct {
	UsageURL string `default:"/api_usage?apikey={apikey}" json:"usage_url"`
}
