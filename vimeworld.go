package vimeworldgo

import (
	"encoding/json"
	"sync/atomic"

	"github.com/valyala/fasthttp"
)

// Client represents a client.
type Client struct {
	client *fasthttp.Client

	last   uint64
	tokens []string
}

// Options represents options for configuring a client.
type Options struct {
	Client *fasthttp.Client
	Tokens []string
}

func (o *Options) setDefaults() {
	if o.Client == nil {
		o.Client = &fasthttp.Client{}
	}
}

// NewClient returns a new Client instance configured with the provided Options.
func NewClient(opts Options) *Client {
	opts.setDefaults()

	return &Client{
		client: opts.Client,
		tokens: opts.Tokens,
	}
}

func (c *Client) newRequest(method, endpoint string, body []byte) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(endpoint)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")

	if len(c.tokens) > 0 {
		req.Header.Set("Access-Token", c.getToken())
	}

	if body != nil {
		req.SetBody(body)
	}

	return req
}

func (c *Client) getToken() string {
	index := atomic.AddUint64(&c.last, 1) - 1
	return c.tokens[index%uint64(len(c.tokens))]
}

// ErrorResponse represents an api error.
type ErrorResponse struct {
	Error Error `json:"error"`
}

func (c *Client) doRequest(method, endpoint string, v interface{}) (*fasthttp.Response, error) {
	req := c.newRequest(method, endpoint, nil)
	defer fasthttp.ReleaseRequest(req)

	return c.do(req, &v)
}

func (c *Client) do(req *fasthttp.Request, v interface{}) (*fasthttp.Response, error) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := c.client.Do(req, resp); err != nil {
		return nil, err
	}

	var errResp ErrorResponse
	_ = json.Unmarshal(resp.Body(), &errResp)

	if errResp.Error.ErrorCode != 0 {
		return nil, errResp.Error
	}

	if err := json.Unmarshal(resp.Body(), &v); err != nil {
		return nil, err
	}

	return resp, nil
}
