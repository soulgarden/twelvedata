package response

import (
	"encoding/json"
	"testing"
)

func TestFloatStringUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		valid   bool
		wantErr bool
	}{
		{
			name:  "number",
			input: `12.34`,
			want:  12.34,
			valid: true,
		},
		{
			name:  "string number",
			input: `"12.34"`,
			want:  12.34,
			valid: true,
		},
		{
			name:  "null",
			input: `null`,
			valid: false,
		},
		{
			name:  "string null",
			input: `"null"`,
			valid: false,
		},
		{
			name:  "empty string",
			input: `""`,
			valid: false,
		},
		{
			name:    "invalid string",
			input:   `"not-a-number"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got FloatString
			err := json.Unmarshal([]byte(tt.input), &got)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got.Valid != tt.valid {
				t.Fatalf("UnmarshalJSON() Valid = %v, want %v", got.Valid, tt.valid)
			}
			if tt.valid && got.Float64 != tt.want {
				t.Fatalf("UnmarshalJSON() Float64 = %v, want %v", got.Float64, tt.want)
			}
		})
	}
}
