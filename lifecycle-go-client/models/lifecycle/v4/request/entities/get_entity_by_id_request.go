package entities

// This file holds the request struct for the GetEntityById operation.

type GetEntityByIdRequest struct {
	// (required) ExtId of the LCM entity.
	ExtId *string
}
