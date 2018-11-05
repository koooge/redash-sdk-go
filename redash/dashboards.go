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

type GetPublicDashboardInput struct {
	Token string
}

type GetPublicDashboardOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetPublicDashboard(input *GetPublicDashboardInput) *GetPublicDashboardOutput {
	path := "/api/dashboards/public/" + input.Token

	resp, err := c.Get(path)
	if err != nil {
		return &GetPublicDashboardOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetPublicDashboardOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDashboardTagsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDashboardTags() *GetDashboardTagsOutput {
	path := "/api/dashboards/tags"

	resp, err := c.Get(path)
	if err != nil {
		return &GetDashboardTagsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDashboardTagsOutput{Body: string(b), StatusCode: resp.StatusCode}
}
