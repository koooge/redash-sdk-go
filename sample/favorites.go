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

	fmt.Println("---ListQueryFavorites()---")
	output := client.ListQueryFavorites(nil)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---ListDashboardFavorites()---")
	output2 := client.ListDashboardFavorites(nil)
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)
}
