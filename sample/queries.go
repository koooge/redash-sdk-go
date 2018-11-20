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

	fmt.Println("---ListQueries()---")
	output := client.ListQueries(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetQuery()---")
	input2 := &redash.GetQueryInput{
		QueryId: queryId,
	}
	output2 := client.GetQuery(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetQuerySearch()---")
	output3 := client.GetQuerySearch(nil)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)

	fmt.Println("---GetQueryRecent()---")
	output4 := client.GetQueryRecent(nil)
	fmt.Println(output4.Body)
	fmt.Println(output4.StatusCode)

	fmt.Println("---GetMyQueries()---")
	output5 := client.GetMyQueries(nil)
	fmt.Println(output5.Body)
	fmt.Println(output5.StatusCode)

	fmt.Println("---GetQueryTags()---")
	output6 := client.GetQueryTags(nil)
	fmt.Println(output6.Body)
	fmt.Println(output6.StatusCode)

	fmt.Println("---CreateQuery()---")
	input7 := &redash.CreateQueryInput{
		DataSourceId: dataSourceId,
		Query:        "SELECT * FROM sql_" + strconv.Itoa(queryId) + ";",
		Name:         "sample PostQueryList",
	}
	output7 := client.CreateQuery(input7)
	fmt.Println(output7.Body)
	fmt.Println(output7.StatusCode)
	queryId7 := output7.QueryId

	fmt.Println("---ModifyQuery()---")
	input8 := &redash.ModifyQueryInput{
		QueryId:      queryId7,
		DataSourceId: dataSourceId,
		Query:        "SELECT * FROM sql_" + strconv.Itoa(queryId) + "|0;",
		Name:         "sample PostQuery",
	}
	output8 := client.ModifyQuery(input8)
	fmt.Println(output8.Body)
	fmt.Println(output8.StatusCode)

	fmt.Println("---DeleteQuery()---")
	input9 := &redash.DeleteQueryInput{
		QueryId: queryId7,
	}
	output9 := client.DeleteQuery(input9)
	fmt.Println(output9.Body)
	fmt.Println(output9.StatusCode)
}
