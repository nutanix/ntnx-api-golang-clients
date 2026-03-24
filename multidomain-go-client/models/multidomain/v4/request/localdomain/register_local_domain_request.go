package localdomain

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/management"
)

// This file holds the request struct for the RegisterLocalDomain operation.

type RegisterLocalDomainRequest struct {
	// (required) Request body to register a domain.
	Body *import3.LocalDomainRegistrationSpec
}
