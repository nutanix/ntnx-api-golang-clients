package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateSnmpUserById operation.

type UpdateSnmpUserByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP user UUID.
	ExtId *string

	// (required) SNMP user to update.
	Body *import1.SnmpUser
}
