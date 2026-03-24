package ovas

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the UpdateOvaById operation.

type UpdateOvaByIdRequest struct {
	// (required) The external identifier for an OVA.
	ExtId *string

	// (required) Update the OVA entity.
	Body *import9.Ova
}
