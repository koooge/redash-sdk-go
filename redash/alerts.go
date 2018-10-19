package redash

import (
	"io/ioutil"
	"strconv"
)

type GetAlertInput struct {
	AlertId int
}

type GetAlertOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlert(input *GetAlertInput) *GetAlertOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId)

	resp, err := c.Get(path)
	if err != nil {
		return &GetAlertOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetAlertListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlertList() *GetAlertListOutput {
	path := "/api/alerts"

	resp, err := c.Get(path)
	if err != nil {
		return &GetAlertListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertListOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetAlertSubscriptionListInput struct {
	AlertId int
}

type GetAlertSubscriptionListOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlertSubscriptionList(input *GetAlertSubscriptionListInput) *GetAlertSubscriptionListOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId) + "/subscriptions"

	resp, err := c.Get(path)
	if err != nil {
		return &GetAlertSubscriptionListOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertSubscriptionListOutput{Body: string(b), StatusCode: resp.StatusCode}
}
