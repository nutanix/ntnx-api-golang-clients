package domainmanager

// This file holds the request struct for the GetProductById operation.

type GetProductByIdRequest struct {
	// (required) The external identifier of the domain manager (Prism Central) resource.
	DomainManagerExtId *string

	// (required) The product ID for a given product. It can be retrieved using the list request.
	ExtId *string
}
