package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQuerySnippetInput struct {
	QuerySnippetId int
}

type GetQuerySnippetOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuerySnippet(input *GetQuerySnippetInput) *GetQuerySnippetOutput {
	path := "/api/query_snippets/" + strconv.Itoa(input.QuerySnippetId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQuerySnippetOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQuerySnippetOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}