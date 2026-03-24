package vpcs

// This file holds the request struct for the GetVpcById operation.

type GetVpcByIdRequest struct {
	// (required) The UUID of the VPC.
	ExtId *string
}
