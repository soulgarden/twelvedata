package response

// OHLCV contains the price bar returned by indicators when include_ohlc is true.
// Volume may be absent for instruments without volume data.
type OHLCV struct {
	Open   string `json:"open,omitempty"`
	High   string `json:"high,omitempty"`
	Low    string `json:"low,omitempty"`
	Close  string `json:"close,omitempty"`
	Volume string `json:"volume,omitempty"`
}
