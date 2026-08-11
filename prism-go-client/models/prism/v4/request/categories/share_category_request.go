package categories

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// This file holds the request struct for the ShareCategory operation.

type ShareCategoryRequest struct {
	// (required) A globally unique identifier of an instance that is suitable for external consumption.
	CategoryExtId *string

	// (required)
	Body *import3.ShareCategoryRequest
}
