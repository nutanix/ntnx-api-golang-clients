package nicprofiles

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the AssociateHostNicToNicProfile operation.

type AssociateHostNicToNicProfileRequest struct {
	// (required) UUID of the NIC Profile.
	ExtId *string

	// (required) Request to associate a Host NIC to a NIC Profile.
	Body *import4.HostNic
}
