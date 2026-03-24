package vmstartuppolicies

// This file holds the request struct for the GetVmStartupPolicyStartConditionConflictById operation.

type GetVmStartupPolicyStartConditionConflictByIdRequest struct {
	// (required) The external ID of the VM startup policy.
	VmStartupPolicyExtId *string

	// (required) The external ID of the start condition conflict of a VM startup policy.
	ExtId *string
}
