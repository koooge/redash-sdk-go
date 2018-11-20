package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const destinationId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---ListDestinations()---")
	output := client.ListDestinations(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetDestination()---")
	input2 := &redash.GetDestinationInput{
		DestinationId: destinationId,
	}
	output2 := client.GetDestination(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetDestinationTypeList()---")
	output3 := client.ListDestinationTypes(nil)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
