package vmguestcustomizationprofiles

import (
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
)

// This file holds the request struct for the UpdateVmGuestCustomizationProfileById operation.

type UpdateVmGuestCustomizationProfileByIdRequest struct {
	// (required) A globally unique identifier of a VM Guest Customization Profile in UUID format.
	ExtId *string

	// (required)
	Body *import19.VmGuestCustomizationProfile
}
