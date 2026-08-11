package templates

import (
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the UpdateTemplateById operation.

type UpdateTemplateByIdRequest struct {
	// (required) The identifier of a template.
	ExtId *string

	// (required) Request to update a template.
	Body *import11.Template
}
