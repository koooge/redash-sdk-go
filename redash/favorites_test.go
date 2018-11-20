package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListQueryFavorites(t *testing.T) {
	const listQueryFavoritesResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/favorites" {
			fmt.Fprint(w, listQueryFavoritesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListQueryFavoritesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listQueryFavoritesResBody},
		{input: &ListQueryFavoritesInput{}, status: 200, body: listQueryFavoritesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListQueryFavorites(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestListDashboardFavorites(t *testing.T) {
	const listDashboardFavoritesResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards/favorites" {
			fmt.Fprint(w, listDashboardFavoritesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDashboardFavoritesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDashboardFavoritesResBody},
		{input: &ListDashboardFavoritesInput{}, status: 200, body: listDashboardFavoritesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDashboardFavorites(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
