package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const queryResultId = 271647

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	input := &redash.GetQueryResultsInput{
		QueryResultId: queryResultId,
	}

	output := client.GetQueryResults(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)
}
