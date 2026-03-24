package bgpsessions

// This file holds the request struct for the DeleteBgpSessionById operation.

type DeleteBgpSessionByIdRequest struct {
	// (required) BGP session UUID.
	ExtId *string
}
