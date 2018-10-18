package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const groupId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	input := &redash.GetGroupInput{
		GroupId: groupId,
	}

	output := client.GetGroup(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	output2 := client.GetGroupList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
