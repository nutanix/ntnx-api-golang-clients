package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the FetchNodeNetworkingDetails operation.

type FetchNodeNetworkingDetailsRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Node specific details required to fetch node networking information.
	Body *import1.NodeDetails
}
