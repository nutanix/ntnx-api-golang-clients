package bmc

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateBmcInfo operation.

type UpdateBmcInfoRequest struct {
	// (required) UUID of the cluster.
	ClusterExtId *string

	// (required) UUID of the host.
	ExtId *string

	// (required) Update BMC info body
	Body *import1.BmcInfo
}
