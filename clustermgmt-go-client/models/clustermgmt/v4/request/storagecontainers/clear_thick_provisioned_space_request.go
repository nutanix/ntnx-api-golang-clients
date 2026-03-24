package storagecontainers

// This file holds the request struct for the ClearThickProvisionedSpace operation.

type ClearThickProvisionedSpaceRequest struct {
	// (required) The external identifier of the Storage Container.
	ExtId *string

	// The external identifier of the remote cluster to which the request is forwarded.
	XClusterId *string
}
