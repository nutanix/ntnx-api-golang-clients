package encryption

// This file holds the request struct for the GetEncryptionConfig operation.

type GetEncryptionConfigRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string
}
