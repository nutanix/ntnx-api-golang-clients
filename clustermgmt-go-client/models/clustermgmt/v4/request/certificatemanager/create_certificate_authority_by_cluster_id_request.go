package certificatemanager

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateCertificateAuthorityByClusterId operation.

type CreateCertificateAuthorityByClusterIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) Description of the Certificate Authority.
	Body *import1.CertificateAuthority
}
