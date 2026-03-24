package networkfunctions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateNetworkFunctionById operation.

type UpdateNetworkFunctionByIdRequest struct {
	// (required) The UUID of the network function.
	ExtId *string

	// (required) Request schema to update the specified network function.
	Body *import4.NetworkFunction
}
