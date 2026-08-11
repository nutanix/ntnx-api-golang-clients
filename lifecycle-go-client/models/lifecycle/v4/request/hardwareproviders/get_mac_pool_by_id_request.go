package hardwareproviders

// This file holds the request struct for the GetMacPoolById operation.

type GetMacPoolByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ConnectionExtId *string

	// (required) External ID of the MAC address pool
	ExtId *string
}
