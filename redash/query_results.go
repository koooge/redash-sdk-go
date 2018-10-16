package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQueryResultsInput struct {
	QueryResultId int
}

type GetQueryResultsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryResults(input *GetQueryResultsInput) *GetQueryResultsOutput {
	path := "/api/query_results/" + strconv.Itoa(input.QueryResultId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryResultsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryResultsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
