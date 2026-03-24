package remoteentities

// This file holds the request struct for the GetRemoteVpnConnectionForClusterById operation.

type GetRemoteVpnConnectionForClusterByIdRequest struct {
	// (required) Reference to the Prism Central cluster.
	ClusterExtId *string

	// (required) Reference to the specified remote VPN connection.
	ExtId *string
}
