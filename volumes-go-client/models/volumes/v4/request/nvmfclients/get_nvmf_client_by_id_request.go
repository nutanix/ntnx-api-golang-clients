package nvmfclients

// This file holds the request struct for the GetNvmfClientById operation.

type GetNvmfClientByIdRequest struct {
	// (required) The external identifier of the NVMe-TCP client.
	ExtId *string
}
