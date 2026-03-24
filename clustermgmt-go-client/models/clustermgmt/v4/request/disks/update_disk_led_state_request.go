package disks

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateDiskLEDState operation.

type UpdateDiskLEDStateRequest struct {
	// (required) The external identifier of the Disk.
	ExtId *string

	// (required) Request model for updating LED state of the Disk.
	Body *import1.LEDStateUpdationSpec
}
