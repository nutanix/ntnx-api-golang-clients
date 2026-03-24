package protectedresources

// This file holds the request struct for the GetProtectedResourceById operation.

type GetProtectedResourceByIdRequest struct {
	// (required) The external identifier of a protected VM or volume group used to retrieve the protected resource.
	ExtId *string
}
