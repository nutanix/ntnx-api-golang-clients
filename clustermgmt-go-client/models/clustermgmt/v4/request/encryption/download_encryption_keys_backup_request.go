package encryption

// This file holds the request struct for the DownloadEncryptionKeysBackup operation.

type DownloadEncryptionKeysBackupRequest struct {
	// (required) ID of the encryption key backup. This ID can be fetched as part of the encryption task response.
	ExtId *string
}
