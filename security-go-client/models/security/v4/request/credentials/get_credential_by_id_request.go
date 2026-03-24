package credentials

// This file holds the request struct for the GetCredentialById operation.

type GetCredentialByIdRequest struct {
	// (required) ExtId of the Credential being operated on.
	ExtId *string
}
