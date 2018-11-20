package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrganizationSettings(t *testing.T) {
	const getOrganizationSettingsResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/settings/organization" {
			fmt.Fprint(w, getOrganizationSettingsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetOrganizationSettingsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getOrganizationSettingsResBody},
		{input: &GetOrganizationSettingsInput{}, status: 200, body: getOrganizationSettingsResBody},
	}

	for _, c := range testCases {
		result := testClient.GetOrganizationSettings(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
