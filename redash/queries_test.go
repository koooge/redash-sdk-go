package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListQueries(t *testing.T) {
	const listQueriesResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries" {
			fmt.Fprint(w, listQueriesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListQueriesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listQueriesResBody},
		{input: &ListQueriesInput{}, status: 200, body: listQueriesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListQueries(c.input)
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
		input  *GetQuerySearchInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getQuerySearchResBody},
		{input: &GetQuerySearchInput{}, status: 200, body: getQuerySearchResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQuerySearch(c.input)
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
		input  *GetQueryRecentInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getQueryRecentResBody},
		{input: &GetQueryRecentInput{}, status: 200, body: getQueryRecentResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryRecent(c.input)
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
		input  *GetMyQueriesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getMyQueriesResBody},
		{input: &GetMyQueriesInput{}, status: 200, body: getMyQueriesResBody},
	}

	for _, c := range testCases {
		result := testClient.GetMyQueries(c.input)
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
		input  *GetQueryTagsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getQueryTagsResBody},
		{input: &GetQueryTagsInput{}, status: 200, body: getQueryTagsResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryTags(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestCreateQuery(t *testing.T) {
	const createQueryResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries" {
			fmt.Fprint(w, createQueryResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *CreateQueryInput
		status int
		body   string
	}{
		{input: &CreateQueryInput{DataSourceId: 1, Query: "something", Name: "something"}, status: 200, body: createQueryResBody},
	}

	for _, c := range testCases {
		result := testClient.CreateQuery(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestModifyQuery(t *testing.T) {
	const modifyQueryResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/1" {
			fmt.Fprint(w, modifyQueryResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ModifyQueryInput
		status int
		body   string
	}{
		{input: &ModifyQueryInput{QueryId: 1, DataSourceId: 1, Query: "something", Name: "something"}, status: 200, body: modifyQueryResBody},
	}

	for _, c := range testCases {
		result := testClient.ModifyQuery(c.input)
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
