package recoveryplans

// This file holds the request struct for the GetNetworkMappingById operation.

type GetNetworkMappingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) Network mapping external identifier.
	ExtId *string
}
