package systemdefinedpolicies

// This file holds the request struct for the GetSdaPolicyById operation.

type GetSdaPolicyByIdRequest struct {
	// (required) Unique ID of the System-Defined Alert Policy.
	ExtId *string
}
