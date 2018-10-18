package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const alertId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	input := &redash.GetAlertInput{
		AlertId: alertId,
	}

	output := client.GetAlert(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	output2 := client.GetAlertList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
