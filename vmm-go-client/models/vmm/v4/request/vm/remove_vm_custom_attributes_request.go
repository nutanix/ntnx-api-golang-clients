package vm

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the RemoveVmCustomAttributes operation.

type RemoveVmCustomAttributesRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string

	// (required) A collection of user-defined key/value pairs as strings in the format 'key:value' representing custom attributes of the
	// VM.
	Body *import19.UpdateCustomAttributesParams
}
