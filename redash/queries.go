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
