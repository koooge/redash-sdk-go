package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const querySnippetId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---ListQuerySnippets()---")
	output := client.ListQuerySnippets(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetQuerySnippet()---")
	input2 := &redash.GetQuerySnippetInput{
		QuerySnippetId: querySnippetId,
	}
	output2 := client.GetQuerySnippet(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
