package vmantiaffinitypolicies

// This file holds the request struct for the DeleteVmAntiAffinityPolicyById operation.

type DeleteVmAntiAffinityPolicyByIdRequest struct {
	// (required) A globally unique identifier of a VM-VM anti-affinity policy of type UUID.
	ExtId *string
}
