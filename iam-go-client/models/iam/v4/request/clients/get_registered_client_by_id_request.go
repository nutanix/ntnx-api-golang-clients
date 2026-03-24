package clients

// This file holds the request struct for the GetRegisteredClientById operation.

type GetRegisteredClientByIdRequest struct {
	// (required) External identifier of the client.
	ExtId *string
}
