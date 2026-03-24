package registereddomains

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the UpdateRegisteredDomainById operation.

type UpdateRegisteredDomainByIdRequest struct {
	// (required) External identifier of the registered domain.
	ExtId *string

	// (required) Request body to update the registered domain.
	Body *import1.RegisteredDomain
}
