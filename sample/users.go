package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const userId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---GetUser()---")
	input := &redash.GetUserInput{
		UserId: userId,
	}
	output := client.GetUser(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetUserList()---")
	output2 := client.GetUserList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
