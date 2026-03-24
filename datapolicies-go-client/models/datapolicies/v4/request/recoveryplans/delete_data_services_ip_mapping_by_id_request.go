package recoveryplans

// This file holds the request struct for the DeleteDataServicesIpMappingById operation.

type DeleteDataServicesIpMappingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) Data services IP mapping external identifier.
	ExtId *string
}
