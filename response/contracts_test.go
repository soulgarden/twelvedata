package response_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/soulgarden/twelvedata/response"
)

// Published examples must decode without silently losing documented data.
func TestDocumentedResponseExamples(t *testing.T) {
	tests := []struct {
		endpoint, fixture string
		target            any
	}{
		{"/time_series", "time_series.json", new(response.TimeSeries)},
		{"/time_series/cross", "time_series_cross.json", new(response.TimeSeriesCross)},
		{"/quote", "quote.json", new(response.Quote)},
		{"/price", "price.json", new(response.Price)},
		{"/eod", "eod.json", new(response.EOD)},
		{"/market_movers/{market}", "market_movers_market.json", new(response.MarketMovers)},
		{"/stocks", "stocks.json", new(response.Stocks)},
		{"/forex_pairs", "forex_pairs.json", new(response.ForexPairs)},
		{"/cryptocurrencies", "cryptocurrencies.json", new(response.Cryptocurrencies)},
		{"/etfs", "etfs.json", new(response.ETFs)},
		{"/funds", "funds.json", new(response.Funds)},
		{"/commodities", "commodities.json", new(response.Commodities)},
		{"/bonds", "bonds.json", new(response.Bonds)},
		{"/symbol_search", "symbol_search.json", new(response.SymbolSearch)},
		{"/cross_listings", "cross_listings.json", new(response.CrossListings)},
		{"/earliest_timestamp", "earliest_timestamp.json", new(response.EarliestTimestamp)},
		{"/exchanges", "exchanges.json", new(response.Exchanges)},
		{"/exchange_schedule", "exchange_schedule.json", new(response.ExchangeSchedule)},
		{"/cryptocurrency_exchanges", "cryptocurrency_exchanges.json", new(response.CryptocurrencyExchanges)},
		{"/market_state", "market_state.json", new([]response.MarketState)},
		{"/countries", "countries.json", new(response.Countries)},
		{"/instrument_type", "instrument_type.json", new(response.InstrumentType)},
		{"/technical_indicators", "technical_indicators.json", new(response.TechnicalIndicators)},
		{"/logo", "logo.json", new(response.Logo)},
		{"/profile", "profile.json", new(response.Profile)},
		{"/key_executives", "key_executives.json", new(response.KeyExecutives)},
		{"/dividends", "dividends.json", new(response.Dividends)},
		{"/dividends_calendar", "dividends_calendar.json", new(response.DividendsCalendar)},
		{"/earnings", "earnings.json", new(response.Earnings)},
		{"/splits", "splits.json", new(response.Splits)},
		{"/splits_calendar", "splits_calendar.json", new(response.SplitsCalendar)},
		{"/statistics", "statistics.json", new(response.Statistics)},
		{"/earnings_calendar", "earnings_calendar.json", new(response.EarningsCalendar)},
		{"/ipo_calendar", "ipo_calendar.json", new(response.IPOCalendar)},
		{"/press_releases", "press_releases.json", new(response.PressReleases)},
		{"/income_statement", "income_statement.json", new(response.IncomeStatements)},
		{"/income_statement/consolidated", "income_statement_consolidated.json", new(response.ConsolidatedIncomeStatements)},
		{"/balance_sheet", "balance_sheet.json", new(response.BalanceSheets)},
		{"/balance_sheet/consolidated", "balance_sheet_consolidated.json", new(response.ConsolidatedBalanceSheets)},
		{"/cash_flow", "cash_flow.json", new(response.CashFlows)},
		{"/cash_flow/consolidated", "cash_flow_consolidated.json", new(response.ConsolidatedCashFlows)},
		{"/market_cap", "market_cap.json", new(response.MarketCap)},
		{"/last_change/{endpoint}", "last_change_endpoint.json", new(response.LastChange)},
		{"/exchange_rate", "exchange_rate.json", new(response.ExchangeRate)},
		{"/currency_conversion", "currency_conversion.json", new(response.CurrencyConversion)},
		{"/etfs/list", "etfs_list.json", new(response.ETFsDirectory)},
		{"/etfs/world", "etfs_world.json", new(response.ETFFullData)},
		{"/etfs/world/summary", "etfs_world_summary.json", new(response.ETFWorldSummary)},
		{"/etfs/world/performance", "etfs_world_performance.json", new(response.ETFPerformance)},
		{"/etfs/world/risk", "etfs_world_risk.json", new(response.ETFRisk)},
		{"/etfs/world/composition", "etfs_world_composition.json", new(response.ETFComposition)},
		{"/etfs/family", "etfs_family.json", new(response.ETFFamilies)},
		{"/etfs/type", "etfs_type.json", new(response.ETFTypes)},
		{"/mutual_funds/list", "mutual_funds_list.json", new(response.MutualFundsDirectory)},
		{"/mutual_funds/world", "mutual_funds_world.json", new(response.MutualFundFullData)},
		{"/mutual_funds/world/summary", "mutual_funds_world_summary.json", new(response.MutualFundSummary)},
		{"/mutual_funds/world/performance", "mutual_funds_world_performance.json", new(response.MutualFundPerformance)},
		{"/mutual_funds/world/risk", "mutual_funds_world_risk.json", new(response.MutualFundRisk)},
		{"/mutual_funds/world/ratings", "mutual_funds_world_ratings.json", new(response.MutualFundRatings)},
		{"/mutual_funds/world/composition", "mutual_funds_world_composition.json", new(response.MutualFundComposition)},
		{"/mutual_funds/world/purchase_info", "mutual_funds_world_purchase_info.json", new(response.MutualFundPurchaseInfo)},
		{"/mutual_funds/world/sustainability", "mutual_funds_world_sustainability.json", new(response.MutualFundSustainability)},
		{"/mutual_funds/family", "mutual_funds_family.json", new(response.MutualFundFamilies)},
		{"/mutual_funds/type", "mutual_funds_type.json", new(response.MutualFundTypes)},
		{"/bbands", "bbands.json", new(response.BBands)},
		{"/sma", "sma.json", new(response.SMA)},
		{"/ema", "ema.json", new(response.EMA)},
		{"/adx", "adx.json", new(response.ADX)},
		{"/macd", "macd.json", new(response.MACD)},
		{"/rsi", "rsi.json", new(response.RSI)},
		{"/stoch", "stoch.json", new(response.Stoch)},
		{"/percent_b", "percent_b.json", new(response.PercentB)},
		{"/atr", "atr.json", new(response.ATR)},
		{"/vwap", "vwap.json", new(response.VWAP)},
		{"/ma", "ma.json", new(response.MA)},
		{"/wma", "wma.json", new(response.WMA)},
		{"/dema", "dema.json", new(response.DEMA)},
		{"/tema", "tema.json", new(response.TEMA)},
		{"/trima", "trima.json", new(response.TRMA)},
		{"/kama", "kama.json", new(response.KAMA)},
		{"/sar", "sar.json", new(response.SAR)},
		{"/cci", "cci.json", new(response.CCI)},
		{"/willr", "willr.json", new(response.WillR)},
		{"/roc", "roc.json", new(response.ROC)},
		{"/mom", "mom.json", new(response.MOM)},
		{"/obv", "obv.json", new(response.OBV)},
		{"/ad", "ad.json", new(response.AD)},
		{"/natr", "natr.json", new(response.NATR)},
		{"/trange", "trange.json", new(response.TR)},
		{"/recommendations", "recommendations.json", new(response.Recommendations)},
		{"/price_target", "price_target.json", new(response.PriceTarget)},
		{"/earnings_estimate", "earnings_estimate.json", new(response.EarningsEstimate)},
		{"/revenue_estimate", "revenue_estimate.json", new(response.RevenueEstimate)},
		{"/eps_trend", "eps_trend.json", new(response.EPSTrend)},
		{"/eps_revisions", "eps_revisions.json", new(response.EPSRevisions)},
		{"/growth_estimates", "growth_estimates.json", new(response.GrowthEstimates)},
		{"/analyst_ratings/light", "analyst_ratings_light.json", new(response.AnalystRatingsSnapshot)},
		{"/analyst_ratings/us_equities", "analyst_ratings_us_equities.json", new(response.AnalystRatingsUSEquities)},
		{"/insider_transactions", "insider_transactions.json", new(response.InsiderTransactions)},
		{"/edgar_filings/archive", "edgar_filings_archive.json", new(response.EDGARFilings)},
		{"/institutional_holders", "institutional_holders.json", new(response.InstitutionalHolders)},
		{"/fund_holders", "fund_holders.json", new(response.FundHolders)},
		{"/direct_holders", "direct_holders.json", new(response.DirectHolders)},
		{"/tax_info", "tax_info.json", new(response.TaxInformation)},
		{"/sanctions/{source}", "sanctions_source.json", new(response.SanctionedEntities)},
		{"/api_usage", "api_usage.json", new(response.Usage)},
		{"/batch", "batch.json", new(response.Batches)},
	}
	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", "contracts", tt.fixture))
			if err != nil {
				t.Fatal(err)
			}
			assertContractRoundTrip(t, data, tt.target)
		})
	}
}

func assertContractRoundTrip(t *testing.T, data []byte, target any) {
	t.Helper()
	if err := json.Unmarshal(data, target); err != nil {
		t.Fatalf("decode documented response: %v", err)
	}
	encoded, err := json.Marshal(target)
	if err != nil {
		t.Fatal(err)
	}
	var want, got any
	for _, item := range []struct {
		data   []byte
		target *any
	}{{data, &want}, {encoded, &got}} {
		decoder := json.NewDecoder(bytes.NewReader(item.data))
		decoder.UseNumber()
		if err := decoder.Decode(item.target); err != nil {
			t.Fatal(err)
		}
	}
	assertDocumentedFields(t, "$", want, got)
}

func assertDocumentedFields(t *testing.T, path string, want, got any) {
	t.Helper()
	switch value := want.(type) {
	case map[string]any:
		object, ok := got.(map[string]any)
		if !ok {
			t.Errorf("%s: want object, got %T", path, got)
			return
		}
		for key, field := range value {
			actual, exists := object[key]
			if !exists {
				t.Errorf("%s.%s: documented field lost", path, key)
				continue
			}
			assertDocumentedFields(t, path+"."+key, field, actual)
		}
	case []any:
		array, ok := got.([]any)
		if !ok || len(array) != len(value) {
			t.Errorf("%s: want %d items, got %v", path, len(value), got)
			return
		}
		for i, field := range value {
			assertDocumentedFields(t, fmt.Sprintf("%s[%d]", path, i), field, array[i])
		}
	default:
		if !equalContractScalar(want, got) {
			t.Errorf("%s: want %v (%T), got %v (%T)", path, want, want, got, got)
		}
	}
}

// Nullable numeric types accept numbers encoded as either JSON numbers or strings.
func equalContractScalar(want, got any) bool {
	if reflect.DeepEqual(want, got) {
		return true
	}
	_, wantNumber := want.(json.Number)
	_, gotNumber := got.(json.Number)
	if !wantNumber && !gotNumber {
		return false
	}
	a, aOK := new(big.Rat).SetString(fmt.Sprint(want))
	b, bOK := new(big.Rat).SetString(fmt.Sprint(got))
	return aOK && bOK && a.Cmp(b) == 0
}
