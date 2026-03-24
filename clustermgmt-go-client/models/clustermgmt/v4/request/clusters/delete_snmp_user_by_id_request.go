package clusters

// This file holds the request struct for the DeleteSnmpUserById operation.

type DeleteSnmpUserByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP user UUID.
	ExtId *string
}
