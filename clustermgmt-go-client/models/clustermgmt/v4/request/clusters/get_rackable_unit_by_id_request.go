package clusters

// This file holds the request struct for the GetRackableUnitById operation.

type GetRackableUnitByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Rackable unit UUID.
	ExtId *string
}
