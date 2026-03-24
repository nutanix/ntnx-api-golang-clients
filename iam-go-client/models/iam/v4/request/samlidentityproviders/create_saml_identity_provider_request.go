package samlidentityproviders

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the CreateSamlIdentityProvider operation.

type CreateSamlIdentityProviderRequest struct {
	// (required) Information for the create SAML identity provider request.
	Body *import4.SamlIdentityProvider
}
