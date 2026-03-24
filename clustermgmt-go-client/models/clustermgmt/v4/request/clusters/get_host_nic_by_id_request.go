package clusters

// This file holds the request struct for the GetHostNicById operation.

type GetHostNicByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Indicates the UUID of a host.
	HostExtId *string

	// (required) Indicates the UUID of a host NIC.
	ExtId *string
}
