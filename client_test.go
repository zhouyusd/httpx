package httpx

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := NewClient(
		WithTimeout(10*time.Second),
		WithCookieJar(),
		WithCheckRedirect(nil),
		WithTransport(nil),
		//WithHttp2(),
	)
	resp := client.Get(
		"http://192.168.2.1:5700/api/crons",
		WithQuery(
			NewBuilder(make(url.Values)).
				Set("searchValue", "ksdt").
				Set("page", "1").
				Set("size", "20").
				Set("filters", "{}").
				Set("queryString", "{\"filters\":null,\"sorts\":null,\"filterRelation\":\"and\"}").
				Set("t", fmt.Sprint(time.Now().UnixMilli())).
				Build()),
		WithHeader(
			NewBuilder(make(http.Header)).
				Set("Authorization", "Bearer eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiclVYcFRsZFFEbGxNMG5ndFF6N0hyb18tRVJXVk1yVjltb0huXzVjdmJpcWd5U2RNM2N0N2dGRlAwT0c3clpmeXBfRWpsYUhnNlZMZjYwZEcyayIsImlhdCI6MTY5ODU2NjA4OSwiZXhwIjoxNzAzNzUwMDg5fQ.jXAHVKZt2CcVgX-TEVKmQ7xC9PXuGAQ6w8Foct483G36I_eZwzlUS5zK_wtOuco9").
				Build()),
	)
	if resp.Error != nil {
		t.Error(resp.Error)
	} else {
		t.Log(resp.Text())
	}
}
