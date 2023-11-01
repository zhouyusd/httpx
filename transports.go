package httpx

import (
	"crypto/tls"
	"golang.org/x/net/http2"
)

var (
	Http2Transport = &http2.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)
