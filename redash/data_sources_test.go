package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDataSources(t *testing.T) {
	const listDataSourcesResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources" {
			fmt.Fprint(w, listDataSourcesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDataSourcesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDataSourcesResBody},
		{input: &ListDataSourcesInput{}, status: 200, body: listDataSourcesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDataSources(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestCreateDataSource(t *testing.T) {
	const createDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/data_sources" {
			fmt.Fprint(w, createDataSourceResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *CreateDataSourceInput
		status int
		body   string
	}{
		{input: &CreateDataSourceInput{Options: &CreateDataSourceInputOptions{Dbname: "something"}, Name: "bar", Type: "pg"}, status: 200, body: createDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.CreateDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestListDataSourcesTypes(t *testing.T) {
	const listDataSourcesTypesResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources/types" {
			fmt.Fprint(w, listDataSourcesTypesResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDataSourcesTypesInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDataSourcesTypesResBody},
		{input: &ListDataSourcesTypesInput{}, status: 200, body: listDataSourcesTypesResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDataSourcesTypes(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDataSource(t *testing.T) {
	const getDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources/1" {
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

func TestDeleteDataSource(t *testing.T) {
	const deleteDataSourceResBody = ``

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete && r.URL.Path == "/api/data_sources/123" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *DeleteDataSourceInput
		status int
		body   string
	}{
		{input: &DeleteDataSourceInput{DataSourceId: 123}, status: 204, body: deleteDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.DeleteDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDataSourceSchema(t *testing.T) {
	const getDataSourceSchemaResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources/1/schema" {
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
