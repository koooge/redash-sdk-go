package redash

import (
	"io/ioutil"
)

type GetSettingsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetOrganizationSettings() *GetSettingsOutput {
	path := "/api/settings/organization"

	resp, err := c.Get(path)
	if err != nil {
		return &GetSettingsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetSettingsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
