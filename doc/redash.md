# redash
--
    import "github.com/koooge/redash-sdk-go/redash"


## Usage

#### type Client

```go
type Client struct {
	Config *Config
}
```


#### func  NewClient

```go
func NewClient(config *Config) *Client
```

#### func (*Client) Delete

```go
func (c *Client) Delete(path string, body string) (*http.Response, error)
```

#### func (*Client) Get

```go
func (c *Client) Get(path string) (*http.Response, error)
```

#### func (*Client) GetDataSources

```go
func (c *Client) GetDataSources(input *GetDataSourcesInput) *GetDataSourcesOutput
```

#### func (*Client) Post

```go
func (c *Client) Post(path string, body string) (*http.Response, error)
```

#### func (*Client) Put

```go
func (c *Client) Put(path string, body string) (*http.Response, error)
```

#### func (*Client) Request

```go
func (c *Client) Request(httpMethod string, path string, body string) (*http.Response, error)
```

#### type Config

```go
type Config struct {
	EndpointUrl string
	ApiKey      string
}
```


#### type GetDataSourcesInput

```go
type GetDataSourcesInput struct {
	Id int
}
```


#### type GetDataSourcesOutput

```go
type GetDataSourcesOutput struct {
	Body       string
	StatusCode int
}
```
