package redash

import (
	"io/ioutil"
)

type ListDashboardsInput struct {
}

type ListDashboardsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDashboards(_ *ListDashboardsInput) *ListDashboardsOutput {
	path := "/api/dashboards"

	resp, err := c.get(path)
	if err != nil {
		return &ListDashboardsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDashboardsOutput{Body: string(b), StatusCode: resp.StatusCode}
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

	resp, err := c.get(path)
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

	resp, err := c.get(path)
	if err != nil {
		return &GetPublicDashboardOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetPublicDashboardOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDashboardTagsInput struct {
}

type GetDashboardTagsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDashboardTags(_ *GetDashboardTagsInput) *GetDashboardTagsOutput {
	path := "/api/dashboards/tags"

	resp, err := c.get(path)
	if err != nil {
		return &GetDashboardTagsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDashboardTagsOutput{Body: string(b), StatusCode: resp.StatusCode}
}
