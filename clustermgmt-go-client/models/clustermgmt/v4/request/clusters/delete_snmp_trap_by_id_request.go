package clusters

// This file holds the request struct for the DeleteSnmpTrapById operation.

type DeleteSnmpTrapByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP trap UUID.
	ExtId *string
}
