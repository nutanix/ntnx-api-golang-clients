package hardwareproviders

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
)

// This file holds the request struct for the CreateConnectionByHardwareProviderId operation.

type CreateConnectionByHardwareProviderIdRequest struct {
	// (required) External ID of the hardware provider
	HardwareProviderExtId *string

	// (required)
	Body *import3.Connection
}
