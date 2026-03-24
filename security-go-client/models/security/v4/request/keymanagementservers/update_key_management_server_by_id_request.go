package keymanagementservers

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the UpdateKeyManagementServerById operation.

type UpdateKeyManagementServerByIdRequest struct {
	// (required) Unique identifier for the key management server of type UUID.
	ExtId *string

	// (required) Parameters to update a key management server.
	Body *import3.KeyManagementServer
}
