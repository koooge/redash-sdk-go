package redash

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Config *Config
}

func NewClient(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

func (c *Client) SetConfig(config *Config) {
	c.Config = config
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.request(http.MethodGet, path, "")
}

func (c *Client) post(path string, body string) (*http.Response, error) {
	return c.request(http.MethodPost, path, body)
}

func (c *Client) put(path string, body string) (*http.Response, error) {
	return c.request(http.MethodPut, path, body)
}

func (c *Client) delete(path string, body string) (*http.Response, error) {
	return c.request(http.MethodDelete, path, body)
}

func (c *Client) request(httpMethod string, path string, body string) (*http.Response, error) {
	url := c.Config.EndpointUrl + path
	resp, err := func() (*http.Response, error) {
		req, err := http.NewRequest(httpMethod, url, strings.NewReader(body))
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", c.Config.ApiKey)

		return http.DefaultClient.Do(req)
	}()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return resp, fmt.Errorf("http response is not OK. code: %d, method: %s, url: %s\nRequestBody: `%s`\nResponseBody: `%s`", resp.StatusCode, httpMethod, url, body, string(b))
	}
	return resp, nil
}
