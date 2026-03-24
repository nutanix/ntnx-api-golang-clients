package vcenterextensions

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UnregisterVcenterExtension operation.

type UnregisterVcenterExtensionRequest struct {
	// (required) The globally unique identifier of vCenter Server extension instance. It should be of type UUID.
	ExtId *string

	// (required)
	Body *import1.VcenterCredentials
}
