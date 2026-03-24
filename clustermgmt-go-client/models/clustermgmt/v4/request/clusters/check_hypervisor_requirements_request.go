package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CheckHypervisorRequirements operation.

type CheckHypervisorRequirementsRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Parameters to get information on whether hypervisor ISO upload is required or not.
	Body *import1.HypervisorUploadParam
}
