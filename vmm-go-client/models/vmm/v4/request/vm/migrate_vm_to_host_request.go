package vm

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the MigrateVmToHost operation.

type MigrateVmToHostRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required) The request body for migrating a VM from one host to another within a cluster.
	Body *import19.VmMigrateToHostParams
}
