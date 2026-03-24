package sslcertificate

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateSSLCertificate operation.

type UpdateSSLCertificateRequest struct {
	// (required) UUID of the cluster to which the host NIC belongs.
	ClusterExtId *string

	// (required)
	Body *import1.SSLCertificate
}
