package redash

import "testing"

func TestSetConfig(t *testing.T) {
	client := NewClient(&Config{EndpointUrl: "foo", ApiKey: "foo"})
	client.SetConfig(&Config{EndpointUrl: "bar", ApiKey: "bar"})
	if client.Config.EndpointUrl != "bar" || client.Config.ApiKey != "bar" {
		t.Errorf("client.Config.EndpointUrl: %+v != \"bar\", client.Config.ApiKey: %+v != \"bar\"", client.Config.EndpointUrl, client.Config.ApiKey)
	}
}
