package networkcontrollers

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateNetworkControllerById operation.

type UpdateNetworkControllerByIdRequest struct {
	// (required) The UUID of the network controller
	ExtId *string

	// (required) Network controller to update
	Body *import4.NetworkController
}
