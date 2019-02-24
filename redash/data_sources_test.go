package redash

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDataSources(t *testing.T) {
	const listDataSourcesResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources" {
			w.Write([]byte(listDataSourcesResBody))
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
			w.Write([]byte(createDataSourceResBody))
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
			w.Write([]byte(listDataSourcesTypesResBody))
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
			w.Write([]byte(getDataSourceResBody))
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

func TestUpdateDataSource(t *testing.T) {
	const updateDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/data_sources/123" {
			w.Write([]byte(updateDataSourceResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *UpdateDataSourceInput
		status int
		body   string
	}{
		{input: &UpdateDataSourceInput{DataSourceId: 123, Options: &UpdateDataSourceInputOptions{Dbname: "something!"}, Name: "baz", Type: "pg"}, status: 200, body: updateDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.UpdateDataSource(c.input)
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

func TestGetDataSourcesSchema(t *testing.T) {
	const getDataSourcesSchemaResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/api/data_sources/1/schema" {
			w.Write([]byte(getDataSourcesSchemaResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDataSourcesSchemaInput
		status int
		body   string
	}{
		{input: &GetDataSourcesSchemaInput{DataSourceId: 1}, status: 200, body: getDataSourcesSchemaResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDataSourcesSchema(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestPauseDataSource(t *testing.T) {
	const pauseDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/data_sources/123/pause" {
			w.Write([]byte(pauseDataSourceResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *PauseDataSourceInput
		status int
		body   string
	}{
		{input: &PauseDataSourceInput{DataSourceId: 123}, status: 200, body: pauseDataSourceResBody},
		{input: &PauseDataSourceInput{DataSourceId: 123, Reason: "something"}, status: 200, body: pauseDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.PauseDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestUnpauseDataSource(t *testing.T) {
	const unpauseDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete && r.URL.Path == "/api/data_sources/123/pause" {
			w.Write([]byte(unpauseDataSourceResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *UnpauseDataSourceInput
		status int
		body   string
	}{
		{input: &UnpauseDataSourceInput{DataSourceId: 123}, status: 200, body: unpauseDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.UnpauseDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestTestDataSource(t *testing.T) {
	const testDataSourceResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/data_sources/123/test" {
			w.Write([]byte(testDataSourceResBody))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *TestDataSourceInput
		status int
		body   string
	}{
		{input: &TestDataSourceInput{DataSourceId: 123}, status: 200, body: testDataSourceResBody},
	}

	for _, c := range testCases {
		result := testClient.TestDataSource(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
