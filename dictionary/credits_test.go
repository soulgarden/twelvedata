package dictionary

import "testing"

func TestRegulatoryCredits(t *testing.T) {
	t.Helper()

	if EDGARFillings != 50 {
		t.Fatalf("EDGARFillings = %d, want %d", EDGARFillings, 50)
	}

	if TaxInformation != 50 {
		t.Fatalf("TaxInformation = %d, want %d", TaxInformation, 50)
	}
}
