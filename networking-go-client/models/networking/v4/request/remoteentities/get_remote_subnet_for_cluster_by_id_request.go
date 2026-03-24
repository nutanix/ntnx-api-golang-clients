package remoteentities

// This file holds the request struct for the GetRemoteSubnetForClusterById operation.

type GetRemoteSubnetForClusterByIdRequest struct {
	// (required) Reference to the Prism Central cluster.
	ClusterExtId *string

	// (required) Reference to the specified remote subnet.
	ExtId *string
}
