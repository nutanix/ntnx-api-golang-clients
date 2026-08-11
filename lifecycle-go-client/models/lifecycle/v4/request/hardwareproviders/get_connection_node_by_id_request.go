package hardwareproviders

// This file holds the request struct for the GetConnectionNodeById operation.

type GetConnectionNodeByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ConnectionExtId *string

	// (required) External ID of the discovered node
	ExtId *string
}
