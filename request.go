package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type (
	Request struct {
		url    string
		method string
		data   interface{}
		header http.Header
		query  url.Values
	}
	ReqOption func(*Request)
)

func WithHeader(header http.Header) ReqOption {
	return func(r *Request) {
		r.header = header
	}
}

func WithQuery(query url.Values) ReqOption {
	return func(r *Request) {
		r.query = query
	}
}

func WithData(data interface{}) ReqOption {
	return func(r *Request) {
		r.data = data
	}
}

func NewRequest(method, url string, options ...ReqOption) *Request {
	req := &Request{url: url, method: method}
	for _, option := range options {
		option(req)
	}
	return req
}

func (r *Request) Url() string {
	if r.query != nil {
		return strings.TrimRight(r.url, "?") + "?" + r.query.Encode()
	}
	return r.url
}

func (r *Request) Data() io.Reader {
	if r.data == nil {
		return nil
	}
	switch v := r.data.(type) {
	case string:
		return strings.NewReader(v)
	case []byte:
		return bytes.NewReader(v)
	case url.Values:
		return strings.NewReader(v.Encode())
	case io.Reader:
		return v
	default:
		marshalJSON, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return bytes.NewReader(marshalJSON)
	}
}

func (r *Request) Request() *http.Request {
	req, _ := http.NewRequest(r.method, r.Url(), r.Data())
	req.Header = r.header
	return req
}
