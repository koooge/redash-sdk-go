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

	input := &redash.GetAlertsInput{
		AlertId: alertId,
	}

	output := client.GetAlerts(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)
}
