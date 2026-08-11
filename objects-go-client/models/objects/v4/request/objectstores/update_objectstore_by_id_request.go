package objectstores

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
)

// This file holds the request struct for the UpdateObjectstoreById operation.

type UpdateObjectstoreByIdRequest struct {
	// (required) The UUID of the Object store.
	ExtId *string

	// (required) Update the Object store.
	// **For UNDEPLOYED_OBJECT_STORE:** - To save configuration changes: Set state to `UNDEPLOYED_OBJECT_STORE` and provide the
	// updated configuration. - To start deployment: Set state to `DEPLOYING_OBJECT_STORE` and ensure all required deployment
	// parameters are provided.
	// **For OBJECT_STORE_DEPLOYMENT_FAILED:** - To retry deployment: Retrieve the current Object store configuration and pass
	// it in the request body with no changes. Only the state field can be modified to `DEPLOYING_OBJECT_STORE` - all other
	// configuration attributes must remain the same as the failed deployment.
	Body *import1.ObjectStore
}
