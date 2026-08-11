package vpcs

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the ShareVpcById operation.

type ShareVpcByIdRequest struct {
	// (required) The UUID of the VPC.
	ExtId *string

	// (required) Request to share a VPC with a project.
	Body *import4.ProjectReference
}
