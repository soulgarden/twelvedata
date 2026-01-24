package twelvedata

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata/dictionary"
	"github.com/valyala/fasthttp"
)

// HTTPCli represents an HTTP client wrapper for API requests.
type HTTPCli struct {
	transport *fasthttp.Client
	cfg       *Conf
	logger    *zerolog.Logger
}

// NewHTTPCli creates a new HTTP client with the specified transport, configuration, and logger.
func NewHTTPCli(transport *fasthttp.Client, cfg *Conf, logger *zerolog.Logger) *HTTPCli {
	return &HTTPCli{transport: transport, cfg: cfg, logger: logger}
}

func (c *HTTPCli) makeRequest(uri string, resp *fasthttp.Response) (int64, int64, error) {
	return c.doRequest(http.MethodGet, uri, nil, nil, resp)
}

func (c *HTTPCli) doRequest(method string, uri string, headers map[string]string, body []byte, resp *fasthttp.Response) (int64, int64, error) {
	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri)
	if method == "" {
		method = http.MethodGet
	}
	req.Header.SetMethod(method)

	for key, val := range headers {
		if val == "" {
			continue
		}
		req.Header.Set(key, val)
	}

	if body != nil {
		req.SetBody(body)
	}

	start := time.Now()

	if err := c.transport.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
		c.logRequest(req, resp, time.Since(start), err)

		if !errors.Is(err, fasthttp.ErrDialTimeout) {
			return 0, 0, fmt.Errorf("http request: %w", err)
		}

		c.logger.Debug().Msg("retrying request after dial timeout")

		// Reset response to ensure clean state for retry
		resp.Reset()

		// Record new start time for retry logging
		retryStart := time.Now()

		if err := c.transport.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
			c.logRequest(req, resp, time.Since(retryStart), err)
			return 0, 0, fmt.Errorf("http cli request: %w", err)
		}

		// Log successful retry
		c.logRequest(req, resp, time.Since(retryStart), nil)
	}

	statusCode := resp.StatusCode()
	if statusCode != http.StatusOK {
		httpErr := NewHTTPError(statusCode, resp.Body(), uri, nil, nil)
		c.logRequest(req, resp, time.Since(start), httpErr)
	} else {
		c.logRequest(req, resp, time.Since(start), nil)
	}

	creditsLeft, creditsUsed, err := c.getCredits(resp)
	if err != nil {
		c.logger.Err(err).Msg("get credits")
	}

	return creditsLeft, creditsUsed, nil
}

func (c *HTTPCli) getCredits(resp *fasthttp.Response) (creditsLeft int64, creditsUsed int64, err error) {
	creditsLeftStr := string(resp.Header.Peek(dictionary.APICreditsLeft))

	if creditsLeftStr != "" {
		creditsLeft, err = strconv.ParseInt(creditsLeftStr, 10, 0)
		if err != nil {
			c.logger.Err(err).Str("val", creditsLeftStr).Msg("str to int")

			return 0, 0, fmt.Errorf("str to int: %w", err)
		}
	}

	creditsUsedStr := string(resp.Header.Peek(dictionary.APICreditsUsed))

	if creditsUsedStr != "" {
		creditsUsed, err = strconv.ParseInt(creditsUsedStr, 10, 0)
		if err != nil {
			c.logger.Err(err).Str("val", creditsUsedStr).Msg("str to int")

			return creditsLeft, 0, fmt.Errorf("str to int: %w", err)
		}
	}

	return creditsLeft, creditsUsed, nil
}

func (c *HTTPCli) logRequest(
	req *fasthttp.Request,
	resp *fasthttp.Response,
	duration time.Duration,
	err error,
) {
	var event *zerolog.Event

	if err == nil {
		event = c.logger.Debug()
	} else {
		event = c.logger.Err(err)
	}

	event.
		Str("request headers", req.Header.String()).
		Int("response code", resp.StatusCode()).
		Dur("duration", duration).
		Msg("request")
}
