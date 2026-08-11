package claimtokens

// This file holds the request struct for the GetSecretByClaimTokenId operation.

type GetSecretByClaimTokenIdRequest struct {
	// (required) External ID of the claim token
	ExtId *string
}
