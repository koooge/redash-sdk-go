package redash

import (
	"io/ioutil"
	"strconv"
)

type GetDataSourceInput struct {
	DataSourceId int
}

type GetDataSourceOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDataSource(input *GetDataSourceInput) *GetDataSourceOutput {
	path := "/api/data_sources/" + strconv.Itoa(input.DataSourceId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetDataSourceOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDataSourceListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDataSourceList() *GetDataSourceListOutput {
	path := "/api/data_sources"

	resp, err := c.Get(path)
	if err != nil {
		return &GetDataSourceListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceListOutput{Body: string(b), StatusCode: resp.StatusCode}
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

	resp, err := c.Get(path)
	if err != nil {
		return &GetDataSourceSchemaOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceSchemaOutput{Body: string(b), StatusCode: resp.StatusCode}
}
