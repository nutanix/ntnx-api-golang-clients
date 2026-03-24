package clusters

// This file holds the request struct for the DeleteRsyslogServerById operation.

type DeleteRsyslogServerByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) RSYSLOG server UUID.
	ExtId *string
}
