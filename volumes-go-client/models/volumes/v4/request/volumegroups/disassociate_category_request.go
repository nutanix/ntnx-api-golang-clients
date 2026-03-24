package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the DisassociateCategory operation.

type DisassociateCategoryRequest struct {
	// (required) The external identifier of a Volume Group.
	ExtId *string

	// (required) The list of categories to be associated/disassociated with the Volume Group. This is a mandatory field.
	Body *import1.CategoryEntityReferences
}
