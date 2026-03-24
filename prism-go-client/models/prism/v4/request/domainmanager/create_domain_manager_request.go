package domainmanager

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// This file holds the request struct for the CreateDomainManager operation.

type CreateDomainManagerRequest struct {
	// (required) Request body to deploy a Prism Central.
	Body *import3.DomainManager
}
