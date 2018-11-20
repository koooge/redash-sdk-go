package redash

import (
	"io/ioutil"
	"strconv"
)

type ListAlertsInput struct {
}

type ListAlertsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListAlerts(_ *ListAlertsInput) *ListAlertsOutput {
	path := "/api/alerts"

	resp, err := c.get(path)
	if err != nil {
		return &ListAlertsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListAlertsOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type GetAlertInput struct {
	AlertId int
}

type GetAlertOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) GetAlert(input *GetAlertInput) *GetAlertOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId)

	resp, err := c.get(path)
	if err != nil {
		return &GetAlertOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &GetAlertOutput{Body: string(b), StatusCode: resp.StatusCode}
}

type ListAlertSubscriptionsInput struct {
	AlertId int
}

type ListAlertSubscriptionsOutput struct {
	Body       string
	StatusCode int
}

func (c *Client) ListAlertSubscriptions(input *ListAlertSubscriptionsInput) *ListAlertSubscriptionsOutput {
	path := "/api/alerts/" + strconv.Itoa(input.AlertId) + "/subscriptions"

	resp, err := c.get(path)
	if err != nil {
		return &ListAlertSubscriptionsOutput{Body: `{"error":"` + err.Error() + `"}`, StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return &ListAlertSubscriptionsOutput{Body: string(b), StatusCode: resp.StatusCode}
}
