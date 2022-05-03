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

type HTTPCli struct {
	transport *fasthttp.Client
	cfg       *Conf
	logger    *zerolog.Logger
}

func NewHTTPCli(transport *fasthttp.Client, cfg *Conf, logger *zerolog.Logger) *HTTPCli {
	return &HTTPCli{transport: transport, cfg: cfg, logger: logger}
}

func (c *HTTPCli) makeRequest(uri string, resp *fasthttp.Response) (int, int, error) {
	req := fasthttp.AcquireRequest()

	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(uri)

	start := time.Now()

	if err := c.transport.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
		c.logRequest(req, resp, time.Since(start), err)

		if !errors.Is(err, fasthttp.ErrDialTimeout) {
			return 0, 0, fmt.Errorf("http request: %w", err)
		}

		if err := c.transport.DoTimeout(req, resp, time.Duration(c.cfg.Timeout)*time.Second); err != nil {
			return 0, 0, fmt.Errorf("http cli request: %w", err)
		}
	}

	if resp.StatusCode() != http.StatusOK {
		c.logRequest(req, resp, time.Since(start), dictionary.ErrBadStatusCode)

		return 0, 0, dictionary.ErrBadStatusCode
	}

	c.logRequest(req, resp, time.Since(start), nil)

	creditsLeft, creditsUsed, err := c.getCredits(resp)
	if err != nil {
		c.logger.Err(err).Msg("get credits")
	}

	return creditsLeft, creditsUsed, nil
}

func (c *HTTPCli) getCredits(resp *fasthttp.Response) (creditsLeft int, creditsUsed int, err error) {
	creditsLeftStr := string(resp.Header.Peek(dictionary.APICreditsLeft))

	if creditsLeftStr != "" {
		creditsLeft, err = strconv.Atoi(creditsLeftStr)
		if err != nil {
			c.logger.Err(err).Str("val", creditsLeftStr).Msg("str to int")

			return 0, 0, fmt.Errorf("str to int: %w", err)
		}
	}

	creditsUsedStr := string(resp.Header.Peek(dictionary.APICreditsUsed))

	if creditsUsedStr != "" {
		creditsUsed, err = strconv.Atoi(creditsUsedStr)
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

	if err != nil {
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
