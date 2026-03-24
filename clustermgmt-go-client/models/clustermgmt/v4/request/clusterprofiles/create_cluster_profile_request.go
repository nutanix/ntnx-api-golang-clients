package clusterprofiles

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateClusterProfile operation.

type CreateClusterProfileRequest struct {
	// (required) The required parameters to create a cluster profile.
	Body *import1.ClusterProfile
}
