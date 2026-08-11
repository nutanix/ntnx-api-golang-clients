package hardwareproviders

// This file holds the request struct for the GetServerIdentityPoolById operation.

type GetServerIdentityPoolByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ConnectionExtId *string

	// (required) External ID of the server identity pool
	ExtId *string
}
