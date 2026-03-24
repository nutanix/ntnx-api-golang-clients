package clusters

// This file holds the request struct for the GetVirtualNicById operation.

type GetVirtualNicByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Indicates the UUID of a host.
	HostExtId *string

	// (required) Virtual NIC UUID.
	ExtId *string
}
