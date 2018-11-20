package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListGroups(t *testing.T) {
	const listGroupsResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/groups" {
			fmt.Fprint(w, listGroupsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListGroupsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listGroupsResBody},
		{input: &ListGroupsInput{}, status: 200, body: listGroupsResBody},
	}

	for _, c := range testCases {
		result := testClient.ListGroups(c.input)
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

func TestListGroupMembers(t *testing.T) {
	const listGroupMembersResBody = `[{"something":"something"}]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/groups/1/members" {
			fmt.Fprint(w, listGroupMembersResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListGroupMembersInput
		status int
		body   string
	}{
		{input: &ListGroupMembersInput{GroupId: 1}, status: 200, body: listGroupMembersResBody},
	}

	for _, c := range testCases {
		result := testClient.ListGroupMembers(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
