package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAlerts(t *testing.T) {
	const listAlertsResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/alerts" {
			fmt.Fprint(w, listAlertsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListAlertsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listAlertsResBody},
		{input: &ListAlertsInput{}, status: 200, body: listAlertsResBody},
	}

	for _, c := range testCases {
		result := testClient.ListAlerts(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetAlert(t *testing.T) {
	const getAlertResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/alerts/1" {
			fmt.Fprint(w, getAlertResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetAlertInput
		status int
		body   string
	}{
		{input: &GetAlertInput{AlertId: 1}, status: 200, body: getAlertResBody},
	}

	for _, c := range testCases {
		result := testClient.GetAlert(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestListAlertSubscriptions(t *testing.T) {
	const listAlertSubscriptionsResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/alerts/1/subscriptions" {
			fmt.Fprint(w, listAlertSubscriptionsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListAlertSubscriptionsInput
		status int
		body   string
	}{
		{input: &ListAlertSubscriptionsInput{AlertId: 1}, status: 200, body: listAlertSubscriptionsResBody},
	}

	for _, c := range testCases {
		result := testClient.ListAlertSubscriptions(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
