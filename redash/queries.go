package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQueriesInput struct {
	QueryId int
}

type GetQueriesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueries(input *GetQueriesInput) *GetQueriesOutput {
	path := "/api/queries/" + strconv.Itoa(input.QueryId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueriesOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueriesOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
