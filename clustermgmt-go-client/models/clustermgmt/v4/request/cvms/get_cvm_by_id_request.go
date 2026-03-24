package cvms

// This file holds the request struct for the GetCvmById operation.

type GetCvmByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) External identifier for the CVM.
	ExtId *string
}
