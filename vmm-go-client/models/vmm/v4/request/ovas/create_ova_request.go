package ovas

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the CreateOva operation.

type CreateOvaRequest struct {
	// (required) Parameters to create an OVA request.
	Body *import9.Ova
}
