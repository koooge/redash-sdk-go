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

	fmt.Println("---ListAlerts()---")
	output := client.ListAlerts(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetAlert()---")
	input2 := &redash.GetAlertInput{
		AlertId: alertId,
	}
	output2 := client.GetAlert(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---ListAlertSubscriptions()---")
	input3 := &redash.ListAlertSubscriptionsInput{
		AlertId: alertId,
	}
	output3 := client.ListAlertSubscriptions(input3)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
