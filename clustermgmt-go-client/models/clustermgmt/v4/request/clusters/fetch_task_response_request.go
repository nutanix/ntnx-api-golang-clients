package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the FetchTaskResponse operation.

type FetchTaskResponseRequest struct {
	// (required) The external identifier of the task.
	ExtId *string

	// (required) Task Response search type.
	TaskResponseType *import1.TaskResponseType
}
