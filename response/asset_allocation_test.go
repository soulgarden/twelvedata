package response_test

import (
	"encoding/json"
	"testing"

	"github.com/soulgarden/twelvedata/response"
)

//nolint:misspell // Fixtures use the provider's exact wire spelling.
func TestAssetAllocationConvertibles(t *testing.T) {
	tests := []struct {
		name, data string
		want       float64
		valid      bool
	}{
		{"documented spelling", `{"convertables":1.5,"cash":2}`, 1.5, true},
		{"legacy spelling", `{"convertibles":2.5,"cash":2}`, 2.5, true},
		{"documented spelling wins", `{"convertables":0,"convertibles":2.5,"cash":2}`, 0, true},
		{"explicit null wins", `{"convertables":null,"convertibles":2.5,"cash":2}`, 0, false},
		{"absent", `{"cash":2}`, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var etf response.ETFAssetAllocation
			var mf response.MutualFundAssetAllocation
			for _, target := range []any{&etf, &mf} {
				if err := json.Unmarshal([]byte(tt.data), target); err != nil {
					t.Fatal(err)
				}
			}
			if etf.Convertibles.Valid != tt.valid || etf.Convertibles.Float64 != tt.want || etf.Cash.Float64 != 2 {
				t.Errorf("ETF allocation = %+v", etf)
			}
			if mf.Convertibles.Valid != tt.valid || mf.Convertibles.Float64 != tt.want || mf.Cash.Float64 != 2 {
				t.Errorf("mutual fund allocation = %+v", mf)
			}
		})
	}
}
