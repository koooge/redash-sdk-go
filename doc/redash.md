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

#### func (*Client) GetAlert

```go
func (c *Client) GetAlert(input *GetAlertInput) *GetAlertOutput
```

#### func (*Client) GetAlertList

```go
func (c *Client) GetAlertList() *GetAlertListOutput
```

#### func (*Client) GetDataSource

```go
func (c *Client) GetDataSource(input *GetDataSourceInput) *GetDataSourceOutput
```

#### func (*Client) GetDataSourceList

```go
func (c *Client) GetDataSourceList() *GetDataSourceListOutput
```

#### func (*Client) GetDestination

```go
func (c *Client) GetDestination(input *GetDestinationInput) *GetDestinationOutput
```

#### func (*Client) GetDestinationList

```go
func (c *Client) GetDestinationList() *GetDestinationListOutput
```

#### func (*Client) GetEvents

```go
func (c *Client) GetEvents() *GetEventsOutput
```

#### func (*Client) GetGroup

```go
func (c *Client) GetGroup(input *GetGroupInput) *GetGroupOutput
```

#### func (*Client) GetGroupList

```go
func (c *Client) GetGroupList() *GetGroupListOutput
```

#### func (*Client) GetOrganizationSettings

```go
func (c *Client) GetOrganizationSettings() *GetOrganizationSettingsOutput
```

#### func (*Client) GetQuery

```go
func (c *Client) GetQuery(input *GetQueryInput) *GetQueryOutput
```

#### func (*Client) GetQueryList

```go
func (c *Client) GetQueryList() *GetQueryListOutput
```

#### func (*Client) GetQueryResult

```go
func (c *Client) GetQueryResult(input *GetQueryResultInput) *GetQueryResultOutput
```

#### func (*Client) GetQuerySnippet

```go
func (c *Client) GetQuerySnippet(input *GetQuerySnippetInput) *GetQuerySnippetOutput
```

#### func (*Client) GetQuerySnippetList

```go
func (c *Client) GetQuerySnippetList() *GetQuerySnippetListOutput
```

#### func (*Client) GetUser

```go
func (c *Client) GetUser(input *GetUserInput) *GetUserOutput
```

#### func (*Client) GetUserList

```go
func (c *Client) GetUserList() *GetUserListOutput
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


#### type GetAlertInput

```go
type GetAlertInput struct {
	AlertId int
}
```


#### type GetAlertListOutput

```go
type GetAlertListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetAlertOutput

```go
type GetAlertOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDataSourceInput

```go
type GetDataSourceInput struct {
	DataSourceId int
}
```


#### type GetDataSourceListOutput

```go
type GetDataSourceListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDataSourceOutput

```go
type GetDataSourceOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDestinationInput

```go
type GetDestinationInput struct {
	DestinationId int
}
```


#### type GetDestinationListOutput

```go
type GetDestinationListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDestinationOutput

```go
type GetDestinationOutput struct {
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


#### type GetGroupInput

```go
type GetGroupInput struct {
	GroupId int
}
```


#### type GetGroupListOutput

```go
type GetGroupListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetGroupOutput

```go
type GetGroupOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetOrganizationSettingsOutput

```go
type GetOrganizationSettingsOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryInput

```go
type GetQueryInput struct {
	QueryId int
}
```


#### type GetQueryListOutput

```go
type GetQueryListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryOutput

```go
type GetQueryOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryResultInput

```go
type GetQueryResultInput struct {
	QueryResultId int
}
```


#### type GetQueryResultOutput

```go
type GetQueryResultOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQuerySnippetInput

```go
type GetQuerySnippetInput struct {
	QuerySnippetId int
}
```


#### type GetQuerySnippetListOutput

```go
type GetQuerySnippetListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQuerySnippetOutput

```go
type GetQuerySnippetOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetUserInput

```go
type GetUserInput struct {
	UserId int
}
```


#### type GetUserListOutput

```go
type GetUserListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetUserOutput

```go
type GetUserOutput struct {
	Body       string
	StatusCode int
}
```
