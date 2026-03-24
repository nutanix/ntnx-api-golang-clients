package clusters

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/operations"
)

// This file holds the request struct for the EnterHostMaintenance operation.

type EnterHostMaintenanceRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Indicates the UUID of a host.
	ExtId *string

	// (required) Property of the host to be put into maintenance mode.
	Body *import4.EnterHostMaintenanceSpec
}
