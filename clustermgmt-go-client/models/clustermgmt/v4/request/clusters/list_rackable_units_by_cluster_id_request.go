package clusters

// This file holds the request struct for the ListRackableUnitsByClusterId operation.

type ListRackableUnitsByClusterIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string
}
