package certificatemanager

// This file holds the request struct for the DeleteCsrById operation.

type DeleteCsrByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate Signing Request (CSR).
	ExtId *string
}
