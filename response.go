package httpx

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Response struct {
	Header     http.Header
	Status     string
	StatusCode int
	Cookies    []*http.Cookie
	Body       bytes.Buffer
	Error      error
}

func NewResponse(httpResp *http.Response, err error) *Response {
	if err != nil {
		return &Response{Error: err}
	}
	var body bytes.Buffer
	if httpResp != nil && httpResp.Body != nil {
		_, _ = body.ReadFrom(httpResp.Body)
		httpResp.Body.Close()
	}
	return &Response{
		httpResp.Header,
		httpResp.Status,
		httpResp.StatusCode,
		httpResp.Cookies(),
		body,
		nil,
	}
}

func (r *Response) Text() string {
	return r.Body.String()
}

func (r *Response) Bytes() []byte {
	return r.Body.Bytes()
}

func (r *Response) Bind(v interface{}) error {
	if r.Error != nil {
		return r.Error
	}
	return json.Unmarshal(r.Body.Bytes(), v)
}
