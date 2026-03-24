package vm

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the MigrateVmDisks operation.

type MigrateVmDisksRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required)
	Body *import19.DiskMigrationParams
}
