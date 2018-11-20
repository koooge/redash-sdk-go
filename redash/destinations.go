package redash

import (
	"io/ioutil"
	"strconv"
)

type ListDestinationsInput struct {
}

type ListDestinationsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDestinations(_ *ListDestinationsInput) *ListDestinationsOutput {
	path := "/api/destinations"

	resp, err := c.get(path)
	if err != nil {
		return &ListDestinationsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDestinationsOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetDestinationInput struct {
	DestinationId int
}

type GetDestinationOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetDestination(input *GetDestinationInput) *GetDestinationOutput {
	path := "/api/destinations/" + strconv.Itoa(input.DestinationId)

	resp, err := c.get(path)
	if err != nil {
		return &GetDestinationOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDestinationOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type ListDestinationTypesInput struct {
}

type ListDestinationTypesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDestinationTypes(_ *ListDestinationTypesInput) *ListDestinationTypesOutput {
	path := "/api/destinations/types"

	resp, err := c.get(path)
	if err != nil {
		return &ListDestinationTypesOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDestinationTypesOutput{Body: string(b), StatusCode: resp.StatusCode}
}
