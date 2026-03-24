package sslcertificate

// This file holds the request struct for the GetSSLCertificate operation.

type GetSSLCertificateRequest struct {
	// (required) UUID of the cluster to which the host NIC belongs.
	ClusterExtId *string
}
