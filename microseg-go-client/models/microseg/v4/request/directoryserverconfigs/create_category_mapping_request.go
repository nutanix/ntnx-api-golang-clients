package directoryserverconfigs

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the CreateCategoryMapping operation.

type CreateCategoryMappingRequest struct {
	// (required) Creates the mapping between a group in Active Directory and the Category.
	Body *import1.CategoryMapping
}
