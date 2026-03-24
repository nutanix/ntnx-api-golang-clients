package bgpsessions

// This file holds the request struct for the GetBgpSessionById operation.

type GetBgpSessionByIdRequest struct {
	// (required) BGP session UUID.
	ExtId *string
}
