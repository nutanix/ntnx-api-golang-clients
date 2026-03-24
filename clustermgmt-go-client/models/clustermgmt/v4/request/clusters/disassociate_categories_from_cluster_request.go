package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the DisassociateCategoriesFromCluster operation.

type DisassociateCategoriesFromClusterRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) The specifications required for updating categories to the entity.
	Body *import1.CategoryEntityReferences
}
