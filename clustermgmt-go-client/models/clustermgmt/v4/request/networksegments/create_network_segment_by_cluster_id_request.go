package networksegments

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateNetworkSegmentByClusterId operation.

type CreateNetworkSegmentByClusterIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Network Segment configuration to create.
	Body *import1.NetworkSegment
}
