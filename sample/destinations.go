package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---ListDestinations()---")
	output1 := client.ListDestinations(nil)
	fmt.Println(output1.StatusCode)
	fmt.Println(output1.Body)

	fmt.Println("---CreateDestination()---")
	input2 := &redash.CreateDestinationInput{
		Options: &redash.CreateDestinationInputOptions{
			Addresses: "foo@bar-baz.qux",
		},
		Name: "something",
		Type: "email",
	}
	output2 := client.CreateDestination(input2)
	fmt.Println(output2.StatusCode)
	fmt.Println(output2.Body)
	destinationId := output2.Id

	fmt.Println("---ListDestinationsTypes()---")
	output3 := client.ListDestinationsTypes(nil)
	fmt.Println(output3.StatusCode)
	fmt.Println(output3.Body)

	fmt.Println("---GetDestination()---")
	input4 := &redash.GetDestinationInput{
		DestinationId: destinationId,
	}
	output4 := client.GetDestination(input4)
	fmt.Println(output4.StatusCode)
	fmt.Println(output4.Body)

	fmt.Println("---UpdateDestination()---")
	input5 := &redash.UpdateDestinationInput{
		DestinationId: destinationId,
		Options: &redash.UpdateDestinationInputOptions{
			Addresses: "foo@bar-baz.qux",
		},
		Name: "somethingsomething",
		Type: "email",
	}
	output5 := client.UpdateDestination(input5)
	fmt.Println(output5.StatusCode)
	fmt.Println(output5.Body)

	fmt.Println("---DeleteDestination()---")
	input6 := &redash.DeleteDestinationInput{
		DestinationId: destinationId,
	}
	output6 := client.DeleteDestination(input6)
	fmt.Println(output6.StatusCode)
	fmt.Println(output6.Body)
}
