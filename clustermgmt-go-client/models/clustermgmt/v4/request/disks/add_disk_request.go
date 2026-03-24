package disks

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the AddDisk operation.

type AddDiskRequest struct {
	// (required) The external identifier of the cluster on which Disk will be added.
	ExtId *string

	// (required) Request model to add a disk to a cluster.
	Body *import1.DiskAdditionSpec
}
