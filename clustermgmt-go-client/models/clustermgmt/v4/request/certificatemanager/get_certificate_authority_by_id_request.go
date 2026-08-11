package certificatemanager

// This file holds the request struct for the GetCertificateAuthorityById operation.

type GetCertificateAuthorityByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate Authority (CA).
	ExtId *string
}
