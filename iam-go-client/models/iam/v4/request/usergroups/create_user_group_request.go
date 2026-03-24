package usergroups

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the CreateUserGroup operation.

type CreateUserGroupRequest struct {
	// (required) Information for the create user group request. It includes fields like name, groupType, idpId, etc.
	Body *import4.UserGroup
}
