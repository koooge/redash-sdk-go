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

#### func (*Client) DeleteQuery

```go
func (c *Client) DeleteQuery(input *DeleteQueryInput) *DeleteQueryOutput
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

#### func (*Client) GetAlertSubscriptionList

```go
func (c *Client) GetAlertSubscriptionList(input *GetAlertSubscriptionListInput) *GetAlertSubscriptionListOutput
```

#### func (*Client) GetDashboard

```go
func (c *Client) GetDashboard(input *GetDashboardInput) *GetDashboardOutput
```

#### func (*Client) GetDashboardFavoriteList

```go
func (c *Client) GetDashboardFavoriteList() *GetDashboardFavoriteListOutput
```

#### func (*Client) GetDashboardList

```go
func (c *Client) GetDashboardList() *GetDashboardListOutput
```

#### func (*Client) GetDashboardTags

```go
func (c *Client) GetDashboardTags() *GetDashboardTagsOutput
```

#### func (*Client) GetDataSource

```go
func (c *Client) GetDataSource(input *GetDataSourceInput) *GetDataSourceOutput
```

#### func (*Client) GetDataSourceList

```go
func (c *Client) GetDataSourceList() *GetDataSourceListOutput
```

#### func (*Client) GetDataSourceSchema

```go
func (c *Client) GetDataSourceSchema(input *GetDataSourceSchemaInput) *GetDataSourceSchemaOutput
```

#### func (*Client) GetDestination

```go
func (c *Client) GetDestination(input *GetDestinationInput) *GetDestinationOutput
```

#### func (*Client) GetDestinationList

```go
func (c *Client) GetDestinationList() *GetDestinationListOutput
```

#### func (*Client) GetDestinationTypeList

```go
func (c *Client) GetDestinationTypeList() *GetDestinationTypeListOutput
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

#### func (*Client) GetGroupMemberList

```go
func (c *Client) GetGroupMemberList(input *GetGroupMemberListInput) *GetGroupMemberListOutput
```

#### func (*Client) GetJob

```go
func (c *Client) GetJob(input *GetJobInput) *GetJobOutput
```

#### func (*Client) GetMyQueries

```go
func (c *Client) GetMyQueries() *GetMyQueriesOutput
```

#### func (*Client) GetOrganizationSettings

```go
func (c *Client) GetOrganizationSettings() *GetOrganizationSettingsOutput
```

#### func (*Client) GetPublicDashboard

```go
func (c *Client) GetPublicDashboard(input *GetPublicDashboardInput) *GetPublicDashboardOutput
```

#### func (*Client) GetQuery

```go
func (c *Client) GetQuery(input *GetQueryInput) *GetQueryOutput
```

#### func (*Client) GetQueryFavoriteList

```go
func (c *Client) GetQueryFavoriteList() *GetQueryFavoriteListOutput
```

#### func (*Client) GetQueryList

```go
func (c *Client) GetQueryList() *GetQueryListOutput
```

#### func (*Client) GetQueryRecent

```go
func (c *Client) GetQueryRecent() *GetQueryRecentOutput
```

#### func (*Client) GetQueryResult

```go
func (c *Client) GetQueryResult(input *GetQueryResultInput) *GetQueryResultOutput
```

#### func (*Client) GetQuerySearch

```go
func (c *Client) GetQuerySearch() *GetQuerySearchOutput
```

#### func (*Client) GetQuerySnippet

```go
func (c *Client) GetQuerySnippet(input *GetQuerySnippetInput) *GetQuerySnippetOutput
```

#### func (*Client) GetQuerySnippetList

```go
func (c *Client) GetQuerySnippetList() *GetQuerySnippetListOutput
```

#### func (*Client) GetQueryTags

```go
func (c *Client) GetQueryTags() *GetQueryTagsOutput
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

#### func (*Client) PostQuery

```go
func (c *Client) PostQuery(input *PostQueryInput) *PostQueryOutput
```

#### func (*Client) PostQueryList

```go
func (c *Client) PostQueryList(input *PostQueryListInput) *PostQueryListOutput
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


#### type DeleteQueryInput

```go
type DeleteQueryInput struct {
	QueryId int
}
```


#### type DeleteQueryOutput

```go
type DeleteQueryOutput struct {
	Body       string
	StatusCode int
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


#### type GetAlertSubscriptionListInput

```go
type GetAlertSubscriptionListInput struct {
	AlertId int
}
```


#### type GetAlertSubscriptionListOutput

```go
type GetAlertSubscriptionListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDashboardFavoriteListOutput

```go
type GetDashboardFavoriteListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDashboardInput

```go
type GetDashboardInput struct {
	DashboardSlug string
}
```


#### type GetDashboardListOutput

```go
type GetDashboardListOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDashboardOutput

```go
type GetDashboardOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDashboardTagsOutput

```go
type GetDashboardTagsOutput struct {
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


#### type GetDataSourceSchemaInput

```go
type GetDataSourceSchemaInput struct {
	DataSourceId int
}
```


#### type GetDataSourceSchemaOutput

```go
type GetDataSourceSchemaOutput struct {
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


#### type GetDestinationTypeListOutput

```go
type GetDestinationTypeListOutput struct {
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


#### type GetGroupMemberListInput

```go
type GetGroupMemberListInput struct {
	GroupId int
}
```


#### type GetGroupMemberListOutput

```go
type GetGroupMemberListOutput struct {
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


#### type GetJobInput

```go
type GetJobInput struct {
	JobId int
}
```


#### type GetJobOutput

```go
type GetJobOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetMyQueriesOutput

```go
type GetMyQueriesOutput struct {
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


#### type GetPublicDashboardInput

```go
type GetPublicDashboardInput struct {
	Token string
}
```


#### type GetPublicDashboardOutput

```go
type GetPublicDashboardOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryFavoriteListOutput

```go
type GetQueryFavoriteListOutput struct {
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


#### type GetQueryRecentOutput

```go
type GetQueryRecentOutput struct {
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


#### type GetQuerySearchOutput

```go
type GetQuerySearchOutput struct {
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


#### type GetQueryTagsOutput

```go
type GetQueryTagsOutput struct {
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


#### type PostQueryInput

```go
type PostQueryInput struct {
	QueryId      int    `json:"-"`
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}
```


#### type PostQueryListInput

```go
type PostQueryListInput struct {
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}
```


#### type PostQueryListOutput

```go
type PostQueryListOutput struct {
	QueryId    int    `json:"id"`
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
}
```


#### type PostQueryOutput

```go
type PostQueryOutput struct {
	Body       string
	StatusCode int
}
```
