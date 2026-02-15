package response

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkTimeSeriesUnmarshal(b *testing.B) {
	benchmarkJSONUnmarshal[TimeSeries](b, unmarshalBenchCases(buildTimeSeriesJSON))
}

func BenchmarkTimeSeriesCrossUnmarshal(b *testing.B) {
	benchmarkJSONUnmarshal[TimeSeriesCross](b, unmarshalBenchCases(buildTimeSeriesCrossJSON))
}

type unmarshalBenchCase struct {
	name        string
	valueCount  int
	payloadFunc func(int) []byte
}

func unmarshalBenchCases(payloadFunc func(int) []byte) []unmarshalBenchCase {
	return []unmarshalBenchCase{
		{name: "values=1", valueCount: 1, payloadFunc: payloadFunc},
		{name: "values=100", valueCount: 100, payloadFunc: payloadFunc},
		{name: "values=1000", valueCount: 1000, payloadFunc: payloadFunc},
	}
}

func benchmarkJSONUnmarshal[T any](b *testing.B, cases []unmarshalBenchCase) {
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			payload := tc.payloadFunc(tc.valueCount)
			b.SetBytes(int64(len(payload)))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var got T
				if err := json.Unmarshal(payload, &got); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func buildTimeSeriesJSON(values int) []byte {
	if values < 1 {
		values = 1
	}

	var b strings.Builder
	b.Grow(256 + values*140)
	b.WriteString(`{"meta":{"symbol":"AAPL","interval":"1min","currency":"USD","exchange_timezone":"America/New_York","exchange":"NASDAQ","mic_code":"XNAS","type":"Common Stock"},"values":[`)

	for i := 0; i < values; i++ {
		if i > 0 {
			b.WriteByte(',')
		}

		fmt.Fprintf(
			&b,
			`{"datetime":"2021-09-16 15:%02d:00","open":"148.%05d","high":"149.%05d","low":"147.%05d","close":"148.%05d","volume":"%d"}`,
			i%60,
			73500+i,
			86000+i,
			73000+i,
			85001+i,
			624277+i,
		)
	}

	b.WriteString(`],"status":"ok"}`)

	return []byte(b.String())
}

func buildTimeSeriesCrossJSON(values int) []byte {
	if values < 1 {
		values = 1
	}

	var b strings.Builder
	b.Grow(256 + values*110)
	b.WriteString(`{"meta":{"base_instrument":"JPY/USD","base_currency":"","base_exchange":"PHYSICAL CURRENCY","interval":"1day","quote_instrument":"BTC/USD","quote_currency":"","quote_exchange":"Coinbase Pro"},"values":[`)

	for i := 0; i < values; i++ {
		if i > 0 {
			b.WriteByte(',')
		}

		fmt.Fprintf(
			&b,
			`{"datetime":"2025-02-28 14:%02d:00","open":"0.0000081115%02d","high":"0.0000081273%02d","low":"0.0000081088%02d","close":"0.0000081268%02d"}`,
			i%60,
			i%100,
			i%100,
			i%100,
			i%100,
		)
	}

	b.WriteString(`]}`)

	return []byte(b.String())
}
