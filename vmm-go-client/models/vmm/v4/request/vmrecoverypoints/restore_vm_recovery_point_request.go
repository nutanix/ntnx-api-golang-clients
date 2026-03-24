package vmrecoverypoints

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the RestoreVmRecoveryPoint operation.

type RestoreVmRecoveryPointRequest struct {
	// (required) A globally unique identifier of a VM recovery point. It should be of type UUID.
	ExtId *string

	// (required)
	Body *import19.RestoreVmRecoveryPointParams
}
