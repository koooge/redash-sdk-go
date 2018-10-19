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
		return &GetQueryResultOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryResultOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetJobInput struct {
	JobId int
}

type GetJobOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetJob(input *GetJobInput) *GetJobOutput {
	path := "/api/jobs/" + strconv.Itoa(input.JobId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetJobOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetJobOutput{Body: string(b), StatusCode: resp.StatusCode}
}
