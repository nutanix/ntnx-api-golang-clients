package vmguestcustomizationprofiles

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the CreateVmGuestCustomizationProfile operation.

type CreateVmGuestCustomizationProfileRequest struct {
	// (required)
	Body *import19.VmGuestCustomizationProfile
}
