package templates

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the CreateTemplate operation.

type CreateTemplateRequest struct {
	// (required) Request to create a template.
	Body *import9.Template
}
