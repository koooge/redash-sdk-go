package redash

import (
	"io/ioutil"
	"strconv"
)

type GetDestinationInput struct {
	DestinationId int
}

type GetDestinationOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDestination(input *GetDestinationInput) *GetDestinationOutput {
	path := "/api/destinations/" + strconv.Itoa(input.DestinationId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetDestinationOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDestinationOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
