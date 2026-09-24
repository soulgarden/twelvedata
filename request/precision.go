package request

// Precision returns a decimal precision for DP and DecimalPlaces parameters.
// A nil pointer omits the parameter; Precision(0) explicitly requests integers.
// Supported values and defaults depend on the endpoint; leave the pointer nil
// to preserve that endpoint's default.
func Precision(places int) *int {
	return &places
}
