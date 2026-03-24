package clusters

// This file holds the request struct for the GetSnmpTrapById operation.

type GetSnmpTrapByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) SNMP trap UUID.
	ExtId *string
}
