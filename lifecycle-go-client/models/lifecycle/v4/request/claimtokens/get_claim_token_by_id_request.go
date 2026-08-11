package claimtokens

// This file holds the request struct for the GetClaimTokenById operation.

type GetClaimTokenByIdRequest struct {
	// (required) External ID of the claim token
	ExtId *string
}
