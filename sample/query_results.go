package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const queryResultId = 271647
const jobId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---GetQueryResult()---")
	input := &redash.GetQueryResultInput{
		QueryResultId: queryResultId,
	}

	output := client.GetQueryResult(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetJob()---")
	input2 := &redash.GetJobInput{
		JobId: jobId,
	}

	output2 := client.GetJob(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
