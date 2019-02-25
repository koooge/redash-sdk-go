package redash

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDestinations(t *testing.T) {
	const listDestinationsResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/destinations" {
			w.Write([]byte(listDestinationsResBody))
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

func TestCreateDestination(t *testing.T) {
	const createDestinationResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/destinations" {
			w.Write([]byte(createDestinationResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *CreateDestinationInput
		status int
		body   string
	}{
		{input: &CreateDestinationInput{Options: &CreateDestinationInputOptions{Addresses: "foo@bar.baz"}, Name: "something", Type: "email"}, status: 200, body: createDestinationResBody},
	}

	for _, c := range testCases {
		result := testClient.CreateDestination(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestListDestinationsTypes(t *testing.T) {
	const listDestinationsTypesResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/destinations/types" {
			w.Write([]byte(listDestinationsTypesResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDestinationsTypesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDestinationsTypesResBody},
		{input: &ListDestinationsTypesInput{}, status: 200, body: listDestinationsTypesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDestinationsTypes(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDestination(t *testing.T) {
	const getDestinationResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/destinations/1" {
			w.Write([]byte(getDestinationResBody))
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

func TestUpdateDestination(t *testing.T) {
	const updateDestinationResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/destinations/1" {
			w.Write([]byte(updateDestinationResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *UpdateDestinationInput
		status int
		body   string
	}{
		{input: &UpdateDestinationInput{DestinationId: 1, Options: &UpdateDestinationInputOptions{Addresses: "hoge@fuga.piyo"}, Name: "awesome", Type: "email"}, status: 200, body: updateDestinationResBody},
	}

	for _, c := range testCases {
		result := testClient.UpdateDestination(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestDeleteDestination(t *testing.T) {
	const deleteDestinationResBody = ``

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete && r.URL.Path == "/api/destinations/1" {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(deleteDestinationResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *DeleteDestinationInput
		status int
		body   string
	}{
		{input: &DeleteDestinationInput{DestinationId: 1}, status: 204, body: deleteDestinationResBody},
	}

	for _, c := range testCases {
		result := testClient.DeleteDestination(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
