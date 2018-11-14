package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGroupList(t *testing.T) {
	const getGroupListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/groups" {
			fmt.Fprint(w, getGroupListResBody)
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
		{status: 200, body: getGroupListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetGroupList()
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetGroup(t *testing.T) {
	const getGroupResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/groups/1" {
			fmt.Fprint(w, getGroupResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetGroupInput
		status int
		body   string
	}{
		{input: &GetGroupInput{GroupId: 1}, status: 200, body: getGroupResBody},
	}

	for _, c := range testCases {
		result := testClient.GetGroup(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetGroupMemberList(t *testing.T) {
	const getGroupMemberListResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/groups/1/members" {
			fmt.Fprint(w, getGroupMemberListResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetGroupMemberListInput
		status int
		body   string
	}{
		{input: &GetGroupMemberListInput{GroupId: 1}, status: 200, body: getGroupMemberListResBody},
	}

	for _, c := range testCases {
		result := testClient.GetGroupMemberList(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
