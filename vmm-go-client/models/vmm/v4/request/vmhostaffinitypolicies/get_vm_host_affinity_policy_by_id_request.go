package vmhostaffinitypolicies

// This file holds the request struct for the GetVmHostAffinityPolicyById operation.

type GetVmHostAffinityPolicyByIdRequest struct {
	// (required) The external ID (UUID) of the VM-host affinity policy.
	ExtId *string
}
