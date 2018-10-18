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
		return &GetDataSourceOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDataSourceOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
