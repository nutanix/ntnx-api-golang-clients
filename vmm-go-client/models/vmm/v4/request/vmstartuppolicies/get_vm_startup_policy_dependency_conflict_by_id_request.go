package vmstartuppolicies

// This file holds the request struct for the GetVmStartupPolicyDependencyConflictById operation.

type GetVmStartupPolicyDependencyConflictByIdRequest struct {
	// (required) The external ID of the VM startup policy.
	VmStartupPolicyExtId *string

	// (required) The external ID of the Dependency conflict of a VM startup policy.
	ExtId *string
}
