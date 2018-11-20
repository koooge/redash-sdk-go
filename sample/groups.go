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

	fmt.Println("---ListGroups()---")
	output := client.ListGroups(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetGroup()---")
	input2 := &redash.GetGroupInput{
		GroupId: groupId,
	}
	output2 := client.GetGroup(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---ListGroupMembers()---")
	input3 := &redash.ListGroupMembersInput{
		GroupId: groupId,
	}
	output3 := client.ListGroupMembers(input3)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
