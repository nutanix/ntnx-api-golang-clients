package approvalpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/management"
)

// This file holds the request struct for the CreateApprovalPolicy operation.

type CreateApprovalPolicyRequest struct {
	// (required) Details used to create an approval policy.
	Body *import1.ApprovalPolicy
}
