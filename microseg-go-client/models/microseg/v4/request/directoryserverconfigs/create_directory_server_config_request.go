package directoryserverconfigs

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the CreateDirectoryServerConfig operation.

type CreateDirectoryServerConfigRequest struct {
	// (required) Configures various aspects of identity categorization.
	Body *import1.DirectoryServerConfig
}
