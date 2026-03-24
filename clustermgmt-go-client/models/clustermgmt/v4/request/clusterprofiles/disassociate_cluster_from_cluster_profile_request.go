package clusterprofiles

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the DisassociateClusterFromClusterProfile operation.

type DisassociateClusterFromClusterProfileRequest struct {
	// (required) UUID of the cluster profile.
	ExtId *string

	// (required) Request body that will accept a list of clusters to disassociate from the cluster profile.
	Body *import1.ClusterReferenceListSpec
}
