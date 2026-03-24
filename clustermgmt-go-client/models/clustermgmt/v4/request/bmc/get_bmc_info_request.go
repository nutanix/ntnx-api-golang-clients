package bmc

// This file holds the request struct for the GetBmcInfo operation.

type GetBmcInfoRequest struct {
	// (required) UUID of the cluster.
	ClusterExtId *string

	// (required) UUID of the host.
	ExtId *string
}
