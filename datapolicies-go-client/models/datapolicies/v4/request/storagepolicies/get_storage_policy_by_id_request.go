package storagepolicies

// This file holds the request struct for the GetStoragePolicyById operation.

type GetStoragePolicyByIdRequest struct {
	// (required) The external identifier of the Storage Policy.
	ExtId *string
}
