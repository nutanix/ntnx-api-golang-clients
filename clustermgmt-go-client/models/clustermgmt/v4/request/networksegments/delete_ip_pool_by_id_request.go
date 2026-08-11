package networksegments

// This file holds the request struct for the DeleteIpPoolById operation.

type DeleteIpPoolByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) IP Pool UUID.
	ExtId *string
}
