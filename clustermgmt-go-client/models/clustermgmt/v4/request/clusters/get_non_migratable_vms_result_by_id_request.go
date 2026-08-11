package clusters

// This file holds the request struct for the GetNonMigratableVmsResultById operation.

type GetNonMigratableVmsResultByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) UUID of the result storing the non-migratable VMs.
	ExtId *string
}
