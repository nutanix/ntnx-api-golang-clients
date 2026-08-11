package entities

// This file holds the request struct for the GetEntityById operation.

type GetEntityByIdRequest struct {
	// (required) The external identifier (UUID) of the LCM entity to retrieve.
	ExtId *string
}
