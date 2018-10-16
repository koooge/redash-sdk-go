package redash

import (
	"io/ioutil"
	"strconv"
)

type GetUsersInput struct {
	UserId int
}

type GetUsersOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetUsers(input *GetUsersInput) *GetUsersOutput {
	path := "/api/users/" + strconv.Itoa(input.UserId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetUsersOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetUsersOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
