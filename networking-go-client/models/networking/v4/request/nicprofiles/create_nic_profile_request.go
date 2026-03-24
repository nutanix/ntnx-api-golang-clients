package nicprofiles

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateNicProfile operation.

type CreateNicProfileRequest struct {
	// (required) Request to create a NIC Profile.
	Body *import4.NicProfile
}
