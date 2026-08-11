package cvms

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateCvmById operation.

type UpdateCvmByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) External identifier for the CVM.
	ExtId *string

	// (required) Request body for updating a specific CVM.
	Body *import1.Cvm
}
