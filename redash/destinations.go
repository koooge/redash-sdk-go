package redash

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// GET /api/destinations
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
		return &ListDestinationsOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDestinationsOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/destinations
type CreateDestinationInput struct {
	Options *CreateDestinationInputOptions `json:"options"`
	Name    string                         `json:"name"`
	Type    string                         `json:"type"`
}

type CreateDestinationInputOptions struct {
	Addresses string `json:"addresses"`
}

type CreateDestinationOutput struct {
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
	Id         int    `json:"id"`
}

func (c *Client) CreateDestination(input *CreateDestinationInput) *CreateDestinationOutput {
	path := "/api/destinations"

	body, err := json.Marshal(input)
	if err != nil {
		return &CreateDestinationOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &CreateDestinationOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	resBody := &CreateDestinationOutput{}
	if err := json.Unmarshal(b, &resBody); err != nil {
		return &CreateDestinationOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}

	return &CreateDestinationOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
		Id:         resBody.Id,
	}
}

// GET /api/destinations/types
type ListDestinationsTypesInput struct {
}

type ListDestinationsTypesOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListDestinationsTypes(_ *ListDestinationsTypesInput) *ListDestinationsTypesOutput {
	path := "/api/destinations/types"

	resp, err := c.get(path)
	if err != nil {
		return &ListDestinationsTypesOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListDestinationsTypesOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// GET /api/destinations/{destination_id}
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
		return &GetDestinationOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetDestinationOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// POST /api/destinations/{destination_id}
type UpdateDestinationInput struct {
	DestinationId int                            `json:"-"`
	Options       *UpdateDestinationInputOptions `json:"options"`
	Name          string                         `json:"name"`
	Type          string                         `json:"type"`
}

type UpdateDestinationInputOptions struct {
	Addresses string `json:"addresses"`
}

type UpdateDestinationOutput struct {
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
}

func (c *Client) UpdateDestination(input *UpdateDestinationInput) *UpdateDestinationOutput {
	path := "/api/destinations/" + strconv.Itoa(input.DestinationId)

	body, err := json.Marshal(input)
	if err != nil {
		return &UpdateDestinationOutput{Body: `{"error":"` + err.Error() + `"}`}
	}

	resp, err := c.post(path, string(body))
	if err != nil {
		return &UpdateDestinationOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &UpdateDestinationOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}

// DELETE /api/destinations/{destination_id}
type DeleteDestinationInput struct {
	DestinationId int
}

type DeleteDestinationOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) DeleteDestination(input *DeleteDestinationInput) *DeleteDestinationOutput {
	path := "/api/destinations/" + strconv.Itoa(input.DestinationId)

	resp, err := c.delete(path)
	if err != nil {
		return &DeleteDestinationOutput{StatusCode: resp.StatusCode, Body: `{"error":"` + err.Error() + `"}`}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &DeleteDestinationOutput{
		StatusCode: resp.StatusCode,
		Body:       string(b),
	}
}
