package objectstores

// This file holds the request struct for the GetCaByCertificateId operation.

type GetCaByCertificateIdRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) The UUID of the certificate of an Object store.
	CertificateExtId *string
}
