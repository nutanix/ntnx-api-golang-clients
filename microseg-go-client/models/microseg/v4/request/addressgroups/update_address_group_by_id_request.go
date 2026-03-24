package addressgroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateAddressGroupById operation.

type UpdateAddressGroupByIdRequest struct {
	// (required) Address group UUID.
	ExtId *string

	// (required) Updates the Address Group with the provided ExtID.
	Body *import1.AddressGroup
}
