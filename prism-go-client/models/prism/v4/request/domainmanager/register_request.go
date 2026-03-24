package domainmanager

import (
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// This file holds the request struct for the Register operation.

type RegisterRequest struct {
	// (required) The external identifier of the domain manager (Prism Central) resource.
	ExtId *string

	// (required) The registration request consists of the remote cluster details. Credentials must be of domain manager (Prism Central)
	// role.
	Body *import5.ClusterRegistrationSpec

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
