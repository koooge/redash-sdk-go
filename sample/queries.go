package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

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
}
