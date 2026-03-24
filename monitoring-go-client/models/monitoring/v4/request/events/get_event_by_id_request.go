package events

// This file holds the request struct for the GetEventById operation.

type GetEventByIdRequest struct {
	// (required) UUID of the generated event.
	ExtId *string
}
