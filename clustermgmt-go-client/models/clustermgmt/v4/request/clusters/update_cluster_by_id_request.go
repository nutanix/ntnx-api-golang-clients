package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateClusterById operation.

type UpdateClusterByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ExtId *string

	// (required) Cluster resource to update.
	Body *import1.Cluster
}
