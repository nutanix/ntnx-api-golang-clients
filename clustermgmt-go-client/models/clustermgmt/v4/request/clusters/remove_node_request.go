package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the RemoveNode operation.

type RemoveNodeRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Parameters to remove nodes from cluster.
	Body *import1.NodeRemovalParams
}
