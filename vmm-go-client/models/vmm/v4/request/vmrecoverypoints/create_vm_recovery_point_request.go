package vmrecoverypoints

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the CreateVmRecoveryPoint operation.

type CreateVmRecoveryPointRequest struct {
	// (required)
	Body *import19.VmRecoveryPoint
}
