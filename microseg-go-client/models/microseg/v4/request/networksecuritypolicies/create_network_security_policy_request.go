package networksecuritypolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the CreateNetworkSecurityPolicy operation.

type CreateNetworkSecurityPolicyRequest struct {
	// (required)
	Body *import1.NetworkSecurityPolicy
}
