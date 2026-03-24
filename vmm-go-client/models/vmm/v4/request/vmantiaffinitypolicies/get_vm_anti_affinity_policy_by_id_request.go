package vmantiaffinitypolicies

// This file holds the request struct for the GetVmAntiAffinityPolicyById operation.

type GetVmAntiAffinityPolicyByIdRequest struct {
	// (required) A globally unique identifier of a VM-VM anti-affinity policy of type UUID.
	ExtId *string
}
