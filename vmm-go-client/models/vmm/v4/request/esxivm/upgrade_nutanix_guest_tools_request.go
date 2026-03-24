package esxivm

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/esxi/config"
)

// This file holds the request struct for the UpgradeNutanixGuestTools operation.

type UpgradeNutanixGuestToolsRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required)
	Body *import4.NutanixGuestToolsUpgradeConfig
}
