package httpx

import (
	"net/http"
	"net/url"
	"testing"
)

func TestHeader(t *testing.T) {
	header := NewBuilder(make(http.Header)).
		Add("ab", "1").
		Add("bc", "2").
		Add("bc", "3").
		Set("cd", "3").
		Add("ab", "4").
		Build()
	t.Log(header)
}

func TestValues(t *testing.T) {
	values := NewBuilder(make(url.Values)).
		Add("ab", "1").
		Add("bc", "2").
		Add("bc", "3").
		Set("cd", "3").
		Add("ab", "4").
		Build()
	t.Log(values)
}
