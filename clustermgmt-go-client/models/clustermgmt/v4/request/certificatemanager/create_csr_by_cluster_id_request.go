package certificatemanager

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateCsrByClusterId operation.

type CreateCsrByClusterIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) Details of a Certificate Signing Request (CSR).
	Body *import1.Csr
}
