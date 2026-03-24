package vmstartuppolicies

import (
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the UpdateVmStartupPolicyById operation.

type UpdateVmStartupPolicyByIdRequest struct {
	// (required) The external ID of the VM startup policy.
	ExtId *string

	// (required)
	Body *import21.VmStartupPolicy
}
