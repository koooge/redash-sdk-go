package redash

import (
	"io/ioutil"
	"strconv"
)

type GetGroupsInput struct {
	GroupId int
}

type GetGroupsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetGroups(input *GetGroupsInput) *GetGroupsOutput {
	path := "/api/groups/" + strconv.Itoa(input.GroupId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetGroupsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetGroupsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
