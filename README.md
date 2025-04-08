# Golang Client Libraries for Nutanix APIs

This project contains Golang Client Libraries for Nutanix APIs  grouped together by their namespace. Clients are currently 
available for the following namespaces.

| Namespace    | Description                                                                                         |
|--------------|-----------------------------------------------------------------------------------------------------|
| vmm          | Manage the life-cycle of virtual machines hosted on Nutanix clusters.                                        |
| prism        | Manage Tasks, Category Associations, Alerts, Alert policies, Events and Audits.|
| clustermgmt  | Manage Hosts, Clusters, and other Nutanix infrastructure.                                                    |
| aiops        | Manage Nutanix infrastructure using Analysis, Reporting, Capacity Planning, What if Analysis, VM Rightsizing, Troubleshooting, App Discovery, Broad Observability, and Ops Automation through Playbooks.|
| storage      | Manage Volume Groups and Storage Containers hosted on Nutanix clusters.                                     |
| iam          | Manage User Identity and Access                                                                     |
| lcm          | Manage Infrastructure, Software and Firmware Upgrades.                                                                     |
| files        | Manage virtual file servers, create and configure shares for client access, protect them using DR and sync policies, provision storage space and administer security controls.|                                                                    |
| networking   | Manage networking configuration on Nutanix clusters, including AHV and advanced networking.|                                                                    |

# Project Structure
Project contains a top level directory corresponding to each namespace as listed above.The directory name would be
{namespace}-go-client

Each namespace directory contains the following sub-directory/files

|Name                  | Description                                       |  
|----------------------|---------------------------------------------------|
|go.mod                | Root dependency management                        |
|go.sum                | Hash of dependencies                              |
|LICENSE.txt           | License for the client                            |
|README.md             | README  for the client                            |
|api                   | Go source files for api                           |
|client                | Go source files for client                        |
|models                | Go source files for models                        |


## Getting Started

Get the clients  

```shell
$ go get nutanix/ntnx-api-golang-clients/vmm-go-client/v4/...
$ go get nutanix/ntnx-api-golang-clients/prism-go-client/v4/...
$ go get nutanix/ntnx-api-golang-clients/iam-go-client/v4/...
... etc ...
```

and use them as:

```go
package main

import (
    "github.com/nutanix/ntnx-api-golang-client/vmm-go-client/v4/client"
)
var (
    ApiClientInstance *client.ApiClient
)

func main() {
    ApiClientInstance = client.NewApiClient()
    ApiClientInstance.Host = "10.19.50.27" // IPv4/IPv6 address or FQDN of the cluster
    ApiClientInstance.Port = 9440 // Port to which to connect to
    ApiClientInstance.Username = "admin" // UserName to connect to the cluster
    ApiClientInstance.Password = "password" // Password to connect to the cluster
    // ...
}
```

For detailed instructions on installing individual clients, please refer to the README documentation for the respective clients in the namespace directories.

### Installing the  latest version
In order to install the latest version of the Go client for a Nutanix namespace

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/...
```

### Installing specific versions
In order to install a specific version of the Go client for a Nutanix namespace, choose the version from the provided tags

In order to install 4.0.1-alpha.1 of the vmm-go-client

```shell
$ go get github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/...@v4.0.1-alpha.1
```

## Status
These are auto generated Golang clients generated from Open API v3.0 yaml specification documents.
Due to the auto-generated nature of these clients, they may contain breaking changes from one release to
the next.

## API Reference
These clients have a full set of [API Reference Documentation](https://developers.nutanix.com/). This documentation is auto-generated, and the location may change.

## License
This library is licensed under Nutanix proprietary license. Full license text is available in [LICENSE](https://developers.nutanix.com/license).

## Contact us
In case of issues, please reach out to us at the [mailing list](mailto:sdk@nutanix.com).
