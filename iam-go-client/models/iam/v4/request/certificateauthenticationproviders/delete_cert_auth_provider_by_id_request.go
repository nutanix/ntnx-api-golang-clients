package certificateauthenticationproviders

// This file holds the request struct for the DeleteCertAuthProviderById operation.

type DeleteCertAuthProviderByIdRequest struct {
	// (required) UUID V5 created for the certificate-based authentication provider.
	ExtId *string
}
