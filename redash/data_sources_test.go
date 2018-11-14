package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDataSourceList(t *testing.T) {
	const getDataSourceListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/data_sources" {
			fmt.Fprint(w, getDataSourceListResBody)
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
		{status: 200, body: getDataSourceListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDataSourceList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDataSource(t *testing.T) {
	const getDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/data_sources/1" {
			fmt.Fprint(w, getDataSourceResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDataSourceInput
		status int
		body   string
	}{
		{input: &GetDataSourceInput{DataSourceId: 1}, status: 200, body: getDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDataSourceSchema(t *testing.T) {
	const getDataSourceSchemaResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/data_sources/1/schema" {
			fmt.Fprint(w, getDataSourceSchemaResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDataSourceSchemaInput
		status int
		body   string
	}{
		{input: &GetDataSourceSchemaInput{DataSourceId: 1}, status: 200, body: getDataSourceSchemaResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDataSourceSchema(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
