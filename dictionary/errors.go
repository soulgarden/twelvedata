package dictionary

import "errors"

const SymbolNotFoundMsg = "**symbol** not found:"
const NewSymbolNotFoundMsg = "**symbol** with specified criteria not found:"

var (
	ErrBadStatusCode             = errors.New("bas status code")
	ErrInvalidTwelveDataResponse = errors.New("invalid twelvedata response")
	ErrTooManyRequests           = errors.New("too many requests")
	ErrForbidden                 = errors.New("forbidden")
	ErrNotFound                  = errors.New("not found")
	ErrUnmarshalResponse         = errors.New("unmarshal error response")
)
