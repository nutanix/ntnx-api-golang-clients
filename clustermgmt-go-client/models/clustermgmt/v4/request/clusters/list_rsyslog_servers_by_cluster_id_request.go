package clusters

// This file holds the request struct for the ListRsyslogServersByClusterId operation.

type ListRsyslogServersByClusterIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string
}
