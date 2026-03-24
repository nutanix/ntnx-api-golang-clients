package networksecuritypolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateNetworkSecurityPolicyById operation.

type UpdateNetworkSecurityPolicyByIdRequest struct {
	// (required) Network security policy UUID.
	ExtId *string

	// (required) Updates the Network Security Policy with the provided ExtID.
	Body *import1.NetworkSecurityPolicy
}
