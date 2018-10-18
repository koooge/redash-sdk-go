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

	input := &redash.GetQueryInput{
		QueryId: queryId,
	}

	output := client.GetQuery(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	output2 := client.GetQueryList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
