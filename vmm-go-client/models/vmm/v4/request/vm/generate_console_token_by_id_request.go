package vm

// This file holds the request struct for the GenerateConsoleTokenById operation.

type GenerateConsoleTokenByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
