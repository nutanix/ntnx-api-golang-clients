package guestrebootpolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the CreateGuestRebootPolicy operation.

type CreateGuestRebootPolicyRequest struct {
	// (required)
	Body *import6.GuestRebootPolicy
}
