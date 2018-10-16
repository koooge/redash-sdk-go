package redash

import (
	"io/ioutil"
	"strconv"
)

type GetAlertsInput struct {
	AlertId int
}

type GetAlertsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlerts(input *GetAlertsInput) *GetAlertsOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetAlertsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
