package redash

import (
	"io/ioutil"
)

type GetDashboardListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDashboardList() *GetDashboardListOutput {
	path := "/api/dashboards"

	resp, err := c.Get(path)
	if err != nil {
		return &GetDashboardListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDashboardListOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDashboardInput struct {
	DashboardSlug string
}

type GetDashboardOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDashboard(input *GetDashboardInput) *GetDashboardOutput {
	path := "/api/dashboards/" + input.DashboardSlug

	resp, err := c.Get(path)
	if err != nil {
		return &GetDashboardOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDashboardOutput{Body: string(b), StatusCode: resp.StatusCode}
}
