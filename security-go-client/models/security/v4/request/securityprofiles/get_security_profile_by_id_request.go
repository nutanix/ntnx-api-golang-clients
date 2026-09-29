package securityprofiles

// This file holds the request struct for the GetSecurityProfileById operation.

type GetSecurityProfileByIdRequest struct {
	// (required) External identifier for cluster security profiles status.
	ExtId *string
}
