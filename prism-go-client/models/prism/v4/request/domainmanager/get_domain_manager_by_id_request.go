package domainmanager

// This file holds the request struct for the GetDomainManagerById operation.

type GetDomainManagerByIdRequest struct {
	// (required) The external identifier of the domain manager (Prism Central) resource.
	ExtId *string
}
