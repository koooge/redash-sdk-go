package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDestinations(t *testing.T) {
	const listDestinationsResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/destinations" {
			fmt.Fprint(w, listDestinationsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDestinationsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDestinationsResBody},
		{input: &ListDestinationsInput{}, status: 200, body: listDestinationsResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDestinations(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDestination(t *testing.T) {
	const getDestinationResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/destinations/1" {
			fmt.Fprint(w, getDestinationResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDestinationInput
		status int
		body   string
	}{
		{input: &GetDestinationInput{DestinationId: 1}, status: 200, body: getDestinationResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDestination(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestListDestinationTypes(t *testing.T) {
	const listDestinationTypesResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/destinations/types" {
			fmt.Fprint(w, listDestinationTypesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDestinationTypesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDestinationTypesResBody},
		{input: &ListDestinationTypesInput{}, status: 200, body: listDestinationTypesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDestinationTypes(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
