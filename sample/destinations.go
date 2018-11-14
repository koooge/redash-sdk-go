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

	fmt.Println("---GetDestination()---")
	input := &redash.GetDestinationInput{
		DestinationId: destinationId,
	}

	output := client.GetDestination(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetDestinationList()---")
	output2 := client.GetDestinationList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetDestinationTypeList()---")
	output3 := client.GetDestinationTypeList()
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
