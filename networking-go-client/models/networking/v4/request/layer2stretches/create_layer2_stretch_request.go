package layer2stretches

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateLayer2Stretch operation.

type CreateLayer2StretchRequest struct {
	// (required) Request schema to create the Layer2Stretch configuration.
	Body *import4.Layer2Stretch
}
