package passwordmanager

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the ChangeSystemUserPasswordById operation.

type ChangeSystemUserPasswordByIdRequest struct {
	// (required) External identifier of the system user password.
	ExtId *string

	// (required) Contains the required specifications for a password update operation.
	Body *import1.ChangePasswordSpec
}
