package vmstartuppolicies

// This file holds the request struct for the GetVmStartupPolicyById operation.

type GetVmStartupPolicyByIdRequest struct {
	// (required) The external ID of the VM startup policy.
	ExtId *string
}
