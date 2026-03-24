package keymanagementservers

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the CreateKeyManagementServer operation.

type CreateKeyManagementServerRequest struct {
	// (required) Parameters required to create the key management server.
	Body *import3.KeyManagementServer
}
