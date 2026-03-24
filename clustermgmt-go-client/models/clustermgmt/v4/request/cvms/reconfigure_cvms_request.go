package cvms

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the ReconfigureCvms operation.

type ReconfigureCvmsRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Request body for reconfiguring CVMs in a cluster.
	Body *import1.CvmReconfigurationSpec
}
