package membership

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
)

// This file holds the request struct for the CreateDomainMembership operation.

type CreateDomainMembershipRequest struct {
	// (required)
	Body *import1.Membership
}
