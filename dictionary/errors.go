package dictionary

import "errors"

const SymbolNotFoundMsg = "**symbol** not found:"
const NewSymbolNotFoundMsg = "**symbol** with specified criteria not found:"

var (
	ErrBadStatusCode             = errors.New("bas status code")
	ErrInvalidTwelveDataResponse = errors.New("invalid twelvedata response")
	ErrTooManyRequests           = errors.New("too many requests")
	ErrNotFound                  = errors.New("not found")
)
