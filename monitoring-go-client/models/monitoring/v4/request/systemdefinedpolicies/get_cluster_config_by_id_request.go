package systemdefinedpolicies

// This file holds the request struct for the GetClusterConfigById operation.

type GetClusterConfigByIdRequest struct {
	// (required) Unique ID of the System-Defined Alert Policy.
	SystemDefinedPolicyExtId *string

	// (required) Cluster UUID.
	ExtId *string
}
