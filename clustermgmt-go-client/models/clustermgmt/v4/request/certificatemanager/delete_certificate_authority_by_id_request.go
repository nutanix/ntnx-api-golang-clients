package certificatemanager

// This file holds the request struct for the DeleteCertificateAuthorityById operation.

type DeleteCertificateAuthorityByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate Authority (CA).
	ExtId *string
}
