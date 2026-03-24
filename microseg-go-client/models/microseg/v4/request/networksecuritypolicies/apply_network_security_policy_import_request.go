package networksecuritypolicies

// This file holds the request struct for the ApplyNetworkSecurityPolicyImport operation.

type ApplyNetworkSecurityPolicyImportRequest struct {
	// (required)
	Path *string

	// The header parameter specifies if the existing policies need to be deleted or retained upon network security policy
	// import.
	NTNXPurgePolicies *bool

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
