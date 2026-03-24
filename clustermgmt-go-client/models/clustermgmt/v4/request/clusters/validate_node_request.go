package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the ValidateNode operation.

type ValidateNodeRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Request body for node validation. It can be OneOf between hypervisor bundle and node uplinks.
	Body *import1.ValidateNodeParam
}
