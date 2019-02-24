[![CircleCI](https://circleci.com/gh/koooge/redash-sdk-go/tree/master.svg?style=svg)](https://circleci.com/gh/koooge/redash-sdk-go/tree/master)
[![Build status](https://ci.appveyor.com/api/projects/status/hv0ofxbhjf5flstm/branch/master?svg=true)](https://ci.appveyor.com/project/koooge/redash-sdk-go/branch/master)
[![Maintainability](https://api.codeclimate.com/v1/badges/87203e61bd46d720e1d1/maintainability)](https://codeclimate.com/github/koooge/redash-sdk-go/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/87203e61bd46d720e1d1/test_coverage)](https://codeclimate.com/github/koooge/redash-sdk-go/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/koooge/redash-sdk-go)](https://goreportcard.com/report/github.com/koooge/redash-sdk-go)
[![GoDoc](https://godoc.org/github.com/koooge/redash-sdk-go?status.svg)](https://godoc.org/github.com/koooge/redash-sdk-go/redash)

# redash-sdk-go
Redash client in go. Many APIs are still WIP.

## Prerequisites
- go1.11+

## Usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const queryId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	input := &redash.GetQueryInput{
		QueryId: queryId,
	}

	output := client.GetQuery(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)
}
```

See more about [sample](https://github.com/koooge/redash-sdk-go/tree/master/sample) and [doc](https://github.com/koooge/redash-sdk-go/blob/master/doc/redash.md).
