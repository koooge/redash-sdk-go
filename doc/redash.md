# redash
--
    import "github.com/koooge/redash-sdk-go/redash"


## Usage

#### type Client

```go
type Client struct {
	Config *Config
}
```


#### func  NewClient

```go
func NewClient(config *Config) *Client
```

#### func (*Client) Delete

```go
func (c *Client) Delete(path string, body string) (*http.Response, error)
```

#### func (*Client) Get

```go
func (c *Client) Get(path string) (*http.Response, error)
```

#### func (*Client) GetAlerts

```go
func (c *Client) GetAlerts(input *GetAlertsInput) *GetAlertsOutput
```

#### func (*Client) GetDataSources

```go
func (c *Client) GetDataSources(input *GetDataSourcesInput) *GetDataSourcesOutput
```

#### func (*Client) GetDestinations

```go
func (c *Client) GetDestinations(input *GetDestinationsInput) *GetDestinationsOutput
```

#### func (*Client) GetEvents

```go
func (c *Client) GetEvents() *GetEventsOutput
```

#### func (*Client) GetGroups

```go
func (c *Client) GetGroups(input *GetGroupsInput) *GetGroupsOutput
```

#### func (*Client) GetOrganizationSettings

```go
func (c *Client) GetOrganizationSettings() *GetSettingsOutput
```

#### func (*Client) GetQueries

```go
func (c *Client) GetQueries(input *GetQueriesInput) *GetQueriesOutput
```

#### func (*Client) GetQueryResults

```go
func (c *Client) GetQueryResults(input *GetQueryResultsInput) *GetQueryResultsOutput
```

#### func (*Client) GetQuerySnippets

```go
func (c *Client) GetQuerySnippets(input *GetQuerySnippetsInput) *GetQuerySnippetsOutput
```

#### func (*Client) GetUsers

```go
func (c *Client) GetUsers(input *GetUsersInput) *GetUsersOutput
```

#### func (*Client) Post

```go
func (c *Client) Post(path string, body string) (*http.Response, error)
```

#### func (*Client) Put

```go
func (c *Client) Put(path string, body string) (*http.Response, error)
```

#### func (*Client) Request

```go
func (c *Client) Request(httpMethod string, path string, body string) (*http.Response, error)
```

#### type Config

```go
type Config struct {
	EndpointUrl string
	ApiKey      string
}
```


#### type GetAlertsInput

```go
type GetAlertsInput struct {
	AlertId int
}
```


#### type GetAlertsOutput

```go
type GetAlertsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDataSourcesInput

```go
type GetDataSourcesInput struct {
	DataSourceId int
}
```


#### type GetDataSourcesOutput

```go
type GetDataSourcesOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDestinationsInput

```go
type GetDestinationsInput struct {
	DestinationId int
}
```


#### type GetDestinationsOutput

```go
type GetDestinationsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetEventsOutput

```go
type GetEventsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetGroupsInput

```go
type GetGroupsInput struct {
	GroupId int
}
```


#### type GetGroupsOutput

```go
type GetGroupsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueriesInput

```go
type GetQueriesInput struct {
	QueryId int
}
```


#### type GetQueriesOutput

```go
type GetQueriesOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryResultsInput

```go
type GetQueryResultsInput struct {
	QueryResultId int
}
```


#### type GetQueryResultsOutput

```go
type GetQueryResultsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQuerySnippetsInput

```go
type GetQuerySnippetsInput struct {
	QuerySnippetId int
}
```


#### type GetQuerySnippetsOutput

```go
type GetQuerySnippetsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetSettingsOutput

```go
type GetSettingsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetUsersInput

```go
type GetUsersInput struct {
	UserId int
}
```


#### type GetUsersOutput

```go
type GetUsersOutput struct {
	Body       string
	StatusCode int
}
```
