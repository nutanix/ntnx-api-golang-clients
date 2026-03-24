package networkcontrollers

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateNetworkController operation.

type CreateNetworkControllerRequest struct {
	// (required) Network controller to create
	Body *import4.NetworkController
}
