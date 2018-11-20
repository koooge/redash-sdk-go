package redash

import (
	"io/ioutil"
	"strconv"
)

type ListQuerySnippetsInput struct {
}

type ListQuerySnippetsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListQuerySnippets(_ *ListQuerySnippetsInput) *ListQuerySnippetsOutput {
	path := "/api/query_snippets"

	resp, err := c.get(path)
	if err != nil {
		return &ListQuerySnippetsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListQuerySnippetsOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetQuerySnippetInput struct {
	QuerySnippetId int
}

type GetQuerySnippetOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuerySnippet(input *GetQuerySnippetInput) *GetQuerySnippetOutput {
	path := "/api/query_snippets/" + strconv.Itoa(input.QuerySnippetId)

	resp, err := c.get(path)
	if err != nil {
		return &GetQuerySnippetOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQuerySnippetOutput{Body: string(b), StatusCode: resp.StatusCode}
}
