package clusterprofiles

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the ApplyClusterProfile operation.

type ApplyClusterProfileRequest struct {
	// (required) UUID of the cluster profile.
	ExtId *string

	// (required) Request body that will contain a list of clusters to apply to a cluster profile.
	Body *import1.ClusterReferenceListSpec

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
