package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuerySnippetList(t *testing.T) {
	const getQuerySnippetListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/query_snippets" {
			fmt.Fprint(w, getQuerySnippetListResBody)
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
		{status: 200, body: getQuerySnippetListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQuerySnippetList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetQuerySnippet(t *testing.T) {
	const getQuerySnippetResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/query_snippets/1" {
			fmt.Fprint(w, getQuerySnippetResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetQuerySnippetInput
		status int
		body   string
	}{
		{input: &GetQuerySnippetInput{QuerySnippetId: 1}, status: 200, body: getQuerySnippetResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQuerySnippet(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
