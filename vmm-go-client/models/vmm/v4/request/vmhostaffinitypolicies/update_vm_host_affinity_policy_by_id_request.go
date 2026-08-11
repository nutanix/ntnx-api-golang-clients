package vmhostaffinitypolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the UpdateVmHostAffinityPolicyById operation.

type UpdateVmHostAffinityPolicyByIdRequest struct {
	// (required) The external ID (UUID) of the VM-host affinity policy.
	ExtId *string

	// (required)
	Body *import6.VmHostAffinityPolicy
}
