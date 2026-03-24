package samlidentityproviders

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UpdateSamlIdentityProviderById operation.

type UpdateSamlIdentityProviderByIdRequest struct {
	// (required) External identifier of the SAML identity provider.
	ExtId *string

	// (required) Information for the update SAML identity provider request.
	Body *import4.SamlIdentityProvider
}
