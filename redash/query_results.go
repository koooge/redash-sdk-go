package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQueryResultInput struct {
	QueryResultId int
}

type GetQueryResultOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryResult(input *GetQueryResultInput) *GetQueryResultOutput {
	path := "/api/query_results/" + strconv.Itoa(input.QueryResultId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryResultOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryResultOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
