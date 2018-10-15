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

func (c *Client) Get(path string) (*http.Response, error) {
	return c.Request(http.MethodGet, path, "")
}

func (c *Client) Post(path string, body string) (*http.Response, error) {
	return c.Request(http.MethodPost, path, body)
}

func (c *Client) Put(path string, body string) (*http.Response, error) {
	return c.Request(http.MethodPut, path, body)
}

func (c *Client) Delete(path string, body string) (*http.Response, error) {
	return c.Request(http.MethodDelete, path, body)
}

func (c *Client) Request(httpMethod string, path string, body string) (*http.Response, error) {
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

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return resp, fmt.Errorf("http response is not OK. code: %d, method: %s, url: %s\nRequestBody: `%s`\nResponseBody: `%s`", resp.StatusCode, httpMethod, url, body, string(b))
	}
	return resp, nil
}
