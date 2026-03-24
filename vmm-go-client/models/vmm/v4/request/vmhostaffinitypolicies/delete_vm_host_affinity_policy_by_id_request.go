package vmhostaffinitypolicies

// This file holds the request struct for the DeleteVmHostAffinityPolicyById operation.

type DeleteVmHostAffinityPolicyByIdRequest struct {
	// (required) The external ID (UUID) of the VM-host affinity policy.
	ExtId *string
}
