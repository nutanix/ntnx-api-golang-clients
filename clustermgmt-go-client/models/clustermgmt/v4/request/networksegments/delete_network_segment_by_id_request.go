package networksegments

// This file holds the request struct for the DeleteNetworkSegmentById operation.

type DeleteNetworkSegmentByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Network Segment UUID.
	ExtId *string
}
