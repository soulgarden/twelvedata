package response

import (
	"encoding/json"

	"github.com/guregu/null/v6"
)

// UnmarshalJSON accepts both the provider's spelling and the legacy key.
func (a *ETFAssetAllocation) UnmarshalJSON(data []byte) error {
	type allocation ETFAssetAllocation
	return unmarshalAssetAllocation(data, (*allocation)(a), &a.Convertibles)
}

// UnmarshalJSON accepts both the provider's spelling and the legacy key.
func (a *MutualFundAssetAllocation) UnmarshalJSON(data []byte) error {
	type allocation MutualFundAssetAllocation
	return unmarshalAssetAllocation(data, (*allocation)(a), &a.Convertibles)
}

func unmarshalAssetAllocation(data []byte, target any, convertibles *null.Float) error {
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}

	var keys struct {
		Documented json.RawMessage `json:"convertables"` //nolint:misspell // Matches the provider's wire key.
		Legacy     json.RawMessage `json:"convertibles"`
	}
	if err := json.Unmarshal(data, &keys); err != nil {
		return err
	}
	// Presence, including explicit null or zero, gives the documented key priority.
	if len(keys.Documented) == 0 && len(keys.Legacy) != 0 {
		return json.Unmarshal(keys.Legacy, convertibles)
	}
	return nil
}
