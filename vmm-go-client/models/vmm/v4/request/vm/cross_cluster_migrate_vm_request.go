package vm

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the CrossClusterMigrateVm operation.

type CrossClusterMigrateVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required) Input on how to migrate a VM across clusters.
	Body *import19.VmCrossClusterMigrateParams

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
