package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQueryInput struct {
	QueryId int
}

type GetQueryOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuery(input *GetQueryInput) *GetQueryOutput {
	path := "/api/queries/" + strconv.Itoa(input.QueryId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
