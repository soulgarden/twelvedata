package response

import (
	"encoding/json"
	"strconv"

	"github.com/guregu/null/v6"
)

// FloatString unmarshals from number, string, or null JSON values into a nullable float.
type FloatString struct {
	null.Float
}

// NewFloatString creates a new FloatString with the specified validity.
func NewFloatString(f float64, valid bool) FloatString {
	return FloatString{Float: null.NewFloat(f, valid)}
}

// FloatStringFrom creates a new valid FloatString.
func FloatStringFrom(f float64) FloatString {
	return NewFloatString(f, true)
}

// FloatStringFromPtr creates a new FloatString that is null if f is nil.
func FloatStringFromPtr(f *float64) FloatString {
	return FloatString{Float: null.FloatFromPtr(f)}
}

// UnmarshalJSON accepts JSON numbers, strings, or null values.
func (f *FloatString) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		if s == "" || s == "null" {
			f.Valid = false
			f.Float64 = 0
			return nil
		}
		value, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		f.Float64 = value
		f.Valid = true
		return nil
	}

	return f.Float.UnmarshalJSON(data)
}
