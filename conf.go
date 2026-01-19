package twelvedata

// Conf holds the configuration settings for the Twelve Data API client,
// including base URLs, API endpoints, authentication, and timeout settings.
type Conf struct {
	BaseURL   string `default:"https://api.twelvedata.com" json:"base_url"`
	BaseWSURL string `default:"ws.twelvedata.com"          json:"base_ws_url"`

	ReferenceData       ReferenceData       `json:"reference_data"`
	CoreData            CoreData            `json:"core_data"`
	Fundamentals        Fundamentals        `json:"fundamentals"`
	TechnicalIndicators TechnicalIndicators `json:"technical_indicators"`
	Currencies          Currencies          `json:"currencies"`
	WebSocket           WebSocket           `json:"web_socket"`
	ETFs                ETFs                `json:"etfs"`
	MutualFunds         MutualFunds         `json:"mutual_funds"`
	Analysis            Analysis            `json:"analysis"`
	Regulatory          Regulatory          `json:"regulatory"`
	Advanced            Advanced            `json:"advanced"`

	APIKey  string `default:"demo" json:"api_key"`
	Timeout int    `default:"15"   json:"timeout"`
}

// ReferenceData contains URL configurations for reference data endpoints including
// asset catalogs, discovery tools, market information, and supporting metadata.
// nolint: lll
type ReferenceData struct {
	// Asset Catalogs
	StocksURL           string `default:"/stocks"           json:"stocks_url"`
	ForexPairsURL       string `default:"/forex_pairs"      json:"forex_pairs_url"`
	CryptocurrenciesURL string `default:"/cryptocurrencies" json:"cryptocurrencies_url"`
	ETFsURL             string `default:"/etfs"             json:"etfs_url"`
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
	MarketStateURL             string `default:"/market_state"             json:"market_state_url"`

	// Supporting Metadata
	CountriesURL           string `default:"/countries"            json:"countries_url"`
	InstrumentTypeURL      string `default:"/instrument_type"      json:"instrument_type_url"`
	TechnicalIndicatorsURL string `default:"/technical_indicators" json:"technical_indicators_url"`
}

// CoreData contains URL configurations for core market data endpoints including
// time series data, quotes, prices, and end-of-day data.
// nolint: lll
type CoreData struct {
	TimeSeriesURL      string `default:"/time_series"       json:"time_series_url"`
	TimeSeriesCrossURL string `default:"/time_series/cross" json:"time_series_cross_url"`
	QuotesURL          string `default:"/quote"             json:"quotes_url"`
	PriceURL           string `default:"/price"             json:"price_url"`
	EODURL             string `default:"/eod"               json:"eod_url"`

	MarketMoversURL string `default:"/market_movers/{market}" json:"market_movers_url"`
}

// Fundamentals contains URL configurations for fundamental data endpoints including
// company profiles, financial statements, dividends, splits, and other corporate data.
// nolint: lll
type Fundamentals struct {
	LogoURL              string `default:"/logo"                    json:"logo_url"`
	EarningsCalendarURL  string `default:"/earnings_calendar"       json:"earnings_calendar_url"`
	IPOCalendarURL       string `default:"/ipo_calendar"            json:"ipo_calendar_url"`
	ProfileURL           string `default:"/profile"                 json:"profile_url"`
	KeyExecutivesURL     string `default:"/key_executives"          json:"key_executives_url"`
	IncomeStatementURL   string `default:"/income_statement"        json:"income_statement_url"`
	BalanceSheetURL      string `default:"/balance_sheet"           json:"balance_sheet_url"`
	CashFlowURL          string `default:"/cash_flow"               json:"cash_flow_url"`
	DividendsURL         string `default:"/dividends"               json:"dividends_url"`
	DividendsCalendarURL string `default:"/dividends_calendar"      json:"dividends_calendar_url"`
	EarningsURL          string `default:"/earnings"                json:"earnings_url"`
	SplitsURL            string `default:"/splits"                  json:"splits_url"`
	SplitsCalendarURL    string `default:"/splits_calendar"         json:"splits_calendar_url"`
	StatisticsURL        string `default:"/statistics"              json:"statistics_url"`
	MarketCapURL         string `default:"/market_cap"              json:"market_cap_url"`
	LastChangeURL        string `default:"/last_change/{endpoint}"  json:"last_change_url"`
}

// WebSocket contains URL configurations for WebSocket endpoints used for real-time data streaming.
type WebSocket struct {
	PriceURL string `default:"/v1/quotes/price" json:"ws_price_url"`
}

// Currencies contains URL configurations for currency-related endpoints including exchange rates and conversions.
type Currencies struct {
	ExchangeRateURL       string `default:"/exchange_rate"        json:"exchange_rate_url"`
	CurrencyConversionURL string `default:"/currency_conversion"  json:"currency_conversion_url"`
}

// TechnicalIndicators contains URL configurations for technical indicator endpoints including
// overlapping studies, momentum indicators, volume indicators, and volatility indicators.
// nolint: lll
type TechnicalIndicators struct {
	// Overlap Studies
	BbandsURL string `default:"/bbands" json:"bbands_url"`
	SMAURL    string `default:"/sma" json:"sma_url"`
	EMAURL    string `default:"/ema" json:"ema_url"`
	MAURL     string `default:"/ma" json:"ma_url"`
	WMAURL    string `default:"/wma" json:"wma_url"`
	VWAPURL   string `default:"/vwap" json:"vwap_url"`
	DEMAURL   string `default:"/dema" json:"dema_url"`
	TEMAURL   string `default:"/tema" json:"tema_url"`
	TRMAURL   string `default:"/trima" json:"trima_url"`
	KAMAURL   string `default:"/kama" json:"kama_url"`
	SARURL    string `default:"/sar" json:"sar_url"`

	// Momentum Indicators
	ADXURL       string `default:"/adx" json:"adx_url"`
	MACDURL      string `default:"/macd" json:"macd_url"`
	RSIURL       string `default:"/rsi" json:"rsi_url"`
	StochURL     string `default:"/stoch" json:"stoch_url"`
	PercentBURL  string `default:"/percent_b" json:"percent_b_url"`
	CCIURL       string `default:"/cci" json:"cci_url"`
	WilliamsRURL string `default:"/willr" json:"williams_r_url"`
	ROCURL       string `default:"/roc" json:"roc_url"`
	MomURL       string `default:"/mom" json:"mom_url"`

	// Volume Indicators
	OBVURL string `default:"/obv" json:"obv_url"`
	ADURL  string `default:"/ad" json:"ad_url"`

	// Volatility Indicators
	ATRURL  string `default:"/atr" json:"atr_url"`
	NATRURL string `default:"/natr" json:"natr_url"`
	TRURL   string `default:"/trange" json:"trange_url"`
}

// Advanced contains URL configurations for advanced API endpoints such as usage tracking and batch operations.
type Advanced struct {
	UsageURL   string `default:"/api_usage" json:"usage_url"`
	BatchesURL string `default:"/batch"     json:"batches_url"`
}

// ETFs contains URL configurations for ETF-related endpoints including directory, full data, and performance metrics.
// nolint: lll
type ETFs struct {
	ETFsDirectoryURL   string `default:"/etfs/list"              json:"etfs_directory_url"`
	ETFsFullDataURL    string `default:"/etfs/world"             json:"etfs_full_data_url"`
	ETFsSummaryURL     string `default:"/etfs/world/summary"     json:"etfs_summary_url"`
	ETFsPerformanceURL string `default:"/etfs/world/performance" json:"etfs_performance_url"`
	ETFsRiskURL        string `default:"/etfs/world/risk"        json:"etfs_risk_url"`
	ETFsCompositionURL string `default:"/etfs/world/composition" json:"etfs_composition_url"`
	ETFsFamiliesURL    string `default:"/etfs/family"            json:"etfs_families_url"`
	ETFsTypesURL       string `default:"/etfs/type"              json:"etfs_types_url"`
}

// MutualFunds contains URL configurations for mutual fund-related endpoints including directory, performance, and ratings.
// nolint: lll
type MutualFunds struct {
	MutualFundsDirectoryURL      string `default:"/funds"                           json:"mutual_funds_directory_url"`
	MutualFundsFullDataURL       string `default:"/mutual_funds/world"              json:"mutual_funds_full_data_url"`
	MutualFundsSummaryURL        string `default:"/mutual_funds/world/summary"      json:"mutual_funds_summary_url"`
	MutualFundsPerformanceURL    string `default:"/mutual_funds/world/performance"  json:"mutual_funds_performance_url"`
	MutualFundsRiskURL           string `default:"/mutual_funds/world/risk"         json:"mutual_funds_risk_url"`
	MutualFundsRatingsURL        string `default:"/mutual_funds/world/ratings"      json:"mutual_funds_ratings_url"`
	MutualFundsCompositionURL    string `default:"/mutual_funds/world/composition"  json:"mutual_funds_composition_url"`
	MutualFundsPurchaseInfoURL   string `default:"/mutual_funds/world/purchase_info" json:"mutual_funds_purchase_info_url"`
	MutualFundsSustainabilityURL string `default:"/mutual_funds/world/sustainability" json:"mutual_funds_sustainability_url"`
	MutualFundsFamiliesURL       string `default:"/mutual_funds/world/families"     json:"mutual_funds_families_url"`
	MutualFundsTypesURL          string `default:"/mutual_funds/world/types"        json:"mutual_funds_types_url"`
}

// Analysis contains URL configurations for analysis endpoints including earnings estimates, recommendations, and price targets.
// nolint: lll
type Analysis struct {
	EarningsEstimateURL         string `default:"/earnings_estimate"            json:"earnings_estimate_url"`
	RevenueEstimateURL          string `default:"/revenue_estimate"             json:"revenue_estimate_url"`
	EPSTrendURL                 string `default:"/eps_trend"                    json:"eps_trend_url"`
	EPSRevisionsURL             string `default:"/eps_revisions"                json:"eps_revisions_url"`
	GrowthEstimatesURL          string `default:"/growth_estimates"             json:"growth_estimates_url"`
	RecommendationsURL          string `default:"/recommendations"              json:"recommendations_url"`
	PriceTargetURL              string `default:"/price_target"                 json:"price_target_url"`
	AnalystRatingsSnapshotURL   string `default:"/analyst_ratings_snapshot"     json:"analyst_ratings_snapshot_url"`
	AnalystRatingsUSEquitiesURL string `default:"/analyst_ratings_us_equities"  json:"analyst_ratings_us_equities_url"`
}

// Regulatory contains URL configurations for regulatory and compliance endpoints including filings and ownership data.
// nolint: lll
type Regulatory struct {
	EDGARFillingsURL        string `default:"/edgar_filings/archive"        json:"edgar_fillings_url"`
	InsiderTransactionsURL  string `default:"/insider_transactions"  json:"insider_transactions_url"`
	InstitutionalHoldersURL string `default:"/institutional_holders" json:"institutional_holders_url"`
	FundHoldersURL          string `default:"/fund_holders"          json:"fund_holders_url"`
	DirectHoldersURL        string `default:"/direct_holders"        json:"direct_holders_url"`
	TaxInformationURL       string `default:"/tax_information"       json:"tax_information_url"`
	SanctionedEntitiesURL   string `default:"/sanctioned_entities"   json:"sanctioned_entities_url"`
}
