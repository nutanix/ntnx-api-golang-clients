package macaddresses

// This file holds the request struct for the GetLearnedMacAddressForLayer2StretchById operation.

type GetLearnedMacAddressForLayer2StretchByIdRequest struct {
	// (required) External ID of the Layer2Stretch configuration.
	Layer2StretchExtId *string

	// (required) External ID of the specified learned MAC address.
	ExtId *string
}
