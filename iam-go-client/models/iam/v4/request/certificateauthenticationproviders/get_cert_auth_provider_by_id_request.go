package certificateauthenticationproviders

// This file holds the request struct for the GetCertAuthProviderById operation.

type GetCertAuthProviderByIdRequest struct {
	// (required) UUID V5 created for the certificate-based authentication provider.
	ExtId *string
}
