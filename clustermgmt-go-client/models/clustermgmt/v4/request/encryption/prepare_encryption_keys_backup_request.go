package encryption

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the PrepareEncryptionKeysBackup operation.

type PrepareEncryptionKeysBackupRequest struct {
	// (required) Encryption key backup specifications.
	Body *import1.EncryptionBackupSpec
}
