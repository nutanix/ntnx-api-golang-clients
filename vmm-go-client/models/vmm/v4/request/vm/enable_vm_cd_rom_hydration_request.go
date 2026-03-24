package vm

// This file holds the request struct for the EnableVmCdRomHydration operation.

type EnableVmCdRomHydrationRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) A globally unique identifier of a CD-ROM of type UUID.
	ExtId *string
}
