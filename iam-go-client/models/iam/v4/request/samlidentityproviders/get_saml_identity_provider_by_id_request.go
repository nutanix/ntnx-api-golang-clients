package samlidentityproviders

// This file holds the request struct for the GetSamlIdentityProviderById operation.

type GetSamlIdentityProviderByIdRequest struct {
	// (required) External identifier of the SAML identity provider.
	ExtId *string
}
