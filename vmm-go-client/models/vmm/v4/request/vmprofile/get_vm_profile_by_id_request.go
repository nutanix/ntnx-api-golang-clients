package vmprofile

// This file holds the request struct for the GetVmProfileById operation.

type GetVmProfileByIdRequest struct {
	// (required) The external ID (UUID) of the VM Profile.
	ExtId *string
}
