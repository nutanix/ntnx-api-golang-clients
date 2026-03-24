package vpcs

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateVpc operation.

type CreateVpcRequest struct {
	// (required) Request schema to create the VPC.
	Body *import4.Vpc
}
