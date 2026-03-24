package vpcs

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateVpcById operation.

type UpdateVpcByIdRequest struct {
	// (required) The UUID of the VPC.
	ExtId *string

	// (required) Request schema to update the specified VPC.
	Body *import4.Vpc
}
