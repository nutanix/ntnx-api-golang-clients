package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateSnmpTrapById operation.

type UpdateSnmpTrapByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP trap UUID.
	ExtId *string

	// (required) SNMP trap to update.
	Body *import1.SnmpTrap
}
