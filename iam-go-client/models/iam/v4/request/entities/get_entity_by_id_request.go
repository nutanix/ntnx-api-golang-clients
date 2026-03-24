package entities

// This file holds the request struct for the GetEntityById operation.

type GetEntityByIdRequest struct {
	// (required) External identifier for the client entity.
	ExtId *string
}
