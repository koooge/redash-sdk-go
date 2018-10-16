package redash

import (
	"io/ioutil"
	"strconv"
)

type GetDestinationsInput struct {
	DestinationId int
}

type GetDestinationsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDestinations(input *GetDestinationsInput) *GetDestinationsOutput {
	path := "/api/destinations/" + strconv.Itoa(input.DestinationId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetDestinationsOutput{Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDestinationsOutput{
		Body:       string(b),
		StatusCode: resp.StatusCode,
	}
}
