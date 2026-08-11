package rolemembership

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// This file holds the request struct for the CreateRoleMembership operation.

type CreateRoleMembershipRequest struct {
	// (required) Information for the create role membership request. It requires the role, identityExtId, identityType,
	// scopeTemplateName, scopeTemplateNameValues, and idpExtId attributes.
	Body *import1.RoleMembership
}
