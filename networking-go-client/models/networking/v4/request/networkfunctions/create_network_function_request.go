package networkfunctions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateNetworkFunction operation.

type CreateNetworkFunctionRequest struct {
	// (required) Request schema to create the network function.
	Body *import4.NetworkFunction
}
