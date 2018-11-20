package redash

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type ListQueriesInput struct {
}

type ListQueriesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListQueries(_ *ListQueriesInput) *ListQueriesOutput {
	path := "/api/queries"

	resp, err := c.get(path)
	if err != nil {
		return &ListQueriesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListQueriesOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetQueryInput struct {
	QueryId int
}

type GetQueryOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuery(input *GetQueryInput) *GetQueryOutput {
	path := "/api/queries/" + strconv.Itoa(input.QueryId)

	resp, err := c.get(path)
	if err != nil {
		return &GetQueryOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetQuerySearchInput struct {
}

type GetQuerySearchOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQuerySearch(_ *GetQuerySearchInput) *GetQuerySearchOutput {
	path := "/api/queries/search"

	resp, err := c.get(path)
	if err != nil {
		return &GetQuerySearchOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQuerySearchOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetQueryRecentInput struct {
}

type GetQueryRecentOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryRecent(_ *GetQueryRecentInput) *GetQueryRecentOutput {
	path := "/api/queries/recent"

	resp, err := c.get(path)
	if err != nil {
		return &GetQueryRecentOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryRecentOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetMyQueriesInput struct {
}

type GetMyQueriesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetMyQueries(_ *GetMyQueriesInput) *GetMyQueriesOutput {
	path := "/api/queries/my"

	resp, err := c.get(path)
	if err != nil {
		return &GetMyQueriesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetMyQueriesOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetQueryTagsInput struct {
}

type GetQueryTagsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryTags(_ *GetQueryTagsInput) *GetQueryTagsOutput {
	path := "/api/queries/tags"

	resp, err := c.get(path)
	if err != nil {
		return &GetQueryTagsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryTagsOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type CreateQueryInput struct {
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}

type CreateQueryOutput struct {
	QueryId    int    `json:"id"`
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
}

func (c *Client) CreateQuery(input *CreateQueryInput) *CreateQueryOutput {
	path := "/api/queries"

	body, err := json.Marshal(input)
	if err != nil {
		return &CreateQueryOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &CreateQueryOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	output := &CreateQueryOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
	if err := json.Unmarshal(b, &output); err != nil {
		return &CreateQueryOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}

	return output
}

type ModifyQueryInput struct {
	QueryId      int    `json:"-"`
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}

type ModifyQueryOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ModifyQuery(input *ModifyQueryInput) *ModifyQueryOutput {
	path := "/api/queries/" + strconv.Itoa(input.QueryId)

	body, err := json.Marshal(input)
	if err != nil {
		return &ModifyQueryOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &ModifyQueryOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ModifyQueryOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type DeleteQueryInput struct {
	QueryId int
}

type DeleteQueryOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) DeleteQuery(input *DeleteQueryInput) *DeleteQueryOutput {
	path := "/api/queries/" + strconv.Itoa(input.QueryId)

	body, err := json.Marshal(input)
	if err != nil {
		return &DeleteQueryOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.delete(path, string(body))
	if err != nil {
		return &DeleteQueryOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &DeleteQueryOutput{Body: string(b), StatusCode: resp.StatusCode}
}
