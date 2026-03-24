package recoverypoints

import (
	import7 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/content"
)

// This file holds the request struct for the DiscoverClusterForRecoveryPointId operation.

type DiscoverClusterForRecoveryPointIdRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	ExtId *string

	// (required) Request body containing recovery point specifications for discovering the cluster.
	Body *import7.ClusterDiscoverSpec
}
