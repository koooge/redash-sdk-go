package redash

import (
	"encoding/json"
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

type GetQueryListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryList() *GetQueryListOutput {
	path := "/api/queries"

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryListOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryListOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type GetQuerySearchOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuerySearch() *GetQuerySearchOutput {
	path := "/api/queries/search"

	resp, err := c.Get(path)
	if err != nil {
		return &GetQuerySearchOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQuerySearchOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type GetQueryRecentOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryRecent() *GetQueryRecentOutput {
	path := "/api/queries/recent"

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryRecentOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryRecentOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type GetMyQueriesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetMyQueries() *GetMyQueriesOutput {
	path := "/api/queries/my"

	resp, err := c.Get(path)
	if err != nil {
		return &GetMyQueriesOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetMyQueriesOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type GetQueryTagsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryTags() *GetQueryTagsOutput {
	path := "/api/queries/tags"

	resp, err := c.Get(path)
	if err != nil {
		return &GetQueryTagsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryTagsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}

type PostQueryListInput struct {
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}

type PostQueryListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) PostQueryList(input *PostQueryListInput) *PostQueryListOutput {
	path := "/api/queries"

	body, err := json.Marshal(input)
	if err != nil {
		return &PostQueryListOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.Post(path, string(body))
	if err != nil {
		return &PostQueryListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &PostQueryListOutput{Body: string(b), StatusCode: resp.StatusCode}
}
