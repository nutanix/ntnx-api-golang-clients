package ovas

// This file holds the request struct for the GetFileByOvaId operation.

type GetFileByOvaIdRequest struct {
	// (required) The external identifier for an OVA.
	OvaExtId *string
}
