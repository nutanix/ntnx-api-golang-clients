package nicprofiles

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the DisassociateHostNicFromNicProfile operation.

type DisassociateHostNicFromNicProfileRequest struct {
	// (required) UUID of the NIC Profile.
	ExtId *string

	// (required) Request to disassociate a Host NIC from a NIC Profile.
	Body *import4.HostNic
}
