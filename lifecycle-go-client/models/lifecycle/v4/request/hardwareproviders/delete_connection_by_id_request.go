package hardwareproviders

// This file holds the request struct for the DeleteConnectionById operation.

type DeleteConnectionByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ExtId *string
}
