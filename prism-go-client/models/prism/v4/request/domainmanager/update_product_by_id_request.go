package domainmanager

import (
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// This file holds the request struct for the UpdateProductById operation.

type UpdateProductByIdRequest struct {
	// (required) The external identifier of the domain manager (Prism Central) resource.
	DomainManagerExtId *string

	// (required) The product ID for a given product. It can be retrieved using the list request.
	ExtId *string

	// (required)
	Body *import5.Product

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
