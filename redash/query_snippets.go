package redash

import (
	"io/ioutil"
	"strconv"
)

type GetQuerySnippetsInput struct {
	QuerySnippetId int
}

type GetQuerySnippetsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuerySnippets(input *GetQuerySnippetsInput) *GetQuerySnippetsOutput {
	path := "/api/query_snippets/" + strconv.Itoa(input.QuerySnippetId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetQuerySnippetsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQuerySnippetsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
