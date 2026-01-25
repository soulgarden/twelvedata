package response

import (
	"encoding/json"
	"testing"
)

func BenchmarkFloatStringUnmarshal(b *testing.B) {
	cases := []struct {
		name    string
		payload []byte
	}{
		{name: "number", payload: []byte("123.45")},
		{name: "string", payload: []byte(`"123.45"`)},
		{name: "string-null", payload: []byte(`"null"`)},
		{name: "null", payload: []byte("null")},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.SetBytes(int64(len(tc.payload)))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var got FloatString
				if err := json.Unmarshal(tc.payload, &got); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
