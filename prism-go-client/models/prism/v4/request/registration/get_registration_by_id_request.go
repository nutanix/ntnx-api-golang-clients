package registration

// This file holds the request struct for the GetRegistrationById operation.

type GetRegistrationByIdRequest struct {
	// (required) The external identifier of the domain manager (Prism Central) resource.
	DomainManagerExtId *string

	// (required) The external identifier of the domain manager (Prism Central) resource.
	ExtId *string
}
