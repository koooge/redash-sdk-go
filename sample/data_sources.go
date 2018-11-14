package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const dataSourceId = 3

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	fmt.Println("---GetDataSource()---")
	input := &redash.GetDataSourceInput{
		DataSourceId: dataSourceId,
	}

	output := client.GetDataSource(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)

	fmt.Println("---GetDataSourceList()---")
	output2 := client.GetDataSourceList()
	fmt.Println(output2.Body)
	fmt.Println(output2.StatusCode)

	fmt.Println("---GetDataSourceSchema()---")
	input3 := &redash.GetDataSourceSchemaInput{
		DataSourceId: dataSourceId,
	}

	output3 := client.GetDataSourceSchema(input3)
	fmt.Println(output3.Body)
	fmt.Println(output3.StatusCode)
}
