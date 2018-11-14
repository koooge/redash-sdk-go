package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDestinationList(t *testing.T) {
	const getDestinationListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/destinations" {
			fmt.Fprint(w, getDestinationListResBody)
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
		{status: 200, body: getDestinationListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDestinationList()
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

func TestGetDestinationTypeList(t *testing.T) {
	const getDestinationTypeListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/destinations/types" {
			fmt.Fprint(w, getDestinationTypeListResBody)
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
		{status: 200, body: getDestinationTypeListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDestinationTypeList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
