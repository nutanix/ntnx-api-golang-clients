package remoteentities

// This file holds the request struct for the GetRemoteVtepGatewayForClusterById operation.

type GetRemoteVtepGatewayForClusterByIdRequest struct {
	// (required) Reference to the Prism Central cluster.
	ClusterExtId *string

	// (required) Reference to the specified remote VTEP gateway.
	ExtId *string
}
