package vmantiaffinitypolicies

import (
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the UpdateVmAntiAffinityPolicyById operation.

type UpdateVmAntiAffinityPolicyByIdRequest struct {
	// (required) A globally unique identifier of a VM-VM anti-affinity policy of type UUID.
	ExtId *string

	// (required)
	Body *import21.VmAntiAffinityPolicy
}
