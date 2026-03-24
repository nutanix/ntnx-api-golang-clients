package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the RemoveSnmpTransport operation.

type RemoveSnmpTransportRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP transports to remove.
	Body *import1.SnmpTransport
}
