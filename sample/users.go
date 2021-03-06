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

	fmt.Println("---ListUsers()---")
	output := client.ListUsers(nil)
	fmt.Println(output.StatusCode)
	fmt.Println(output.Body)

	fmt.Println("---CreateUser()---")

	fmt.Println("---GetUser()---")
	input2 := &redash.GetUserInput{
		UserId: userId,
	}
	output2 := client.GetUser(input2)
	fmt.Println(output2.StatusCode)
	fmt.Println(output2.Body)

	fmt.Println("---UpdateUser()---")
	fmt.Println("---InviteUser()---")
	fmt.Println("---ResetPassword()---")
	fmt.Println("---GenerateApiKey()---")
	fmt.Println("---DisableUser()---")
	fmt.Println("---EnableUser()---")

	fmt.Println("---DeleteUser()---")
}
