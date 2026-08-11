package claimtokens

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
)

// This file holds the request struct for the UpdateClaimTokenById operation.

type UpdateClaimTokenByIdRequest struct {
	// (required) External ID of the claim token
	ExtId *string

	// (required)
	Body *import3.ClaimToken
}
