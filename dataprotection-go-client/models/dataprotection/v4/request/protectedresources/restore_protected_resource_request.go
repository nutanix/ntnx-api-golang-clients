package protectedresources

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
)

// This file holds the request struct for the RestoreProtectedResource operation.

type RestoreProtectedResourceRequest struct {
	// (required) The external identifier of a protected VM or volume group used to retrieve the protected resource.
	ExtId *string

	// (required) Restore action specifications for a minutely scheduled protected resource.
	Body *import1.ProtectedResourceRestoreSpec
}
