package protectionpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the CreateProtectionPolicy operation.

type CreateProtectionPolicyRequest struct {
	// (required) A protection policy automates the process of creating and replicating recovery points. When a protection policy is
	// configured to create local recovery points, the request includes:<br> - The recovery point objective<br> - The retention
	// policy<br> - The entities that need to be protected by specifying the categories in which they are tagged.<br> To
	// automate recovery point replication, you can also specify the replication location(s). Only users with legacy roles,
	// such as admin, can create a Cross-AZ protection policy.
	Body *import1.ProtectionPolicy
}
