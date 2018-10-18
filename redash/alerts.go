package redash

import (
	"io/ioutil"
	"strconv"
)

type GetAlertInput struct {
	AlertId int
}

type GetAlertOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlert(input *GetAlertInput) *GetAlertOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetAlertOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
