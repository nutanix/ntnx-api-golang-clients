package guestrebootpolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the UpdateGuestRebootPolicyById operation.

type UpdateGuestRebootPolicyByIdRequest struct {
	// (required) A globally unique identifier of a Guest Reboot Policy in UUID format.
	ExtId *string

	// (required)
	Body *import6.GuestRebootPolicy
}
