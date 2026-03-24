package objectstores

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
)

// This file holds the request struct for the CreateObjectstore operation.

type CreateObjectstoreRequest struct {
	// (required) Object store schema to be deployed.
	Body *import1.ObjectStore
}
