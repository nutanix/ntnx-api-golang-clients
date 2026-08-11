package certificatemanager

// This file holds the request struct for the GetCsrById operation.

type GetCsrByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate Signing Request (CSR).
	ExtId *string
}
