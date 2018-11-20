package redash

import (
	"io/ioutil"
)

type ListQueryFavoritesInput struct {
}

type ListQueryFavoritesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListQueryFavorites(_ *ListQueryFavoritesInput) *ListQueryFavoritesOutput {
	path := "/api/queries/favorites"

	resp, err := c.get(path)
	if err != nil {
		return &ListQueryFavoritesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListQueryFavoritesOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type ListDashboardFavoritesInput struct {
}

type ListDashboardFavoritesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDashboardFavorites(_ *ListDashboardFavoritesInput) *ListDashboardFavoritesOutput {
	path := "/api/dashboards/favorites"

	resp, err := c.get(path)
	if err != nil {
		return &ListDashboardFavoritesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDashboardFavoritesOutput{Body: string(b), StatusCode: resp.StatusCode}
}
