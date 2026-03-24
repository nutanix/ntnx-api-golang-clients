package vcenterextensions

// This file holds the request struct for the GetVcenterExtensionById operation.

type GetVcenterExtensionByIdRequest struct {
	// (required) The globally unique identifier of vCenter Server extension instance. It should be of type UUID.
	ExtId *string
}
