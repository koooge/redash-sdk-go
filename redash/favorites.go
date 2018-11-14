package redash

import (
	"io/ioutil"
)

type GetQueryFavoriteListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetQueryFavoriteList() *GetQueryFavoriteListOutput {
	path := "/api/queries/favorites"

	resp, err := c.get(path)
	if err != nil {
		return &GetQueryFavoriteListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetQueryFavoriteListOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDashboardFavoriteListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDashboardFavoriteList() *GetDashboardFavoriteListOutput {
	path := "/api/dashboards/favorites"

	resp, err := c.get(path)
	if err != nil {
		return &GetDashboardFavoriteListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDashboardFavoriteListOutput{Body: string(b), StatusCode: resp.StatusCode}
}
