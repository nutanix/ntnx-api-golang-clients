package esxivm

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/esxi/config"
)

// This file holds the request struct for the RevertVm operation.

type RevertVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required) Input for the VM revert operation. Specify the VM Recovery Point ID to which the VM would be reverted.
	Body *import4.RevertParams
}
