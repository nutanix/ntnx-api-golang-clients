package recoverypointstores

// This file holds the request struct for the GetRecoveryPointStoreById operation.

type GetRecoveryPointStoreByIdRequest struct {
	// (required) External identifier of the recovery point store to which the recovery point should be replicated. This is used for
	// Multicloud Snapshot Technology (MST) replication to object storage.
	ExtId *string
}
