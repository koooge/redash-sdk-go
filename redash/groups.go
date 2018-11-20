package redash

import (
	"io/ioutil"
	"strconv"
)

type ListGroupsInput struct {
}

type ListGroupsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListGroups(_ *ListGroupsInput) *ListGroupsOutput {
	path := "/api/groups"

	resp, err := c.get(path)
	if err != nil {
		return &ListGroupsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListGroupsOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetGroupInput struct {
	GroupId int
}

type GetGroupOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetGroup(input *GetGroupInput) *GetGroupOutput {
	path := "/api/groups/" + strconv.Itoa(input.GroupId)

	resp, err := c.get(path)
	if err != nil {
		return &GetGroupOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetGroupOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type ListGroupMembersInput struct {
	GroupId int
}

type ListGroupMembersOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListGroupMembers(input *ListGroupMembersInput) *ListGroupMembersOutput {
	path := "/api/groups/" + strconv.Itoa(input.GroupId) + "/members"

	resp, err := c.get(path)
	if err != nil {
		return &ListGroupMembersOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListGroupMembersOutput{Body: string(b), StatusCode: resp.StatusCode}
}
