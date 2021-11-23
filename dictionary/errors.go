package dictionary

import "errors"

var (
	ErrBadStatusCode             = errors.New("bas status code")
	ErrInvalidTwelveDataResponse = errors.New("invalid twelvedata response")
	ErrTooManyRequests           = errors.New("too many requests")
)
