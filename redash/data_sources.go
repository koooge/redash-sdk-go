package redash

import (
	"io/ioutil"
	"strconv"
)

type ListDataSourcesInput struct {
}

type ListDataSourcesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDataSources(_ *ListDataSourcesInput) *ListDataSourcesOutput {
	path := "/api/data_sources"

	resp, err := c.get(path)
	if err != nil {
		return &ListDataSourcesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDataSourcesOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDataSourceInput struct {
	DataSourceId int
}

type GetDataSourceOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDataSource(input *GetDataSourceInput) *GetDataSourceOutput {
	path := "/api/data_sources/" + strconv.Itoa(input.DataSourceId)

	resp, err := c.get(path)
	if err != nil {
		return &GetDataSourceOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDataSourceSchemaInput struct {
	DataSourceId int
}

type GetDataSourceSchemaOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDataSourceSchema(input *GetDataSourceSchemaInput) *GetDataSourceSchemaOutput {
	path := "/api/data_sources/" + strconv.Itoa(input.DataSourceId) + "/schema"

	resp, err := c.get(path)
	if err != nil {
		return &GetDataSourceSchemaOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceSchemaOutput{Body: string(b), StatusCode: resp.StatusCode}
}
