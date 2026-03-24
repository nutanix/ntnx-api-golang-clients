package clusters

// This file holds the request struct for the GetHostById operation.

type GetHostByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Indicates the UUID of a host.
	ExtId *string
}
