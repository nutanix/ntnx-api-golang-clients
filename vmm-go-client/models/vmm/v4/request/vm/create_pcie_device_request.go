package vm

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the CreatePcieDevice operation.

type CreatePcieDeviceRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required)
	Body *import19.PcieDevice
}
