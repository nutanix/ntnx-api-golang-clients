# Go Client For Nutanix Nutanix Prism Versioned APIs

The Go client for Nutanix Prism Versioned APIs is designed for Go client application developers offering them simple and flexible access to APIs that task Management, Category Associations, Prism Central DR, Alerts, Alert policies, Events and Audits.

## Features
- Invoke Nutanix APIs with a simple interface.
- Handle Authentication seamlessly.
- Reduce boilerplate code implementation.
- Use standard methods for installation.

## Version
- API version: v4.0.a1
- Package version: v4.0.1-alpha.1

## Requirements.
Go 1.11 or above are fully supported and tested.


## Installation & Usage

### Installation

#### Using go get

##### Install the latest version

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/...
```

##### Install a specific version

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/...@v4.0.1-alpha.1
```

#### Using go modules

##### Install the latest version

In go.mod file add the following

```
module your-module

go {GO_VERSION}

require (
	github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4 latest
)
```

##### Install a specific version

In go.mod file add the following

```
module your-module

go {GO_VERSION}

require (
	github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4 v4.0.1-alpha.1
)
```


## Configuration
The Go client for Nutanix Prism Versioned APIs can be configured with the following parameters

| Parameter | Description                                                                      | Required | Default Value|
|-----------|----------------------------------------------------------------------------------|----------|--------------|
| Host      | IPv4/IPv6 address or FQDN of the cluster to which the client will connect to     | Yes      | N/A          |
| Port      | Port on the cluster to which the client will connect to                          | No       | 9440         |
| Username  | Username to connect to a cluster                                                 | Yes      | N/A          |
| Password  | Password to connect to a cluster                                                 | Yes      | N/A          |
| Debug     | Runs the client in debug mode if specified                                       | No       | False        |
| VerifySSL | Verify SSL certificate of cluster, the client will connect to                    | No       | True         |
| MaxRetryAttempts| Maximum number of retry attempts while connecting to the cluster           | No       | 5            |
| RetryInterval| Interval in milliseconds at which retry attempts are made                     | No       | 3000         |
| LoggerFile | File location to which debug logs are written to                                | No       | N/A|
| Timeout | Global timeout in milliseconds for all operations                                  | No       | 30000        |

### Sample Configuration
```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
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


### Authentication
Nutanix APIs currently support HTTP Basic Authentication only, and the Go client can be configured using the username and password parameters to send Basic headers along with every request.

### Retry Mechanism
The Go client can be configured to retry requests that fail with the following status codes. The numbers of seconds before which the next retry is attempted is determined by the RetryInterval.

- [408 - Request Timeout](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408)
- [502 - Bad Gateway](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502)
- [503 - Service Unavailable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
)
var (
	ApiClientInstance *client.ApiClient
)

ApiClientInstance = client.NewApiClient()
ApiClientInstance.MaxRetryAttempts = 5 // Max retry attempts while reconnecting on a loss of connection
ApiClientInstance.RetryInterval = 5000 // Interval in ms to use during retry attempts
```

## Usage

### Invoking an operation
```go

import (
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/api"
)

var (
	ApiClientInstance *client.ApiClient
	SystemDefinedPoliciesApiInstance *api.SystemDefinedPoliciesApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client as shown in the previous step
// ...

// Initialize the API
SystemDefinedPoliciesApiInstance = api.NewSystemDefinedPoliciesApi(ApiClientInstance)
entityUid := "string_sample_data"
globalConfig := true

// 
getResponse, err := SystemDefinedPoliciesApiInstance.GetSdaPolicyById(entityUid, globalConfig)
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
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
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
import (
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/api"
    // import request body DTO for put api
)

var (
	ApiClientInstance *client.ApiClient
	SystemDefinedPoliciesApiInstance *api.SystemDefinedPoliciesApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client as shown in a previous step
// ...

// Initialize the API
SystemDefinedPoliciesApiInstance = api.NewSystemDefinedPoliciesApi(ApiClientInstance)
entityUid := "string_sample_data"
globalConfig := true

// 
getResponse, err := SystemDefinedPoliciesApiInstance.GetSdaPolicyById(entityUid, globalConfig)
if err != nil {
    ....
}

// Extract E-Tag Header
etagValue := ApiClientInstance.GetEtag(getResponse)

args := make(map[string] interface {})
args["If-Match"] = etagValue
// ...
// Perform update call with received E-Tag reference
// initialize/change parameters for update
}
systemDefined := getResponse.GetData().(import1.SystemDefined)

// The body parameter in the following operation is received from the previous GET request's response which needs to be updated.
response, err := SystemDefinedPoliciesApiInstance.UpdateSdaPolicy(systemDefined, entityUid, args)
if err != nil {
....
}
```

### List Operations
List Operations for Nutanix APIs support pagination, filtering, sorting and projections. The table below details the parameters that can be used to set the options for pagination etc.

| Parameter | Description
|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| _page     | specifies the page number of the result set. Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will be set to its nearest bound. In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page.|
| _limit    | specifies the total number of records returned in the result set. Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. |
| _filter   | allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the [OData V4.01 URL](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter) conventions. |
| _orderby  | allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. |
| _select   | allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. |
| _expand   | allows clients to request related resources when a resource that satisfies a particular request is retrieved. Each expand item is evaluated relative to the entity containing the property being expanded. Other query options can be applied to an expanded property by appending a semicolon-separated list of query options, enclosed in parentheses, to the property name. Allowed system query options are $filter,$select, $orderby. |

```go
import (
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/api"
)
var (
	ApiClientInstance *client.ApiClient
	AlertsApiInstance *api.AlertsApi
)

ApiClientInstance = client.NewApiClient()
// Configure the client
// ...

// Initialize the API
AlertsApiInstance = api.NewAlertsApi(ApiClientInstance)
$page := 0
$limit := 50
$filter := "string_sample_data"
$orderby := "string_sample_data"

// 
response, err := AlertsApiInstance.GetAlerts($page, $limit, $filter, $orderby)
if err != nil {
    ....
}


```
The list of filterable and sortable fields with expansion keys can be found in the documentation [here](https://developers.nutanix.com/).

## API Reference

This library has a full set of [API Reference Documentation](https://developers.nutanix.com/). This documentation is auto-generated, and the location may change.

## License
This library is licensed under Nutanix proprietary license. Full license text is available in [LICENSE](https://developers.nutanix.com/license).

## Contact us
In case of issues please reach out to us at the [mailing list](@sdk@nutanix.com)