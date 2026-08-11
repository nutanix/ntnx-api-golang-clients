package networksegments

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateIpPoolById operation.

type UpdateIpPoolByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) IP Pool UUID.
	ExtId *string

	// (required) IP Pool configuration to be updated.
	Body *import1.IpPool
}
