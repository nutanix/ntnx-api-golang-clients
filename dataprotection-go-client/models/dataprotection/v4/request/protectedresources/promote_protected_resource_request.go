package protectedresources

// This file holds the request struct for the PromoteProtectedResource operation.

type PromoteProtectedResourceRequest struct {
	// (required) The external identifier of a protected VM or volume group used to retrieve the protected resource.
	ExtId *string
}
