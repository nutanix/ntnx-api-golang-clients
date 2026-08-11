package bundles

// This file holds the request struct for the DeleteBundleById operation.

type DeleteBundleByIdRequest struct {
	// (required) The external identifier (UUID) of the LCM bundle.
	ExtId *string
}
