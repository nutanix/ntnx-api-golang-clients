package encryption

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the RotateEncryptionKeys operation.

type RotateEncryptionKeysRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) Rotate the encryption key on a cluster. This can be performed for SOFTWARE encryption type.
	Body *import1.EncryptionKeyRotationSpec
}
