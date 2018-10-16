package redash

import (
	"io/ioutil"
	"strconv"
)

type GetDataSourcesInput struct {
	DataSourceId int
}

type GetDataSourcesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDataSources(input *GetDataSourcesInput) *GetDataSourcesOutput {
	path := "/api/data_sources/" + strconv.Itoa(input.DataSourceId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetDataSourcesOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourcesOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
