package iscsiclients

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the UpdateIscsiClientById operation.

type UpdateIscsiClientByIdRequest struct {
	// (required) The external identifier of an iSCSI client.
	ExtId *string

	// (required) A model that represents an iSCSI client that can be associated with a Volume Group as an external attachment.
	Body *import1.IscsiClient
}
