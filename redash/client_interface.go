package redash

type IClient interface {
	// client
	SetConfig(*Config)

	// alerts
	ListAlerts(*ListAlertsInput) *ListAlertOutput
	GetAlert(*GetAlertInput) *GetAlertOutput
	ListAlertSubscriptions(*ListAlertSubscriptionsInput) *ListAlertSubscriptionsOutput

	// dashboards
	ListDashboards(*ListDashboardsInput) *ListDashboardsOutput
	GetDashboard(*GetDashboardInput) *GetDashboardOutput
	GetPublicDashboard(*GetPublicDashboardInput) *GetPublicDashboardOutput
	GetDashboardTags(*GetDashboardTagsInput) *GetDashboardTagsOutput

	// data_sources
	ListDataSources(*ListDataSourcesInput) *ListDataSourcesOutput
	GetDataSource(*GetDataSourceInput) *GetDataSourceOutput
	GetDataSourceSchema(*GetDataSourceSchemaInput) *GetDataSourceSchemaOutput

	// destinations
	ListDestinations(*ListDestinationsInput) *ListDestinationOutput
	GetDestination(*GetDestinationInput) *GetDestinationOutput
	ListDestinationTypes(*ListDestinationTypesInput) *ListDestinationTypesOutput

	// events
	GetEvents(*GetEventsInput) *GetEventsOutput

	// favorites
	ListQueryFavorites(*ListQueryFavoritesInput) *ListQueryFavoritesOutput
	ListGetDashboardFavorites(*ListGetDashboardFavoritesInput) *ListDashboardFavoritesOutput

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
	PostQuery(*PostQueryInput) *PostQueryOutput
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
	ListUsers(*ListUsersInput) *GetUserListOutput
	GetUser(*GetUserInput) *GetUserOutput
}
