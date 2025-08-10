package twelvedata

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/soulgarden/twelvedata/response"
	"github.com/valyala/fasthttp"
)

var encoder = schema.NewEncoder()

type Endpoint[Request any, Response any, Credits response.Credits, Error error] struct {
	httpCli *HTTPCli
	URL     string
}

func NewEndpoint[Request any, Response any, Credits response.Credits, Error error](httpCli *HTTPCli, URI string) *Endpoint[Request, Response, Credits, Error] {
	return &Endpoint[Request, Response, Credits, Error]{
		httpCli: httpCli,
		URL:     URI,
	}
}

func (endpoint Endpoint[Request, Response, Credits, ErrorResponse]) Call(req Request) (resp Response, creds response.Credits, err Error) {
	httpResp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(httpResp)

	var (
		creditsLeft, creditsUsed int64
		innerErr                 error
	)

	values := url.Values{}

	if innerErr = encoder.Encode(req, values); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("encoding url params: %w", innerErr), nil)
	}

	var uri *url.URL

	if uri, innerErr = url.Parse(endpoint.URL); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("parse uri: %w", innerErr), nil)
	}

	uri.RawQuery = values.Encode()

	if creditsLeft, creditsUsed, innerErr = endpoint.httpCli.makeRequest(uri.String(), httpResp); innerErr != nil {
		return resp, creds, NewError[Error](innerErr, nil)
	}

	creds = &response.CreditsImpl{}

	creds.SetCreditsLeft(creditsLeft)
	creds.SetCreditsUsed(creditsUsed)

	var respErr response.Error

	if innerErr := json.Unmarshal(httpResp.Body(), &respErr); innerErr == nil && respErr.Status == "error" {
		return resp, creds, NewError[Error](fmt.Errorf("error received: %s", respErr.Error()), respErr)
	}

	if innerErr := json.Unmarshal(httpResp.Body(), &resp); innerErr != nil {
		return resp, creds, NewError[Error](fmt.Errorf("unmarshall json: %w", innerErr), nil)
	}

	// TODO: handle 404/400/500/timeout errors

	return resp, creds, err
}

func NewError[T error](err error, t T) ErrImpl[T] {
	return ErrImpl[T]{
		generic: t,
		inner:   err,
	}
}

type Error interface {
	Error() string
}

type ErrImpl[Err error] struct {
	generic Err
	inner   error
}

func (e ErrImpl[Err]) Error() string {
	return e.inner.Error()
}
