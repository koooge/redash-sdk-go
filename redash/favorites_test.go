package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQueryFavoriteList(t *testing.T) {
	const getQueryFavoriteListResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/queries/favorites" {
			fmt.Fprint(w, getQueryFavoriteListResBody)
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
		{status: 200, body: getQueryFavoriteListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetQueryFavoriteList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDashboardFavoriteList(t *testing.T) {
	const getDashboardFavoriteListResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards/favorites" {
			fmt.Fprint(w, getDashboardFavoriteListResBody)
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
		{status: 200, body: getDashboardFavoriteListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDashboardFavoriteList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
