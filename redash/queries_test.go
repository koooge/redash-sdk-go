package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const getQueryListResBody = `{"something":"something"}`

func TestGetQueryList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries" {
			fmt.Fprint(w, getQueryListResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		status int
		body   string
	}{
		{status: 200, body: getQueryListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
