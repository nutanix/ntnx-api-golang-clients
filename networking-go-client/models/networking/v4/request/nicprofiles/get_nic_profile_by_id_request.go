package nicprofiles

// This file holds the request struct for the GetNicProfileById operation.

type GetNicProfileByIdRequest struct {
	// (required) UUID of the NIC Profile.
	ExtId *string
}
