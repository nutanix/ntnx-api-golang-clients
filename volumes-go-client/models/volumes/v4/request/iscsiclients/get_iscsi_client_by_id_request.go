package iscsiclients

// This file holds the request struct for the GetIscsiClientById operation.

type GetIscsiClientByIdRequest struct {
	// (required) The external identifier of an iSCSI client.
	ExtId *string
}
