package clusters

// This file holds the request struct for the GetSnmpUserById operation.

type GetSnmpUserByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP user UUID.
	ExtId *string
}
