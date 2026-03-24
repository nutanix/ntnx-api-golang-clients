package clusters

// This file holds the request struct for the GetRsyslogServerById operation.

type GetRsyslogServerByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) RSYSLOG server UUID.
	ExtId *string
}
