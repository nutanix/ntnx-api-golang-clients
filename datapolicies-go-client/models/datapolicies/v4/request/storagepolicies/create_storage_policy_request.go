package storagepolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the CreateStoragePolicy operation.

type CreateStoragePolicyRequest struct {
	// (required) A model that represents a Storage Policy resource.
	Body *import1.StoragePolicy
}
