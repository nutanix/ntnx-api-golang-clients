package vmprofile

import (
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the UpdateVmProfileById operation.

type UpdateVmProfileByIdRequest struct {
	// (required) The external ID (UUID) of the VM Profile.
	ExtId *string

	// (required)
	Body *import21.VmProfile
}
