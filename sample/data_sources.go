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

	fmt.Println("---ListDataSources()---")
	output1 := client.ListDataSources(nil)
	fmt.Println(output1.StatusCode)
	fmt.Println(output1.Body)

	fmt.Println("---CreateDataSource()---")
	input2 := &redash.CreateDataSourceInput{
		Options: &redash.CreateDataSourceInputOptions{
			Dbname: "foobar",
		},
		Name: "CreateDataSourceSample",
		Type: "pg",
	}
	output2 := client.CreateDataSource(input2)
	fmt.Println(output2.StatusCode)
	fmt.Println(output2.Body)
	dataSourceId := output2.Id

	fmt.Println("---ListDataSourcesTypes()---")
	output3 := client.ListDataSourcesTypes(nil)
	fmt.Println(output3.StatusCode)
	fmt.Println(output3.Body)

	fmt.Println("---GetDataSource()---")
	input4 := &redash.GetDataSourceInput{
		DataSourceId: dataSourceId,
	}
	output4 := client.GetDataSource(input4)
	fmt.Println(output4.StatusCode)
	fmt.Println(output4.Body)

	fmt.Println("---UpdateDataSource()---")
	input5 := &redash.UpdateDataSourceInput{
		DataSourceId: dataSourceId,
		Options: &redash.UpdateDataSourceInputOptions{
			Dbname: "hogefuga",
		},
		Name: "UpdateDataSourceSample",
		Type: "pg",
	}
	output5 := client.UpdateDataSource(input5)
	fmt.Println(output5.StatusCode)
	fmt.Println(output5.Body)

	fmt.Println("---GetDataSourceSchema()---")
	input7 := &redash.GetDataSourceSchemaInput{
		DataSourceId: dataSourceId,
	}
	output7 := client.GetDataSourceSchema(input7)
	fmt.Println(output7.StatusCode)
	fmt.Println(output7.Body)

	// fmt.Println("---PauseDataSource()---")
	// input8 := &redash.PauseDataSourceInput{
	// 	DataSourceId: dataSourceId,
	// }
	// output8 := client.PauseDataSource(input8)
	// fmt.Println(output8.StatusCode)
	// fmt.Println(output8.Body)

	// fmt.Println("---UnpauseDataSource()---")
	// input9 := &redash.UnpauseDataSourceInput{
	// 	DataSourceId: dataSourceId,
	// }
	// output9 := client.UnpauseDataSource(input9)
	// fmt.Println(output9.StatusCode)
	// fmt.Println(output9.Body)

	// fmt.Println("---UnpauseDataSource()---")
	// input10 := &redash.TestDataSourceInput{
	// 	DataSourceId: dataSourceId,
	// }
	// output10 := client.TestDataSource(input10)
	// fmt.Println(output10.StatusCode)
	// fmt.Println(output10.Body)

	fmt.Println("---DeleteDataSource()---")
	input6 := &redash.DeleteDataSourceInput{
		DataSourceId: dataSourceId,
	}
	output6 := client.DeleteDataSource(input6)
	fmt.Println(output6.StatusCode)
	fmt.Println(output6.Body)
}
