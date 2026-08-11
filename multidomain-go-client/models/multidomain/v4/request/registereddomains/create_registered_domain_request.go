package registereddomains

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the CreateRegisteredDomain operation.

type CreateRegisteredDomainRequest struct {
	// (required) Request body for creating a registered domain.
	Body *import3.RegisteredDomain
}
