package hardwareproviders

// This file holds the request struct for the GetIpPoolById operation.

type GetIpPoolByIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required) External ID of the connection
	ConnectionExtId *string

	// (required) External ID of the IP address pool
	ExtId *string
}
