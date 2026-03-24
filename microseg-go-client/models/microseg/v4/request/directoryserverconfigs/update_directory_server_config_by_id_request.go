package directoryserverconfigs

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateDirectoryServerConfigById operation.

type UpdateDirectoryServerConfigByIdRequest struct {
	// (required) UUID to specify Directory Server.
	ExtId *string

	// (required) Updates the Directory Server Config with the provided ExtID.
	Body *import1.DirectoryServerConfig
}
