package recoveryplans

// This file holds the request struct for the GetDataServicesIpMappingById operation.

type GetDataServicesIpMappingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) Data services IP mapping external identifier.
	ExtId *string
}
