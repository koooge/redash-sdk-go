package redash

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// GET /api/data_sources
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
		return &ListDataSourcesOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDataSourcesOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/data_sources
type CreateDataSourceInput struct {
	Options *CreateDataSourceInputOptions `json:"options"`
	Name    string                        `json:"name"`
	Type    string                        `json:"type"`
}

type CreateDataSourceInputOptions struct {
	Dbname string `json:"dbname"`
}

type CreateDataSourceOutput struct {
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
	Id         int    `json:"id"`
}

func (c *Client) CreateDataSource(input *CreateDataSourceInput) *CreateDataSourceOutput {
	path := "/api/data_sources"

	body, err := json.Marshal(input)
	if err != nil {
		return &CreateDataSourceOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &CreateDataSourceOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	resBody := &CreateDataSourceOutput{}
	if err := json.Unmarshal(b, &resBody); err != nil {
		return &CreateDataSourceOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	return &CreateDataSourceOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
		Id:         resBody.Id,
	}
}

// GET /api/data_sources/{data_source_id}
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
		return &GetDataSourceOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/data_sources/{data_source_id}

// DELETE /api/data_sources/{data_source_id}
type DeleteDataSourceInput struct {
	DataSourceId int
}

type DeleteDataSourceOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) DeleteDataSource(input *DeleteDataSourceInput) *DeleteDataSourceOutput {
	path := "/api/data_sources/" + strconv.Itoa(input.DataSourceId)

	resp, err := c.delete(path)
	if err != nil {
		return &DeleteDataSourceOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &DeleteDataSourceOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// GET /api/data_sources/{data_source_id}/types

// GET /api/data_sources/{data_source_id}/schema
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
		return &GetDataSourceSchemaOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceSchemaOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/data_sources/{data_source_id}/pause

// DELETE /api/data_sources/{data_source_id}/pause

// POST /api/data_sources/{data_source_id}/test
