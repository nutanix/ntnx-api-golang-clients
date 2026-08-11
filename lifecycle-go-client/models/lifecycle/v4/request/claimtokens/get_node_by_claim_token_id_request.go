package claimtokens

// This file holds the request struct for the GetNodeByClaimTokenId operation.

type GetNodeByClaimTokenIdRequest struct {
	// (required) External ID of the claim token
	ClaimTokenExtId *string

	// (required) External ID of the node
	ExtId *string
}
