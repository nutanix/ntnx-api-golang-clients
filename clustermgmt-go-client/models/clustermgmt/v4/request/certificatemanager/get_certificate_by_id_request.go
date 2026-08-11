package certificatemanager

// This file holds the request struct for the GetCertificateById operation.

type GetCertificateByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate.
	ExtId *string
}
