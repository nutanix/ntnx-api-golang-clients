package credentials

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the CreateCredential operation.

type CreateCredentialRequest struct {
	// (required) Response object containing the created Credential entity.
	Body *import3.Credential
}
