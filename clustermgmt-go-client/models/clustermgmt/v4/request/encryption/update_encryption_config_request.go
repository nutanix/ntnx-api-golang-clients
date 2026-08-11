package encryption

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateEncryptionConfig operation.

type UpdateEncryptionConfigRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Configure encryption on a cluster.
	Body *import1.EncryptionConfig
}
