package layer2stretches

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateLayer2StretchById operation.

type UpdateLayer2StretchByIdRequest struct {
	// (required) External ID of the Layer2Stretch configuration.
	ExtId *string

	// (required) Request schema to update the specified Layer2Stretch configuration.
	Body *import4.Layer2Stretch
}
