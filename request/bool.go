package request

// Bool returns a pointer for optional boolean request parameters.
// A nil pointer omits the parameter; Bool(false) explicitly sends false.
func Bool(value bool) *bool {
	return &value
}
