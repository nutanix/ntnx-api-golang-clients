package inventory

import (
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
)

// This file holds the request struct for the PerformInventory operation.

type PerformInventoryRequest struct {
	// The cluster UUID on which the resource is present or the operation is being performed.
	XClusterId *string

	//
	Body *import7.InventorySpec

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
