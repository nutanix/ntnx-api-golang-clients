package samlidentityproviders

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the ShareSamlIdentityProvider operation.

type ShareSamlIdentityProviderRequest struct {
	// (required) External identifier of the SAML identity provider.
	ExtId *string

	// (required) Shares a SAML identity provider with specified project.
	Body *import4.SamlIdentityProviderShareRequest
}
