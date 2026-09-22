package twelvedata

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/soulgarden/twelvedata/request"
)

func TestDocumentedQueryParameters(t *testing.T) {
	tests := []struct {
		endpoint string
		request  any
		params   map[string]string
	}{
		{"/stocks", new(request.GetStock), map[string]string{"outputsize": "50", "page": "2"}},
		{"/forex_pairs", new(request.GetForexPairs), map[string]string{"outputsize": "50", "page": "2"}},
		{"/cryptocurrencies", new(request.GetCryptocurrencies), map[string]string{"outputsize": "50", "page": "2"}},
		{"/etfs", new(request.GetETFs), map[string]string{"outputsize": "50", "page": "2"}},
		{"/funds", new(request.GetFunds), map[string]string{"mic_code": "XNAS"}},
		{"/commodities", new(request.GetCommodities), map[string]string{"outputsize": "50", "page": "2"}},
		{"/bonds", new(request.GetBonds), map[string]string{"mic_code": "XNAS"}},
		{"/press_releases", new(request.GetPressReleases), map[string]string{"page": "2"}},
		{"/etfs/list", new(request.GetETFsDirectory), map[string]string{"delimiter": ";", "dp": "0", "format": "JSON"}},
		{"/mutual_funds/list", new(request.GetMutualFundsDirectory), map[string]string{"delimiter": ";", "dp": "0", "format": "JSON"}},
		{"/recommendations", new(request.GetRecommendations), map[string]string{"mic_code": "XNAS"}},
		{"/price_target", new(request.GetPriceTarget), map[string]string{"mic_code": "XNAS"}},
		{"/earnings_estimate", new(request.GetEarningsEstimate), map[string]string{"mic_code": "XNAS"}},
		{"/revenue_estimate", new(request.GetRevenueEstimate), map[string]string{"mic_code": "XNAS"}},
		{"/eps_trend", new(request.GetEPSTrend), map[string]string{"mic_code": "XNAS"}},
		{"/eps_revisions", new(request.GetEPSRevisions), map[string]string{"mic_code": "XNAS"}},
		{"/growth_estimates", new(request.GetGrowthEstimates), map[string]string{"mic_code": "XNAS"}},
		{"/analyst_ratings/light", new(request.GetAnalystRatingsSnapshot), map[string]string{"mic_code": "XNAS"}},
		{"/analyst_ratings/us_equities", new(request.GetAnalystRatingsUSEquities), map[string]string{"mic_code": "XNAS"}},
	}
	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			for key, value := range tt.params {
				setQueryParameter(t, tt.request, key, value)
			}
			got, err := buildQueryParams(tt.request)
			if err != nil {
				t.Fatal(err)
			}
			for key, want := range tt.params {
				if got.Get(key) != want {
					t.Errorf("%s: want %q, got %q", key, want, got.Get(key))
				}
			}
		})
	}
}

func TestDecimalPrecisionQuery(t *testing.T) {
	requests := []any{
		new(request.GetAD),
		new(request.GetADX),
		new(request.GetATR),
		new(request.GetBBands),
		new(request.GetCCI),
		new(request.GetCurrencyConversion),
		new(request.GetDEMA),
		new(request.GetEarnings),
		new(request.GetEarningsCalendar),
		new(request.GetEMA),
		new(request.GetEOD),
		new(request.GetETFComposition),
		new(request.GetETFFullData),
		new(request.GetETFPerformance),
		new(request.GetETFRisk),
		new(request.GetETFSummary),
		new(request.GetExchangeRate),
		new(request.GetKAMA),
		new(request.GetMA),
		new(request.GetMACD),
		new(request.GetMarketMovers),
		new(request.GetMOM),
		new(request.GetMutualFundComposition),
		new(request.GetMutualFundFullData),
		new(request.GetMutualFundPerformance),
		new(request.GetMutualFundPurchaseInfo),
		new(request.GetMutualFundRatings),
		new(request.GetMutualFundRisk),
		new(request.GetMutualFundSummary),
		new(request.GetMutualFundSustainability),
		new(request.GetNATR),
		new(request.GetOBV),
		new(request.GetPercentB),
		new(request.GetPrice),
		new(request.GetQuote),
		new(request.GetRevenueEstimate),
		new(request.GetROC),
		new(request.GetRSI),
		new(request.GetSAR),
		new(request.GetSMA),
		new(request.GetStoch),
		new(request.GetTEMA),
		new(request.GetTimeSeries),
		new(request.GetTimeSeriesCross),
		new(request.GetTR),
		new(request.GetTRMA),
		new(request.GetVWAP),
		new(request.GetWillR),
		new(request.GetWMA),
		new(request.GetETFsDirectory),
		new(request.GetMutualFundsDirectory),
	}
	for _, req := range requests {
		t.Run(reflect.TypeOf(req).Elem().Name(), func(t *testing.T) {
			// A nil precision must leave the provider's default unchanged.
			initial, err := buildQueryParams(req)
			if err != nil {
				t.Fatal(err)
			}
			if initial.Has("dp") {
				t.Errorf("unset precision sent: %s", initial.Encode())
			}
			for _, value := range []string{"0", "-1", "2", "11"} {
				setQueryParameter(t, req, "dp", value)
				got, err := buildQueryParams(req)
				if err != nil {
					t.Fatal(err)
				}
				if got.Get("dp") != value {
					t.Errorf("precision %s: query = %s", value, got.Encode())
				}
			}
		})
	}
}

func setQueryParameter(t *testing.T, req any, key, value string) {
	t.Helper()
	object := reflect.ValueOf(req).Elem()
	for i := 0; i < object.NumField(); i++ {
		if strings.Split(object.Type().Field(i).Tag.Get("schema"), ",")[0] != key {
			continue
		}
		field := object.Field(i)
		if field.Kind() == reflect.Pointer {
			field.Set(reflect.New(field.Type().Elem()))
			field = field.Elem()
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int:
			number, err := strconv.Atoi(value)
			if err != nil {
				t.Fatal(err)
			}
			field.SetInt(int64(number))
		default:
			t.Fatalf("unsupported query type %s", field.Type())
		}
		return
	}
	t.Fatalf("%T has no %s parameter", req, key)
}
