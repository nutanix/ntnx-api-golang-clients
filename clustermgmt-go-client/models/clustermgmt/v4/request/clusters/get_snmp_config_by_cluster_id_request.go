package clusters

// This file holds the request struct for the GetSnmpConfigByClusterId operation.

type GetSnmpConfigByClusterIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string
}
