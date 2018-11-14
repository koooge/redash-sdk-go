package redash

import (
	"io/ioutil"
)

type GetEventsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetEvents() *GetEventsOutput {
	path := "/api/events"

	resp, err := c.get(path)
	if err != nil {
		return &GetEventsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetEventsOutput{Body: string(b), StatusCode: resp.StatusCode}
}
