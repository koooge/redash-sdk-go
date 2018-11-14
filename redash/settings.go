package redash

import (
	"io/ioutil"
)

type GetOrganizationSettingsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetOrganizationSettings() *GetOrganizationSettingsOutput {
	path := "/api/settings/organization"

	resp, err := c.get(path)
	if err != nil {
		return &GetOrganizationSettingsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetOrganizationSettingsOutput{Body: string(b), StatusCode: resp.StatusCode}
}
