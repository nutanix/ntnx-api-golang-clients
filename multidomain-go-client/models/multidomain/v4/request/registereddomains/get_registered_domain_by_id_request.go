package registereddomains

// This file holds the request struct for the GetRegisteredDomainById operation.

type GetRegisteredDomainByIdRequest struct {
	// (required) External identifier of the registered domain.
	ExtId *string

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned.
	Select_ *string
}
