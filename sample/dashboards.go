package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const dashboardSlug = "something"

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	output := client.GetDashboardList()
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	input2 := &redash.GetDashboardInput{
		DashboardSlug: dashboardSlug,
	}
	output2 := client.GetDashboard(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

}
