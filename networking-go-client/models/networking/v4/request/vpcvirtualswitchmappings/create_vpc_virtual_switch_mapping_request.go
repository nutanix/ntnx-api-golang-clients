package vpcvirtualswitchmappings

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateVpcVirtualSwitchMapping operation.

type CreateVpcVirtualSwitchMappingRequest struct {
	// (required) Request body for setting VPC for virtual switch mappings traffic config.
	Body *[]import4.VpcVirtualSwitchMapping
}
