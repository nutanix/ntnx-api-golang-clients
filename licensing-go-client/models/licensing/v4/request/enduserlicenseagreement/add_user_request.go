package enduserlicenseagreement

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/agreements"
)

// This file holds the request struct for the AddUser operation.

type AddUserRequest struct {
	// (required)
	Body *import1.EndUser
}
