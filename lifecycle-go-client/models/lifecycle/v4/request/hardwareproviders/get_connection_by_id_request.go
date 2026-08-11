package hardwareproviders

// This file holds the request struct for the GetConnectionById operation.

type GetConnectionByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ExtId *string
}
