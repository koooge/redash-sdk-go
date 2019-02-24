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

#### func (*Client) CreateDataSource

```go
func (c *Client) CreateDataSource(input *CreateDataSourceInput) *CreateDataSourceOutput
```

#### func (*Client) CreateQuery

```go
func (c *Client) CreateQuery(input *CreateQueryInput) *CreateQueryOutput
```

#### func (*Client) DeleteDataSource

```go
func (c *Client) DeleteDataSource(input *DeleteDataSourceInput) *DeleteDataSourceOutput
```

#### func (*Client) DeleteQuery

```go
func (c *Client) DeleteQuery(input *DeleteQueryInput) *DeleteQueryOutput
```

#### func (*Client) GetAlert

```go
func (c *Client) GetAlert(input *GetAlertInput) *GetAlertOutput
```

#### func (*Client) GetDashboard

```go
func (c *Client) GetDashboard(input *GetDashboardInput) *GetDashboardOutput
```

#### func (*Client) GetDashboardTags

```go
func (c *Client) GetDashboardTags(_ *GetDashboardTagsInput) *GetDashboardTagsOutput
```

#### func (*Client) GetDataSource

```go
func (c *Client) GetDataSource(input *GetDataSourceInput) *GetDataSourceOutput
```

#### func (*Client) GetDataSourceSchema

```go
func (c *Client) GetDataSourceSchema(input *GetDataSourceSchemaInput) *GetDataSourceSchemaOutput
```

#### func (*Client) GetDestination

```go
func (c *Client) GetDestination(input *GetDestinationInput) *GetDestinationOutput
```

#### func (*Client) GetEvents

```go
func (c *Client) GetEvents(_ *GetEventsInput) *GetEventsOutput
```

#### func (*Client) GetGroup

```go
func (c *Client) GetGroup(input *GetGroupInput) *GetGroupOutput
```

#### func (*Client) GetJob

```go
func (c *Client) GetJob(input *GetJobInput) *GetJobOutput
```

#### func (*Client) GetMyQueries

```go
func (c *Client) GetMyQueries(_ *GetMyQueriesInput) *GetMyQueriesOutput
```

#### func (*Client) GetOrganizationSettings

```go
func (c *Client) GetOrganizationSettings(_ *GetOrganizationSettingsInput) *GetOrganizationSettingsOutput
```

#### func (*Client) GetPublicDashboard

```go
func (c *Client) GetPublicDashboard(input *GetPublicDashboardInput) *GetPublicDashboardOutput
```

#### func (*Client) GetQuery

```go
func (c *Client) GetQuery(input *GetQueryInput) *GetQueryOutput
```

#### func (*Client) GetQueryRecent

```go
func (c *Client) GetQueryRecent(_ *GetQueryRecentInput) *GetQueryRecentOutput
```

#### func (*Client) GetQueryResult

```go
func (c *Client) GetQueryResult(input *GetQueryResultInput) *GetQueryResultOutput
```

#### func (*Client) GetQuerySearch

```go
func (c *Client) GetQuerySearch(_ *GetQuerySearchInput) *GetQuerySearchOutput
```

#### func (*Client) GetQuerySnippet

```go
func (c *Client) GetQuerySnippet(input *GetQuerySnippetInput) *GetQuerySnippetOutput
```

#### func (*Client) GetQueryTags

```go
func (c *Client) GetQueryTags(_ *GetQueryTagsInput) *GetQueryTagsOutput
```

#### func (*Client) GetUser

```go
func (c *Client) GetUser(input *GetUserInput) *GetUserOutput
```

#### func (*Client) ListAlertSubscriptions

```go
func (c *Client) ListAlertSubscriptions(input *ListAlertSubscriptionsInput) *ListAlertSubscriptionsOutput
```

#### func (*Client) ListAlerts

```go
func (c *Client) ListAlerts(_ *ListAlertsInput) *ListAlertsOutput
```

#### func (*Client) ListDashboardFavorites

```go
func (c *Client) ListDashboardFavorites(_ *ListDashboardFavoritesInput) *ListDashboardFavoritesOutput
```

#### func (*Client) ListDashboards

```go
func (c *Client) ListDashboards(_ *ListDashboardsInput) *ListDashboardsOutput
```

#### func (*Client) ListDataSources

```go
func (c *Client) ListDataSources(_ *ListDataSourcesInput) *ListDataSourcesOutput
```

#### func (*Client) ListDataSourcesTypes

```go
func (c *Client) ListDataSourcesTypes(_ *ListDataSourcesTypesInput) *ListDataSourcesTypesOutput
```

#### func (*Client) ListDestinationTypes

```go
func (c *Client) ListDestinationTypes(_ *ListDestinationTypesInput) *ListDestinationTypesOutput
```

#### func (*Client) ListDestinations

```go
func (c *Client) ListDestinations(_ *ListDestinationsInput) *ListDestinationsOutput
```

#### func (*Client) ListGroupMembers

```go
func (c *Client) ListGroupMembers(input *ListGroupMembersInput) *ListGroupMembersOutput
```

#### func (*Client) ListGroups

```go
func (c *Client) ListGroups(_ *ListGroupsInput) *ListGroupsOutput
```

#### func (*Client) ListQueries

```go
func (c *Client) ListQueries(_ *ListQueriesInput) *ListQueriesOutput
```

#### func (*Client) ListQueryFavorites

```go
func (c *Client) ListQueryFavorites(_ *ListQueryFavoritesInput) *ListQueryFavoritesOutput
```

#### func (*Client) ListQuerySnippets

```go
func (c *Client) ListQuerySnippets(_ *ListQuerySnippetsInput) *ListQuerySnippetsOutput
```

#### func (*Client) ListUsers

```go
func (c *Client) ListUsers(_ *ListUsersInput) *ListUsersOutput
```

#### func (*Client) ModifyQuery

```go
func (c *Client) ModifyQuery(input *ModifyQueryInput) *ModifyQueryOutput
```

#### func (*Client) PauseDataSource

```go
func (c *Client) PauseDataSource(input *PauseDataSourceInput) *PauseDataSourceOutput
```

#### func (*Client) SetConfig

```go
func (c *Client) SetConfig(config *Config)
```

#### func (*Client) TestDataSource

```go
func (c *Client) TestDataSource(input *TestDataSourceInput) *TestDataSourceOutput
```

#### func (*Client) UnpauseDataSource

```go
func (c *Client) UnpauseDataSource(input *UnpauseDataSourceInput) *UnpauseDataSourceOutput
```

#### func (*Client) UpdateDataSource

```go
func (c *Client) UpdateDataSource(input *UpdateDataSourceInput) *UpdateDataSourceOutput
```

#### type Config

```go
type Config struct {
	EndpointUrl string
	ApiKey      string
}
```


#### type CreateDataSourceInput

```go
type CreateDataSourceInput struct {
	Options *CreateDataSourceInputOptions `json:"options"`
	Name    string                        `json:"name"`
	Type    string                        `json:"type"`
}
```

POST /api/data_sources

#### type CreateDataSourceInputOptions

```go
type CreateDataSourceInputOptions struct {
	Dbname string `json:"dbname"`
}
```


#### type CreateDataSourceOutput

```go
type CreateDataSourceOutput struct {
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
	Id         int    `json:"id"`
}
```


#### type CreateQueryInput

```go
type CreateQueryInput struct {
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}
```


#### type CreateQueryOutput

```go
type CreateQueryOutput struct {
	QueryId    int    `json:"id"`
	Body       string `json:"-"`
	StatusCode int    `json:"-"`
}
```


#### type DeleteDataSourceInput

```go
type DeleteDataSourceInput struct {
	DataSourceId int
}
```

DELETE /api/data_sources/{data_source_id}

#### type DeleteDataSourceOutput

```go
type DeleteDataSourceOutput struct {
	Body       string
	StatusCode int
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


#### type GetAlertOutput

```go
type GetAlertOutput struct {
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


#### type GetDashboardOutput

```go
type GetDashboardOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetDashboardTagsInput

```go
type GetDashboardTagsInput struct {
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

GET /api/data_sources/{data_source_id}

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

GET /api/data_sources/{data_source_id}/schema

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


#### type GetDestinationOutput

```go
type GetDestinationOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetEventsInput

```go
type GetEventsInput struct {
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


#### type GetMyQueriesInput

```go
type GetMyQueriesInput struct {
}
```


#### type GetMyQueriesOutput

```go
type GetMyQueriesOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetOrganizationSettingsInput

```go
type GetOrganizationSettingsInput struct {
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


#### type GetQueryInput

```go
type GetQueryInput struct {
	QueryId int
}
```


#### type GetQueryOutput

```go
type GetQueryOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryRecentInput

```go
type GetQueryRecentInput struct {
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


#### type GetQuerySearchInput

```go
type GetQuerySearchInput struct {
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


#### type GetQuerySnippetOutput

```go
type GetQuerySnippetOutput struct {
	Body       string
	StatusCode int
}
```


#### type GetQueryTagsInput

```go
type GetQueryTagsInput struct {
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


#### type GetUserOutput

```go
type GetUserOutput struct {
	Body       string
	StatusCode int
}
```


#### type IClient

```go
type IClient interface {
	// client
	SetConfig(*Config)

	// alerts
	ListAlerts(*ListAlertsInput) *ListAlertsOutput
	GetAlert(*GetAlertInput) *GetAlertOutput
	ListAlertSubscriptions(*ListAlertSubscriptionsInput) *ListAlertSubscriptionsOutput

	// dashboards
	ListDashboards(*ListDashboardsInput) *ListDashboardsOutput
	GetDashboard(*GetDashboardInput) *GetDashboardOutput
	GetPublicDashboard(*GetPublicDashboardInput) *GetPublicDashboardOutput
	GetDashboardTags(*GetDashboardTagsInput) *GetDashboardTagsOutput

	// data_sources
	ListDataSources(*ListDataSourcesInput) *ListDataSourcesOutput
	CreateDataSource(*CreateDataSourceInput) *CreateDataSourceOutput
	ListDataSourcesTypes(*ListDataSourcesTypesInput) *ListDataSourcesTypesOutput
	GetDataSource(*GetDataSourceInput) *GetDataSourceOutput
	UpdateDataSource(*UpdateDataSourceInput) *UpdateDataSourceOutput
	DeleteDataSource(*DeleteDataSourceInput) *DeleteDataSourceOutput
	GetDataSourceSchema(*GetDataSourceSchemaInput) *GetDataSourceSchemaOutput
	PauseDataSource(*PauseDataSourceInput) *PauseDataSourceOutput
	UnpauseDataSource(*UnpauseDataSourceInput) *UnpauseDataSourceOutput
	TestDataSource(*TestDataSourceInput) *TestDataSourceOutput

	// destinations
	ListDestinations(*ListDestinationsInput) *ListDestinationsOutput
	GetDestination(*GetDestinationInput) *GetDestinationOutput
	ListDestinationTypes(*ListDestinationTypesInput) *ListDestinationTypesOutput

	// events
	GetEvents(*GetEventsInput) *GetEventsOutput

	// favorites
	ListQueryFavorites(*ListQueryFavoritesInput) *ListQueryFavoritesOutput
	ListDashboardFavorites(*ListDashboardFavoritesInput) *ListDashboardFavoritesOutput

	// groups
	ListGroups(*ListGroupsInput) *ListGroupsOutput
	GetGroup(*GetGroupInput) *GetGroupOutput
	ListGroupMembers(*ListGroupMembersInput) *ListGroupMembersOutput

	// queries
	ListQueries(*ListQueriesInput) *ListQueriesOutput
	GetQuery(*GetQueryInput) *GetQueryOutput
	GetQuerySearch(*GetQuerySearchInput) *GetQuerySearchOutput
	GetQueryRecent(*GetQueryRecentInput) *GetQueryRecentOutput
	GetMyQueries(*GetMyQueriesInput) *GetMyQueriesOutput
	GetQueryTags(*GetQueryTagsInput) *GetQueryTagsOutput
	CreateQuery(*CreateQueryInput) *CreateQueryOutput
	ModifyQuery(*ModifyQueryInput) *ModifyQueryOutput
	DeleteQuery(*DeleteQueryInput) *DeleteQueryOutput

	// query_results
	GetQueryResult(*GetQueryResultInput) *GetQueryResultOutput
	GetJob(*GetJobInput) *GetJobOutput

	// query_snippets
	ListQuerySnippets(*ListQuerySnippetsInput) *ListQuerySnippetsOutput
	GetQuerySnippet(*GetQuerySnippetInput) *GetQuerySnippetOutput

	// settings
	GetOrganizationSettings(*GetOrganizationSettingsInput) *GetOrganizationSettingsOutput

	// users
	ListUsers(*ListUsersInput) *ListUsersOutput
	GetUser(*GetUserInput) *GetUserOutput
}
```


#### type ListAlertSubscriptionsInput

```go
type ListAlertSubscriptionsInput struct {
	AlertId int
}
```


#### type ListAlertSubscriptionsOutput

```go
type ListAlertSubscriptionsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListAlertsInput

```go
type ListAlertsInput struct {
}
```


#### type ListAlertsOutput

```go
type ListAlertsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDashboardFavoritesInput

```go
type ListDashboardFavoritesInput struct {
}
```


#### type ListDashboardFavoritesOutput

```go
type ListDashboardFavoritesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDashboardsInput

```go
type ListDashboardsInput struct {
}
```


#### type ListDashboardsOutput

```go
type ListDashboardsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDataSourcesInput

```go
type ListDataSourcesInput struct {
}
```

GET /api/data_sources

#### type ListDataSourcesOutput

```go
type ListDataSourcesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDataSourcesTypesInput

```go
type ListDataSourcesTypesInput struct {
}
```

GET /api/data_sources/types

#### type ListDataSourcesTypesOutput

```go
type ListDataSourcesTypesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDestinationTypesInput

```go
type ListDestinationTypesInput struct {
}
```


#### type ListDestinationTypesOutput

```go
type ListDestinationTypesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListDestinationsInput

```go
type ListDestinationsInput struct {
}
```


#### type ListDestinationsOutput

```go
type ListDestinationsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListGroupMembersInput

```go
type ListGroupMembersInput struct {
	GroupId int
}
```


#### type ListGroupMembersOutput

```go
type ListGroupMembersOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListGroupsInput

```go
type ListGroupsInput struct {
}
```


#### type ListGroupsOutput

```go
type ListGroupsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListQueriesInput

```go
type ListQueriesInput struct {
}
```


#### type ListQueriesOutput

```go
type ListQueriesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListQueryFavoritesInput

```go
type ListQueryFavoritesInput struct {
}
```


#### type ListQueryFavoritesOutput

```go
type ListQueryFavoritesOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListQuerySnippetsInput

```go
type ListQuerySnippetsInput struct {
}
```


#### type ListQuerySnippetsOutput

```go
type ListQuerySnippetsOutput struct {
	Body       string
	StatusCode int
}
```


#### type ListUsersInput

```go
type ListUsersInput struct {
}
```


#### type ListUsersOutput

```go
type ListUsersOutput struct {
	Body       string
	StatusCode int
}
```


#### type ModifyQueryInput

```go
type ModifyQueryInput struct {
	QueryId      int    `json:"-"`
	DataSourceId int    `json:"data_source_id"`
	Query        string `json:"query"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Schedule     string `json:"schedule,omitempty"`
}
```


#### type ModifyQueryOutput

```go
type ModifyQueryOutput struct {
	Body       string
	StatusCode int
}
```


#### type PauseDataSourceInput

```go
type PauseDataSourceInput struct {
	DataSourceId int    `json:"-"`
	Reason       string `json:"reason,omitempty"`
}
```

POST /api/data_sources/{data_source_id}/pause

#### type PauseDataSourceOutput

```go
type PauseDataSourceOutput struct {
	Body       string
	StatusCode int
}
```


#### type TestDataSourceInput

```go
type TestDataSourceInput struct {
	DataSourceId int
}
```

POST /api/data_sources/{data_source_id}/test

#### type TestDataSourceOutput

```go
type TestDataSourceOutput struct {
	Body       string
	StatusCode int
}
```


#### type UnpauseDataSourceInput

```go
type UnpauseDataSourceInput struct {
	DataSourceId int
}
```

DELETE /api/data_sources/{data_source_id}/pause

#### type UnpauseDataSourceOutput

```go
type UnpauseDataSourceOutput struct {
	Body       string
	StatusCode int
}
```


#### type UpdateDataSourceInput

```go
type UpdateDataSourceInput struct {
	DataSourceId int                           `json:"-"`
	Options      *UpdateDataSourceInputOptions `json:"options"`
	Name         string                        `json:"name"`
	Type         string                        `json:"type"`
}
```

POST /api/data_sources/{data_source_id}

#### type UpdateDataSourceInputOptions

```go
type UpdateDataSourceInputOptions struct {
	Dbname string `json:"dbname"`
}
```


#### type UpdateDataSourceOutput

```go
type UpdateDataSourceOutput struct {
	Body       string
	StatusCode int
}
```
