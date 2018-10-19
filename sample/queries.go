package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/koooge/redash-sdk-go/redash"
)

const dataSourceId = 16 // Query Results
const queryId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	// GetQuery()
	input := &redash.GetQueryInput{
		QueryId: queryId,
	}

	output := client.GetQuery(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	// GetQueryList()
	output2 := client.GetQueryList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	// GetQuerySearch()
	output3 := client.GetQuerySearch()
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)

	// GetQueryRecent()
	output4 := client.GetQueryRecent()
	fmt.Println(output4.Body)
	fmt.Println(output4.StatusCode)

	// GetMyQueries()
	output5 := client.GetMyQueries()
	fmt.Println(output5.Body)
	fmt.Println(output5.StatusCode)

	// GetQueryTags()
	output6 := client.GetQueryTags()
	fmt.Println(output6.Body)
	fmt.Println(output6.StatusCode)

	// PostQueryList()
	input7 := &redash.PostQueryListInput{
		DataSourceId: dataSourceId,
		Query:        "SELECT * FROM sql_" + strconv.Itoa(queryId) + ";",
		Name:         "sample PostQueryList",
	}

	output7 := client.PostQueryList(input7)
	fmt.Println(output7.Body)
	fmt.Println(output7.StatusCode)
	queryId7 := output7.QueryId

	// PostQuery()
	input8 := &redash.PostQueryInput{
		QueryId:      queryId7,
		DataSourceId: dataSourceId,
		Query:        "SELECT * FROM sql_" + strconv.Itoa(queryId) + "|0;",
		Name:         "sample PostQuery",
	}

	output8 := client.PostQuery(input8)
	fmt.Println(output8.Body)
	fmt.Println(output8.StatusCode)

	// DeleteQuery()
	input9 := &redash.DeleteQueryInput{
		QueryId: queryId7,
	}

	output9 := client.DeleteQuery(input9)
	fmt.Println(output9.Body)
	fmt.Println(output9.StatusCode)
}
