package objectstores

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
)

// This file holds the request struct for the UpdateObjectstoreById operation.

type UpdateObjectstoreByIdRequest struct {
	// (required) The UUID of the Object store.
	ExtId *string

	// (required) Update the Object store. To retry the deployment of an Object store, get the Object store and pass the Object store in
	// the request body unchanged.
	Body *import1.ObjectStore
}
