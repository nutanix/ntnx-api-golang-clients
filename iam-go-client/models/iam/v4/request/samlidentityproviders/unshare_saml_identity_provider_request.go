package samlidentityproviders

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UnshareSamlIdentityProvider operation.

type UnshareSamlIdentityProviderRequest struct {
	// (required) External identifier of the SAML identity provider.
	ExtId *string

	// (required) Unshares a SAML identity provider from a specific project.
	Body *import4.SamlIdentityProviderUnshareRequest
}
