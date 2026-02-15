package dictionary

import "testing"

func TestRegulatoryCredits(t *testing.T) {
	t.Helper()

	if EDGARFilings != 50 {
		t.Fatalf("EDGARFilings = %d, want %d", EDGARFilings, 50)
	}

	if TaxInformation != 50 {
		t.Fatalf("TaxInformation = %d, want %d", TaxInformation, 50)
	}
}
