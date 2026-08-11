package membership

// This file holds the request struct for the GetDomainMembershipById operation.

type GetDomainMembershipByIdRequest struct {
	// (required) External identifier of the membership.
	ExtId *string
}
