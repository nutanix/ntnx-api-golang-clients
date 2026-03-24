package templates

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the CompleteGuestUpdate operation.

type CompleteGuestUpdateRequest struct {
	// (required) The identifier of a template.
	ExtId *string

	// (required) Request to complete the ongoing guest update on a template.
	Body *import9.CompleteGuestUpdateSpec
}
