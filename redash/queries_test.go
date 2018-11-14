package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQueryList(t *testing.T) {
	const getQueryListResBody = `{"something":"something","results":[{"something":"something"}]}`

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

func TestGetQuery(t *testing.T) {
	const getQueryResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/1" {
			fmt.Fprint(w, getQueryResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetQueryInput
		status int
		body   string
	}{
		{input: &GetQueryInput{QueryId: 1}, status: 200, body: getQueryResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQuery(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetQuerySearch(t *testing.T) {
	const getQuerySearchResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/search" {
			fmt.Fprint(w, getQuerySearchResBody)
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
		{status: 200, body: getQuerySearchResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQuerySearch()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetQueryRecent(t *testing.T) {
	const getQueryRecentResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/recent" {
			fmt.Fprint(w, getQueryRecentResBody)
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
		{status: 200, body: getQueryRecentResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryRecent()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetMyQueries(t *testing.T) {
	const getMyQueriesResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/my" {
			fmt.Fprint(w, getMyQueriesResBody)
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
		{status: 200, body: getMyQueriesResBody},
	}

	for _, c := range testCases {
		result := testClient.GetMyQueries()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetQueryTags(t *testing.T) {
	const getQueryTagsResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/tags" {
			fmt.Fprint(w, getQueryTagsResBody)
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
		{status: 200, body: getQueryTagsResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryTags()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestPostQueryList(t *testing.T) {
	const postQueryListResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries" {
			fmt.Fprint(w, postQueryListResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *PostQueryListInput
		status int
		body   string
	}{
		{input: &PostQueryListInput{DataSourceId: 1, Query: "something", Name: "something"}, status: 200, body: postQueryListResBody},
	}

	for _, c := range testCases {
		result := testClient.PostQueryList(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestPostQuery(t *testing.T) {
	const postQueryResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/1" {
			fmt.Fprint(w, postQueryResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *PostQueryInput
		status int
		body   string
	}{
		{input: &PostQueryInput{QueryId: 1, DataSourceId: 1, Query: "something", Name: "something"}, status: 200, body: postQueryResBody},
	}

	for _, c := range testCases {
		result := testClient.PostQuery(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestDeleteQuery(t *testing.T) {
	const deleteQueryResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/1" {
			fmt.Fprint(w, deleteQueryResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *DeleteQueryInput
		status int
		body   string
	}{
		{input: &DeleteQueryInput{QueryId: 1}, status: 200, body: deleteQueryResBody},
	}

	for _, c := range testCases {
		result := testClient.DeleteQuery(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
