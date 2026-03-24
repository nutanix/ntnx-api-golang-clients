package recoverypoints

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
)

// This file holds the request struct for the ReplicateRecoveryPoint operation.

type ReplicateRecoveryPointRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	ExtId *string

	// (required) External identifier of the cluster and the Prism Central where the recovery point is to be replicated. The recovery
	// point identified by `extId` is replicated from the current Prism Central to the remote Prism Central with the external
	// identifier `pcExtId` . This replication allows the data-protection service to decide on which cluster to perform the
	// automatic replication. However, the user also has the option to choose the cluster identified by `clusterExtId` to which
	// the recovery point identified by `extId` should be replicated. Cross-AZ replication can be performed only by users
	// having legacy roles (e.g. `admin`).
	Body *import1.RecoveryPointReplicationSpec
}
