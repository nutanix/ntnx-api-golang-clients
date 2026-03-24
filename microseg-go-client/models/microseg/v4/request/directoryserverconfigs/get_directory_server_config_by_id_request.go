package directoryserverconfigs

// This file holds the request struct for the GetDirectoryServerConfigById operation.

type GetDirectoryServerConfigByIdRequest struct {
	// (required) UUID to specify Directory Server.
	ExtId *string
}
