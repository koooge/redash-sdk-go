package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const alertId = 1
const subscriberId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---GetAlert()---")
	input := &redash.GetAlertInput{
		AlertId: alertId,
	}

	output := client.GetAlert(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetAlertList()---")
	output2 := client.GetAlertList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetAlertSubscriptionList()---")
	input3 := &redash.GetAlertSubscriptionListInput{
		AlertId: alertId,
	}

	output3 := client.GetAlertSubscriptionList(input3)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
