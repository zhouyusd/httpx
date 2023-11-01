package httpx

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type (
	Client struct {
		client *http.Client
	}

	Option func(*http.Client)
)

func WithTransport(transport http.RoundTripper) Option {
	return func(c *http.Client) {
		c.Transport = transport
	}
}

func WithHttp2() Option {
	return func(c *http.Client) {
		c.Transport = Http2Transport
	}
}

func WithCookieJar() Option {
	return func(c *http.Client) {
		jar, _ := cookiejar.New(nil)
		c.Jar = jar
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *http.Client) {
		c.Timeout = timeout
	}
}

func WithCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) Option {
	return func(c *http.Client) {
		c.CheckRedirect = checkRedirect
	}
}

func NewClient(options ...Option) *Client {
	c := &http.Client{}
	for _, option := range options {
		option(c)
	}
	return &Client{c}
}

func (c *Client) Do(req *Request) *Response {
	httpResp, err := c.client.Do(req.Request())
	return NewResponse(httpResp, err)
}

func (c *Client) Get(url string, options ...ReqOption) *Response {
	return c.Do(NewRequest(http.MethodGet, url, options...))
}

func (c *Client) Post(url string, options ...ReqOption) *Response {
	return c.Do(NewRequest(http.MethodPost, url, options...))
}

func (c *Client) Put(url string, options ...ReqOption) *Response {
	return c.Do(NewRequest(http.MethodPut, url, options...))
}

func (c *Client) Patch(url string, options ...ReqOption) *Response {
	return c.Do(NewRequest(http.MethodPatch, url, options...))
}

func (c *Client) Delete(url string, options ...ReqOption) *Response {
	return c.Do(NewRequest(http.MethodDelete, url, options...))
}
