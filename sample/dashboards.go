package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const dashboardSlug = "something"
const token = "somethingToken"

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---ListDashboards()---")
	output := client.ListDashboards(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetDashboard()---")
	input2 := &redash.GetDashboardInput{
		DashboardSlug: dashboardSlug,
	}
	output2 := client.GetDashboard(input2)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetPublicDashboard()---")
	input3 := &redash.GetPublicDashboardInput{
		Token: token,
	}
	output3 := client.GetPublicDashboard(input3)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)

	fmt.Println("---GetDashboardTags()---")
	output4 := client.GetDashboardTags(nil)
	fmt.Println(output4.Body)
	fmt.Println(output4.StatusCode)
}
