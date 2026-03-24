package alerts

// This file holds the request struct for the GetAlertById operation.

type GetAlertByIdRequest struct {
	// (required) UUID of the generated alert.
	ExtId *string
}
