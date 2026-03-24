package recoveryplans

// This file holds the request struct for the DeleteNetworkMappingById operation.

type DeleteNetworkMappingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) Network mapping external identifier.
	ExtId *string
}
