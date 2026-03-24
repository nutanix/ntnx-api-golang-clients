package vmhostaffinitypolicies

// This file holds the request struct for the ReEnforceVmHostAffinityPolicyById operation.

type ReEnforceVmHostAffinityPolicyByIdRequest struct {
	// (required) The external ID (UUID) of the VM-host affinity policy.
	ExtId *string
}
