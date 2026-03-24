package credentials

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the UpdateCredentialById operation.

type UpdateCredentialByIdRequest struct {
	// (required) ExtId of the Credential being operated on.
	ExtId *string

	// (required) Body containing the updated Credential entity.
	Body *import3.Credential
}
