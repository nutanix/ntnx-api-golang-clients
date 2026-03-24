package approvalpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/management"
)

// This file holds the request struct for the UpdateApprovalPolicyByExtId operation.

type UpdateApprovalPolicyByExtIdRequest struct {
	// (required) Approval policy external identifier.
	ExtId *string

	// (required)
	Body *import1.ApprovalPolicy
}
