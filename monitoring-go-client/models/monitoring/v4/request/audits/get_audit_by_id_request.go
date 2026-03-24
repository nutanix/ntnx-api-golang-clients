package audits

// This file holds the request struct for the GetAuditById operation.

type GetAuditByIdRequest struct {
	// (required) UUID of the generated audit.
	ExtId *string
}
