# Go Client For Nutanix Lifecycle Management APIs

The Go client for Nutanix Lifecycle Management APIs is designed for Go client application developers offering them simple and flexible access to APIs that manage Infrastructure, Software and Firmware Upgrades.

## Features
- Invoke Nutanix APIs with a simple interface.
- Handle Authentication seamlessly.
- Reduce boilerplate code implementation.
- Use standard methods for installation.

## Version
- API version: v4.1
- Package version: v4.1.1

## Requirements.
Go 1.17 or above are fully supported and tested.


## Installation & Usage

### Installation

#### Using go get

##### Install the latest version

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/...
```

##### Install a specific version

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/...@v4.1.1
```

#### Using go modules

##### Install the latest version

In go.mod file add the following

```
module your-module

go {GO_VERSION}

require (
	github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4 latest
)
```

##### Install a specific version

In go.mod file add the following

```
module your-module

go {GO_VERSION}

require (
	github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4 v4.1.1
)
```


## Configuration
The Go client for Nutanix Lifecycle Management APIs can be configured with the following parameters

| Parameter | Description                                                                      | Required | Default Value|
|-----------|----------------------------------------------------------------------------------|----------|--------------|
| Scheme    | URI scheme for connecting to the cluster (HTTP or HTTPS using SSL/TLS)           | No       | https        |
| Host      | IPv4/IPv6 address or FQDN of the cluster to which the client will connect to     | Yes      | N/A          |
| Port      | Port on the cluster to which the client will connect to                          | No       | 9440         |
| Username  | Username to connect to a cluster                                                 | Yes      | N/A          |
| Password  | Password to connect to a cluster                                                 | Yes      | N/A          |
| Debug     | Runs the client in debug mode if specified                                       | No       | False        |
| VerifySSL | Verify SSL certificate of cluster, the client will connect to                    | No       | True         |
| Proxy     | Configure a proxy, the client will connect to                                    | No       | N/A          |
| MaxRetryAttempts| Maximum number of retry attempts while connecting to the cluster           | No       | 5            |
| RetryInterval | Interval (in time.Duration) at which retry attempts are made                  | No       | 3 * time.Second |
| LoggerFile | File location to which debug logs are written to                                | No       | N/A          |
| ConnectTimeout | Connection timeout (in time.Duration) for all operations                    | No       | 30 * time.Second |
| ReadTimeout | Read timeout (in time.Duration) for all operations                             | No       | 30 * time.Second |
| DownloadDirectory | Directory location for files to download                                 | No       | Current Directory |
| DownloadChunkSize | Chunk size in bytes for files to download                                | No       | 8*1024 bytes |

A Proxy can be configured with the following parameters

| Parameter      | Description                                                            | Required | Default Value|
|----------------|------------------------------------------------------------------------|----------|--------------|
| Proxy.Scheme   | URI Scheme for connecting to the proxy ("http", "https" or "socks5")   | Yes      | N/A          |
| Proxy.Host     | Host of the proxy to which the client will connect to                  | Yes      | N/A          |
| Proxy.Port     | Port of the proxy to which the client will connect to                  | Yes      | N/A          |
| Proxy.Username | Username to connect to the proxy                                       | -        | N/A          |
| Proxy.Password | Password to connect to the proxy                                       | -        | N/A          |

### Sample Configuration
```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
)
var (
	ApiClientInstance *client.ApiClient
)

ApiClientInstance = client.NewApiClient()
ApiClientInstance.Host = "10.19.50.27" // IPv4/IPv6 address or FQDN of the cluster
ApiClientInstance.Port = 9440 // Port to which to connect to
ApiClientInstance.Username = "admin" // UserName to connect to the cluster
ApiClientInstance.Password = "password" // Password to connect to the cluster

```

### Proxy Configuration
```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
)
var (
	ApiClientInstance *client.ApiClient
)

ApiClientInstance = client.NewApiClient()
// Configure the client as shown in the previous step
// ...

ApiClientInstance.Proxy = new(client.Proxy)

ApiClientInstance.Proxy.Scheme = "socks5"
ApiClientInstance.Proxy.Username = "proxy_admin"
ApiClientInstance.Proxy.Password = "proxy_password"
ApiClientInstance.Proxy.Host = "127.0.0.1"
ApiClientInstance.Proxy.Port = 1080

```



### Authentication
Nutanix APIs currently support two type of authentication schemes:

- **HTTP Basic Authentication**
      - The Go client can be configured using the username and password parameters to send Basic headers along with every request.
- **API Key Authentication**
      - The Go client can be configured to set an API key to send "**X-ntnx-api-key**" header with every request.
 ```go
  import (
  	"github.com/nutanix-core/ntnx-api-go-sdk-internal/iam-api-external-go-client/v16/client"
  )

  var (
  	ApiClientInstance *client.ApiClient
  )

  ApiClientInstance = client.NewApiClient()
  ApiClientInstance.SetApiKey("abcde12345")
  ```

### Retry Mechanism
The client can be configured to retry requests that fail with the following status codes. The numbers of seconds before which the next retry is attempted is determined by the retryInterval:

- [408 - Request Timeout](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408)
- [429 - Too Many Requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)
- [502 - Bad Gateway](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502)
- [503 - Service Unavailable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

The client will also redirect requests that fail with [302 - Found](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302) to the new location specified in the response header `Location`.
! Note : Within Golang SDK maximum redirect attempts are limited to 10 by default and cannot be changed tro a custom value.

```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
    "time"
)
var (
	ApiClientInstance *client.ApiClient
)

ApiClientInstance = client.NewApiClient()
ApiClientInstance.MaxRetryAttempts = 5 // Max retry attempts while reconnecting on a loss of connection
ApiClientInstance.RetryInterval = 5 * time.Second // Interval (in time.Duration) to use during retry attempts - integer values are treated as time.Nanosecond by default
```

## Usage

### Invoking an operation
```go
// The following sample code is an example and does not reflect the real APIs provided by this client.
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/api"
)

var (
	ApiClientInstance *client.ApiClient
	SampleApiInstance *api.SampleApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client as shown in the previous step
// ...

// Initialize the API
SampleApiInstance = api.SampleApi(ApiClientInstance)
var extId string = '8a17d0bb-3147-4f3a-bbbd-48ad2a4c19fc' // UUID.

// Get sample entity by ID
response, err := SampleApiInstance.GetSampleEntityById(&extId)
if err != nil {
....
}
```

### Request Options
The library provides the ability to specify additional options that can be applied directly on the 'ApiClient' object used to make network calls to the API. The library also provides a mechanism to specify operation specific headers.
#### Client headers
The 'ApiClient' can be configured to send additional headers on each request.

```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
)

var (
	ApiClientInstance *client.ApiClient
)

ApiClientInstance = client.NewApiClient()
ApiClientInstance.AddDefaultHeader("Accept-Encoding", "gzip, deflate, br")
```
You can also modify the headers sent with each individual operation:

#### Operation specific headers
Nutanix APIs require that concurrent updates are protected using [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) headers. This would mean that the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header received in the response of a fetch (GET) operation should be used as an [If-Match](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match) header for the modification (PUT) operation.
```go
// The following sample code is an example and does not reflect the real APIs provided by this client.
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/api"
)

var (
	ApiClientInstance *client.ApiClient
	SampleApiInstance *api.SampleApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client as shown in a previous step
// ...

// Initialize the API
SampleApiInstance = api.SampleApi(ApiClientInstance)
var extId string = '8a17d0bb-3147-4f3a-bbbd-48ad2a4c19fc' // UUID.

// Get sample entity by ID
response, err := SampleApiInstance.GetSampleEntityById(&extId)
if err != nil {
....
}

// Extract E-Tag Header
etagValue := ApiClientInstance.GetEtag(response)
    
// The following sample code is an example and does not reflect the real APIs provided by this client.

// Update sample entity by ID
args := make(map[string] interface {})
args["If-Match"] = etagValue
// The body parameter in the following operation is received from the previous GET request's response which needs to be updated.
response, err := SampleApiInstance.UpdateSampleEntityById(&body, &extId, args)
if err != nil {
    ....
}
```

### List Operations
List Operations for Nutanix APIs support pagination, filtering, sorting and projections. The table below details the parameters that can be used to set the options for pagination etc.

| Parameter | Description
|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| _page     | specifies the page number of the result set. Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.|
| _limit    | specifies the total number of records returned in the result set. Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set|
| _filter   | allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the [OData V4.01 URL](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter) conventions. |
| _orderby  | allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. |
| _select   | allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource will be returned. |
| _expand   | allows clients to request related resources when a resource that satisfies a particular request is retrieved. Each expanded item is evaluated relative to the entity containing the property being expanded. Other query options can be applied to an expanded property by appending a semicolon-separated list of query options, enclosed in parentheses, to the property name. Permissible system query options are $filter,$select and $orderby. |

```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/api"
)
var (
	ApiClientInstance *client.ApiClient
	BundlesApiInstance *api.BundlesApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client
// ...

// Initialize the API
BundlesApiInstance = api.NewBundlesApi(ApiClientInstance)
page_ := 0
limit_ := 50
filter_ := "string_sample_data"
orderby_ := "string_sample_data"
select_ := "string_sample_data"

// 
response, err := BundlesApiInstance.ListBundles(&page_, &limit_, &filter_, &orderby_, &select_)
if err != nil {
    ....
}


```
The list of filterable and sortable fields with expansion keys can be found in the documentation [here](https://developers.nutanix.com/).

## API Reference

This library has a full set of [API Reference Documentation](https://developers.nutanix.com/sdk-reference?namespace=lifecycle&version=v4.1&language=go). This documentation is auto-generated, and the location may change.

## License
This library is licensed under Apache 2.0 license. Full license text is available in [LICENSE](https://www.apache.org/licenses/LICENSE-2.0.txt).

## Contact us
In case of issues please reach out to us at the [mailing list](mailto:sdk@nutanix.com)