package redash

type IClient interface {
	// client
	SetConfig(*Config)

	// alerts
	GetAlertList() *GetAlertListOutput
	GetAlert(*GetAlertInput) *GetAlertOutput
	GetAlertSubscriptionList(*GetAlertSubscriptionListInput) *GetAlertSubscriptionListOutput

	// dashboards
	GetDashboardList() *GetDashboardListOutput
	GetDashboard(*GetDashboardInput) *GetDashboardOutput
	GetPublicDashboard(*GetPublicDashboardInput) *GetPublicDashboardOutput
	GetDashboardTags() *GetDashboardTagsOutput

	// data_sources
	GetDataSourceList() *GetDataSourceListOutput
	GetDataSource(*GetDataSourceInput) *GetDataSourceOutput
	GetDataSourceSchema(*GetDataSourceSchemaInput) *GetDataSourceSchemaOutput

	// destinations
	GetDestinationList() *GetDestinationListOutput
	GetDestination(*GetDestinationInput) *GetDestinationOutput
	GetDestinationTypeList() *GetDestinationTypeListOutput

	// events
	GetEvents() *GetEventsOutput

	// favorites
	GetQueryFavoriteList() *GetQueryFavoriteListOutput
	GetDashboardFavoriteList() *GetDashboardFavoriteListOutput

	// groups
	GetGroupList() *GetGroupListOutput
	GetGroup(*GetGroupInput) *GetGroupOutput
	GetGroupMemberList(*GetGroupMemberListInput) *GetGroupMemberListOutput

	// queries
	GetQueryList() *GetQueryListOutput
	GetQuery(*GetQueryInput) *GetQueryOutput
	GetQuerySearch() *GetQuerySearchOutput
	GetQueryRecent() *GetQueryRecentOutput
	GetMyQueries() *GetMyQueriesOutput
	GetQueryTags() *GetQueryTagsOutput
	PostQueryList(*PostQueryListInput) *PostQueryListOutput
	PostQuery(*PostQueryInput) *PostQueryOutput
	DeleteQuery(*DeleteQueryInput) *DeleteQueryOutput

	// query_results
	GetQueryResult(*GetQueryResultInput) *GetQueryResultOutput
	GetJob(*GetJobInput) *GetJobOutput

	// query_snippets
	GetQuerySnippetList() *GetQuerySnippetListOutput
	GetQuerySnippet(*GetQuerySnippetInput) *GetQuerySnippetOutput

	// settings
	GetOrganizationSettings() *GetOrganizationSettingsOutput

	// users
	GetUserList() *GetUserListOutput
	GetUser(*GetUserInput) *GetUserOutput
}
